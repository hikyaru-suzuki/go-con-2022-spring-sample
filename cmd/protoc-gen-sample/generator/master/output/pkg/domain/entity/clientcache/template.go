package clientcache

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
)

//go:embed entity-gen.go.tpl
var templateFileBytes []byte

type creator struct {
	tpl *template.Template
}

func New() output.EachTemplateCreator {
	tpl := template.Must(core.GetBaseTemplate().Parse(string(templateFileBytes)))
	return &creator{tpl: tpl}
}

func (c *creator) Create(
	message *input.Message,
	fkParentMap map[string]map[string]*output.FK,
	fkChildMap map[string]map[string][]*output.FK,
) (*output.TemplateInfo, error) {
	type Column struct {
		CamelName string
		SnakeName string
		Comment   string
		Type      string
	}
	type Table struct {
		PkgName   string
		CamelName string
		Comment   string
		Columns   []*Column
	}

	if !input.ClientAccessorSet.Contains(message.Option.AccessorType) {
		return nil, nil
	}

	data := &Table{
		PkgName:   core.ToPkgName(message.SnakeName),
		CamelName: core.ToPascalCase(message.SnakeName),
		Comment:   message.Comment,
		Columns:   make([]*Column, 0, len(message.Fields)),
	}

	for _, field := range message.Fields {
		if !input.ClientAccessorSet.Contains(field.Option.AccessorType) {
			continue
		}

		if !field.Option.DDL.PK {
			continue
		}

		typeName := field.Type
		if field.TypeKind == input.TypeKind_Enum {
			typeName = "enums." + field.Type
		}
		column := &Column{
			CamelName: core.ToPascalCase(field.SnakeName),
			SnakeName: field.SnakeName,
			Comment:   field.Comment,
			Type:      typeName,
		}

		data.Columns = append(data.Columns, column)
	}
	data.Columns = append(data.Columns, &Column{
		CamelName: "Data",
		Comment:   "データ",
		Type:      "[]byte",
		SnakeName: "data",
	})

	buf := &bytes.Buffer{}
	if err := c.tpl.Execute(buf, data); err != nil {
		return nil, perrors.Stack(err)
	}

	return &output.TemplateInfo{
		Data:     buf.Bytes(),
		FilePath: core.JoinPath("pkg/domain/entity/clientcache", message.SnakeName+"-gen.go"),
	}, nil
}
