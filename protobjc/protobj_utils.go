package protobjc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

type FileContent struct {
	fileName string
	content  string
}

func NewFileContent(fileName, content string) *FileContent {
	return &FileContent{
		fileName: fileName,
		content:  content,
	}
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
	filePath = filepath.Join(outputDir, content.fileName)

	err = os.WriteFile(filePath, []byte(content.content), os.ModePerm)
	if err != nil {
		PrintErrorAndExit(err.Error())
		return
	} else {
		fmt.Printf("write file :%s \n", filePath)
	}
}

type Void struct{}

var Empty = Void{}

func PrintErrorAndExit(err string) {
	stack := debug.Stack()

	println(err + "\n" + string(stack) + "\n")
	os.Exit(-1)
}
