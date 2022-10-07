package java

import (
	"fmt"
	. "io.protobj/protobjc/protobjc"
	"os"
	"strings"
	"sync"
)

type Generator struct {
	BaseGenerator
	fieldWriterMap map[Modifier2FieldType]IFieldWriter
	fieldReaderMap map[Modifier2FieldType]IFieldReader
}

func NewGenerator(messageMap map[string]*MessageConfig, config ParsedArgs) *Generator {

	generator := &Generator{
		BaseGenerator: BaseGenerator{
			MessageConfigMap: messageMap,
			Config:           config,
		},
		fieldReaderMap: map[Modifier2FieldType]IFieldReader{},
		fieldWriterMap: map[Modifier2FieldType]IFieldWriter{},
	}
	generator.addFieldWriter(NewArrEnumFieldWriter())
	generator.addFieldWriter(NewArrMessageFieldWriter())
	generator.addFieldWriter(NewArrPrimitiveFieldWriter())
	generator.addFieldWriter(&DftEnumFieldWriter{})
	generator.addFieldWriter(NewDftMapFieldWriter())
	generator.addFieldWriter(&DftMessageFieldWriter{})
	generator.addFieldWriter(&DftPrimitiveFieldWriter{})
	generator.addFieldWriter(&ExtMessageFieldWriter{})
	generator.addFieldWriter(NewLstEnumFieldWriter())
	generator.addFieldWriter(NewLstMessageFieldWriter())
	generator.addFieldWriter(NewLstPrimitiveFieldWriter())
	generator.addFieldWriter(NewSetEnumFieldWriter())
	generator.addFieldWriter(NewSetMessageFieldWriter())
	generator.addFieldWriter(NewSetPrimitiveFieldWriter())

	generator.addFieldReader(&ArrEnumFieldReader{})
	generator.addFieldReader(&ArrMessageFieldReader{})
	generator.addFieldReader(NewArrPrimitiveFieldReader())
	generator.addFieldReader(&DftEnumFieldReader{})
	generator.addFieldReader(NewDftMapFieldReader())
	generator.addFieldReader(&DftMessageFieldReader{})
	generator.addFieldReader(&DftPrimitiveFieldReader{})
	generator.addFieldReader(&ExtMessageFieldReader{})
	generator.addFieldReader(NewLstEnumFieldReader())
	generator.addFieldReader(NewLstMessageFieldReader())
	generator.addFieldReader(NewLstPrimitiveFieldReader())
	generator.addFieldReader(NewSetEnumFieldReader())
	generator.addFieldReader(NewSetMessageFieldReader())
	generator.addFieldReader(NewSetPrimitiveFieldReader())

	return generator
}

func (generator *Generator) addFieldReader(fieldReader IFieldReader) {
	modifier := fieldReader.Modifier()
	for focusType, _ := range fieldReader.FocusTypes() {
		modifier2FieldType := NewModifier2FieldType(modifier, focusType)
		if old, ok := generator.fieldReaderMap[modifier2FieldType]; ok {
			PrintErrorAndExit(fmt.Sprintf("fieldReader duplicated %T %T [%s,%s]", fieldReader, old, modifier.Name(), focusType.Value().Name))
		}
		generator.fieldReaderMap[modifier2FieldType] = fieldReader
	}
}

func (generator *Generator) addFieldWriter(fieldWriter IFieldWriter) {
	modifier := fieldWriter.Modifier()
	for focusType, _ := range fieldWriter.FocusTypes() {
		modifier2FieldType := NewModifier2FieldType(modifier, focusType)
		if old, ok := generator.fieldWriterMap[modifier2FieldType]; ok {
			PrintErrorAndExit(fmt.Sprintf("fieldWriter duplicated %T %T [%s,%s]", fieldWriter, old, modifier.Name(), focusType.Value().Name))
		}
		generator.fieldWriterMap[modifier2FieldType] = fieldWriter
	}
}

func (generator *Generator) GetWriter(modifier2FieldType Modifier2FieldType) IFieldWriter {
	writer, ok := generator.fieldWriterMap[modifier2FieldType]
	if !ok {
		PrintErrorAndExit(fmt.Sprintf("fieldWriter not exists [%s,%s]", ModifierName[int32(modifier2FieldType.Modifier)], modifier2FieldType.FieldType.Value().Name))
	}
	return writer
}
func (generator *Generator) GetReader(modifier2FieldType Modifier2FieldType) IFieldReader {
	reader, ok := generator.fieldReaderMap[modifier2FieldType]
	if !ok {
		PrintErrorAndExit(fmt.Sprintf("fieldReader not exists [%s,%s]", ModifierName[int32(modifier2FieldType.Modifier)], modifier2FieldType.FieldType.Value().Name))
	}
	return reader
}

func (generator *Generator) LanguageType() LanguageType {
	return JAVA
}
func (generator *Generator) Generate() {
	messageConfigMap := generator.MessageConfigMap
	var waitGroup sync.WaitGroup
	var fileContentsChan = make(chan *FileContent)
	for _, messageConfig := range messageConfigMap {
		if generator.Config.OutputType != O_SCHEMA {
			waitGroup.Add(1)
			go func(message *MessageConfig) {
				defer waitGroup.Done()
				fileContent := generator.createMessage(message)
				if fileContent != nil {
					fileContentsChan <- fileContent
				}
			}(messageConfig)
		}
		if generator.Config.OutputType != O_MESSAGE {
			waitGroup.Add(1)
			go func(message *MessageConfig) {
				defer waitGroup.Done()
				fileContent := generator.createSchema(message)
				if fileContent != nil {
					fileContentsChan <- fileContent
				}

			}(messageConfig)
		}

	}
	go func() {
		waitGroup.Wait()
		close(fileContentsChan)
	}()
	for {
		fileContent, ok := <-fileContentsChan
		if !ok {
			break
		}
		waitGroup.Add(1)
		go func(content *FileContent) {
			defer waitGroup.Done()
			WriteFile(generator.Config.OutputDir, fileContent)
		}(fileContent)
	}
	waitGroup.Wait()
}

func (generator *Generator) createMessage(m *MessageConfig) *FileContent {
	switch m.MessageType {
	case ENUM:
		return generator.createEnumClass(m)
	case MESSAGE:
		return generator.createMessageClass(m)
	default:
		return nil
	}
}

func (generator *Generator) createEnumClass(m *MessageConfig) *FileContent {
	header := NewCodeBuilder()
	header.Add(pkg(m.Pkg)).NewLine(2)
	if len(m.Note) > 0 {
		header.Add(I("//${0}", m.Note)).NewLine()
	}
	header.Add(I("public enum ${0} {", m.Name)).NewLine()
	for _, fieldConfig := range m.GetSortedFields() {
		header.Add(I("//${0} = ${1};", fieldConfig.FieldName, fieldConfig.FieldNum))
		if len(fieldConfig.Note) > 0 {
			header.Add(fieldConfig.Note)
		}
		header.NewLine()
		header.Add(I("${0},", fieldConfig.FieldName)).NewLine()
	}
	header.Add(";").NewLine()
	header.Add("}").NewLine()
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName(), ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())
}

func (generator *Generator) createMessageClass(m *MessageConfig) *FileContent {
	header := NewCodeBuilder()
	header.Add(pkg(m.Pkg)).NewLine(2)

	fields := NewCodeBuilder()
	methods := NewCodeBuilder()
	methods.SetCurrent(1)

	if len(m.Note) > 0 {
		fields.Add(I("//${0}", m.Note)).NewLine()
	}
	parent := m.ExtMessage

	if parent != nil {
		extField := m.ExtField
		fields.Add("//").Add(extField.GetDefinition()).NewLine()
		fields.Add(I("public class ${0} extends ${1} {", m.Name, parent.Name)).NewLine(2)
		fields.AddImportMessage(parent.GetFullName())
	} else {
		fields.Add(I("public class ${0} {", m.Name)).NewLine(2)
	}
	var fieldList = m.GetSortedFields()
	for _, field := range fieldList {
		modifier := field.Modifier
		typeName := field.TypeName
		fieldType, err := FieldTypeValueOf(typeName)
		var typeAndImport *TypeAndImport
		if err == nil {
			typeAndImport = getTypeAndImportFromBuiltinType(modifier, fieldType)
			if typeAndImport == nil && fieldType == MAP {
				keyFieldType, err := FieldTypeValueOf(field.KeyType)
				if err != nil {
					PrintErrorAndExit(err.Error())
				}
				valueFieldType, err := FieldTypeValueOf(field.ValueTypeName)
				if err == nil {
					if keyFieldType == STRING && valueFieldType == STRING {
						typeAndImport = NewTypeAndImport("Map<String,String>", "java.util.Map")
					} else if keyFieldType == STRING {
						mapType := I("Object2${0}Map<String>", FirstUpper(valueFieldType.Value().JavaType))
						typeAndImport = NewTypeAndImport(mapType, I("it.unimi.dsi.fastutil.objects.Object2${0}Map", FirstUpper(valueFieldType.Value().JavaType)))
					} else if valueFieldType == STRING {
						mapType := I("${0}2ObjectMap<String>", FirstUpper(keyFieldType.Value().JavaType))
						typeAndImport = NewTypeAndImport(mapType, I("it.unimi.dsi.fastutil.${0}s.${1}2ObjectMap", keyFieldType.Value().JavaType, FirstUpper(keyFieldType.Value().JavaType)))
					} else {
						mapType := FirstUpper(keyFieldType.Value().JavaType) + "2" + FirstUpper(valueFieldType.Value().JavaType) + "Map"
						typeAndImport = NewTypeAndImport(mapType, I("it.unimi.dsi.fastutil.${0}s.${1}", keyFieldType.Value().JavaType, mapType))
					}
				} else {
					if keyFieldType == STRING {
						mapType := I("Map<String,${0}>", field.ValueTypeName)
						typeAndImport = NewTypeAndImport(mapType, "java.util.Map", field.ValueTypeFullName)
					} else {
						mapType := I("${0}2ObjectMap<${1}>", FirstUpper(keyFieldType.Value().JavaType), field.ValueTypeName)
						mapType0 := I("${0}2ObjectMap", FirstUpper(keyFieldType.Value().JavaType))
						typeAndImport = NewTypeAndImport(mapType, I("it.unimi.dsi.fastutil.${0}s.${1}",
							keyFieldType.Value().JavaType, mapType0), field.ValueTypeFullName)
					}
				}
			}
		} else {
			var typeFullName = field.TypeFullName
			message, _ := generator.FindMessage(m, typeFullName)
			switch modifier {
			case DFT:
				typeAndImport = NewTypeAndImport(field.TypeName, field.TypeFullName)
			case SET:
				switch message.MessageType {
				case ENUM:
					typeAndImport = NewTypeAndImport(I("EnumSet<${0}>", field.TypeName), field.TypeFullName, "java.util.EnumSet")
				case MESSAGE:
					typeAndImport = NewTypeAndImport(I("Set<${0}>", field.TypeName), field.TypeFullName, "java.util.Set")
				}
			case LST:
				typeAndImport = NewTypeAndImport(I("List<${0}>", field.TypeName), field.TypeFullName, "java.util.List")
			case ARR:
				typeAndImport = NewTypeAndImport(I("${0}[]", field.TypeName), field.TypeFullName)
			case EXT:
				typeAndImport = nil
			}
			if message.MessageType == MESSAGE && modifier == EXT {
				continue
			}
		}

		if typeAndImport == nil {
			PrintErrorAndExit(I("field type not found ${0} ${1}", m.GetFullName(), field.TypeFullName))
		}
		createField(m, fields, methods, field, typeAndImport)
	}
	header.AddImportMessages(fields.ImportMessages)
	header.AddImportMessages(methods.ImportMessages)
	appendImportMessagesForJava(m.Pkg, header)
	header.AddBuilder(fields).AddBuilder(methods).Add("}").NewLine()
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName(), ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())
}

func createField(m *MessageConfig, fields *CodeBuilder, methods *CodeBuilder, field *FieldConfig, typeAndImport *TypeAndImport) {
	if len(typeAndImport.Imports) > 0 {
		for _, importMessage := range typeAndImport.Imports {
			if importMessage == typeAndImport.Type {
				continue
			}
			fields.AddImportMessage(importMessage)
		}
	}
	fields.Add("//").Add(field.GetDefinition()).NewLine()

	fields.Add(I("private ${0} ${1};", typeAndImport.Type, field.FieldName)).NewLine(2)
	firstUpper := FirstUpper(field.FieldName)
	methods.Add(I("public void set${0}(${1} ${2}) {", firstUpper, typeAndImport.Type, field.FieldName)).NewLine()
	methods.Add(I("this.${0} = ${1};", field.FieldName, field.FieldName)).NewLine()
	methods.Add("}").NewLine(2)

	var getPrefix string
	if typeAndImport.Type == "boolean" {
		getPrefix = "is"
	} else {
		getPrefix = "get"
	}
	methods.Add(I("public ${0} ${1}${2}() {", typeAndImport.Type, getPrefix, firstUpper)).NewLine()
	methods.Add(I("return ${0};", field.FieldName)).NewLine()
	methods.Add("}").NewLine(2)
}

func getTypeAndImportFromBuiltinType(modifier Modifier, fieldType FieldType) *TypeAndImport {
	switch fieldType {
	case BOOL:
		switch modifier {
		case DFT:
			return NewTypeAndImport("boolean")
		case LST:
			return NewTypeAndImport("BooleanList", "it.unimi.dsi.fastutil.booleans.BooleanList")
		case SET:
			return NewTypeAndImport("BooleanSet", "it.unimi.dsi.fastutil.booleans.BooleanSet")
		case ARR:
			return NewTypeAndImport("boolean[]")
		default:
			return nil
		}
	case I8, U8:
		switch modifier {
		case DFT:
			return NewTypeAndImport("byte")
		case LST:
			return NewTypeAndImport("ByteList", "it.unimi.dsi.fastutil.bytes.ByteList")
		case SET:
			return NewTypeAndImport("ByteSet", "it.unimi.dsi.fastutil.bytes.ByteSet")
		case ARR:
			return NewTypeAndImport("byte[]")
		default:
			return nil
		}
	case I16, U16:
		switch modifier {
		case DFT:
			return NewTypeAndImport("short")
		case LST:
			return NewTypeAndImport("ShortList", "it.unimi.dsi.fastutil.shorts.ShortList")
		case SET:
			return NewTypeAndImport("ShortSet", "it.unimi.dsi.fastutil.shorts.ShortSet")
		case ARR:
			return NewTypeAndImport("short[]")
		default:
			return nil
		}
	case I32, U32, S32, F32, SF32:
		switch modifier {
		case DFT:
			return NewTypeAndImport("int")
		case LST:
			return NewTypeAndImport("IntList", "it.unimi.dsi.fastutil.ints.IntList")
		case SET:
			return NewTypeAndImport("IntSet", "it.unimi.dsi.fastutil.ints.IntSet")
		case ARR:
			return NewTypeAndImport("int[]")
		default:
			return nil
		}
	case I64, U64, S64, F64, SF64:
		switch modifier {
		case DFT:
			return NewTypeAndImport("long")
		case LST:
			return NewTypeAndImport("LongList", "it.unimi.dsi.fastutil.longs.LongList")
		case SET:
			return NewTypeAndImport("LongSet", "it.unimi.dsi.fastutil.longs.LongSet")
		case ARR:
			return NewTypeAndImport("long[]")
		default:
			return nil
		}
	case STRING:
		switch modifier {
		case DFT:
			return NewTypeAndImport("String")
		case LST:
			return NewTypeAndImport("List<String>", "java.util.List")
		case SET:
			return NewTypeAndImport("Set<String>", "java.util.Set")
		case ARR:
			return NewTypeAndImport("String[]")
		default:
			return nil
		}
	case FLOAT:
		switch modifier {
		case DFT:
			return NewTypeAndImport("float")
		case LST:
			return NewTypeAndImport("FloatList", "it.unimi.dsi.fastutil.floats.FloatList")
		case SET:
			return NewTypeAndImport("FloatSet", "it.unimi.dsi.fastutil.floats.FloatSet")
		case ARR:
			return NewTypeAndImport("float[]")
		default:
			return nil
		}
	case DOUBLE:
		switch modifier {
		case DFT:
			return NewTypeAndImport("double")
		case LST:
			return NewTypeAndImport("DoubleList", "it.unimi.dsi.fastutil.doubles.DoubleList")
		case SET:
			return NewTypeAndImport("DoubleSet", "it.unimi.dsi.fastutil.doubles.DoubleSet")
		case ARR:
			return NewTypeAndImport("double[]")
		default:
			return nil
		}
	default:
		return nil
	}
}

func (generator *Generator) createSchema(m *MessageConfig) *FileContent {
	switch m.MessageType {
	case ENUM:
		return generator.createEnumSchema(m)
	case MESSAGE:
		return generator.createMessageSchema(m)
	}
	return nil
}

func (generator *Generator) createEnumSchema(m *MessageConfig) *FileContent {
	p := m.Pkg
	header := NewCodeBuilder()
	header.Add(pkg(p)).NewLine(2)

	header.AddImportMessage("io.protobj.core.Input")
	header.AddImportMessage("io.protobj.core.Output")
	header.AddImportMessage("io.protobj.core.Schema")
	header.AddImportMessage("java.io.IOException")

	var writeBody = generator.createEnumWriteBody(m, false)
	var writeWithFieldNumberBody = generator.createEnumWriteBody(m, true)
	var readBody = generator.createEnumReadBody(m)

	body := NewCodeBuilder()
	body.Add(N(EnumSchemaTemplate, map[string]interface{}{
		"class":                    m.Name,
		"writeBody":                writeBody.String(),
		"writeWithFieldNumberBody": writeWithFieldNumberBody.String(),
		"readBody":                 readBody.String(),
	})).NewLine()
	appendImportMessagesForJava(p, header)
	header.AddBuilder(body)
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName()+"Schema", ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())
}

func (generator *Generator) createEnumWriteBody(m *MessageConfig, withFieldNum bool) *CodeBuilder {
	writeBody := NewCodeBuilder()
	writeBody.SetCurrent(2)
	writeBody.Add(isNull("message")).Add(LC).NewLine()
	writeBody.Add("output.writeI32(0);").NewLine()
	writeBody.Add("return;").NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add("switch (message) ").Add(LC).NewLine()
	for _, value := range m.GetSortedFields() {
		writeBody.Add(I("case ${0}: ", value.FieldName)).Add(LC).NewLine()
		if withFieldNum {
			writeBody.Add(I("output.writeI32(fieldNum,${0});", value.FieldNum)).NewLine()
		} else {
			writeBody.Add(I("output.writeI32(${0});", value.FieldNum)).NewLine()
		}
		writeBody.Add("break;").NewLine()
		writeBody.Add(RC).NewLine()
	}
	writeBody.Add("default: ").Add(LC).NewLine()
	writeBody.Add("throw new RuntimeException(\"undefine enum \" + message);").NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add(RC).NewLine()
	return writeBody
}

func (generator *Generator) createEnumReadBody(m *MessageConfig) *CodeBuilder {
	readBody := NewCodeBuilder()
	readBody.SetCurrent(2)
	readBody.Add("int value = input.readI32();").NewLine()
	readBody.Add("switch (value) ").Add(LC).NewLine()
	for _, value := range m.GetSortedFields() {
		readBody.Add(I("case ${0}: ", value.FieldNum)).Add(LC).NewLine()
		readBody.Add(I("return ${0}.${1};", m.Name, value.FieldName)).NewLine()
		readBody.Add(RC).NewLine()
	}
	readBody.Add("default: ").Add(LC).NewLine()
	readBody.Add("return null;").NewLine()
	readBody.Add(RC).NewLine()
	readBody.Add(RC).NewLine()
	return readBody
}

func (generator *Generator) createMessageSchema(m *MessageConfig) *FileContent {
	p := m.Pkg
	header := NewCodeBuilder()
	header.Add(pkg(p)).NewLine(2)

	header.AddImportMessage("io.protobj.core.Input")
	header.AddImportMessage("io.protobj.core.Output")
	header.AddImportMessage("io.protobj.core.Schema")
	header.AddImportMessage("java.io.IOException")

	var writeBody = generator.createWriteBody(m)
	var readBody = generator.createReadBody(m)
	body := NewCodeBuilder()
	body.Add(N(MessageSchemaTemplate, map[string]interface{}{
		"class":        m.Name,
		"writeBody":    writeBody.String(),
		"readBody":     readBody.String(),
		"messageIndex": m.MessageIndex,
	})).NewLine()

	header.AddImportMessages(writeBody.ImportMessages)
	header.AddImportMessages(readBody.ImportMessages)
	appendImportMessagesForJava(p, header)
	header.AddBuilder(body)
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName()+"Schema", ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())

}

func (generator *Generator) createWriteBody(m *MessageConfig) *CodeBuilder {
	writeBody := NewCodeBuilder()
	writeBody.SetCurrent(2)
	for _, field := range m.GetSortedFields() {
		modifier := field.Modifier
		fieldType, _ := generator.GetFieldType(m, field.TypeName, field.TypeFullName)
		var fieldWriter = generator.GetWriter(NewModifier2FieldType(modifier, fieldType))
		getValue := I("message.get${0}()", field.FirstUpperFieldName())
		writeBody.Add("//").Add(field.GetDefinition()).NewLine()
		fieldWriter.Write(generator, writeBody, m, field, getValue)
	}
	return writeBody
}

func (generator *Generator) createReadBody(m *MessageConfig) *CodeBuilder {
	readBody := NewCodeBuilder()
	readBody.SetCurrent(4)
	for _, field := range m.GetSortedFields() {
		modifier := field.Modifier
		fieldType, _ := generator.GetFieldType(m, field.TypeName, field.TypeFullName)
		reader := generator.GetReader(NewModifier2FieldType(modifier, fieldType))
		getValue := I("message.get${0}()", field.FirstUpperFieldName())
		setValue := fmt.Sprintf("message.set%s(${value})", field.FirstUpperFieldName())
		readBody.Add("//").Add(field.GetDefinition()).NewLine()
		readBody.Add(I("case ${0}: ", field.FieldNum)).Add(LC).NewLine()
		reader.Read(generator, readBody, m, field, getValue, setValue)
		readBody.Add("break;").NewLine()
		readBody.Add(RC).NewLine()
	}
	return readBody
}

type TypeAndImport struct {
	Type    string
	Imports []string
}

func NewTypeAndImport(Type string, imports ...string) *TypeAndImport {
	return &TypeAndImport{
		Type:    Type,
		Imports: imports,
	}
}
