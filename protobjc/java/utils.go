package java

import (
	. "io.protobj/protobjc"
	"sort"
	"strings"
)

const LC = "{"
const RC = "}"
const IF = "if "
const ELSE_IF = "} else if "
const ELSE = "} else { "
const CONTINUE = "continue;"
const LAMBDA = "() -> "

func pkg(pkg string) string {
	return I("package ${0};", pkg)
}

func notNull(value string) string {
	return I("if (${0} != null) ", value)
}
func isNull(value string) string {
	return I("if (${0} == null) ", value)
}

func collectionNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.size() > 0) ", "value", value)
}
func arrayNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.length > 0) ", "value", value)
}

func stringNotEmpty(value string) string {
	return NI("if (${value} != null && ${value}.length() > 0) ", "value", value)
}

func isNull2(value1, value2 string) string {
	return I("if (${0} == null || ${1} == null)", value1, value2)
}

func classEquals(object, Class string) string {
	return I("(${0}.getClass() == ${1}.class) ", object, Class)
}

func appendImportMessages(sourcePackage string, header *CodeBuilder) {
	importList := distinctPackageForJava(sourcePackage, header.ImportMessages)
	var javaCoreLibPrefix = "java."
	sort.Slice(importList, func(i, j int) bool {
		o1 := importList[i]
		o2 := importList[j]
		if strings.HasPrefix(o1, javaCoreLibPrefix) && !strings.HasPrefix(o2, javaCoreLibPrefix) {
			return false
		}
		if strings.HasPrefix(o2, javaCoreLibPrefix) && !strings.HasPrefix(o1, javaCoreLibPrefix) {
			return true
		}
		return strings.Compare(o1, o2) < 0
	})
	var importCoreLib = false
	for i, s := range importList {
		if strings.HasPrefix(s, javaCoreLibPrefix) && !importCoreLib {
			if i != 0 {
				header.NewLine()
			}
			importCoreLib = true
		}
		header.Add(I("import ${0};", s)).NewLine()
	}
	header.NewLine()
}

func distinctPackageForJava(pkg string, importMessages map[string]Void) []string {
	var importList []string
	for k := range importMessages {
		lastIndex := strings.LastIndex(k, ".")
		if lastIndex < 0 || k[:lastIndex] != pkg {
			importList = append(importList, k)
		}
	}
	return importList
}

func readMessageStart() string {
	return "final int oldLimit = input.readMessageStart();"
}

func readMessageStop() string {
	return "input.readMessageStop(oldLimit);"
}

func isNotDefault(value string, fieldType FieldType) string {
	if fieldType == BOOL {
		return NI("if (${value} != false) ", "value", value)
	} else if fieldType == STRING {
		return notNull(value)
	} else {
		return NI("if (${value} != 0) ", "value", value)
	}
}

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
