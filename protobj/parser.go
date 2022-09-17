package protobj

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	selfAntlr "io.protobj/protobj-go/antlr"
	"strconv"
	"strings"
	"sync"
)

func Load(fileList []string) map[string]*MessageConfig {
	var lock sync.Mutex
	m := map[string]*MessageConfig{}
	if len(fileList) == 0 {
		return m
	}
	var countDown sync.WaitGroup
	countDown.Add(len(fileList))
	for _, fileName := range fileList {
		go func(file string) {
			stream, err := antlr.NewFileStream(file)
			if nil != err {
				panic(fmt.Sprintf("read file err : %s", file))
			}
			protobjLexer := selfAntlr.NewProtobjLexer(stream)
			tokenStream := antlr.NewCommonTokenStream(protobjLexer, 0)
			protobjParser := selfAntlr.NewProtobjParser(tokenStream)
			protobjFileReader := newProtobjFileReader(file, tokenStream)
			protobjParser.AddErrorListener(&ProtobjErrorListener{fileName: file})
			walker := antlr.NewParseTreeWalker()
			walker.Walk(protobjFileReader, protobjParser.Protobj())
			for name, config := range protobjFileReader.messageConfigMap {
				putMessageConfig(&lock, m, name, config)
			}
			countDown.Done()
		}(fileName)
	}
	countDown.Wait()

	for _, messageConfig := range m {
		if messageConfig.MessageType == MESSAGE {
			countDown.Add(1)
			go func(message *MessageConfig) {
				check(m, message)
				afterCheck(message)
				countDown.Done()
			}(messageConfig)
		}
	}
	countDown.Wait()
	return m
}

func check(messageConfigMap map[string]*MessageConfig, messageConfig *MessageConfig) {
	//import message check
	for k, _ := range messageConfig.ImportMessages {
		_, ok := messageConfigMap[k]
		if !ok {
			PrintErrorAndExit(fmt.Sprintf("import message not exists：%s in file:%s", k, messageConfig.FileName))
		}
	}
	for _, fieldConfig := range messageConfig.FieldConfigMap {
		typeName := fieldConfig.TypeName
		var messageFullName string
		_, err := FieldTypeValueOf(typeName)
		if err == nil {
			if typeName == "map" {
				messageFullName = checkMessageFieldExists(messageConfigMap, messageConfig, fieldConfig, fieldConfig.ValueTypeName)
				if len(messageFullName) > 0 {
					fieldConfig.ValueTypeFullName = messageFullName
				}
			}
		} else {
			messageFullName = checkMessageFieldExists(messageConfigMap, messageConfig, fieldConfig, fieldConfig.TypeName)
			if len(messageFullName) > 0 {
				fieldConfig.TypeFullName = messageFullName
			}
		}

		if fieldConfig.IsPolymorphic() {
			if len(messageFullName) == 0 {
				PrintErrorAndExit(fmt.Sprintf("field cant add option polymorphic %s in %s", fieldConfig.FieldName, messageConfig.GetFullName()))
			}
			message := messageConfigMap[messageFullName]
			if message.MessageIndex < 0 {
				PrintErrorAndExit(fmt.Sprintf("field cant add option polymorphic %s in %s", fieldConfig.FieldName, messageConfig.GetFullName()))
			}
		}
		if fieldConfig.Modifier == EXT {
			message := messageConfigMap[messageFullName]
			if message == messageConfig {
				PrintErrorAndExit(fmt.Sprintf("field is not parent %s %s", fieldConfig.FieldName, messageConfig.GetFullName()))
			}
			err := messageConfig.setParent(message, fieldConfig)
			if err != nil {
				PrintErrorAndExit(err.Error())
			}
			err = message.addChild(messageConfig)
			if err != nil {
				PrintErrorAndExit(err.Error())
			}
		}
	}
}

func checkMessageFieldExists(configMap map[string]*MessageConfig, messageConfig *MessageConfig, fieldConfig *FieldConfig, typeName string) string {
	_, err := FieldTypeValueOf(typeName)
	if err != nil {
		fullName := messageConfig.Pkg + "." + typeName
		_, exists := configMap[fullName]
		if exists {
			return fullName
		}
		for key, _ := range messageConfig.ImportMessages {
			if strings.HasSuffix(key, "."+typeName) {
				return key
			}
		}
		_, exists = configMap[typeName]
		if exists {
			return typeName
		}
		err := fmt.Sprintf("file:%s message not exists in %s : %s ", messageConfig.FileName, messageConfig.GetFullName(), fieldConfig.FieldName)
		PrintErrorAndExit(err)
	}
	return ""
}

func afterCheck(messageConfig *MessageConfig) {
	if messageConfig.MessageIndex > -1 {
		messageConfig.AddSelfToChildMap()
	}
}

type ProtobjErrorListener struct {
	antlr.DefaultErrorListener
	fileName string
}

func (c *ProtobjErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("%s line %d:%d %s", c.fileName, line, column, msg))
}

func putMessageConfig(lock *sync.Mutex, messageMap map[string]*MessageConfig,
	name string, message *MessageConfig) {
	lock.Lock()
	defer lock.Unlock()
	old, ok := messageMap[name]
	if ok {
		PrintErrorAndExit(fmt.Sprintf("%s and %s duplicate message def :%s", old.FileName, message.FileName, message.Name))
	}
	messageMap[name] = message
}

type protobjFileReader struct {
	*selfAntlr.BaseProtobjListener
	fileName          string
	pkg               string
	importMessages    map[string]Void
	messageConfigMap  map[string]*MessageConfig
	commonTokenStream *antlr.CommonTokenStream
	exitsExtendField  *FieldConfig
}

func newProtobjFileReader(fileName string, tokenStream *antlr.CommonTokenStream) *protobjFileReader {
	return &protobjFileReader{
		fileName:          fileName,
		commonTokenStream: tokenStream,
		messageConfigMap:  map[string]*MessageConfig{},
		importMessages:    map[string]Void{},
	}
}
func (s *protobjFileReader) EnterPackageStatement(ctx *selfAntlr.PackageStatementContext) {
	s.pkg = ctx.GetChild(1).(antlr.ParseTree).GetText()
}

func (s *protobjFileReader) EnterImportStatement(ctx *selfAntlr.ImportStatementContext) {
	importMessage := ctx.GetChild(1).(antlr.ParseTree).GetText()
	s.importMessages[importMessage] = Empty
}
func (s *protobjFileReader) EnterMessageDef(ctx *selfAntlr.MessageDefContext) {
	messageName := ctx.MessageName().GetText()
	s.checkMessageName(ctx.MessageName().GetStart(), messageName)
	messageConfig := NewMessageConfig(s.fileName, s.importMessages, s.pkg, MESSAGE, messageName)
	messageConfig.Note = s.getLeftNote(ctx.GetStart())
	messageIndex := ctx.MessageIndex()
	if messageIndex != nil {
		index := messageIndex.GetText()
		intIndex, err := strconv.Atoi(index)
		if err != nil {
			s.printErrorAndExit(messageIndex.GetStart(),
				fmt.Sprintf("message index not number :%s ", messageConfig.GetFullName()))
		}
		if intIndex < 0 {
			s.printErrorAndExit(messageIndex.GetStart(),
				fmt.Sprintf("message index must >= 0 :%s", messageConfig.GetFullName()))
		}
		messageConfig.MessageIndex = int32(intIndex)
	}
	context := (ctx.MessageBody()).(*selfAntlr.MessageBodyContext)
	for _, e := range context.AllMessageElement() {
		var elementContext = e.(*selfAntlr.MessageElementContext)
		var fieldConfig *FieldConfig
		if elementContext.Field() != nil {
			fieldConfig = s.parseField(messageConfig, elementContext.Field())
		}
		if elementContext.MapField() != nil {
			fieldConfig = s.parseMapField(messageConfig, elementContext.MapField())
		}
		if elementContext.ExtendsField() != nil {
			fieldConfig = s.parseExtendField(messageConfig, elementContext.ExtendsField())
		}
		if fieldConfig == nil {
			continue
		}
		s.putField(messageConfig, fieldConfig, elementContext.GetStart())
	}
	s.putMessage(messageConfig, ctx.GetStart())
	s.exitsExtendField = nil
}

func (s *protobjFileReader) EnterEnumDef(ctx *selfAntlr.EnumDefContext) {
	enumName := ctx.EnumName().GetText()
	s.checkMessageName(ctx.EnumName().GetStart(), enumName)
	messageConfig := NewMessageConfig(s.fileName, s.importMessages, s.pkg, ENUM, enumName)
	messageConfig.Note = s.getLeftNote(ctx.GetStart())
	enumBodyContext := ctx.EnumBody().(*selfAntlr.EnumBodyContext)
	for _, iEnumElementContext := range enumBodyContext.AllEnumElement() {
		enumFieldContext := iEnumElementContext.(*selfAntlr.EnumElementContext).EnumField().(*selfAntlr.EnumFieldContext)
		fieldConfig := NewFieldConfig()
		s.setFieldNum(enumFieldContext.Ident().GetStart(), enumFieldContext.IntLit().GetText(), messageConfig, fieldConfig)
		fieldConfig.Note = s.getRightNote(enumFieldContext.SEMI().GetSymbol())
		fieldConfig.FieldName = enumFieldContext.Ident().GetText()
		s.putField(messageConfig, fieldConfig, enumFieldContext.GetStart())
	}
	s.putMessage(messageConfig, ctx.GetStart())
}

func (s *protobjFileReader) checkMessageName(token antlr.Token, messageName string) {
	_, err := FieldTypeValueOf(messageName)
	if err == nil {
		s.printErrorAndExit(token, "message name is builtin type")
	}
}
func (s *protobjFileReader) printErrorAndExit(token antlr.Token, err string) {
	line := token.GetLine()
	column := token.GetColumn()
	PrintErrorAndExit(fmt.Sprintf("%s line[%d:%d] %s \n", s.fileName, line, column, err))
}

func (s *protobjFileReader) getLeftNote(token antlr.Token) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...", r)
		}
	}()
	tokenIndex := token.GetTokenIndex()
	left := s.commonTokenStream.GetHiddenTokensToLeft(tokenIndex, 1)
	var note string
	if left != nil && len(left) > 0 {
		note += left[0].GetText()[2:]
	}
	return note
}

func (s *protobjFileReader) getRightNote(token antlr.Token) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...", r)
		}
	}()
	tokenIndex := token.GetTokenIndex()
	left := s.commonTokenStream.GetHiddenTokensToRight(tokenIndex, 1)
	var note string
	if left != nil && len(left) > 0 {
		note += left[0].GetText()[2:]
	}
	return note
}

func (s *protobjFileReader) parseField(messageConfig *MessageConfig, field selfAntlr.IFieldContext) *FieldConfig {
	var fieldConfig = NewFieldConfig()
	fieldContext := field.(*selfAntlr.FieldContext)
	deprecated := s.parseOption(fieldConfig, fieldContext.FieldOptions())
	if deprecated {
		return nil
	}
	fieldConfig.Modifier = DFT
	if fieldContext.Modifier() != nil {
		modifierStr := fieldContext.Modifier().GetText()
		if len(modifierStr) > 0 {
			modifier, err := ModifierValueOf(modifierStr)
			if err != nil {
				s.printErrorAndExit(fieldContext.Modifier().GetStart(), err.Error())
			}
			fieldConfig.Modifier = modifier
		}
	}
	fieldConfig.Note = s.getRightNote(fieldContext.SEMI().GetSymbol())
	fieldConfig.InitDefinition(field)
	//fieldConfig
	fieldConfig.TypeName = fieldContext.Type_().GetText()
	fieldConfig.FieldName = fieldContext.FieldName().GetText()
	fieldNumberStr := fieldContext.FieldNumber().GetText()
	s.setFieldNum(fieldContext.FieldNumber().GetStart(), fieldNumberStr, messageConfig, fieldConfig)
	return fieldConfig
}

func (s *protobjFileReader) parseOption(config *FieldConfig, options selfAntlr.IFieldOptionsContext) bool {
	if options == nil {
		return false
	}
	optionsContext := options.(*selfAntlr.FieldOptionsContext)
	for _, context := range optionsContext.AllFieldOption() {
		optionContext := context.(*selfAntlr.FieldOptionContext)
		optionName := optionContext.OptionName().GetText()
		optionValue := optionContext.Constant().GetText()
		option, err := ParseOption(optionName)
		if err != nil {
			s.printErrorAndExit(optionContext.OptionName().GetStart(), err.Error())
		}
		value, err := option.parseValue(optionValue)
		if err != nil {
			s.printErrorAndExit(optionContext.Constant().GetStart(), err.Error())
		}
		if option == deprecated && value.(bool) {
			return true
		}
		if option == polymorphic && !value.(bool) {
			continue
		}
		config.Options[option] = value
	}
	return false
}

func (s *protobjFileReader) setFieldNum(token antlr.Token, fieldNumberStr string, messageConfig *MessageConfig, fieldConfig *FieldConfig) {
	err := fieldConfig.SetFieldNum(fieldNumberStr)

	if err != nil {
		s.printErrorAndExit(token, fmt.Sprintf("field num not number %s", messageConfig.GetFullName()))
	}
	var lower int32
	if messageConfig.MessageType == MESSAGE {
		lower = 0
	} else {
		lower = -1
	}
	if fieldConfig.FieldNum <= lower {
		s.printErrorAndExit(token, fmt.Sprintf("field num must > %d in %s", lower, messageConfig.GetFullName()))
	}
}

func (s *protobjFileReader) putField(messageConfig *MessageConfig, fieldConfig *FieldConfig, token antlr.Token) {
	for _, fieldCfg := range messageConfig.FieldConfigMap {
		if fieldCfg.FieldName == fieldConfig.FieldName {
			s.printErrorAndExit(token, fmt.Sprintf("field name duplicate :%s in %s", fieldConfig.FieldName, messageConfig.GetFullName()))
		}
	}
	_, ok := messageConfig.FieldConfigMap[fieldConfig.FieldNum]
	if ok {
		s.printErrorAndExit(token, fmt.Sprintf("field num duplicate: %d in %s", fieldConfig.FieldNum, messageConfig.GetFullName()))
	}
	messageConfig.FieldConfigMap[fieldConfig.FieldNum] = fieldConfig
}

func (s *protobjFileReader) putMessage(messageConfig *MessageConfig, token antlr.Token) {
	old, ok := s.messageConfigMap[messageConfig.GetFullName()]
	if ok {
		s.printErrorAndExit(token, fmt.Sprintf("message def duplicate %s", old.GetFullName()))
	}
	s.messageConfigMap[messageConfig.GetFullName()] = messageConfig
}

func (s *protobjFileReader) parseMapField(messageConfig *MessageConfig, mapField selfAntlr.IMapFieldContext) *FieldConfig {
	fieldConfig := NewFieldConfig()
	mapFieldContext := mapField.(*selfAntlr.MapFieldContext)
	deprecated := s.parseOption(fieldConfig, mapFieldContext.FieldOptions())
	if deprecated {
		return nil
	}
	fieldConfig.Note = s.getRightNote(mapFieldContext.SEMI().GetSymbol())
	fieldConfig.InitDefinition(mapFieldContext)
	mapTypeContext := mapFieldContext.MapType().(*selfAntlr.MapTypeContext)
	keyType := mapTypeContext.KeyType().GetText()
	fieldConfig.TypeName = "map"
	fieldConfig.KeyType = keyType
	fieldConfig.ValueTypeName = mapTypeContext.Type_().GetText()
	fieldConfig.FieldName = mapFieldContext.MapName().GetText()
	s.setFieldNum(mapField.GetStart(), mapFieldContext.FieldNumber().GetText(), messageConfig, fieldConfig)
	return fieldConfig
}

func (s *protobjFileReader) parseExtendField(messageConfig *MessageConfig, extendsField selfAntlr.IExtendsFieldContext) *FieldConfig {
	extendsFieldContext := extendsField.(*selfAntlr.ExtendsFieldContext)
	if s.exitsExtendField != nil {
		s.printErrorAndExit(extendsField.GetStart(), fmt.Sprintf("[%s] exists many extend field[%s,%s]",
			messageConfig.GetFullName(), s.exitsExtendField.FieldName, extendsFieldContext.FieldName().GetText()))
		return nil
	}
	fieldConfig := NewFieldConfig()
	deprecated := s.parseOption(fieldConfig, extendsFieldContext.FieldOptions())
	if deprecated {
		return nil
	}
	fieldConfig.Note = s.getRightNote(extendsFieldContext.SEMI().GetSymbol())
	fieldConfig.InitDefinition(extendsFieldContext)
	s.exitsExtendField = fieldConfig
	fieldConfig.Modifier = EXT
	fieldConfig.TypeName = extendsFieldContext.MessageType().GetText()
	fieldConfig.FieldName = extendsFieldContext.FieldName().GetText()
	s.setFieldNum(extendsFieldContext.GetStart(), extendsFieldContext.FieldNumber().GetText(), messageConfig, fieldConfig)
	return fieldConfig
}
