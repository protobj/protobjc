package protobj

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"sort"
	"strconv"
	"strings"
)

type FieldType int

const (
	BOOL     FieldType = 0
	I8       FieldType = 1
	U8       FieldType = 2
	I16      FieldType = 3
	U16      FieldType = 4
	I32      FieldType = 5
	U32      FieldType = 6
	S32      FieldType = 7
	F32      FieldType = 8
	SF32     FieldType = 9
	I64      FieldType = 10
	U64      FieldType = 11
	S64      FieldType = 12
	F64      FieldType = 13
	SF64     FieldType = 14
	STRING   FieldType = 15
	DOUBLE   FieldType = 16
	FLOAT    FieldType = 17
	MAP      FieldType = 18
	FEnum    FieldType = 19
	FMessage FieldType = 20
)

func FieldTypeValueOf(fieldType string) (FieldType, error) {
	switch fieldType {
	case "bool":
		return BOOL, nil
	case "i8":
		return I8, nil
	case "u8":
		return U8, nil
	case "i16":
		return I16, nil
	case "u16":
		return U16, nil
	case "i32":
		return I32, nil
	case "u32":
		return U32, nil
	case "s32":
		return S32, nil
	case "f32":
		return F32, nil
	case "sf32":
		return SF32, nil
	case "i64":
		return I64, nil
	case "u64":
		return U64, nil
	case "s64":
		return S64, nil
	case "f64":
		return F64, nil
	case "sf64":
		return SF64, nil
	case "string":
		return STRING, nil
	case "double":
		return DOUBLE, nil
	case "float":
		return FLOAT, nil
	case "map":
		return MAP, nil
	}
	return 0, errors.New("error fieldType:" + fieldType)
}

func (receiver FieldType) Value() FieldTypeValue {
	return FieldTypeMap[int(receiver)]
}

type FieldTypeValue struct {
	FieldType FieldType
	JavaType  string
	GoType    string
}

var FieldTypeMap = map[int]FieldTypeValue{}

func init() {

	FieldTypeMap[0] = FieldTypeValue{BOOL, "boolean", "bool"}
	FieldTypeMap[1] = FieldTypeValue{I8, "byte", "int8"}
	FieldTypeMap[2] = FieldTypeValue{U8, "byte", "uint8"}
	FieldTypeMap[3] = FieldTypeValue{I16, "short", "int16"}
	FieldTypeMap[4] = FieldTypeValue{U16, "short", "uint16"}
	FieldTypeMap[5] = FieldTypeValue{I32, "int", "int32"}
	FieldTypeMap[6] = FieldTypeValue{U32, "int", "uint32"}
	FieldTypeMap[7] = FieldTypeValue{S32, "int", "int32"}
	FieldTypeMap[8] = FieldTypeValue{F32, "int", "int32"}
	FieldTypeMap[9] = FieldTypeValue{SF32, "int", "int32"}
	FieldTypeMap[10] = FieldTypeValue{I64, "long", "int64"}
	FieldTypeMap[11] = FieldTypeValue{U64, "long", "uint64"}
	FieldTypeMap[12] = FieldTypeValue{S64, "long", "int64"}
	FieldTypeMap[13] = FieldTypeValue{F64, "long", "int64"}
	FieldTypeMap[14] = FieldTypeValue{SF64, "long", "int64"}
	FieldTypeMap[15] = FieldTypeValue{STRING, "String", "string"}
	FieldTypeMap[16] = FieldTypeValue{DOUBLE, "double", "float64"}
	FieldTypeMap[17] = FieldTypeValue{FLOAT, "float", "float32"}
	FieldTypeMap[18] = FieldTypeValue{MAP, "null", "null"}
	FieldTypeMap[19] = FieldTypeValue{FEnum, "null", "null"}
	FieldTypeMap[20] = FieldTypeValue{FMessage, "null", "null"}
}

type MessageType int32

const (
	MESSAGE MessageType = 0
	ENUM    MessageType = 1
)

func (receiver MessageType) toFieldType() FieldType {
	if receiver == MESSAGE {
		return FMessage
	}
	return FEnum
}

type Modifier int32

const (
	DFT Modifier = 0
	LST Modifier = 1
	SET Modifier = 2
	ARR Modifier = 3
	EXT Modifier = 4
)

var ModifierName = map[int32]string{
	0: "DFT",
	1: "LST",
	2: "SET",
	3: "ARR",
	4: "EXT",
}
var ModifierValue = reverseMap(ModifierName)

func ModifierValueOf(str string) (Modifier, error) {
	value, ok := ModifierValue[strings.ToUpper(str)]

	if ok {
		return Modifier(value), nil
	}
	return 0, errors.New("error modifier " + str)
}

func (o FieldOption) parseValue(value string) (any, error) {
	switch o {
	case polymorphic:
		return strconv.ParseBool(value)
	case deprecated:
		return strconv.ParseBool(value)
	}
	return deprecated, errors.New(fmt.Sprintf("option: %s value:%s error", o, value))
}

type FieldOption int32

const (
	polymorphic FieldOption = 0
	deprecated  FieldOption = 1
)

var FieldOptionName = map[int32]string{
	0: "polymorphic",
	1: "deprecated",
}
var FieldOptionValue = reverseMap(FieldOptionName)

func (o FieldOption) String() string {
	return FieldOptionName[int32(o)]
}
func ParseOption(value string) (FieldOption, error) {
	i, ok := FieldOptionValue[value]
	if ok {
		return FieldOption(i), nil
	}
	return polymorphic, errors.New("unknown option " + value)
}

type MessageConfig struct {
	FileName              string
	Pkg                   string
	ImportMessages        map[string]Void
	MessageType           MessageType
	Name                  string
	Note                  string
	MessageIndex          int32
	FieldConfigMap        map[int32]*FieldConfig
	ChildMessageConfigMap map[int32]*MessageConfig
	ExtMessage            *MessageConfig
	ExtField              *FieldConfig
}

func NewMessageConfig(fileName string, importMessages map[string]Void, pkg string,
	messageType MessageType, messageName string) *MessageConfig {
	return &MessageConfig{
		FileName:              fileName,
		Pkg:                   pkg,
		ImportMessages:        importMessages,
		MessageType:           messageType,
		Name:                  messageName,
		FieldConfigMap:        map[int32]*FieldConfig{},
		ChildMessageConfigMap: map[int32]*MessageConfig{},
	}
}
func (m *MessageConfig) GetSortedFields() []*FieldConfig {
	var fieldList []*FieldConfig
	for _, config := range m.FieldConfigMap {
		fieldList = append(fieldList, config)
	}
	sort.Slice(fieldList, func(i, j int) bool {
		return fieldList[i].FieldNum-fieldList[j].FieldNum < 0
	})
	return fieldList
}

func (m *MessageConfig) GetFullName() string {
	return fmt.Sprintf("%s.%s", m.Pkg, m.Name)
}

func (m *MessageConfig) AddSelfToChildMap() {
	m.ChildMessageConfigMap[m.MessageIndex] = m
}

func (m *MessageConfig) setParent(extMessage *MessageConfig, extField *FieldConfig) error {
	if m.ExtMessage != nil {
		return errors.New(fmt.Sprintf("modifier:[ext]  only has one in message: %s field:%s", m.GetFullName(), extField.GetDefinition()))
	}
	m.ExtMessage = extMessage
	m.ExtField = extField
	if len(m.ChildMessageConfigMap) > 0 {
		for _, messageConfig := range m.ChildMessageConfigMap {
			err := m.ExtMessage.addChild(messageConfig)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MessageConfig) addChild(childMessage *MessageConfig) error {
	if m.MessageIndex < 0 {
		return errors.New(fmt.Sprintf("parent message index must > 0 :%s", m.GetFullName()))
	}
	if childMessage.MessageIndex <= m.MessageIndex {
		return errors.New(fmt.Sprintf("child index must gt parent index parent:%s:%d child:%s:%d", m.GetFullName(), m.MessageIndex, childMessage.GetFullName(), childMessage.MessageIndex))
	}
	old, ok := m.ChildMessageConfigMap[childMessage.MessageIndex]
	if ok {
		return errors.New(fmt.Sprintf("%s child has same index %s : %d", m.GetFullName(), old.GetFullName(), childMessage.MessageIndex))
	}
	if m.ExtMessage != nil {
		err := m.ExtMessage.addChild(childMessage)
		if err != nil {
			return err
		}
	}
	return nil
}

type FieldConfig struct {
	Modifier          Modifier
	TypeName          string
	TypeFullName      string
	KeyType           string
	ValueTypeName     string
	ValueTypeFullName string
	FieldName         string
	FieldNum          int32
	Options           map[FieldOption]any //FieldOptionValue
	Note              string
	Definition        string
}

func NewFieldConfig() *FieldConfig {
	return &FieldConfig{
		Options:  map[FieldOption]any{},
		Modifier: DFT,
	}
}
func (receiver *FieldConfig) GetDefinition() string {
	s := new(strings.Builder)
	if len(receiver.Note) != 0 {
		_, err := fmt.Fprintf(s, "%s %s", receiver.Definition, receiver.Note)
		if err != nil {
			panic(err)
		}
	} else {
		return receiver.Definition
	}
	return s.String()
}

func (receiver *FieldConfig) SetFieldNum(fieldNum string) error {
	value, err := strconv.ParseInt(fieldNum, 10, 32)
	if err != nil {
		return err
	}
	receiver.FieldNum = int32(value)
	return nil
}

func (receiver *FieldConfig) InitDefinition(context antlr.ParserRuleContext) {
	if context.GetChildCount() == 0 {
		receiver.Definition = ""
		return
	}
	s := new(strings.Builder)
	for i := 0; i < context.GetChildCount(); i++ {
		child := context.GetChild(i)
		s.WriteString(child.(antlr.ParseTree).GetText())
		s.WriteString(" ")
	}
	receiver.Definition = s.String()
	return
}

func (receiver *FieldConfig) MessageName() string {
	if receiver.TypeName == "map" {
		return receiver.ValueTypeName
	}
	return receiver.TypeName
}
func (receiver *FieldConfig) FullMessageName() string {
	if receiver.TypeName == "map" {
		return receiver.ValueTypeFullName
	}
	return receiver.TypeFullName
}

func (receiver *FieldConfig) IsPolymorphic() bool {
	_, ok := receiver.Options[polymorphic]
	return ok
}
func (receiver *FieldConfig) FirstUpperFieldName() string {
	return FirstUpper(receiver.FieldName)
}

func reverseMap(nameMap map[int32]string) map[string]int32 {
	var result = map[string]int32{}
	for k, v := range nameMap {
		result[v] = k
	}
	return result
}
