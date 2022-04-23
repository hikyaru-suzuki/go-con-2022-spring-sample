package master

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
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

func (c *creator) Create(
	message *input.Message,
	fkParentMap map[string]map[string]*output.FK,
	fkChildMap map[string]map[string][]*output.FK,
) (*output.TemplateInfo, error) {
	type ValidateOption struct {
		Key   string
		Value string
	}
	type Column struct {
		GoName          string
		SnakeName       string
		Comment         string
		Type            string
		ValidateOptions []*ValidateOption
	}
	type Table struct {
		PkgName   string
		GoName    string
		SnakeName string
		Comment   string
		Columns   []*Column
	}

	if !input.AdminAccessorSet.Contains(message.Option.AccessorType) {
		return nil, nil
	}

	data := &Table{
		PkgName:   core.ToPkgName(message.SnakeName),
		GoName:    core.ToGolangPascalCase(message.SnakeName),
		SnakeName: message.SnakeName,
		Comment:   message.Comment,
		Columns:   make([]*Column, 0, len(message.Fields)),
	}

	for _, field := range message.Fields {
		if !input.AdminAccessorSet.Contains(field.Option.AccessorType) {
			continue
		}

		/* 型の整形 */
		typeName := field.Type
		isTime := core.IsTimeField(field.SnakeName)
		if isTime {
			typeName = "time.Time"
		}
		nullable := false
		if isTime {
			nullable = true
		} else if field.Option.DDL.Nullable {
			nullable = true
		}
		if nullable {
			typeName = "*" + typeName
		}
		isEnum := field.TypeKind == input.TypeKind_Enum
		isList := field.IsList
		if isEnum {
			typeName = "enum." + typeName
		} else if isList {
			typeName = "string"
		}

		column := &Column{
			GoName:          core.ToGolangPascalCase(field.SnakeName),
			SnakeName:       field.SnakeName,
			Comment:         field.Comment,
			Type:            typeName,
			ValidateOptions: nil,
		}

		/* validate */
		column.ValidateOptions = make([]*ValidateOption, 0, len(field.Option.Validates))
		if field.Option.DDL.FK != nil {
			values := make([]string, 0, 2+len(field.Option.DDL.FK.ParentColumnSnakeNames))
			values = append(values, core.ToGolangPascalCase(field.Option.DDL.FK.TableSnakeName), core.ToGolangPascalCase(field.Option.DDL.FK.ColumnSnakeName))
			for _, name := range field.Option.DDL.FK.ParentColumnSnakeNames {
				values = append(values, core.ToGolangPascalCase(name))
			}
			value := strings.Join(values, "_")

			if isList {
				column.ValidateOptions = append(column.ValidateOptions, &ValidateOption{
					Key:   "fk-csv",
					Value: value,
				})
			} else {
				column.ValidateOptions = append(column.ValidateOptions, &ValidateOption{
					Key:   "fk",
					Value: value,
				})
			}
		}
		if isEnum {
			column.ValidateOptions = append(column.ValidateOptions, &ValidateOption{
				Key:   "enum",
				Value: "",
			})
		}
		for _, validate := range field.Option.Validates {
			column.ValidateOptions = append(column.ValidateOptions, &ValidateOption{
				Key:   validate.Key,
				Value: validate.Value,
			})
		}

		data.Columns = append(data.Columns, column)
	}

	buf := &bytes.Buffer{}
	if err := c.tpl.Execute(buf, data); err != nil {
		return nil, perrors.Stack(err)
	}

	return &output.TemplateInfo{
		Data:     buf.Bytes(),
		FilePath: core.JoinPath("pkg/domain/entity/master", data.SnakeName+"-gen.go"),
	}, nil
}
