package output

import (
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/input"
)

type FK struct {
	TableSnakeName  string
	ColumnSnakeName string
}

type TemplateInfo struct {
	Data     []byte
	FilePath string
}

type EachTemplateCreator interface {
	Create(
		message *input.Message,
		fkParentMap map[string]map[string]*FK,
		fkChildMap map[string]map[string][]*FK,
	) (*TemplateInfo, error)
}

type BulkTemplateCreator interface {
	Create(
		messages []*input.Message,
		fkParentMap map[string]map[string]*FK,
		fkChildMap map[string]map[string][]*FK,
	) (*TemplateInfo, error)
}
