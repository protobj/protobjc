package protobj

import (
	"errors"
	"fmt"
	"strings"
)

type IGenerator interface {
	FindMessage(source *MessageConfig, messageFullName string) (*MessageConfig, error)
	GetFieldType(sourceMessage *MessageConfig, typeName string, typeFullName string) (FieldType, error)
	LanguageType() LanguageType
	Generate()
	GetReader(modifier2FieldType Modifier2FieldType) IFieldReader
	GetWriter(modifier2FieldType Modifier2FieldType) IFieldWriter
}

type BaseGenerator struct {
	MessageConfigMap map[string]*MessageConfig
	Config           ParsedArgs
}

func (b *BaseGenerator) FindMessage(source *MessageConfig, messageFullName string) (*MessageConfig, error) {
	if len(messageFullName) == 0 {
		return nil, errors.New(fmt.Sprintf("message not found:%s in %s:%s", "nil", source.FileName, source.Name))
	}
	messageConfig, ok := b.MessageConfigMap[messageFullName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("message not found:%s in %s:%s", messageFullName, source.FileName, source.Name))
	}
	return messageConfig, nil
}
func (b *BaseGenerator) GetFieldType(sourceMessage *MessageConfig, typeName string, typeFullName string) (FieldType, error) {
	fieldType, err := FieldTypeValueOf(typeName)
	if err == nil {
		return fieldType, nil
	}
	message, err := b.FindMessage(sourceMessage, typeFullName)
	if err != nil {
		return 0, err
	}
	return message.MessageType.ToFieldType(), nil
}
func (b *BaseGenerator) LanguageType() LanguageType {
	panic("unsupported func")
}
func (b *BaseGenerator) Generate() {
	panic("unsupported func")
}

var intent = " "
var tab = strings.Repeat(intent, 4)

type CodeBuilder struct {
	current        int
	builder        strings.Builder
	lineBuilder    strings.Builder
	ImportMessages map[string]Void
	lastChar       byte
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		current:        0,
		ImportMessages: map[string]Void{},
	}
}
func (b *CodeBuilder) SetCurrent(current int) {
	b.current = current
}
func (b *CodeBuilder) Add(value string) *CodeBuilder {
	b.lineBuilder.WriteString(value)
	return b
}
func (b *CodeBuilder) append0(value string) {
	if b.current > 0 {
		b.builder.WriteString(strings.Repeat(tab, b.current))
	}
	b.builder.WriteString(value)
}

func (b *CodeBuilder) AddBuilder(builder *CodeBuilder) *CodeBuilder {
	b.builder.WriteString(builder.builder.String())
	return b
}

func (b *CodeBuilder) String() string {
	return b.builder.String()
}

func (b *CodeBuilder) AddImportMessage(importMessage string) {
	if len(importMessage) > 0 {
		b.ImportMessages[importMessage] = Empty
	}
}
func (b *CodeBuilder) AddImportMessages(importMessages map[string]Void) {
	for k, _ := range importMessages {
		b.AddImportMessage(k)
	}
}

func (b *CodeBuilder) NewLine(count ...int) {
	if len(count) == 0 {
		count = []int{1}
	}
	for range count {
		b.lineBuilder.WriteString(strings.Repeat("\n", count[0]))
	}
	s := b.lineBuilder.String()
	b.lineBuilder.Reset()
	b.addLine(s)
}

func (b *CodeBuilder) addLine(value string) {
	trim := strings.Trim(strings.ReplaceAll(value, "\n", ""), " ")
	if strings.HasPrefix(trim, "}") && strings.HasSuffix(trim, "{") {
		b.current--
		if b.current < 0 {
			b.current = 0
		}
		b.append0(value)
		b.current++
	} else if strings.HasSuffix(trim, "{") {
		b.append0(value)
		b.current++
	} else if strings.HasSuffix(trim, "}") || strings.HasSuffix(trim, "});") || strings.HasPrefix(trim, "} while") {
		b.current--
		if b.current < 0 {
			b.current = 0
		}
		b.append0(value)
	} else if strings.HasSuffix(trim, ":") {
		b.append0(value)
		b.current++
		b.lastChar = ':'
	} else if b.lastChar == ':' {
		b.append0(value)
		b.current--
		if b.current < 0 {
			b.current = 0
		}
		b.lastChar = 'a'
	} else {
		b.append0(value)
	}

}

type IFieldWriter interface {
	Modifier() Modifier
	FocusTypes() map[FieldType]Void
	Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string)
	WritePacked(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType, value string)
}

type FieldWriter struct {
}

func (f *FieldWriter) FocusTypes() map[FieldType]Void {
	panic("UnsupportedOperation")
}

func (f *FieldWriter) Modifier() Modifier {
	panic("UnsupportedOperation")
}
func (f *FieldWriter) Write(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, value string) {
	panic("UnsupportedOperation")
}
func (f *FieldWriter) WritePacked(generator IGenerator, writeBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType, value string) {
	panic("UnsupportedOperation")
}

type IFieldReader interface {
	Modifier() Modifier
	FocusTypes() map[FieldType]Void
	Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string)
	ReadPacked(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType) string
}

type FieldReader struct {
}

func (f *FieldReader) FocusTypes() map[FieldType]Void {
	panic("UnsupportedOperation")
}

func (f *FieldReader) Modifier() Modifier {
	panic("UnsupportedOperation")
}

func (f *FieldReader) Read(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, getValue, setValue string) {
	panic("UnsupportedOperation")
}

func (f *FieldReader) ReadPacked(generator IGenerator, readBody *CodeBuilder, sourceMessage *MessageConfig, fieldConfig *FieldConfig, fieldType FieldType) string {
	panic("UnsupportedOperation")
}

type MapKeyValueFieldType struct {
	KeyType   FieldType
	ValueType FieldType
}

func NewMapKeyValueFieldType(keyType, valueType FieldType) MapKeyValueFieldType {
	return MapKeyValueFieldType{KeyType: keyType, ValueType: valueType}
}

type Modifier2FieldType struct {
	Modifier  Modifier
	FieldType FieldType
}

func NewModifier2FieldType(modifier Modifier, fieldType FieldType) Modifier2FieldType {
	return Modifier2FieldType{
		Modifier:  modifier,
		FieldType: fieldType,
	}
}
