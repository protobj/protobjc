package java

import (
	"fmt"
	. "io.protobj/protobj-go/protobj"
)

type SetPrimitiveFieldReader struct {
	*FieldReader
}

func (s *SetPrimitiveFieldReader) Modifier() Modifier {
	return SET
}

func (s *SetPrimitiveFieldReader) FocusTypes() map[FieldType]Void {
	m := map[FieldType]Void{}
	for _, v := range FieldTypeMap {
		if v.FieldType == FEnum || v.FieldType == FMessage || v.FieldType == MAP {
			continue
		}
		m[v.FieldType] = Empty
	}
	return m
}

func (s *SetPrimitiveFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	fieldType, _ := generator.GetFieldType(sourceMessage, fieldConfig.TypeName, fieldConfig.TypeName)
	value := I("input.read${0}${1}()", fieldType.Value().Name, s.readSuffix())
	readBody.Add(NI(setValue, "value", value)).Add(";").NewLine()
}

func (s *SetPrimitiveFieldReader) readSuffix() string {
	return "Set"
}

type SetMessageFieldReader struct {
	*FieldReader
}

func (s *SetMessageFieldReader) Modifier() Modifier {
	return SET
}

func (s *SetMessageFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}

func (s *SetMessageFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	readBody.Add(isNull(getValue)).Add(LC).NewLine()
	readBody.Add(NI(setValue, "value", s.New(readBody))).Add(";").NewLine()
	readBody.Add(RC).NewLine()

	readBody.Add(readMessageStart()).NewLine()
	reader := generator.GetReader(NewModifier2FieldType(DFT, FMessage))
	message, _ := generator.FindMessage(sourceMessage, fieldConfig.FullMessageName())
	value := reader.ReadPacked(generator, readBody, message, fieldConfig, FMessage)
	setValue = NI("${getValue}.add(${value})", "getValue", getValue, "value", value)
	readBody.Add(setValue).Add(";").NewLine()
	readBody.Add(readMessageStop()).NewLine()
}

func (s *SetMessageFieldReader) New(readBody *CodeBuilder) string {
	readBody.AddImportMessage("java.util.HashSet")
	return "new HashSet<>()"
}

type SetEnumFieldReader struct {
	*FieldReader
}

func (s *SetEnumFieldReader) Modifier() Modifier {
	return SET
}

func (s *SetEnumFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FEnum: Empty,
	}
}

func (s *SetEnumFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	readBody.Add(isNull(getValue)).Add(LC).NewLine()
	readBody.Add(NI(setValue, "value", s.New(readBody, fieldConfig.TypeName))).Add(";").NewLine()
	readBody.Add(RC).NewLine()
	reader := generator.GetReader(NewModifier2FieldType(DFT, FEnum))
	value := reader.ReadPacked(generator, readBody, sourceMessage, fieldConfig, FEnum)
	readBody.Add(NI("${getValue}.add(${value});", "getValue", getValue, "value", value)).Add(";").NewLine()
}

func (s *SetEnumFieldReader) New(readBody *CodeBuilder, fieldTypeName string) string {
	readBody.AddImportMessage("java.util.EnumSet")
	return I("EnumSet.noneOf(${0}.class)", fieldTypeName)
}

type LstPrimitiveFieldReader struct {
	*SetPrimitiveFieldReader
}

func (s *LstPrimitiveFieldReader) readSuffix() string {
	return "List"
}

func (s *LstPrimitiveFieldReader) Modifier() Modifier {
	return LST
}

type LstMessageFieldReader struct {
	*SetMessageFieldReader
}

func (s *LstMessageFieldReader) Modifier() Modifier {
	return LST
}

func (s *LstMessageFieldReader) New(readBody *CodeBuilder) string {
	readBody.AddImportMessage("java.util.ArrayList")
	return "new ArrayList<>()"
}

type LstEnumFieldReader struct {
	*SetEnumFieldReader
}

func (s *LstEnumFieldReader) Modifier() Modifier {
	return LST
}

func (s *LstEnumFieldReader) New(readBody *CodeBuilder, fieldTypeName string) string {
	readBody.AddImportMessage("java.util.ArrayList")
	return "new ArrayList<>()"
}

type ExtMessageFieldReader struct {
	*FieldReader
}

func (e *ExtMessageFieldReader) Modifier() Modifier {
	return EXT
}

func (e *ExtMessageFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}

func (e *ExtMessageFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	readBody.AddImportMessage(sourceMessage.GetFullName())
	readBody.AddImportMessage(I("%sSchema", sourceMessage.GetFullName()))
	readBody.Add(readMessageStart()).NewLine()
	readBody.Add(I("${0}Schema.mergeFrom(input,message)", fieldConfig.TypeName)).Add(";").NewLine()
	readBody.Add(readMessageStop()).NewLine()
}

type ArrPrimitiveFieldReader struct {
	*SetPrimitiveFieldReader
}

func (a *ArrPrimitiveFieldReader) Modifier() Modifier {
	return ARR
}

func (a *ArrPrimitiveFieldReader) readSuffix() string {
	return "Array"
}

type ArrMessageFieldReader struct {
	*FieldReader
}

func (a *ArrMessageFieldReader) Modifier() Modifier {
	return ARR
}

func (a *ArrMessageFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}

func (a *ArrMessageFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	readBody.Add(readMessageStart()).NewLine()
	a.read0(generator, readBody, sourceMessage, fieldConfig, getValue, setValue)
	readBody.Add(readMessageStop()).NewLine()

}

func (a *ArrMessageFieldReader) read0(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue string, setValue string) {
	indexReader := generator.GetReader(NewModifier2FieldType(DFT, I32))
	indexValue := indexReader.ReadPacked(generator, readBody, sourceMessage, fieldConfig, I32)
	readBody.Add("int index = ").Add(indexValue).Add(";").NewLine()
	readBody.Add(isNull(getValue)).Add(LC).NewLine()
	instance := I("new ${0}[index + 1]", fieldConfig.TypeName)
	readBody.Add(NI(setValue, "value", instance)).Add(";").NewLine()
	readBody.Add(RC).NewLine()
	setArrValue := I("${0}[index] = ", getValue)

	reader := generator.GetReader(NewModifier2FieldType(DFT, FMessage))
	message, _ := generator.FindMessage(sourceMessage, fieldConfig.FullMessageName())
	value := reader.ReadPacked(generator, readBody, message, fieldConfig, FMessage)
	readBody.Add(setArrValue).Add(value).Add(";").NewLine()
}

type ArrEnumFieldReader struct {
	*FieldReader
}

func (a *ArrEnumFieldReader) Modifier() Modifier {
	return ARR
}

func (a *ArrEnumFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FEnum: Empty,
	}
}

func (a *ArrEnumFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	indexReader := generator.GetReader(NewModifier2FieldType(DFT, I32))
	indexValue := indexReader.ReadPacked(generator, readBody, sourceMessage, fieldConfig, I32)
	readBody.Add("int index = ").Add(indexValue).Add(";").NewLine()
	readBody.Add(isNull(getValue)).Add(LC).NewLine()
	instance := I("new ${0}[index + 1]", fieldConfig.TypeName)
	readBody.Add(NI(setValue, "value", instance)).Add(";").NewLine()
	readBody.Add(RC).NewLine()
	readBody.Add("do").Add(LC).NewLine()
	setArrValue := I("${0}[index] = ", getValue)

	reader := generator.GetReader(NewModifier2FieldType(DFT, FMessage))
	value := reader.ReadPacked(generator, readBody, sourceMessage, fieldConfig, FMessage)
	readBody.Add(setArrValue).Add(value).Add(";").NewLine()
	readBody.Add("if (index == 0)").Add(LC).NewLine()
	readBody.Add("break;").NewLine()
	readBody.Add(RC).NewLine()
	readBody.Add(NI("} while ((index = ${indexValue}) >=0);", "indexValue", indexValue)).NewLine()

}

type DftPrimitiveFieldReader struct {
	*FieldReader
}

func (d *DftPrimitiveFieldReader) Modifier() Modifier {
	return DFT
}

func (d *DftPrimitiveFieldReader) FocusTypes() map[FieldType]Void {
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

func (d *DftPrimitiveFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	fieldType, _ := generator.GetFieldType(sourceMessage, fieldConfig.TypeName, fieldConfig.TypeName)
	value := d.ReadPacked(generator, readBody, sourceMessage, fieldConfig, fieldType)
	readBody.Add(NI(setValue, "value", value)).Add(";").NewLine()
}

func (d *DftPrimitiveFieldReader) ReadPacked(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType) string {
	return I("input.read${0}()", fieldType.Value().Name)
}

type DftMessageFieldReader struct {
	*FieldReader
}

func (d *DftMessageFieldReader) Modifier() Modifier {
	return DFT
}

func (d *DftMessageFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FMessage: Empty,
	}
}

func (d *DftMessageFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	readBody.Add(readMessageStart()).NewLine()
	packValue := d.ReadPacked(generator, readBody, sourceMessage, fieldConfig, FMessage)
	readBody.Add(NI(setValue, "value", packValue)).Add(";").NewLine()
	readBody.Add(readMessageStop()).NewLine()

}

func (d *DftMessageFieldReader) ReadPacked(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType) string {
	messageName := fieldConfig.MessageName()
	messageFullName := fieldConfig.FullMessageName()
	polymorphic := fieldConfig.IsPolymorphic()
	if polymorphic {
		reader := generator.GetReader(NewModifier2FieldType(DFT, I32))
		readBody.Add("int msgIndex = ").Add(reader.ReadPacked(generator, readBody, sourceMessage, fieldConfig, I32)).Add(";").NewLine()
		readBody.Add(I("${0} packValue = null;", messageName)).NewLine()
		readBody.Add("switch(msgIndex) ").Add(LC).NewLine()
		for _, polyMessage := range sourceMessage.ChildMessageConfigMap {
			readBody.Add(I("case ${0} :", polyMessage.MessageIndex)).Add(LC).NewLine()
			readBody.Add("packValue = ").Add(d.readPacked0(readBody, polyMessage.Name, polyMessage.GetFullName())).Add(";").NewLine()
			readBody.Add("break;").NewLine()
			readBody.Add(RC).NewLine()
		}

		readBody.Add("default: ").Add(LC).NewLine()
		readBody.Add("input.handleUnknownPolymorphicField(msgIndex);").NewLine()
		readBody.Add("break;").NewLine()
		readBody.Add(RC).NewLine()

		readBody.Add(RC).NewLine()
		return "packValue"
	} else {
		return d.readPacked0(readBody, messageName, messageFullName)
	}
}
func (d *DftMessageFieldReader) readPacked0(readBody *CodeBuilder, messageName string, messageFullName string) string {
	readBody.AddImportMessage(messageFullName)
	readBody.AddImportMessage(messageFullName + "Schema")
	return I("${0}Schema.mergeFrom(input,null)", messageName)
}

type DftEnumFieldReader struct {
	*DftMessageFieldReader
}

func (d *DftEnumFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		FEnum: Empty,
	}
}

func (d *DftEnumFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	value := d.readPacked0(readBody, fieldConfig.TypeName, fieldConfig.TypeFullName)
	readBody.Add(NI(setValue, "value", value)).Add(";").NewLine()
}

func (d *DftEnumFieldReader) ReadPacked(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType) string {
	return d.readPacked0(readBody, fieldConfig.MessageName(), fieldConfig.FullMessageName())
}

type IMapFieldReader interface {
	FocusTypes() map[MapKeyValueFieldType]Void
	Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, getValue, setValue string)
	UnsupportedKeyType() map[FieldType]Void
	UnsupportedValueType() map[FieldType]Void
}
type MapFieldReader struct {
}

func (m *MapFieldReader) FocusTypes() map[MapKeyValueFieldType]Void {
	//TODO implement me
	panic("implement me")
}

func (m *MapFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, getValue, setValue string) {
	panic("implement me")
}

func (m *MapFieldReader) UnsupportedKeyType() map[FieldType]Void {
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

func (m *MapFieldReader) UnsupportedValueType() map[FieldType]Void {
	return map[FieldType]Void{
		MAP: Empty,
	}
}

type DftMapFieldReader struct {
	*FieldReader
	mapFieldReaderMap map[MapKeyValueFieldType]IMapFieldReader
}

func (d *DftMapFieldReader) addFieldReader(reader IMapFieldReader) {
	for keyValueFieldType, _ := range reader.FocusTypes() {
		_, ok := d.mapFieldReaderMap[keyValueFieldType]
		if ok {
			PrintErrorAndExit(fmt.Sprintf("mapFieldReader duplicated %v", keyValueFieldType))
		}
		d.mapFieldReaderMap[keyValueFieldType] = reader
	}
}

func NewDftMapFieldReader() *DftMapFieldReader {
	reader := DftMapFieldReader{}
	reader.addFieldReader(&Primitive2PrimitiveMapFieldReader{})
	reader.addFieldReader(&Primitive2MessageMapFieldReader{})
	return &reader
}

func (d *DftMapFieldReader) Modifier() Modifier {
	return DFT
}

func (d *DftMapFieldReader) FocusTypes() map[FieldType]Void {
	return map[FieldType]Void{
		MAP: Empty,
	}
}

func (d *DftMapFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	keyType, _ := generator.GetFieldType(sourceMessage, fieldConfig.KeyType, fieldConfig.KeyType)
	valueType, _ := generator.GetFieldType(sourceMessage, fieldConfig.ValueTypeName, fieldConfig.ValueTypeFullName)
	keyValueFieldType := NewMapKeyValueFieldType(keyType, valueType)
	mapFieldReader, ok := d.mapFieldReaderMap[keyValueFieldType]
	if !ok {
		PrintErrorAndExit(fmt.Sprintf("unsupported Map<%v,%v>", keyType, valueType))
	}
	mapFieldReader.Read(generator, readBody, sourceMessage, fieldConfig, keyValueFieldType, getValue, setValue)
}

type Primitive2MessageMapFieldReader struct {
	MapFieldReader
}

func (p *Primitive2MessageMapFieldReader) FocusTypes() map[MapKeyValueFieldType]Void {
	m := map[MapKeyValueFieldType]Void{}
	for _, keyType := range FieldTypeMap {
		if _, ok := p.UnsupportedKeyType()[keyType.FieldType]; ok {
			continue
		}
		m[NewMapKeyValueFieldType(keyType.FieldType, FMessage)] = Empty
		m[NewMapKeyValueFieldType(keyType.FieldType, FEnum)] = Empty
	}
	return m
}

func (p *Primitive2MessageMapFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, getValue, setValue string) {
	fieldMessage, _ := generator.FindMessage(sourceMessage, fieldConfig.ValueTypeFullName)

	readBody.Add(isNull(getValue)).Add(LC).NewLine()
	keyType := keyValueType.KeyType
	var newMap string
	if keyType != STRING {
		readBody.AddImportMessage(NI("it.unimi.dsi.fastutil.${keyType}s.${KeyType}2ObjectOpenHashMap",
			"keyType", keyType.Value().JavaType, "KeyType", FirstUpper(keyType.Value().JavaType)))
		newMap = NI("new ${KeyType}2ObjectOpenHashMap<>()", "KeyType", FirstUpper(keyType.Value().JavaType))
	} else {
		readBody.AddImportMessage("java.util.HashMap")
		newMap = "new HashMap<>()"
	}
	readBody.Add(NI(setValue, "value", newMap)).Add(";").NewLine()
	readBody.Add(RC).NewLine()

	readBody.Add(readMessageStart()).NewLine()
	keyReader := generator.GetReader(NewModifier2FieldType(DFT, keyType))
	keyValue := keyReader.ReadPacked(generator, readBody, sourceMessage, fieldConfig, keyType)
	readBody.Add(NI("${keyType} key = ", "keyType", keyType.Value().JavaType)).Add(keyValue).Add(";").NewLine()
	valueFieldType := fieldMessage.MessageType.ToFieldType()
	valueReader := generator.GetReader(NewModifier2FieldType(DFT, valueFieldType))
	mapValueValue := valueReader.ReadPacked(generator, readBody, fieldMessage, fieldConfig, valueFieldType)
	readBody.Add(getValue).Add(I(".put(key, ${0});", mapValueValue)).NewLine()
	readBody.Add(readMessageStop()).NewLine()
}

type Primitive2PrimitiveMapFieldReader struct {
	*MapFieldReader
}

func (p *Primitive2PrimitiveMapFieldReader) FocusTypes() map[MapKeyValueFieldType]Void {
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

func (p *Primitive2PrimitiveMapFieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, keyValueType MapKeyValueFieldType, getValue, setValue string) {
	keyType := keyValueType.KeyType
	valueType := keyValueType.ValueType
	mapValue := NI("input.read${keyType}${valueType}Map()", "keyType", keyType.Value().Name, "valueType", valueType.Value().Name)
	readBody.Add(NI(setValue, "value", mapValue)).Add(";").NewLine()
}
