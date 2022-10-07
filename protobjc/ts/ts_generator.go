package ts

import . "io.protobj/protobjc"

type Generator struct {
	BaseGenerator
}

func NewGenerator(messageMap map[string]*MessageConfig, config ParsedArgs) *Generator {
	return &Generator{BaseGenerator{
		MessageConfigMap: messageMap,
		Config:           config,
	}}
}

func (b *Generator) LanguageType() LanguageType {
	return TS
}

func (b *Generator) Generate() {
	//TODO
}
