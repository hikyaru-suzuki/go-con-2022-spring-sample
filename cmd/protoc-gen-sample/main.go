package main

import (
	"strings"
	"time"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/core"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/enum"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/master"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/generator/transaction"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/plogging"
)

type flagKind string

const (
	flagKindGenEnum        flagKind = "gen_enum"
	flagKindGenMaster      flagKind = "gen_master"
	flagKindGenTransaction flagKind = "gen_transaction"
)

func main() {
	locationName := "Asia/Tokyo"
	location, err := time.LoadLocation(locationName)
	if err != nil {
		location = time.FixedZone(locationName, 9*60*60)
	}
	time.Local = location

	startTime := time.Now()
	logger := plogging.GetLogger()
	logger.Infof("protoc-gen-sample start\n")

	generatorBuilder := core.NewGeneratorBuilder()

	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		generatorMap := createGeneratorMap(plugin)
		kinds := make([]string, 0, len(generatorMap))
		for kind, generator := range generatorMap {
			kinds = append(kinds, string(kind))
			generatorBuilder.AppendGenerator(generator)
		}
		logger.Infof("flag %s\n", strings.Join(kinds, ","))

		if err := generatorBuilder.Generate(); err != nil {
			return perrors.Stack(err)
		}
		return nil
	})

	endTime := time.Now()
	logger.Infof("protoc-gen-sample end, elapsed: %s\n", endTime.Sub(startTime).String())
}

func createGeneratorMap(plugin *protogen.Plugin) map[flagKind]core.Generator {
	generatorMap := make(map[flagKind]core.Generator)

	for _, param := range strings.Split(plugin.Request.GetParameter(), ",") {
		s := strings.Split(param, "=")

		switch flagKind(s[0]) {
		case flagKindGenEnum:
			generatorMap[flagKindGenEnum] = enum.NewGenerator(plugin)
		case flagKindGenMaster:
			generatorMap[flagKindGenMaster] = master.NewGenerator(plugin)
		case flagKindGenTransaction:
			generatorMap[flagKindGenTransaction] = transaction.NewGenerator(plugin)
		default:
			continue
		}
	}

	return generatorMap
}
