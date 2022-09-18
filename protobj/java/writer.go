package java

import (
	. "io.protobj/protobj-go/protobj"
	"strings"
)

type SetPrimitiveFieldWriter struct {
	*FieldWriter
}

func (s SetPrimitiveFieldWriter) Modifier() Modifier {
	return SET
}

func (s SetPrimitiveFieldWriter) FocusTypes() map[FieldType]Void {
	m := map[FieldType]Void{}
	for _, v := range FieldTypeMap {
		if v.FieldType == FEnum || v.FieldType == FMessage || v.FieldType == MAP {
			continue
		}
		m[v.FieldType] = Empty
	}
	return m
}

func (s SetPrimitiveFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	writeBody.Add(s.notNull(value)).Add(LC).NewLine()
	writeBody.Add(N("output.write${type}${suffix}(${fieldNum}, ${value});", map[string]interface{}{
		"type":     strings.ToUpper(fieldConfig.TypeName),
		"suffix":   s.suffix(),
		"fieldNum": fieldConfig.FieldNum,
		"value":    value,
	})).NewLine()
	writeBody.Add(RC).NewLine()
}
func (s SetPrimitiveFieldWriter) notNull(value string) string {
	return collectionNotEmpty(value)
}

func (s SetPrimitiveFieldWriter) suffix() string {
	return "Set"
}

type SetMessageFieldWriter struct {
	*FieldWriter
}

func (s SetMessageFieldWriter) Modifier() Modifier {
	return SET
}

func (s SetMessageFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}
func (s SetMessageFieldWriter) getFieldType() FieldType {
	return FMessage
}
func (s SetMessageFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	writeBody.Add(collectionNotEmpty(value)).Add(LC).NewLine()
	params := map[string]interface{}{
		"typeName":  fieldConfig.TypeName,
		"fieldName": fieldConfig.FieldName,
		"value":     value,
	}
	writeBody.Add(N("for(${typeName} ${fieldName}Unit : ${value}){", params)).NewLine()
	value = N("${fieldName}Unit", params)

}
