package input

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
	options "github.com/hikyaru-suzuki/go-con-2022-spring-sample/pkg/domain/proto/server/options/enum"
)

func ConvertMessageFromProto(file *protogen.File) ([]*Enum, error) {
	ret := make([]*Enum, 0, len(file.Enums))

	for _, enum := range file.Enums {
		define, ok := proto.GetExtension(enum.Desc.Options(), options.E_Enum).(*options.EnumOption)
		if !ok {
			return nil, perrors.Newf("type assertion failed")
		}

		var accessorType AccessorType
		t := define.GetAccessorType()
		switch t {
		case options.EnumOption_Unknown:
			return nil, perrors.Newf("unsupported AccessorType: %v", t)
		case options.EnumOption_OnlyServer:
			accessorType = AccessorType_OnlyServer
		case options.EnumOption_ServerAndClient:
			accessorType = AccessorType_ServerAndClient
		default:
			return nil, perrors.Newf("unsupported AccessorType: %v", t)
		}

		elements := make([]*Element, 0, len(enum.Values))
		for _, value := range enum.Values {
			names := strings.Split(value.GoIdent.GoName, "_")
			elements = append(elements, &Element{
				RawName: names[len(names)-1],
				Value:   int32(value.Desc.Number()),
				Comment: core.CommentReplacer.Replace(value.Comments.Leading.String()),
			})
		}

		ret = append(ret, &Enum{
			AccessorType: accessorType,
			SnakeName:    core.ToSnakeCase(string(enum.Desc.FullName().Name())),
			Comment:      core.CommentReplacer.Replace(enum.Comments.Leading.String()),
			Elements:     elements,
		})
	}

	return ret, nil
}
