package core

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/parallel"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
)

type GeneratorBase struct {
	genFiles []GenFile
}

func NewGeneratorBase() *GeneratorBase {
	return &GeneratorBase{
		genFiles: make([]GenFile, 0),
	}
}

func (g *GeneratorBase) SetGenFiles(genFiles []GenFile) {
	g.genFiles = genFiles
}

func (g *GeneratorBase) Format() error {
	pg, ctx := parallel.NewGroupWithContext(context.Background(), parallel.DefaultSize)

	for _, file := range g.genFiles {
		file := file

		if !strings.HasSuffix(file.GetFilePath(), "-gen.go") {
			continue
		}

		pg.Go(ctx, func(_ context.Context) error {
			if err := file.Format(); err != nil {
				return perrors.Stack(err)
			}

			return nil
		})
	}
	if err := pg.Wait(); err != nil {
		return perrors.Stack(err)
	}

	return nil
}

func (g *GeneratorBase) Generate() error {
	pg, ctx := parallel.NewGroupWithContext(context.Background(), parallel.DefaultSize)

	for _, file := range g.genFiles {
		file := file

		pg.Go(ctx, func(_ context.Context) error {
			outputDir := filepath.Dir(file.GetFilePath())
			if err := os.MkdirAll(outputDir, 0777); err != nil {
				return perrors.Stack(err)
			}

			if err := file.CreateOrWrite(); err != nil {
				return perrors.Stack(err)
			}

			return nil
		})
	}
	if err := pg.Wait(); err != nil {
		return perrors.Stack(err)
	}

	return nil
}

func (g *GeneratorBase) GetGeneratedFilePaths() []string {
	paths := make([]string, 0, len(g.genFiles))

	for _, file := range g.genFiles {
		paths = append(paths, file.GetFilePath())
	}

	return paths
}
