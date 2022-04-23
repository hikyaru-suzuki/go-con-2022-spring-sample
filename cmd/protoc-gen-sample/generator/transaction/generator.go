package transaction

import (
	_ "embed"
	"sort"

	"google.golang.org/protobuf/compiler/protogen"

	entity "github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction/output/pkg/domain/entity/transaction"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction/output"
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
	entity.New(),
}

var bulkCreators = []output.BulkTemplateCreator{}

func (g *generator) Build() ([]core.GenFile, error) {
	messages := make([]*input.Message, 0)

	for _, file := range g.plugin.Files {
		if !file.Generate {
			continue
		}
		if file.Proto.GetPackage() != "server.transaction" {
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

	genFiles := make([]core.GenFile, 0)
	for _, creator := range eachCreators {
		for _, message := range messages {
			info, err := creator.Create(message)
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
		info, err := creator.Create(messages)
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
