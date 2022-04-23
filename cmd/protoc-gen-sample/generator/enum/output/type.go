package output

import (
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum/input"
)

type TemplateInfo struct {
	Data     []byte
	FilePath string
}

type EachTemplateCreator interface {
	Create(enum *input.Enum) (*TemplateInfo, error)
}

type BulkTemplateCreator interface {
	Create(enums []*input.Enum) (*TemplateInfo, error)
}
