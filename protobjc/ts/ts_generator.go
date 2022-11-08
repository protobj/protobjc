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
	return &Generator{
		BaseGenerator: BaseGenerator{
			MessageConfigMap: messageMap,
			Config:           config,
			FieldReaderMap:   map[Modifier2FieldType]IFieldReader{},
			FieldWriterMap:   map[Modifier2FieldType]IFieldWriter{},
		},
	}
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
	appendImportMessages(m.Pkg, header)
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
	return nil
}

func (generator *Generator) createEnumWriteBody(m *MessageConfig, withFieldNum bool) *CodeBuilder {
	return nil
}

func (generator *Generator) createEnumReadBody(m *MessageConfig) *CodeBuilder {
	return nil
}

func (generator *Generator) createMessageSchema(m *MessageConfig) *FileContent {
	return nil

}

func (generator *Generator) createWriteBody(m *MessageConfig) *CodeBuilder {
	return nil
}

func (generator *Generator) createReadBody(m *MessageConfig) *CodeBuilder {
	return nil
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
