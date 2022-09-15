// Code generated from D:/code/protobj/protobj-java/src/main/resources\Protobj.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr // Protobj
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProtobjListener is a complete listener for a parse tree produced by ProtobjParser.
type BaseProtobjListener struct{}

var _ ProtobjListener = &BaseProtobjListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProtobjListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProtobjListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProtobjListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProtobjListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProtobj is called when production protobj is entered.
func (s *BaseProtobjListener) EnterProtobj(ctx *ProtobjContext) {}

// ExitProtobj is called when production protobj is exited.
func (s *BaseProtobjListener) ExitProtobj(ctx *ProtobjContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BaseProtobjListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BaseProtobjListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseProtobjListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseProtobjListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BaseProtobjListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BaseProtobjListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterField is called when production field is entered.
func (s *BaseProtobjListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseProtobjListener) ExitField(ctx *FieldContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseProtobjListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseProtobjListener) ExitModifier(ctx *ModifierContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseProtobjListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseProtobjListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *BaseProtobjListener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *BaseProtobjListener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterFieldNumber is called when production fieldNumber is entered.
func (s *BaseProtobjListener) EnterFieldNumber(ctx *FieldNumberContext) {}

// ExitFieldNumber is called when production fieldNumber is exited.
func (s *BaseProtobjListener) ExitFieldNumber(ctx *FieldNumberContext) {}

// EnterExtendsField is called when production extendsField is entered.
func (s *BaseProtobjListener) EnterExtendsField(ctx *ExtendsFieldContext) {}

// ExitExtendsField is called when production extendsField is exited.
func (s *BaseProtobjListener) ExitExtendsField(ctx *ExtendsFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BaseProtobjListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BaseProtobjListener) ExitMapField(ctx *MapFieldContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseProtobjListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseProtobjListener) ExitMapType(ctx *MapTypeContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseProtobjListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseProtobjListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseProtobjListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseProtobjListener) ExitType_(ctx *Type_Context) {}

// EnterTopLevelDef is called when production topLevelDef is entered.
func (s *BaseProtobjListener) EnterTopLevelDef(ctx *TopLevelDefContext) {}

// ExitTopLevelDef is called when production topLevelDef is exited.
func (s *BaseProtobjListener) ExitTopLevelDef(ctx *TopLevelDefContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseProtobjListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseProtobjListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseProtobjListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseProtobjListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumElement is called when production enumElement is entered.
func (s *BaseProtobjListener) EnterEnumElement(ctx *EnumElementContext) {}

// ExitEnumElement is called when production enumElement is exited.
func (s *BaseProtobjListener) ExitEnumElement(ctx *EnumElementContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseProtobjListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseProtobjListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterMessageDef is called when production messageDef is entered.
func (s *BaseProtobjListener) EnterMessageDef(ctx *MessageDefContext) {}

// ExitMessageDef is called when production messageDef is exited.
func (s *BaseProtobjListener) ExitMessageDef(ctx *MessageDefContext) {}

// EnterMessageIndex is called when production messageIndex is entered.
func (s *BaseProtobjListener) EnterMessageIndex(ctx *MessageIndexContext) {}

// ExitMessageIndex is called when production messageIndex is exited.
func (s *BaseProtobjListener) ExitMessageIndex(ctx *MessageIndexContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *BaseProtobjListener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *BaseProtobjListener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterMessageElement is called when production messageElement is entered.
func (s *BaseProtobjListener) EnterMessageElement(ctx *MessageElementContext) {}

// ExitMessageElement is called when production messageElement is exited.
func (s *BaseProtobjListener) ExitMessageElement(ctx *MessageElementContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseProtobjListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseProtobjListener) ExitConstant(ctx *ConstantContext) {}

// EnterBlockLit is called when production blockLit is entered.
func (s *BaseProtobjListener) EnterBlockLit(ctx *BlockLitContext) {}

// ExitBlockLit is called when production blockLit is exited.
func (s *BaseProtobjListener) ExitBlockLit(ctx *BlockLitContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseProtobjListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseProtobjListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterIdent is called when production ident is entered.
func (s *BaseProtobjListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseProtobjListener) ExitIdent(ctx *IdentContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseProtobjListener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseProtobjListener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *BaseProtobjListener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *BaseProtobjListener) ExitMessageName(ctx *MessageNameContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseProtobjListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseProtobjListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseProtobjListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseProtobjListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterMapName is called when production mapName is entered.
func (s *BaseProtobjListener) EnterMapName(ctx *MapNameContext) {}

// ExitMapName is called when production mapName is exited.
func (s *BaseProtobjListener) ExitMapName(ctx *MapNameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BaseProtobjListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BaseProtobjListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseProtobjListener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseProtobjListener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterIntLit is called when production intLit is entered.
func (s *BaseProtobjListener) EnterIntLit(ctx *IntLitContext) {}

// ExitIntLit is called when production intLit is exited.
func (s *BaseProtobjListener) ExitIntLit(ctx *IntLitContext) {}

// EnterStrLit is called when production strLit is entered.
func (s *BaseProtobjListener) EnterStrLit(ctx *StrLitContext) {}

// ExitStrLit is called when production strLit is exited.
func (s *BaseProtobjListener) ExitStrLit(ctx *StrLitContext) {}

// EnterBoolLit is called when production boolLit is entered.
func (s *BaseProtobjListener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production boolLit is exited.
func (s *BaseProtobjListener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterFloatLit is called when production floatLit is entered.
func (s *BaseProtobjListener) EnterFloatLit(ctx *FloatLitContext) {}

// ExitFloatLit is called when production floatLit is exited.
func (s *BaseProtobjListener) ExitFloatLit(ctx *FloatLitContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseProtobjListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseProtobjListener) ExitKeywords(ctx *KeywordsContext) {}
