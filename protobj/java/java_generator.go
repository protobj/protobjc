package java

import (
	. "io.protobj/protobj-go/protobj"
	"sync"
)

type Generator struct {
	BaseGenerator
}

func NewGenerator(messageMap map[string]*MessageConfig, config ParsedArgs) *Generator {
	return &Generator{BaseGenerator: BaseGenerator{
		MessageConfigMap: messageMap,
		Config:           config,
	}}
}

func (b *Generator) LanguageType() LanguageType {
	return JAVA
}
func (b *Generator) Generate() {
	messageConfigMap := b.MessageConfigMap
	var waitGroup sync.WaitGroup
	var fileContentsChan = make(chan *FileContent)
	for _, messageConfig := range messageConfigMap {
		if b.Config.OutputType != O_SCHEMA {
			waitGroup.Add(1)
			go func(message *MessageConfig) {
				fileContent := b.createMessage(message)
				fileContentsChan <- fileContent
			}(messageConfig)
		}
		if b.Config.OutputType != O_MESSAGE {
			waitGroup.Add(1)
			go func(message *MessageConfig) {
				fileContent := b.createSchema(message)
				fileContentsChan <- fileContent
			}(messageConfig)
		}

	}
	waitGroup.Wait()
	close(fileContentsChan)
	waitGroup.Add(len(fileContentsChan))
	for {
		fileContent, ok := <-fileContentsChan
		if !ok {
			break
		}
		go func(content *FileContent) {
			WriteFile(b.Config.OutputDir, fileContent)
		}(fileContent)
	}
	waitGroup.Wait()
}

func (b *Generator) createMessage(m *MessageConfig) *FileContent {
	switch m.MessageType {
	case ENUM:
		return b.createEnumClass(m)
	case MESSAGE:
		return b.createMessageClass(m)
	default:
		return nil
	}
}

func (b *Generator) createEnumClass(m *MessageConfig) *FileContent {
	header := NewCodeBuilder()
	header.Add(pkg(m.Pkg)).NewLine(2)
	if len(m.Note) > 0 {
		header.Add(I("//${0}", m.Note)).NewLine()
	}

	return nil
}

func (b *Generator) createMessageClass(m *MessageConfig) *FileContent {
	return nil
}

func (b *Generator) createSchema(m *MessageConfig) *FileContent {

	return nil
}
