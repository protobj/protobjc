package ts

import (
	. "io.protobj/protobjc"
	"sort"
	"strings"
)

func AddImportMessage(b *CodeBuilder, importMessage string) {
	if len(importMessage) > 0 {
		b.ImportMessages[importMessage] = Empty
	}
}
func AddImportMessages(b *CodeBuilder, importMessages map[string]Void) {
	for k, _ := range importMessages {
		AddImportMessage(b, k)
	}
}

func appendImportMessages(sourcePackage string, header *CodeBuilder) {
	importList := distinctPackage(sourcePackage, header.ImportMessages)
	var tsCoreLib = "\"protobj-ts\""
	sort.Slice(importList, func(i, j int) bool {
		o1 := importList[i]
		o2 := importList[j]
		if strings.HasSuffix(o1, tsCoreLib) && !strings.HasSuffix(o2, tsCoreLib) {
			return false
		}
		if strings.HasSuffix(o2, tsCoreLib) && !strings.HasSuffix(o1, tsCoreLib) {
			return true
		}
		return strings.Compare(o1, o2) < 0
	})
	var importCoreLib = false
	for i, s := range importList {
		if strings.HasPrefix(s, tsCoreLib) && !importCoreLib {
			if i != 0 {
				header.NewLine()
			}
			importCoreLib = true
		}
		header.Add(I("import ${0};", s)).NewLine()
	}
	header.NewLine()
}
func distinctPackage(pkg string, importMessages map[string]Void) []string {
	var importList []string
	for k := range importMessages {
		lastIndex := strings.LastIndex(k, ".")
		if lastIndex < 0 || k[:lastIndex] != pkg {
			importList = append(importList, k)
		}
	}
	//lastIndex := strings.LastIndex(importMessage, ".")
	//name := importMessage[lastIndex+1:]
	//path := importMessage[:lastIndex]
	//path = strings.ReplaceAll(path, ".", "/")
	//path = "./" + path
	//importMessage = fmt.Sprintf("import { %s } from \"%s\"", name, path)
	return importList
}
