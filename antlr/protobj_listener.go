// Code generated from D:/code/protobj/protobj-java/src/main/resources\Protobj.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr // Protobj
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProtobjListener is a complete listener for a parse tree produced by ProtobjParser.
type ProtobjListener interface {
	antlr.ParseTreeListener

	// EnterProtobj is called when entering the protobj production.
	EnterProtobj(c *ProtobjContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterOptionName is called when entering the optionName production.
	EnterOptionName(c *OptionNameContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterFieldOption is called when entering the fieldOption production.
	EnterFieldOption(c *FieldOptionContext)

	// EnterFieldNumber is called when entering the fieldNumber production.
	EnterFieldNumber(c *FieldNumberContext)

	// EnterExtendsField is called when entering the extendsField production.
	EnterExtendsField(c *ExtendsFieldContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTopLevelDef is called when entering the topLevelDef production.
	EnterTopLevelDef(c *TopLevelDefContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumElement is called when entering the enumElement production.
	EnterEnumElement(c *EnumElementContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterMessageDef is called when entering the messageDef production.
	EnterMessageDef(c *MessageDefContext)

	// EnterMessageIndex is called when entering the messageIndex production.
	EnterMessageIndex(c *MessageIndexContext)

	// EnterMessageBody is called when entering the messageBody production.
	EnterMessageBody(c *MessageBodyContext)

	// EnterMessageElement is called when entering the messageElement production.
	EnterMessageElement(c *MessageElementContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterBlockLit is called when entering the blockLit production.
	EnterBlockLit(c *BlockLitContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterMessageName is called when entering the messageName production.
	EnterMessageName(c *MessageNameContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterMapName is called when entering the mapName production.
	EnterMapName(c *MapNameContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterIntLit is called when entering the intLit production.
	EnterIntLit(c *IntLitContext)

	// EnterStrLit is called when entering the strLit production.
	EnterStrLit(c *StrLitContext)

	// EnterBoolLit is called when entering the boolLit production.
	EnterBoolLit(c *BoolLitContext)

	// EnterFloatLit is called when entering the floatLit production.
	EnterFloatLit(c *FloatLitContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// ExitProtobj is called when exiting the protobj production.
	ExitProtobj(c *ProtobjContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitOptionName is called when exiting the optionName production.
	ExitOptionName(c *OptionNameContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitFieldOption is called when exiting the fieldOption production.
	ExitFieldOption(c *FieldOptionContext)

	// ExitFieldNumber is called when exiting the fieldNumber production.
	ExitFieldNumber(c *FieldNumberContext)

	// ExitExtendsField is called when exiting the extendsField production.
	ExitExtendsField(c *ExtendsFieldContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTopLevelDef is called when exiting the topLevelDef production.
	ExitTopLevelDef(c *TopLevelDefContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumElement is called when exiting the enumElement production.
	ExitEnumElement(c *EnumElementContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitMessageDef is called when exiting the messageDef production.
	ExitMessageDef(c *MessageDefContext)

	// ExitMessageIndex is called when exiting the messageIndex production.
	ExitMessageIndex(c *MessageIndexContext)

	// ExitMessageBody is called when exiting the messageBody production.
	ExitMessageBody(c *MessageBodyContext)

	// ExitMessageElement is called when exiting the messageElement production.
	ExitMessageElement(c *MessageElementContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitBlockLit is called when exiting the blockLit production.
	ExitBlockLit(c *BlockLitContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitMessageName is called when exiting the messageName production.
	ExitMessageName(c *MessageNameContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitMapName is called when exiting the mapName production.
	ExitMapName(c *MapNameContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitIntLit is called when exiting the intLit production.
	ExitIntLit(c *IntLitContext)

	// ExitStrLit is called when exiting the strLit production.
	ExitStrLit(c *StrLitContext)

	// ExitBoolLit is called when exiting the boolLit production.
	ExitBoolLit(c *BoolLitContext)

	// ExitFloatLit is called when exiting the floatLit production.
	ExitFloatLit(c *FloatLitContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)
}
