package ts

import (
	. "io.protobj/protobjc"
	"strings"
)

const LC = "{"
const RC = "}"
const IF = "if "
const ELSE_IF = "} else if "
const ELSE = "} else { "
const CONTINUE = "continue;"
const LAMBDA = "() => "

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

func appendImportMessages(sourcePackage string, sourceFullMessageName string, header *CodeBuilder) {
	importList := mapKeyToSlice(header.ImportMessages)
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
		var coreLibStr = strings.Join(mapKeyToSlice(coreLib), ",")
		header.Add(I("import { ${0} } from \"${1}\"", coreLibStr, tsCoreLib)).NewLine()
	}

	for _, s := range importList {
		hasPrefix := strings.Contains(s, tsCoreLib)
		if !hasPrefix && s != sourceFullMessageName {
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
func mapKeyToSlice(importMessages map[string]Void) []string {
	var importList []string
	for k := range importMessages {
		importList = append(importList, k)
	}
	return importList
}

func notNull(value string) string {
	return I("if (${0} != null) ", value)
}
func isNull(value string) string {
	return I("if (${0} == null) ", value)
}

func collectionNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.length > 0) ", "value", value)
}
func mapNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.size > 0) ", "value", value)
}
func arrayNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.length > 0) ", "value", value)
}

func stringNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.length > 0) ", "value", value)
}

func isNull2(value1, value2 string) string {
	return I("if (${0} == null || ${1} == null)", value1, value2)
}

func classEquals(object, Class string) string {
	return I("(Object.getPrototypeOf(${0}) == ${1}.prototype) ", object, Class)
}

func readMessageStart() string {
	return "const oldLimit = input.readMessageStart();"
}

func readMessageStop() string {
	return "input.readMessageStop(oldLimit);"
}
func isNotDefault(value string, fieldType FieldType) string {
	if fieldType == BOOL {
		return NI("if (${value} != false) ", "value", value)
	} else if fieldType == STRING || fieldType == I64 || fieldType == U64 || fieldType == S64 || fieldType == F64 || fieldType == SF64 {
		return notNull(value)
	} else {
		return NI("if (${value} != 0) ", "value", value)
	}
}
