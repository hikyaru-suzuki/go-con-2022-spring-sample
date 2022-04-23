package master

import (
	"bytes"
	_ "embed"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
	"text/template"
)

//go:embed entity-gen.go.tmpl
var templateFileBytes []byte

type creator struct {
	tpl *template.Template
}

func New() output.EachTemplateCreator {
	tpl := template.Must(core.GetBaseTemplate().Parse(string(templateFileBytes)))
	return &creator{tpl: tpl}
}

func (c *creator) Create(message *input.Message) (*output.TemplateInfo, error) {
	if !input.ServerMessageAccessorSet.Contains(message.Option.AccessorType) {
		return nil, nil
	}

	type Column struct {
		GoName  string
		Type    string
		Comment string
	}
	type Table struct {
		GoName    string
		SnakeName string
		Comment   string
		Columns   []*Column
	}

	data := &Table{
		GoName:    core.ToGolangPascalCase(message.SnakeName),
		SnakeName: message.SnakeName,
		Comment:   message.Comment,
		Columns:   make([]*Column, 0, len(message.Fields)),
	}

	for _, field := range message.Fields {
		if !input.ServerFieldAccessorSet.Contains(field.Option.AccessorType) {
			continue
		}

		typeName := field.Type

		if core.IsTimeField(field.SnakeName) {
			typeName = "time.Time"
		}

		if field.TypeKind == input.TypeKind_Enum && field.IsList {
			typeName = "enum." + typeName + "s"
		} else {
			if field.TypeKind == input.TypeKind_Enum {
				typeName = "enum." + typeName
			}
			if field.IsList {
				typeName = "[]" + typeName
			}
		}

		column := &Column{
			GoName:  core.ToGolangPascalCase(field.SnakeName),
			Type:    typeName,
			Comment: field.Comment,
		}

		data.Columns = append(data.Columns, column)
	}

	buf := &bytes.Buffer{}
	if err := c.tpl.Execute(buf, data); err != nil {
		return nil, perrors.Stack(err)
	}

	return &output.TemplateInfo{
		Data:     buf.Bytes(),
		FilePath: core.JoinPath("pkg/domain/entity/transaction", data.SnakeName+"-gen.go"),
	}, nil
}
