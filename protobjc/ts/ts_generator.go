package ts

import (
	. "io.protobj/protobjc"
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
	return nil
}

func (generator *Generator) createMessageClass(m *MessageConfig) *FileContent {
	return nil
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
