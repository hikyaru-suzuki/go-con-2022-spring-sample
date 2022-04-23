package enum

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
)

//go:embed enum-gen.go.tmpl
var templateFileBytes []byte

type creator struct {
	tpl *template.Template
}

func New() output.EachTemplateCreator {
	tpl := template.Must(core.GetBaseTemplate().Parse(string(templateFileBytes)))
	return &creator{tpl: tpl}
}

func (c *creator) Create(enum *input.Enum) (*output.TemplateInfo, error) {
	type Element struct {
		PascalName string
		LowerName  string
		Value      int32
		Comment    string
	}
	type Enum struct {
		PascalName string
		Comment    string
		Elements   []*Element
	}

	data := &Enum{
		PascalName: core.ToPascalCase(enum.SnakeName),
		Comment:    enum.Comment,
		Elements:   make([]*Element, 0, len(enum.Elements)),
	}
	for _, element := range enum.Elements {
		data.Elements = append(data.Elements, &Element{
			PascalName: element.RawName,
			Value:      element.Value,
			Comment:    element.Comment,
		})
	}

	buf := &bytes.Buffer{}
	if err := c.tpl.Execute(buf, data); err != nil {
		return nil, perrors.Stack(err)
	}

	return &output.TemplateInfo{
		Data:     buf.Bytes(),
		FilePath: core.JoinPath("pkg/domain/enum", enum.SnakeName+"-gen.go"),
	}, nil
}
