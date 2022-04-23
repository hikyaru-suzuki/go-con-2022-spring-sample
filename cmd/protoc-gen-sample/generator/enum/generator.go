package enum

import (
	_ "embed"
	"sort"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum/input"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum/output"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum/output/pkg/domain/enum"
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
	enum.New(),
}

var bulkCreators = []output.BulkTemplateCreator{}

func (g *generator) Build() ([]core.GenFile, error) {
	var enums []*input.Enum

	for _, f := range g.plugin.Files {
		if !f.Generate {
			continue
		}
		if f.Proto.GetPackage() != "server.enums" {
			continue
		}

		ret, err := input.ConvertMessageFromProto(f)
		if err != nil {
			return nil, perrors.Stack(err)
		}

		enums = ret
	}

	// 入力ファイルの順番に左右されないようソートする
	sort.SliceStable(enums, func(i, j int) bool {
		return enums[i].SnakeName < enums[j].SnakeName
	})

	genFiles := make([]core.GenFile, 0)
	for _, creator := range eachCreators {
		for _, enum := range enums {
			info, err := creator.Create(enum)
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
		info, err := creator.Create(enums)
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
