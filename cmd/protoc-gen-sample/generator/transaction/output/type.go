package output

import (
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction/input"
)

type TemplateInfo struct {
	Data     []byte
	FilePath string
}

type EachTemplateCreator interface {
	Create(message *input.Message) (*TemplateInfo, error)
}

type BulkTemplateCreator interface {
	Create(messages []*input.Message) (*TemplateInfo, error)
}
