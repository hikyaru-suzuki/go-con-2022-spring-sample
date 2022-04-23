package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/scylladb/go-set/strset"
)

func main() {
	locationName := "Asia/Tokyo"
	location, err := time.LoadLocation(locationName)
	if err != nil {
		location = time.FixedZone(locationName, 9*60*60)
	}
	time.Local = location

	startTime := time.Now()
	Infof("gen-file-remover start\n")

	err = remove()
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	Infof("gen-file-remover end, elapsed: %s\n", endTime.Sub(startTime).String())
}

func Infof(format string, a ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, time.Now().Format("15:04:05.000000000")+" "+format, a...)
	if err != nil {
		panic(err)
	}
}

func remove() error {
	file, err := os.Open("/tmp/generated_files.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	pathSet := strset.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		pathSet.Add(scanner.Text())
	}

	rootPaths := []string{
		"proto",
		"cmd",
		"pkg",
	}
	f := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// pathSetにあれば削除しない
		if pathSet.Has(path) {
			return nil
		}
		// 自動生成物だけが対象
		if !strings.HasSuffix(path, "-gen.go") {
			return nil
		}

		return os.Remove(path)
	}
	for _, rootPath := range rootPaths {
		if err := filepath.WalkDir(rootPath, f); err != nil {
			return err
		}
	}

	return nil
}
