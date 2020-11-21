package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/fajrirahmat/readcv/search"

	"code.sajari.com/docconv"
)

var sourceDir string
var targetDir string
var keywordPath string

func init() {
	flag.StringVar(&sourceDir, "source", ".", "Source directory that cv save")
	flag.StringVar(&targetDir, "target", ".", "Target directory for qualified cv save")
	flag.StringVar(&keywordPath, "keyword", "keyword.txt", "Location keyword file")
	flag.Parse()
}

type fileInfo struct {
	Name string
	Path string
}

func main() {
	var files []fileInfo
	var qualified []string
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, fileInfo{info.Name(), path})
		return nil
	})
	for _, theFile := range files {

		res, err := docconv.ConvertPath(theFile.Path)
		if err != nil {
			log.Fatal(err)
		}
		ok, err := search.FoundKeyword(res.Body, getKeywords())
		if err != nil {
			log.Print(err)
		}
		if ok {
			qualified = append(qualified, theFile.Name)
			copy(theFile.Path, fmt.Sprintf("%s/%s", targetDir, theFile.Name))
		}
	}
	fmt.Printf("Found: %d\n", len(qualified))
}

func getKeywords() []string {
	f, err := os.Open(keywordPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var keywords []string
	for scanner.Scan() {
		keywords = append(keywords, scanner.Text())
	}
	return keywords
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
