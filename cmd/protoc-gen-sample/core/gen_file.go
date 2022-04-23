package core

import (
	"go/format"
	"os"
	"sync"

	"golang.org/x/tools/imports"

	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/perrors"
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/cmd/protoc-gen-sample/plogging"
)

type GenFile interface {
	Format() error
	CreateOrWrite() error
	GetFilePath() string
}

type Generator interface {
	Build() ([]GenFile, error)

	SetGenFiles(genFiles []GenFile)
	Format() error
	Generate() error
	GetGeneratedFilePaths() []string
}

type GeneratorBuilder interface {
	AppendGenerator(generator Generator) GeneratorBuilder
	Generate() error
}

type genFile struct {
	filePath     string
	newData      []byte
	oldDataCache map[string][]byte
	mu           *sync.Mutex
}

func NewGenFile(filePath string, newData []byte) GenFile {
	return &genFile{
		filePath:     filePath,
		newData:      newData,
		oldDataCache: make(map[string][]byte),
		mu:           &sync.Mutex{},
	}
}

func (g *genFile) Format() error {
	importsData, err := imports.Process("", g.newData, &imports.Options{
		Fragment:   true,
		AllErrors:  false,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: false,
	})
	if err != nil {
		return perrors.Newf("%v, %s", err, g.GetFilePath())
	}

	fmtData, err := format.Source(importsData)
	if err != nil {
		return perrors.Newf("%v, %s", err, g.GetFilePath())
	}
	g.newData = fmtData

	return nil
}

func (g *genFile) CreateOrWrite() error {
	path := g.filePath
	file, err := os.Create(path)
	if err != nil {
		return perrors.Stack(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			plogging.GetLogger().Infof("Failed to close file %s: %+v\n", path, err)
		}
	}()

	if _, err := file.Write(g.newData); err != nil {
		return perrors.Stack(err)
	}

	return nil
}

func (g *genFile) GetFilePath() string {
	return g.filePath
}
