package input

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"

	options "github.com/hikyaru-suzuki/go-con-2022-spring-sample/pkg/domain/proto/server/options/transaction"
)

func ConvertMessageFromProto(file *protogen.File) (*Message, error) {
	if len(file.Messages) != 1 {
		return nil, perrors.Newf("proto file must have only one message")
	}

	return convert(file.Messages[0])
}

func convert(message *protogen.Message) (*Message, error) {
	messageOption, ok := proto.GetExtension(message.Desc.Options(), options.E_Message).(*options.MessageOption)
	if !ok {
		return nil, perrors.Newf("type assertion failed")
	}

	messageAccessorType, err := ConvertMessageAccessorTypeFromProto(messageOption.GetAccessorType())
	if err != nil {
		return nil, perrors.Stack(err)
	}

	var indexes []*Index
	var interleave *Interleave
	if messageOption.GetDdl() != nil {
		indexes = make([]*Index, 0, len(messageOption.GetDdl().GetIndexes()))
		for _, index := range messageOption.GetDdl().GetIndexes() {
			keys := make([]*IndexKey, 0, len(index.GetKeys()))
			for _, key := range index.GetKeys() {
				keys = append(keys, &IndexKey{
					SnakeName: core.ToSnakeCase(key.Column),
					Desc:      key.Desc,
				})
			}
			storing := make([]string, 0, len(index.GetStoring()))
			for _, s := range index.GetStoring() {
				storing = append(storing, core.ToSnakeCase(s))
			}

			indexes = append(indexes, &Index{
				Keys:         keys,
				Unique:       index.Unique,
				SnakeStoring: storing,
			})
		}

		if messageOption.GetDdl().GetInterleave() != nil {
			interleave = &Interleave{
				TableSnakeName: core.ToSnakeCase(messageOption.GetDdl().GetInterleave().GetTable()),
			}
		}
	}

	ret := &Message{
		Messages:  make([]*Message, 0, message.Desc.Messages().Len()),
		SnakeName: core.ToSnakeCase(string(message.Desc.FullName().Name())),
		Comment:   core.CommentReplacer.Replace(message.Comments.Leading.String()),
		Fields:    nil,
		Option: &MessageOption{
			AccessorType: messageAccessorType,
			DDL: &MessageOptionDDL{
				Indexes:    indexes,
				Interleave: interleave,
			},
		},
	}

	inputFields := make([]*Field, 0, len(message.Fields)+2)
	for _, field := range message.Fields {
		var typeName string
		var typeKind TypeKind

		fieldOption, ok := proto.GetExtension(field.Desc.Options(), options.E_Field).(*options.FieldOption)
		if !ok {
			return nil, perrors.Newf("type assertion failed")
		}
		fieldAccessorType, err := ConvertFieldAccessorTypeFromProto(fieldOption.GetAccessorType())
		if err != nil {
			return nil, perrors.Stack(err)
		}

		switch field.Desc.Kind() {
		case protoreflect.BoolKind:
			typeName = FieldType_Bool
			typeKind = TypeKind_Bool
		case protoreflect.Int32Kind:
			typeName = FieldType_Int64
			typeKind = TypeKind_Int64
		case protoreflect.Int64Kind:
			typeName = FieldType_Int64
			typeKind = TypeKind_Int64
		case protoreflect.StringKind:
			typeName = FieldType_String
			typeKind = TypeKind_String
		case protoreflect.EnumKind:
			typeName = string(field.Desc.Enum().Name())
			typeKind = TypeKind_Enum
		case protoreflect.MessageKind:
			// 型がmessageのfieldはクライアントでしか使えない
			if messageAccessorType == MessageAccessorType_OnlyServer ||
				(ServerMessageAccessorSet.Contains(messageAccessorType) && fieldAccessorType != FieldAccessorType_OnlyClient) {
				return nil, perrors.Newf("message type cannot use as server column. table name: %s, column name: %s", ret.SnakeName, core.ToSnakeCase(field.Desc.TextName()))
			}

			m, err := convert(field.Message)
			if err != nil {
				return nil, perrors.Stack(err)
			}
			ret.Messages = append(ret.Messages, m)
		case protoreflect.BytesKind, protoreflect.DoubleKind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
			protoreflect.FloatKind, protoreflect.GroupKind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind:
			return nil, perrors.Newf("unsupported Kind: %v", field.Desc.Kind().String())
		default:
			return nil, perrors.Newf("unsupported Kind: %v", field.Desc.Kind().String())
		}

		inputField := &Field{
			SnakeName: core.ToSnakeCase(field.Desc.TextName()),
			Comment:   core.CommentReplacer.Replace(field.Comments.Leading.String()),
			Type:      typeName,
			TypeKind:  typeKind,
			IsList:    field.Desc.IsList(),
			Number:    int32(field.Desc.Number()),
			Option:    nil,
		}

		var pk bool
		var masterRef *MasterRef
		if fieldOption.GetDdl() != nil {
			pk = fieldOption.GetDdl().GetPk()
			if fieldOption.GetDdl().GetMasterRef() != nil {
				masterRef = &MasterRef{
					TableSnakeName:  core.ToSnakeCase(fieldOption.GetDdl().GetMasterRef().GetTable()),
					ColumnSnakeName: core.ToSnakeCase(fieldOption.GetDdl().GetMasterRef().GetColumn()),
				}
			}
		}

		option := &FieldOption{
			AccessorType: fieldAccessorType,
			DDL: &FieldOptionDDL{
				PK:        pk,
				MasterRef: masterRef,
			},
		}

		inputField.Option = option

		inputFields = append(inputFields, inputField)
	}
	inputFields = append(inputFields,
		&Field{
			SnakeName: "created_time",
			Comment:   "作成日時",
			Type:      "int64",
			TypeKind:  TypeKind_Int64,
			IsList:    false,
			Option: &FieldOption{
				AccessorType: FieldAccessorType_OnlyServer,
				DDL: &FieldOptionDDL{
					PK:        false,
					MasterRef: nil,
				},
			},
		},
		&Field{
			SnakeName: "updated_time",
			Comment:   "更新日時",
			Type:      "int64",
			TypeKind:  TypeKind_Int64,
			IsList:    false,
			Option: &FieldOption{
				AccessorType: FieldAccessorType_OnlyServer,
				DDL: &FieldOptionDDL{
					PK:        false,
					MasterRef: nil,
				},
			},
		},
	)
	ret.Fields = inputFields

	return ret, nil
}

func ConvertMessageAccessorTypeFromProto(in options.MessageOption_AccessorType) (MessageAccessorType, error) {
	var out MessageAccessorType

	switch in {
	case options.MessageOption_OnlyServer:
		out = MessageAccessorType_OnlyServer
	case options.MessageOption_ServerAndClient:
		out = MessageAccessorType_ServerAndClient
	default:
		return 0, perrors.Newf("unsupported AccessorType: %v", in)
	}

	return out, nil
}

func ConvertFieldAccessorTypeFromProto(in options.FieldOption_AccessorType) (FieldAccessorType, error) {
	var out FieldAccessorType

	switch in {
	case options.FieldOption_All:
		out = FieldAccessorType_All
	case options.FieldOption_OnlyServer:
		out = FieldAccessorType_OnlyServer
	case options.FieldOption_OnlyClient:
		out = FieldAccessorType_OnlyClient
	default:
		return 0, perrors.Newf("unsupported AccessorType: %v", in)
	}

	return out, nil
}
