package protobj

import (
	"os"
	"path/filepath"
	"strings"
)

type FileContent struct {
	fileName string
	content  string
}

func WriteFile(outputDir string, content *FileContent) {
	index := strings.LastIndex(content.fileName, string(os.PathSeparator))
	subDir := content.fileName[0:index]
	filePath := filepath.Join(outputDir, subDir)
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		PrintErrorAndExit(err.Error())
		return
	}
	filePath = filepath.Join(filePath, content.fileName)

	err = os.WriteFile(filePath, []byte(content.content), os.ModePerm)
	if err != nil {
		PrintErrorAndExit(err.Error())
		return
	}
}

type Void struct{}

var Empty = Void{}

func PrintErrorAndExit(err string) {
	println(err)
	os.Exit(-1)
}