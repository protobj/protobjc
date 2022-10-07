package ts

import "io.protobj/protobjc/protobjc"

type Generator struct {
	protobjc.BaseGenerator
}

func NewGenerator(messageMap map[string]*protobjc.MessageConfig, config protobjc.ParsedArgs) *Generator {
	return &Generator{protobjc.BaseGenerator{
		MessageConfigMap: messageMap,
		Config:           config,
	}}
}

func (b *Generator) LanguageType() protobjc.LanguageType {
	return protobjc.TS
}

func (b *Generator) Generate() {
	//TODO
}
