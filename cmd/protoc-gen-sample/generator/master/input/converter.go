package input

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"

	options "github.com/hikyaru-suzuki/go-con-2022-spring-sample/pkg/domain/proto/server/options/master"
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
	if messageOption.GetDdl() != nil {
		indexes = make([]*Index, 0, len(messageOption.GetDdl().GetIndexes()))
		for _, index := range messageOption.GetDdl().GetIndexes() {
			snakeNameKeys := make([]string, 0, len(index.GetKeys()))
			for _, key := range index.GetKeys() {
				snakeNameKeys = append(snakeNameKeys, core.ToSnakeCase(key))
			}

			indexes = append(indexes, &Index{SnakeNameKeys: snakeNameKeys})
		}
	}
	ret := &Message{
		Messages:  make([]*Message, 0, message.Desc.Messages().Len()),
		SnakeName: core.ToSnakeCase(string(message.Desc.FullName().Name())),
		Comment:   core.CommentReplacer.Replace(message.Comments.Leading.String()),
		Fields:    nil,
		Option: &MessageOption{
			AccessorType: messageAccessorType,
			DDL:          &MessageOptionDDL{Indexes: indexes},
		},
	}

	inputFields := make([]*Field, 0, len(message.Fields)+1)
	for _, field := range message.Fields {
		var typeName string
		var typeKind TypeKind
		switch field.Desc.Kind() {
		case protoreflect.FloatKind:
			typeName = FieldType_Float32
			typeKind = TypeKind_Float32
		case protoreflect.BoolKind:
			typeName = FieldType_Bool
			typeKind = TypeKind_Bool
		case protoreflect.Int32Kind:
			typeName = FieldType_Int32
			typeKind = TypeKind_Int32
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
			typeName = string(field.Desc.Message().Name())
			typeKind = TypeKind_Message
			m, err := convert(field.Message)
			if err != nil {
				return nil, perrors.Stack(err)
			}
			ret.Messages = append(ret.Messages, m)
		case protoreflect.BytesKind, protoreflect.DoubleKind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
			protoreflect.GroupKind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind:
			return nil, perrors.Newf("unsupported Kind: %v", field.Desc.Kind().String())
		default:
			return nil, perrors.Newf("unsupported Kind: %v", field.Desc.Kind().String())
		}

		inputField := &Field{
			SnakeName: core.ToSnakeCase(field.Desc.TextName()),
			Comment:   core.CommentReplacer.Replace(field.Comments.Leading.String()),
			Number:    int32(field.Desc.Number()),
			Type:      typeName,
			TypeKind:  typeKind,
			IsList:    field.Desc.IsList(),
			Option:    nil,
		}

		fieldOption, ok := proto.GetExtension(field.Desc.Options(), options.E_Field).(*options.FieldOption)
		if !ok {
			return nil, perrors.Newf("type assertion failed")
		}
		ddlOption := fieldOption.GetDdl()
		validationOptions := fieldOption.GetValidations()
		fieldAccessorType, err := ConvertFieldAccessorTypeFromProto(fieldOption.GetAccessorType())
		if err != nil {
			return nil, perrors.Stack(err)
		}

		option := &FieldOption{
			AccessorType: fieldAccessorType,
			DDL: &FieldOptionDDL{
				PK:       ddlOption.GetPk(),
				FK:       nil,
				Size:     ddlOption.GetSize(),
				Nullable: ddlOption.GetNullable(),
			},
			Validates: make([]*FieldOptionValidate, 0, len(validationOptions)),
		}
		if ddlOption.GetFk() != nil {
			parentColumns := make([]string, 0, len(ddlOption.GetFk().GetParentColumns()))
			for _, name := range ddlOption.GetFk().GetParentColumns() {
				parentColumns = append(parentColumns, core.ToSnakeCase(name))
			}
			option.DDL.FK = &FieldOptionDDLFK{
				TableSnakeName:         core.ToSnakeCase(ddlOption.GetFk().GetTable()),
				ColumnSnakeName:        core.ToSnakeCase(ddlOption.GetFk().GetColumn()),
				ParentColumnSnakeNames: parentColumns,
			}
		}
		for _, validates := range validationOptions {
			option.Validates = append(option.Validates, &FieldOptionValidate{
				Key:   validates.GetKey(),
				Value: validates.GetValue(),
			})
		}
		inputField.Option = option

		inputFields = append(inputFields, inputField)
	}
	inputFields = append(inputFields, &Field{
		SnakeName: "master_version",
		Comment:   "マスタバージョン",
		Type:      "int32",
		TypeKind:  TypeKind_Int32,
		IsList:    false,
		Option: &FieldOption{
			AccessorType: AccessorType_All,
			DDL: &FieldOptionDDL{
				PK:       false,
				FK:       nil,
				Size:     0,
				Nullable: false,
			},
			Validates: []*FieldOptionValidate{
				{Key: "min", Value: "0"},
			},
		},
	})
	ret.Fields = inputFields

	return ret, nil
}

func ConvertMessageAccessorTypeFromProto(in options.MessageOption_AccessorType) (AccessorType, error) {
	var out AccessorType

	switch in {
	case options.MessageOption_All:
		out = AccessorType_All
	case options.MessageOption_OnlyAdmin:
		out = AccessorType_OnlyAdmin
	case options.MessageOption_OnlyServer:
		out = AccessorType_OnlyServer
	case options.MessageOption_AdminAndServer:
		out = AccessorType_AdminAndServer
	case options.MessageOption_AdminAndClient:
		out = AccessorType_AdminAndClient
	default:
		return 0, perrors.Newf("unsupported AccessorType: %v", in)
	}

	return out, nil
}

func ConvertFieldAccessorTypeFromProto(in options.FieldOption_AccessorType) (AccessorType, error) {
	var out AccessorType

	switch in {
	case options.FieldOption_All:
		out = AccessorType_All
	case options.FieldOption_OnlyAdmin:
		out = AccessorType_OnlyAdmin
	case options.FieldOption_OnlyServer:
		out = AccessorType_OnlyServer
	case options.FieldOption_OnlyClient:
		out = AccessorType_OnlyClient
	case options.FieldOption_AdminAndServer:
		out = AccessorType_AdminAndServer
	case options.FieldOption_AdminAndClient:
		out = AccessorType_AdminAndClient
	case options.FieldOption_ServerAndClient:
		out = AccessorType_ServerAndClient
	default:
		return 0, perrors.Newf("unsupported AccessorType: %v", in)
	}

	return out, nil
}
