package main

import (
	"errors"
	"fmt"
	"io.protobj/protobjc"
	"io.protobj/protobjc/java"
	"io.protobj/protobjc/ts"
	"os"
	"path/filepath"
	"strings"
)

type ArgsType int32

const (
	Help         ArgsType = 0
	SourceDir    ArgsType = 1
	LanguageType ArgsType = 2
	OutputDir    ArgsType = 3
	OutputType   ArgsType = 4
)

type ArgsInfo struct {
	argsType    ArgsType
	names       []string
	description string
	required    bool
}

var argsList = []ArgsInfo{
	{argsType: Help, names: []string{"-help", "-h"}, description: "帮助", required: false},
	{argsType: SourceDir, names: []string{"-source_dir", "-s"}, description: "协议文件目录", required: true},
	{argsType: LanguageType, names: []string{"-language", "-lang"}, description: "生成的语言：Java,Go,Ts", required: true},
	{argsType: OutputDir, names: []string{"-output_dir", "-o"}, description: "输出目录", required: true},
	{argsType: OutputType, names: []string{"-output_type", "-o_type"}, description: "输出选项:all,schema,message 默认:all", required: false},
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		return
	}
	parsedArgs, done := parseArgs(args)
	if done {
		return
	}

	files := getFiles(parsedArgs.SourceDir)
	messageMap := protobjc.Load(files)
	println("start generate")
	switch parsedArgs.LanguageType {
	case protobjc.JAVA:
		java.NewGenerator(messageMap, parsedArgs).Generate()
	case protobjc.TS:
		ts.NewGenerator(messageMap, parsedArgs).Generate()
	}
	println("generated...")
}

func getFiles(dir string) (files []string) {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, entry := range readDir {
		fileName := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			return getFiles(fileName)
		} else if strings.HasSuffix(entry.Name(), ".protobj") {
			files = append(files, fileName)
		}
	}
	return files
}

func parseArgs(args []string) (protobjc.ParsedArgs, bool) {
	var sourceDir *string
	var languageType *protobjc.LanguageType
	var outputDir *string
	var outputType protobjc.OutputType
	for i := 0; i < len(args); i++ {
		name := args[i]
		argsInfo, err := findArgsInfo(name)
		if err != nil {
			continue
		}
		argsType := argsInfo.argsType
		switch argsType {
		case Help:
			printUsage()
			return protobjc.ParsedArgs{}, true
		case SourceDir:
			sourceDir = &args[i+1]
			i++
		case LanguageType:
			langType, err := protobjc.ToLanguageType(args[i+1])
			if err != nil {
				println(err.Error())
				return protobjc.ParsedArgs{}, true
			}
			languageType = &langType
			i++
		case OutputDir:
			outputDir = &args[i+1]
			i++
		case OutputType:
			outputType = protobjc.ToOutputType(args[i+1])
			i++
		}
	}
	if sourceDir == nil {
		println("source_dir is required")
		return protobjc.ParsedArgs{}, true
	}
	if languageType == nil {
		println("language is required")
		return protobjc.ParsedArgs{}, true
	}
	if outputDir == nil {
		println("output_dir is required")
		return protobjc.ParsedArgs{}, true
	}

	return protobjc.ParsedArgs{
		SourceDir:    *sourceDir,
		OutputDir:    *outputDir,
		LanguageType: *languageType,
		OutputType:   outputType,
	}, false
}

func printUsage() {
	for _, info := range argsList {
		fmt.Println(protobjc.I("names:${0} desc:${1} required:${2}", info.names, info.description, info.required))
	}
}

func findArgsInfo(name string) (ArgsInfo, error) {
	for _, info := range argsList {
		for i := range info.names {
			if info.names[i] == name {
				return info, nil
			}
		}
	}
	return ArgsInfo{}, errors.New("unknown args:" + name)
}
