package servercache

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
)

//go:embed entity-gen.go.tmpl
var templateFileBytes []byte

//go:embed message.tpl
var messageTemplateFileBytes []byte

type creator struct {
	tpl *template.Template
}

func New() output.EachTemplateCreator {
	msgTpl := template.Must(core.GetBaseTemplate().Parse(string(messageTemplateFileBytes)))
	tpl := template.Must(msgTpl.Parse(string(templateFileBytes)))
	return &creator{tpl: tpl}
}

type Column struct {
	GoName  string
	Comment string
	Type    string
}
type Data struct {
	PkgName  string
	GoName   string
	Comment  string
	Columns  []*Column
	Messages []*Data
}

func createData(message *input.Message, goNamePrefix string) *Data {
	goName := goNamePrefix + core.ToGolangPascalCase(message.SnakeName)
	data := &Data{
		PkgName:  core.ToPkgName(message.SnakeName),
		GoName:   goName,
		Comment:  message.Comment,
		Columns:  make([]*Column, 0, len(message.Fields)),
		Messages: make([]*Data, 0, len(message.Messages)),
	}

	// パッケージ内のstruct名衝突回避のためnested messageは「親のstruct名_」をprefixとする
	if goNamePrefix == "" {
		goNamePrefix = goName + "_"
	}

	for _, m := range message.Messages {
		data.Messages = append(data.Messages, createData(m, goNamePrefix))
	}

	for _, field := range message.Fields {
		if !input.ServerAccessorSet.Contains(field.Option.AccessorType) {
			continue
		}
		if core.IsMasterVersion(field.SnakeName) {
			continue
		}

		/* 型の整形 */
		typeName := field.Type
		if core.IsTimeField(field.SnakeName) {
			typeName = "time.Time"
		}

		if field.TypeKind == input.TypeKind_Enum && field.IsList {
			typeName = "enum." + typeName + "s"
		} else if field.TypeKind == input.TypeKind_Message && field.IsList {
			typeName = goNamePrefix + typeName + "Slice"
		} else {
			switch field.TypeKind {
			case input.TypeKind_Enum:
				typeName = "enum." + typeName
			case input.TypeKind_Message:
				typeName = "*" + goNamePrefix + typeName
			}
			if field.IsList {
				typeName = "[]" + typeName
			}
		}

		column := &Column{
			GoName:  core.ToGolangPascalCase(field.SnakeName),
			Comment: field.Comment,
			Type:    typeName,
		}

		data.Columns = append(data.Columns, column)
	}

	return data
}

func (c *creator) Create(
	message *input.Message,
	fkParentMap map[string]map[string]*output.FK,
	fkChildMap map[string]map[string][]*output.FK,
) (*output.TemplateInfo, error) {
	if !input.ServerAccessorSet.Contains(message.Option.AccessorType) {
		return nil, nil
	}

	data := createData(message, "")

	buf := &bytes.Buffer{}
	if err := c.tpl.Execute(buf, data); err != nil {
		return nil, perrors.Stack(err)
	}

	return &output.TemplateInfo{
		Data:     buf.Bytes(),
		FilePath: core.JoinPath("pkg/domain/entity/servercache", message.SnakeName+"-gen.go"),
	}, nil
}
