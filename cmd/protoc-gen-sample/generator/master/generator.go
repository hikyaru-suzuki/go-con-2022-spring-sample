package master

import (
	"google.golang.org/protobuf/compiler/protogen"
	"sort"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output/pkg/domain/entity/clientcache"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output/pkg/domain/entity/master"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master/output/pkg/domain/entity/servercache"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
)

type generator struct {
	*core.GeneratorBase
	plugin *protogen.Plugin
}

func NewGenerator(plugin *protogen.Plugin) core.Generator {
	return &generator{
		GeneratorBase: core.NewGeneratorBase(),
		plugin:        plugin,
	}
}

var eachCreators = []output.EachTemplateCreator{
	master.New(),
	servercache.New(),
	clientcache.New(),
}

var bulkCreators = []output.BulkTemplateCreator{}

func (g *generator) Build() ([]core.GenFile, error) {
	messages := make([]*input.Message, 0)

	for _, file := range g.plugin.Files {
		if !file.Generate {
			continue
		}
		if file.Proto.GetPackage() != "server.master" {
			continue
		}

		message, err := input.ConvertMessageFromProto(file)
		if err != nil {
			return nil, perrors.Stack(err)
		}
		messages = append(messages, message)
	}

	// 入力ファイルの順番に左右されないようソートする
	sort.SliceStable(messages, func(i, j int) bool {
		return messages[i].SnakeName < messages[j].SnakeName
	})

	fkParentMap := make(map[string]map[string]*output.FK)
	fkChildMap := make(map[string]map[string][]*output.FK)
	for _, message := range messages {
		for _, field := range message.Fields {
			if field.Option.DDL.FK == nil {
				continue
			}

			fk := field.Option.DDL.FK
			tableName := core.ToSnakeCase(fk.TableSnakeName)
			columnName := core.ToSnakeCase(fk.ColumnSnakeName)

			if _, ok := fkParentMap[message.SnakeName]; !ok {
				fkParentMap[message.SnakeName] = make(map[string]*output.FK)
			}
			fkParentMap[message.SnakeName][field.SnakeName] = &output.FK{
				TableSnakeName:  tableName,
				ColumnSnakeName: columnName,
			}

			if _, ok := fkChildMap[tableName]; !ok {
				fkChildMap[tableName] = make(map[string][]*output.FK)
			}
			fkChildMap[tableName][columnName] = append(fkChildMap[tableName][columnName], &output.FK{
				TableSnakeName:  message.SnakeName,
				ColumnSnakeName: field.SnakeName,
			})
		}
	}

	genFiles := make([]core.GenFile, 0)
	for _, creator := range eachCreators {
		for _, message := range messages {
			info, err := creator.Create(message, fkParentMap, fkChildMap)
			if err != nil {
				return nil, perrors.Stack(err)
			}
			if info == nil {
				continue
			}

			genFiles = append(genFiles, core.NewGenFile(info.FilePath, info.Data))
		}
	}
	for _, creator := range bulkCreators {
		info, err := creator.Create(messages, fkParentMap, fkChildMap)
		if err != nil {
			return nil, perrors.Stack(err)
		}
		if info == nil {
			continue
		}

		genFiles = append(genFiles, core.NewGenFile(info.FilePath, info.Data))
	}

	return genFiles, nil
}
