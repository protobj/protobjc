package java

import (
	"fmt"
	. "io.protobj/protobj-go/protobj"
	"strings"
)

type SetPrimitiveFieldWriter struct {
	*FieldWriter
}

func (s *SetPrimitiveFieldWriter) Modifier() Modifier {
	return SET
}

func (s *SetPrimitiveFieldWriter) FocusTypes() map[FieldType]Void {
	m := map[FieldType]Void{}
	for _, v := range FieldTypeMap {
		if v.FieldType == FEnum || v.FieldType == FMessage || v.FieldType == MAP {
			continue
		}
		m[v.FieldType] = Empty
	}
	return m
}

func (s *SetPrimitiveFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	writeBody.Add(s.notNull(value)).Add(LC).NewLine()
	writeBody.Add(N("output.write${type}${suffix}(${fieldNum}, ${value});", map[string]interface{}{
		"type":     strings.ToUpper(fieldConfig.TypeName),
		"suffix":   s.suffix(),
		"fieldNum": fieldConfig.FieldNum,
		"value":    value,
	})).NewLine()
	writeBody.Add(RC).NewLine()
}
func (s *SetPrimitiveFieldWriter) notNull(value string) string {
	return collectionNotEmpty(value)
}

func (s *SetPrimitiveFieldWriter) suffix() string {
	return "Set"
}

type SetMessageFieldWriter struct {
	FieldWriter
}

func (s *SetMessageFieldWriter) Modifier() Modifier {
	return SET
}

func (s *SetMessageFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		s.getFieldType(): Empty,
	}
}
func (s *SetMessageFieldWriter) getFieldType() FieldType {
	return FMessage
}

func (s *SetMessageFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	writeBody.Add(collectionNotEmpty(value)).Add(LC).NewLine()
	params := map[string]interface{}{
		"typeName":  fieldConfig.TypeName,
		"fieldName": fieldConfig.FieldName,
		"value":     value,
	}
	writeBody.Add(N("for(${typeName} ${fieldName}Unit : ${value}){", params)).NewLine()
	value = N("${fieldName}Unit", params)
	writer := generator.GetWriter(NewModifier2FieldType(DFT, s.getFieldType()))
	writer.Write(generator, writeBody, sourceMessage, fieldConfig, value)
	writeBody.Add(RC).NewLine()
	writeBody.Add(RC).NewLine()
}

type SetEnumFieldWriter struct {
	SetMessageFieldWriter
}

func (s *SetEnumFieldWriter) getFieldType() FieldType {
	return FEnum
}

type LstEnumFieldWriter struct {
	SetEnumFieldWriter
}

func (receiver *LstEnumFieldWriter) Modifier() Modifier {
	return LST
}

type LstPrimitiveFieldWriter struct {
	*SetPrimitiveFieldWriter
}

func (l *LstPrimitiveFieldWriter) Modifier() Modifier {
	return LST
}
func (l *LstPrimitiveFieldWriter) suffix() string {
	return "List"
}

type LstMessageFieldWriter struct {
	*SetMessageFieldWriter
}

func (l *LstMessageFieldWriter) Modifier() Modifier {
	return LST
}

type ExtMessageFieldWriter struct {
	*FieldWriter
}

func (e *ExtMessageFieldWriter) Modifier() Modifier {
	return EXT
}

func (e *ExtMessageFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}

func (e *ExtMessageFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	params := map[string]interface{}{
		"fieldNum": fieldConfig.FieldNum,
		"typeName": fieldConfig.TypeName,
	}
	writeBody.Add(N("output.writeMessage(${fieldNum}, () -> ${typeName}Schema.writeTo(output, message,false));", params)).NewLine()
	typeFullName := fieldConfig.TypeFullName
	writeBody.AddImportMessage(I("${0}Schema", typeFullName))
}

type ArrPrimitiveFieldWriter struct {
	*SetPrimitiveFieldWriter
}

func (a *ArrPrimitiveFieldWriter) Modifier() Modifier {
	return ARR
}
func (a *ArrPrimitiveFieldWriter) notNull(value string) string {
	return arrayNotEmpty(value)
}

func (a *ArrPrimitiveFieldWriter) suffix() string {
	return "Array"
}

type ArrMessageFieldWriter struct {
	*FieldWriter
}

func (rcvr *ArrMessageFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		rcvr.getFieldType(): Empty,
	}
}

func (rcvr *ArrMessageFieldWriter) getFieldType() FieldType {
	return FMessage
}

func (rcvr *ArrMessageFieldWriter) Modifier() Modifier {
	return ARR
}

func (rcvr *ArrMessageFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	fieldType, _ := generator.GetFieldType(sourceMessage, fieldConfig.TypeName, fieldConfig.TypeFullName)
	writeBody.Add(arrayNotEmpty(value)).Add(LC).NewLine()
	params := map[string]interface{}{
		"typeName":  fieldConfig.TypeName,
		"fieldName": fieldConfig.FieldName,
		"value":     value,
	}
	writeBody.Add(N("${typeName}[] ${fieldName} = ${value};", params)).NewLine()
	writeBody.Add(N("for (int i = ${value}.length-1; i >= 0; i--) ", params)).Add(LC).NewLine()
	writeBody.Add("int index = i;").NewLine()
	value = N("${fieldName}[index]", params)
	writeBody.Add(isNull(value)).Add(LC).NewLine()
	writeBody.Add(CONTINUE).NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add(I("output.writeMessage(${0}, ", fieldConfig.FieldNum)).Add(LAMBDA).Add(LC).NewLine()
	fieldWriter := generator.GetWriter(NewModifier2FieldType(DFT, I32))
	fieldWriter.WritePacked(generator, writeBody, sourceMessage, fieldConfig, I32, "index")
	fieldWriter = generator.GetWriter(NewModifier2FieldType(DFT, rcvr.getFieldType()))
	fieldWriter.WritePacked(generator, writeBody, sourceMessage, fieldConfig, fieldType, value)
	writeBody.Add(RC).Add(");").NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add(RC).NewLine()
}

type ArrEnumFieldWriter struct {
	*ArrMessageFieldWriter
}

func (a *ArrEnumFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		a.getFieldType(): Empty,
	}
}

func (a *ArrEnumFieldWriter) getFieldType() FieldType {
	return FEnum
}

func (a *ArrEnumFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	fieldType, _ := generator.GetFieldType(sourceMessage, fieldConfig.TypeName, fieldConfig.TypeFullName)
	writeBody.Add(arrayNotEmpty(value)).Add(LC).NewLine()
	params := map[string]interface{}{
		"typeName": fieldConfig.TypeName, "fieldName": fieldConfig.FieldName, "value": value,
	}
	writeBody.Add(N("${typeName}[] ${fieldName} = ${value};", params)).NewLine()
	writeBody.Add(I("output.writeMessage(${0}, ", fieldConfig.FieldNum)).Add(LAMBDA).Add(LC).NewLine()
	writeBody.Add(N("for (int i = ${value}.length-1; i >= 0; i--) ", params)).Add(LC).NewLine()
	value = N("${fieldName}[i]", params)
	writeBody.Add(isNull(value)).Add(LC).NewLine()
	writeBody.Add(CONTINUE).NewLine()
	writeBody.Add(RC).NewLine()
	fieldWriter := generator.GetWriter(NewModifier2FieldType(DFT, I32))
	fieldWriter.WritePacked(generator, writeBody, sourceMessage, fieldConfig, I32, "i")
	fieldWriter = generator.GetWriter(NewModifier2FieldType(DFT, a.getFieldType()))
	fieldWriter.WritePacked(generator, writeBody, sourceMessage, fieldConfig, fieldType, value)
	writeBody.Add(RC).NewLine()
	writeBody.Add(RC).Add(");").NewLine()
	writeBody.Add(RC).NewLine()
}

type DftPrimitiveFieldWriter struct {
	*FieldWriter
}

func (d *DftPrimitiveFieldWriter) Modifier() Modifier {
	return DFT
}

func (d *DftPrimitiveFieldWriter) FocusTypes() map[FieldType]Void {
	m := map[FieldType]Void{}
	for _, value := range FieldTypeMap {
		if value.FieldType == FEnum {
			continue
		}
		if value.FieldType == FMessage {
			continue
		}
		if value.FieldType == MAP {
			continue
		}
		m[value.FieldType] = Empty
	}
	return m
}

func (d *DftPrimitiveFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	fieldType, _ := generator.GetFieldType(sourceMessage, fieldConfig.TypeName, fieldConfig.TypeName)
	if fieldType == BOOL {
		value = I("message.is${0}()", fieldConfig.FirstUpperFieldName())
	}
	params := map[string]interface{}{
		"type":     strings.ToUpper(fieldConfig.TypeName),
		"fieldNum": fieldConfig.FieldNum,
		"value":    value,
	}

	writeBody.Add(isNotDefault(value, fieldType)).Add(LC).NewLine()
	writeBody.Add(N("output.write${type}(${fieldNum}, ${value});", params)).NewLine()
	writeBody.Add(RC).NewLine()
}

func (d *DftPrimitiveFieldWriter) WritePacked(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType, value string) {
	params := map[string]interface{}{
		"type":  fieldType.Value().Name,
		"value": value,
	}
	writeBody.Add(N("output.write${type}(${value});", params)).NewLine()
}

type DftMessageFieldWriter struct {
}

func (d *DftMessageFieldWriter) Modifier() Modifier {
	return DFT
}

func (d *DftMessageFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}

func (d *DftMessageFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {

	writeBody.Add(notNull(value)).Add(LC).NewLine()

	messageField, _ := generator.FindMessage(sourceMessage, fieldConfig.TypeFullName)
	if fieldConfig.IsPolymorphic() {
		index := 0
		for _, field := range messageField.ChildMessageConfigMap {
			writeBody.AddImportMessage(field.GetFullName())
			writeBody.AddImportMessage(field.GetFullName() + "Schema")
			var ifstr string
			if index == 0 {
				ifstr = IF
			} else {
				ifstr = ELSE_IF
			}

			index++
			writeBody.Add(ifstr).Add(classEquals(value, field.Name)).Add(LC).NewLine()
			params := map[string]interface{}{"fieldNum": fieldConfig.FieldNum, "typeName": field.Name, "value": value}
			if field.Name == fieldConfig.TypeName {
				writeBody.Add(N("output.writeMessage(${fieldNum}, () -> ${typeName}Schema.writeTo(output, ${value}, true));", params)).NewLine()
			} else {
				writeBody.Add(N("output.writeMessage(${fieldNum}, () -> ${typeName}Schema.writeTo(output, (${typeName})${value}, true));", params)).NewLine()
			}
		}
		writeBody.Add(ELSE).NewLine()
		writeBody.Add(I("throw new RuntimeException(\"undefine message\"+ ${0}.getClass().getName());", value))
		writeBody.Add(RC).NewLine()
	} else {
		writeBody.AddImportMessage(fieldConfig.TypeFullName)
		writeBody.AddImportMessage(fieldConfig.TypeFullName + "Schema")
		writeBody.Add(I("output.writeMessage(${0}, () -> ${1}Schema.writeTo(output, ${2}, false));", fieldConfig.FieldNum, fieldConfig.TypeName, value)).NewLine()
	}

	writeBody.Add(RC).NewLine()
}

func (d *DftMessageFieldWriter) WritePacked(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType, value string) {

	messageField, _ := generator.FindMessage(sourceMessage, fieldConfig.TypeFullName)
	if fieldConfig.IsPolymorphic() {
		index := 0
		for _, field := range messageField.ChildMessageConfigMap {
			writeBody.AddImportMessage(field.GetFullName())
			writeBody.AddImportMessage(field.GetFullName() + "Schema")
			var ifstr string
			if index == 0 {
				ifstr = IF
			} else {
				ifstr = ELSE_IF
			}

			index++
			writeBody.Add(ifstr).Add(classEquals(value, field.Name)).Add(LC).NewLine()
			params := map[string]interface{}{"typeName": field.Name, "value": value}
			if field.Name == fieldConfig.TypeName {
				writeBody.Add(N("${typeName}Schema.writeTo(output, ${value}, true);", params)).NewLine()
			} else {
				writeBody.Add(N("${typeName}Schema.writeTo(output, (${typeName})${value}, true);", params)).NewLine()
			}
		}
		writeBody.Add(ELSE).NewLine()
		writeBody.Add(I("throw new RuntimeException(\"undefine message\"+ ${0}.getClass().getName());", value))
		writeBody.Add(RC).NewLine()
	} else {
		writeBody.AddImportMessage(fieldConfig.TypeFullName)
		writeBody.AddImportMessage(fieldConfig.TypeFullName + "Schema")
		writeBody.Add(I("${0}Schema.writeTo(output, ${1}, false);", fieldConfig.TypeName, value)).NewLine()
	}
}

type DftEnumFieldWriter struct {
}

func (d *DftEnumFieldWriter) Modifier() Modifier {
	return DFT
}

func (d *DftEnumFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FEnum: Empty,
	}
}

func (d *DftEnumFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	writeBody.Add(notNull(value)).Add(LC).NewLine()
	writeBody.AddImportMessage(fieldConfig.TypeFullName)
	writeBody.AddImportMessage(fieldConfig.TypeFullName + "Schema")
	writeBody.Add(I("${1}Schema.writeWithFieldNumber(${0}, output, ${2});", fieldConfig.FieldNum, fieldConfig.TypeName, value)).NewLine()
	writeBody.Add(RC).NewLine()
}

func (d *DftEnumFieldWriter) WritePacked(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType, value string) {
	messageName := fieldConfig.MessageName
	writeBody.AddImportMessage(fieldConfig.FullMessageName())
	writeBody.AddImportMessage(fieldConfig.FullMessageName() + "Schema")
	writeBody.Add(I("${0}Schema.writeTo(output, ${1}, false);", messageName, value)).NewLine()
}

type IMapFieldWriter interface {
	FocusTypes() map[MapKeyValueFieldType]Void
	Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, value string)
	UnsupportedKeyType() map[FieldType]Void
	UnsupportedValueType() map[FieldType]Void
}

type MapFieldWriter struct {
}

func (m *MapFieldWriter) FocusTypes() map[MapKeyValueFieldType]Void {
	panic("implement me")
}

func (m *MapFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, value string) {
	panic("implement me")
}

func (m *MapFieldWriter) UnsupportedKeyType() map[FieldType]Void {
	m2 := map[FieldType]Void{}
	for _, value := range FieldTypeMap {
		if value.FieldType == MAP {
			continue
		}
		if value.FieldType == BOOL {
			continue
		}
		if value.FieldType == FEnum {
			continue
		}
		if value.FieldType == FMessage {
			continue
		}
		m2[value.FieldType] = Empty
	}
	return m2
}

func (m *MapFieldWriter) UnsupportedValueType() map[FieldType]Void {
	return map[FieldType]Void{
		MAP: Empty,
	}
}

type DftMapFieldWriter struct {
	*FieldWriter
	mapFieldWriterMap map[MapKeyValueFieldType]IMapFieldWriter
}

func (w *DftMapFieldWriter) addFieldWriter(writer IMapFieldWriter) {
	for keyValueFieldType, _ := range writer.FocusTypes() {
		_, ok := w.mapFieldWriterMap[keyValueFieldType]
		if ok {
			PrintErrorAndExit(fmt.Sprintf("mapFieldWriter duplicated %v", keyValueFieldType))
		}
		w.mapFieldWriterMap[keyValueFieldType] = writer
	}
}

func NewDftMapFieldWriter() *DftMapFieldWriter {
	writer := &DftMapFieldWriter{
		mapFieldWriterMap: map[MapKeyValueFieldType]IMapFieldWriter{},
	}
	writer.addFieldWriter(&Primitive2PrimitiveMapFieldWriter{})
	writer.addFieldWriter(&Primitive2MessageMapFieldWriter{})
	writer.addFieldWriter(&String2MessageMapFieldWriter{})
	return writer
}

func (w *DftMapFieldWriter) Modifier() Modifier {
	return DFT
}

func (w *DftMapFieldWriter) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		MAP: Empty,
	}
}

func (w *DftMapFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	keyType, _ := generator.GetFieldType(sourceMessage, fieldConfig.KeyType, fieldConfig.KeyType)
	valueType, _ := generator.GetFieldType(sourceMessage, fieldConfig.ValueTypeName, fieldConfig.ValueTypeFullName)
	keyValueFieldType := NewMapKeyValueFieldType(keyType, valueType)
	mapFieldWriter, ok := w.mapFieldWriterMap[keyValueFieldType]
	if !ok {
		PrintErrorAndExit(fmt.Sprintf("unsupported Map<%v,%v>", keyType, valueType))
	}
	mapFieldWriter.Write(generator, writeBody, sourceMessage, fieldConfig, keyValueFieldType, value)
}

type Primitive2PrimitiveMapFieldWriter struct {
	*MapFieldWriter
}

func (p *Primitive2PrimitiveMapFieldWriter) FocusTypes() map[MapKeyValueFieldType]Void {
	m := map[MapKeyValueFieldType]Void{}
	for _, keyType := range FieldTypeMap {
		if _, ok := p.UnsupportedKeyType()[keyType.FieldType]; ok {
			continue
		}
		for _, valueType := range FieldTypeMap {
			if _, ok := p.UnsupportedValueType()[valueType.FieldType]; ok {
				continue
			}
			if valueType.FieldType == FMessage || valueType.FieldType == FEnum {
				continue
			}
			m[NewMapKeyValueFieldType(keyType.FieldType, valueType.FieldType)] = Empty
		}
	}
	return m
}

func (p *Primitive2PrimitiveMapFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, value string) {
	writeBody.Add(collectionNotEmpty(value)).Add(LC).NewLine()
	params := map[string]interface{}{
		"keyType":     keyValueType.KeyType.Value().Name,
		"valueType":   keyValueType.ValueType.Value().Name,
		"fieldNumber": fieldConfig.FieldNum,
		"value":       value,
	}
	writeBody.Add(N("output.write${keyType}${valueType}Map(${fieldNumber}, ${value});", params)).NewLine()
	writeBody.Add(RC).NewLine()
}

type MapFieldParam struct {
	entrySetType  string
	entrySetValue string
	keyValue      string
	valueValue    string
}

type Primitive2MessageMapFieldWriter struct {
	*MapFieldWriter
}

func (p *Primitive2MessageMapFieldWriter) FocusTypes() map[MapKeyValueFieldType]Void {
	m := map[MapKeyValueFieldType]Void{}
	for _, keyType := range FieldTypeMap {

		if _, ok := p.UnsupportedKeyType()[keyType.FieldType]; ok {
			continue
		}

		if keyType.FieldType == STRING {
			continue
		}

		m[NewMapKeyValueFieldType(keyType.FieldType, FMessage)] = Empty
		m[NewMapKeyValueFieldType(keyType.FieldType, FEnum)] = Empty
	}

	return m
}

func (p *Primitive2MessageMapFieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, value string) {

	writeBody.Add(collectionNotEmpty(value)).Add(LC).NewLine()
	keyType := keyValueType.KeyType
	mapFieldParam := p.getMapFieldParam(writeBody, fieldConfig, value, keyType)
	params := map[string]interface{}{
		"entrySetType":  mapFieldParam.entrySetType,
		"entrySetValue": mapFieldParam.entrySetValue,
	}
	writeBody.Add(N("for(${entrySetType} entry : ${entrySetValue})", params)).Add(LC).NewLine()
	if keyType == STRING {
		writeBody.Add(isNull2(mapFieldParam.keyValue, mapFieldParam.valueValue)).Add(LC).NewLine()
	} else {
		writeBody.Add(isNull(mapFieldParam.valueValue)).Add(LC).NewLine()
	}
	writeBody.Add(CONTINUE).NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add(I("output.writeMessage(${0},", fieldConfig.FieldNum)).Add(LAMBDA).Add(LC).NewLine()

	keyWriter := generator.GetWriter(NewModifier2FieldType(DFT, keyType))
	keyWriter.WritePacked(generator, writeBody, sourceMessage, fieldConfig, keyType, mapFieldParam.keyValue)

	valueWriter := generator.GetWriter(NewModifier2FieldType(DFT, keyValueType.ValueType))
	valueWriter.WritePacked(generator, writeBody, sourceMessage, fieldConfig, keyValueType.ValueType, mapFieldParam.valueValue)
	writeBody.Add(RC).Add(");").NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add(RC).NewLine()
}

func (p *Primitive2MessageMapFieldWriter) getMapFieldParam(writeBody *CodeBuilder, fieldConfig *FieldConfig, value string, keyType FieldType) MapFieldParam {
	writeBody.AddImportMessage(fmt.Sprintf("it.unimi.dsi.fastutil.%ss.%s2ObjectMap", keyType.Value().JavaType, FirstUpper(keyType.Value().JavaType)))
	keyJavaType := keyType.Value().JavaType
	KeyJavaType := FirstUpper(keyJavaType)
	params := map[string]interface{}{
		"KeyJavaType": KeyJavaType,
		"keyJavaType": keyJavaType,
		"valueType":   fieldConfig.ValueTypeName,
		"value":       value,
	}
	return MapFieldParam{
		N("${KeyJavaType}2ObjectMap.Entry<${valueType}>", params),
		N("${value}.${keyJavaType}2ObjectEntrySet()", params),
		N("entry.get${KeyJavaType}Key()", params),
		"entry.getValue()",
	}
}

type String2MessageMapFieldWriter struct {
	*Primitive2MessageMapFieldWriter
}

func (s *String2MessageMapFieldWriter) FocusTypes() map[MapKeyValueFieldType]Void {
	m := map[MapKeyValueFieldType]Void{}
	m[NewMapKeyValueFieldType(STRING, FMessage)] = Empty
	m[NewMapKeyValueFieldType(STRING, FEnum)] = Empty
	return m
}

func (s *String2MessageMapFieldWriter) getMapFieldParam(writeBody *CodeBuilder, fieldConfig *FieldConfig, value string, keyType FieldType) MapFieldParam {
	writeBody.AddImportMessage("java.util.Map")
	keyJavaType := keyType.Value().JavaType
	KeyJavaType := FirstUpper(keyJavaType)
	params := map[string]interface{}{
		"KeyJavaType": KeyJavaType,
		"keyJavaType": keyJavaType,
		"valueType":   fieldConfig.ValueTypeName,
		"value":       value,
	}
	return MapFieldParam{
		N("Map.Entry<${KeyJavaType},${valueType}>", params),
		N("${value}.entrySet()", params),
		"entry.getKey()",
		"entry.getValue()",
	}
}
