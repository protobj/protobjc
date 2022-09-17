package ts

import "io.protobj/protobj-go/protobj"

type Generator struct {
	protobj.BaseGenerator
}

func NewGenerator(messageMap map[string]*protobj.MessageConfig, config protobj.ParsedArgs) *Generator {
	return &Generator{protobj.BaseGenerator{
		MessageConfigMap: messageMap,
		Config:           config,
	}}
}

func (b *Generator) LanguageType() protobj.LanguageType {
	return protobj.TS
}

func (b *Generator) Generate() {
	//TODO
}
