package ts

import (
	"fmt"
	. "io.protobj/protobjc"
	"os"
	"strings"
	"sync"
)

type Generator struct {
	BaseGenerator
}

func NewGenerator(messageMap map[string]*MessageConfig, config ParsedArgs) *Generator {
	generator := Generator{
		BaseGenerator: BaseGenerator{
			MessageConfigMap: messageMap,
			Config:           config,
			FieldReaderMap:   map[Modifier2FieldType]IFieldReader{},
			FieldWriterMap:   map[Modifier2FieldType]IFieldWriter{},
		},
	}

	generator.AddFieldWriter(NewArrEnumFieldWriter())
	generator.AddFieldWriter(NewArrMessageFieldWriter())
	generator.AddFieldWriter(NewArrPrimitiveFieldWriter())
	generator.AddFieldWriter(&DftEnumFieldWriter{})
	generator.AddFieldWriter(NewDftMapFieldWriter())
	generator.AddFieldWriter(&DftMessageFieldWriter{})
	generator.AddFieldWriter(&DftPrimitiveFieldWriter{})
	generator.AddFieldWriter(&ExtMessageFieldWriter{})
	generator.AddFieldWriter(NewLstEnumFieldWriter())
	generator.AddFieldWriter(NewLstMessageFieldWriter())
	generator.AddFieldWriter(NewLstPrimitiveFieldWriter())
	generator.AddFieldWriter(NewSetEnumFieldWriter())
	generator.AddFieldWriter(NewSetMessageFieldWriter())
	generator.AddFieldWriter(NewSetPrimitiveFieldWriter())

	generator.AddFieldReader(&ArrEnumFieldReader{})
	generator.AddFieldReader(&ArrMessageFieldReader{})
	generator.AddFieldReader(NewArrPrimitiveFieldReader())
	generator.AddFieldReader(&DftEnumFieldReader{})
	generator.AddFieldReader(NewDftMapFieldReader())
	generator.AddFieldReader(&DftMessageFieldReader{})
	generator.AddFieldReader(&DftPrimitiveFieldReader{})
	generator.AddFieldReader(&ExtMessageFieldReader{})
	generator.AddFieldReader(NewLstEnumFieldReader())
	generator.AddFieldReader(NewLstMessageFieldReader())
	generator.AddFieldReader(NewLstPrimitiveFieldReader())
	generator.AddFieldReader(NewSetEnumFieldReader())
	generator.AddFieldReader(NewSetMessageFieldReader())
	generator.AddFieldReader(NewSetPrimitiveFieldReader())
	return &generator
}

func (b *Generator) LanguageType() LanguageType {
	return TS
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
	if len(m.Note) > 0 {
		header.Add(I("//${0}", m.Note)).NewLine()
	}
	header.Add(I("export enum ${0} {", m.Name)).NewLine()
	for _, fieldConfig := range m.GetSortedFields() {
		header.Add(I("//${0} = ${1};", fieldConfig.FieldName, fieldConfig.FieldNum))
		if len(fieldConfig.Note) > 0 {
			header.Add(fieldConfig.Note)
		}
		header.NewLine()
		header.Add(I("${0},", fieldConfig.FieldName)).NewLine()
	}
	header.Add("}").NewLine()
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName(), ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())
}

func (generator *Generator) createMessageClass(m *MessageConfig) *FileContent {
	header := NewCodeBuilder()
	fields := NewCodeBuilder()

	if len(m.Note) > 0 {
		fields.Add(I("//${0}", m.Note)).NewLine()
	}
	parent := m.ExtMessage

	if parent != nil {
		extField := m.ExtField
		fields.Add("//").Add(extField.GetDefinition()).NewLine()
		fields.Add(I("export class ${0} extends ${1} {", m.Name, parent.Name)).NewLine(2)
		AddImportMessage(fields, parent.GetFullName())
	} else {
		fields.Add(I("export class ${0} {", m.Name)).NewLine(2)
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
						typeAndImport = NewTypeAndImport("Map<string,string>")
					} else if keyFieldType == STRING {
						mapType := I("Map<string,${0}>", valueFieldType.Value().LowerName())
						typeAndImport = NewTypeAndImport(mapType, I("{ ${0} } from \"protobj-ts\"", valueFieldType.Value().LowerName()))
					} else if valueFieldType == STRING {
						mapType := I("Map<${0},string>", keyFieldType.Value().LowerName())
						typeAndImport = NewTypeAndImport(mapType, I("{ ${0} } from \"protobj-ts\"", keyFieldType.Value().LowerName()))
					} else {
						mapType := I("Map<${0},${1}>", keyFieldType.Value().LowerName(), valueFieldType.Value().LowerName())
						typeAndImport = NewTypeAndImport(mapType, I("{ ${0},${1} } from \"protobj-ts\"", keyFieldType.Value().LowerName(), valueFieldType.Value().LowerName()))
					}
				} else {
					if keyFieldType == STRING {
						mapType := I("Map<string,${0}>", field.ValueTypeName)
						typeAndImport = NewTypeAndImport(mapType, field.ValueTypeFullName)
					} else {
						mapType := I("Map<${0},${1}>", keyFieldType.Value().LowerName(), field.ValueTypeName)
						typeAndImport = NewTypeAndImport(mapType, I("{ ${0} } from \"protobj-ts\"", keyFieldType.Value().LowerName()), field.ValueTypeFullName)
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
				typeAndImport = NewTypeAndImport(I("Set<${0}>", field.TypeName), field.TypeFullName)
			case LST:
				typeAndImport = NewTypeAndImport(I("Array<${0}>", field.TypeName), field.TypeFullName)
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
		createField(m, fields, field, typeAndImport)
	}
	AddImportMessages(header, fields.ImportMessages)
	appendImportMessages(m.Pkg, m.GetFullName(), header)
	header.AddBuilder(fields).Add("}").NewLine()
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName(), ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())
}

func getTypeAndImportFromBuiltinType(modifier Modifier, fieldType FieldType) *TypeAndImport {
	if fieldType == MAP {
		return nil
	}
	name := fieldType.Value().LowerName()
	switch modifier {
	case DFT:
		if fieldType == STRING {
			return NewTypeAndImport(name)
		}
		return NewTypeAndImport(name, fmt.Sprintf("{ %s } from \"protobj-ts\"", name))
	case LST:
		if fieldType == STRING {
			return NewTypeAndImport("Array<" + name + ">")
		}
		return NewTypeAndImport("Array<"+name+">", fmt.Sprintf("{ %s } from \"protobj-ts\"", name))
	case SET:
		if fieldType == STRING {
			return NewTypeAndImport("Set<" + name + ">")
		}
		return NewTypeAndImport("Set<"+name+">", fmt.Sprintf("{ %s } from \"protobj-ts\"", name))
	case ARR:
		if fieldType == STRING {
			return NewTypeAndImport(name + "[]")
		}
		return NewTypeAndImport(name+"[]", fmt.Sprintf("{ %s } from \"protobj-ts\"", name))
	}
	return nil
}
func createField(m *MessageConfig, fields *CodeBuilder, field *FieldConfig, typeAndImport *TypeAndImport) {
	if len(typeAndImport.Imports) > 0 {
		for _, importMessage := range typeAndImport.Imports {
			if importMessage == typeAndImport.Type {
				continue
			}
			AddImportMessage(fields, importMessage)
		}
	}
	fields.Add("//").Add(field.GetDefinition()).NewLine()
	fields.Add(I("${1}:${0}", typeAndImport.Type, field.FieldName)).NewLine(2)
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
	header := NewCodeBuilder()
	AddImportMessage(header, "{ Input } from \"protobj-ts\"")
	AddImportMessage(header, "{ Output } from \"protobj-ts\"")
	AddImportMessage(header, "{ Schema } from \"protobj-ts\"")
	AddImportMessage(header, m.GetFullName())
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
	appendImportMessages(m.Pkg, m.GetFullName()+"Schema", header)
	header.AddBuilder(body)
	suffix, _ := generator.LanguageType().FileSuffix()
	fileName := strings.ReplaceAll(m.GetFullName()+"Schema", ".", string(os.PathSeparator)) + "." + suffix
	return NewFileContent(fileName, header.String())
}

func (generator *Generator) createEnumWriteBody(m *MessageConfig, withFieldNum bool) *CodeBuilder {
	writeBody := NewCodeBuilder()
	writeBody.SetCurrent(2)
	writeBody.Add(isNull("message")).Add(LC).NewLine()
	writeBody.Add("output.writeI8_Packed(0);").NewLine()
	writeBody.Add("return;").NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add("switch (message) ").Add(LC).NewLine()
	for _, value := range m.GetSortedFields() {
		writeBody.Add(I("case ${0}.${1}: ", m.Name, value.FieldName)).Add(LC).NewLine()
		if withFieldNum {
			writeBody.Add(I("output.writeI32(fieldNum,${0});", value.FieldNum)).NewLine()
		} else {
			writeBody.Add(I("output.writeI8_Packed(${0});", value.FieldNum)).NewLine()
		}
		writeBody.Add("break;").NewLine()
		writeBody.Add(RC).NewLine()
	}
	writeBody.Add("default: ").Add(LC).NewLine()
	writeBody.Add("throw new Error(\"undefine enum \" + message);").NewLine()
	writeBody.Add(RC).NewLine()
	writeBody.Add(RC).NewLine()
	return writeBody
}

func (generator *Generator) createEnumReadBody(m *MessageConfig) *CodeBuilder {
	readBody := NewCodeBuilder()
	readBody.SetCurrent(2)
	readBody.Add("const value = input.readI32();").NewLine()
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

	AddImportMessage(header, "{ Input } from \"protobj-ts\"")
	AddImportMessage(header, "{ Output } from \"protobj-ts\"")
	AddImportMessage(header, "{ Schema } from \"protobj-ts\"")
	AddImportMessage(header, m.GetFullName())

	var writeBody = generator.createWriteBody(m)
	var readBody = generator.createReadBody(m)
	body := NewCodeBuilder()
	body.Add(N(MessageSchemaTemplate, map[string]interface{}{
		"class":        m.Name,
		"writeBody":    writeBody.String(),
		"readBody":     readBody.String(),
		"messageIndex": m.MessageIndex,
	})).NewLine()

	AddImportMessages(header, writeBody.ImportMessages)
	AddImportMessages(header, readBody.ImportMessages)
	appendImportMessages(p, m.GetFullName()+"Schema", header)
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
		getValue := I("message.${0}", field.FieldName)
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
		getValue := I("message.${0}", field.FieldName)
		setValue := fmt.Sprintf("message.%s=${value}", field.FieldName)
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
