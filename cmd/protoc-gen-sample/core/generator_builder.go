package core

import (
	"context"
	"os"
	"strings"
	"sync"

	"github.com/scylladb/go-set/strset"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/parallel"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/plogging"
)

type generatorBuilder struct {
	generators []Generator
}

func NewGeneratorBuilder() GeneratorBuilder {
	return &generatorBuilder{
		generators: make([]Generator, 0),
	}
}

func (g *generatorBuilder) AppendGenerator(generator Generator) GeneratorBuilder {
	g.generators = append(g.generators, generator)

	return g
}

func (g *generatorBuilder) Generate() error {
	pg, ctx := parallel.NewGroupWithContext(context.Background(), parallel.DefaultSize)

	mu := &sync.Mutex{}
	pathSet := strset.New()

	for _, generator := range g.generators {
		generator := generator

		pg.Go(ctx, func(_ context.Context) error {
			genFileDirectories, err := generator.Build()
			if err != nil {
				return perrors.Stack(err)
			}
			generator.SetGenFiles(genFileDirectories)

			if err := generator.Format(); err != nil {
				return perrors.Stack(err)
			}

			if err := generator.Generate(); err != nil {
				return perrors.Stack(err)
			}

			mu.Lock()
			pathSet.Add(generator.GetGeneratedFilePaths()...)
			mu.Unlock()

			return nil
		})
	}
	if err := pg.Wait(); err != nil {
		return perrors.Stack(err)
	}

	data := "\n" + strings.Join(pathSet.List(), "\n") + "\n"

	file, err := os.OpenFile("/tmp/generated_files.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return perrors.Stack(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			plogging.GetLogger().Infof("Failed to close file %s: %+v\n", "log.txt", err)
		}
	}()

	if _, err = file.WriteString(data); err != nil {
		return perrors.Stack(err)
	}

	return nil
}
