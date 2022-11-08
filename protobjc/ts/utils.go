package ts

import (
	. "io.protobj/protobjc"
	"strings"
)

func AddImportMessage(b *CodeBuilder, importMessage string) {
	if len(importMessage) > 0 {
		b.ImportMessages[importMessage] = Empty
	}
}
func AddImportMessages(b *CodeBuilder, importMessages map[string]Void) {
	for k := range importMessages {
		AddImportMessage(b, k)
	}
}

func appendImportMessages(sourcePackage string, header *CodeBuilder) {
	importList := mapKeyToSlice(sourcePackage, header.ImportMessages)
	var tsCoreLib = "protobj-ts"
	var coreLib = map[string]Void{}
	for _, v := range importList {
		if strings.Contains(v, tsCoreLib) {
			v = v[strings.IndexAny(v, "{")+1 : strings.IndexAny(v, "}")]
			v = strings.TrimSpace(v)
			split := strings.Split(v, ",")
			for _, s := range split {
				coreLib[s] = Empty
			}
		}
	}
	if len(coreLib) > 0 {
		var coreLibStr = strings.Join(mapKeyToSlice("", coreLib), ",")
		header.Add(I("import { ${0} } from \"${1}\"", coreLibStr, tsCoreLib)).NewLine()
	}

	for _, s := range importList {
		hasPrefix := strings.Contains(s, tsCoreLib)
		if !hasPrefix {
			split := strings.Split(sourcePackage, "\\.")
			length := len(split)
			className := s[strings.LastIndex(s, ".")+1:]
			s = strings.ReplaceAll(s, ".", "/")
			s = strings.Repeat("../", length) + s

			header.Add(I("import { ${0} } from \"${1}\"", className, s)).NewLine()
		}
	}
	header.NewLine()
}
func mapKeyToSlice(pkg string, importMessages map[string]Void) []string {
	var importList []string
	for k := range importMessages {
		importList = append(importList, k)
	}
	return importList
}
