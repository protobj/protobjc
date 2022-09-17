grammar Protobj;

protobj
  : packageStatement
    (
        importStatement
      | topLevelDef
      | emptyStatement_
    )* EOF
  ;

// Package

packageStatement
  : PACKAGE fullIdent SEMI
  ;


// Import Statement

importStatement
  : IMPORT  fullIdent  SEMI
  ;


// Option

optionName
  : fullIdent
  ;

// Normal Field

field
  : modifier type_ fieldName EQ fieldNumber ( LB fieldOptions RB )? SEMI
  ;

modifier
  :( REPEATED|SET|ARRAY|DEFAULT )?
  ;

fieldOptions
  : fieldOption ( COMMA  fieldOption )*
  ;

fieldOption
  : optionName EQ constant
  ;

fieldNumber
  : intLit
  ;

extendsField
  :  EXTEND messageType fieldName EQ fieldNumber ( LB fieldOptions RB )? SEMI
  ;
// Map field

mapField
  : ( DEFAULT )? mapType mapName EQ fieldNumber ( LB fieldOptions RB )? SEMI
  ;
mapType
  :MAP LT keyType COMMA type_ GT
  ;

keyType
  : I8
  | U8
  | I16
  | U16
  | I32
  | U32
  | S32
  | F32
  | SF32
  | I64
  | U64
  | S64
  | F64
  | SF64
  | STRING
  | DOUBLE
  | FLOAT
  ;

// field types

type_
  : keyType
  | messageType
  | enumType
  ;

// Top Level definitions

topLevelDef
  : messageDef
  | enumDef
  ;

// enum

enumDef
  : ENUM enumName enumBody
  ;

enumBody
  : LC enumElement* RC
  ;

enumElement
  : enumField
  | emptyStatement_
  ;

enumField
  : ident EQ intLit SEMI
  ;

// message

messageDef
  : MESSAGE messageName(LB messageIndex RB)? messageBody
  ;

messageIndex
  :INT_LIT
  ;

messageBody
  : LC messageElement* RC
  ;

messageElement
  : field
  | extendsField
  | mapField
  | emptyStatement_
  ;
// lexical

constant
  : fullIdent
  | (MINUS | PLUS )? intLit
  | ( MINUS | PLUS )? floatLit
  | strLit
  | boolLit
  | blockLit
  ;

// not specified in specification but used in tests
blockLit
  : LC ( ident COLON constant )* RC
  ;

emptyStatement_: SEMI;

// Lexical elements

ident: IDENTIFIER | keywords;
fullIdent: ident ( DOT ident )*;
messageName: ident;
enumName: ident;
fieldName: ident;
mapName: ident;
messageType: ( DOT )? ( ident DOT )* messageName;
enumType: ( DOT )? ( ident DOT )* enumName;

intLit: INT_LIT;
strLit: STR_LIT;
boolLit: BOOL_LIT;
floatLit: FLOAT_LIT;

// keywords
IMPORT: 'import';
PACKAGE: 'package';
REPEATED: 'lst';
MAP: 'map';
BOOL: 'bool';
I8:'i8';
U8:'u8';
I16: 'i16';
U16: 'u16';
I32: 'i32';
U32: 'u32';
S32: 's32';
F32: 'f32';
SF32: 'sf32';
I64: 'i64';
U64: 'u64';
S64: 's64';
F64: 'f64';
SF64: 'sf64';
STRING: 'string';
DOUBLE: 'double';
FLOAT: 'float';
ENUM: 'enum';
MESSAGE: 'message';
EXTEND: 'ext';
SET: 'set';
ARRAY: 'arr';
DEFAULT:'dft';
// symbols

SEMI: ';';
EQ: '=';
LP: '(';
RP: ')';
LB: '[';
RB: ']';
LC: '{';
RC: '}';
LT: '<';
GT: '>';
DOT: '.';
COMMA: ',';
COLON: ':';
PLUS: '+';
MINUS: '-';

STR_LIT: ( '\'' ( CHAR_VALUE )*? '\'' ) |  ( '"' ( CHAR_VALUE )*? '"' );
fragment CHAR_VALUE: HEX_ESCAPE | OCT_ESCAPE | CHAR_ESCAPE | ~[\u0000\n\\];
fragment HEX_ESCAPE: '\\' ( 'x' | 'X' ) HEX_DIGIT HEX_DIGIT;
fragment OCT_ESCAPE: '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT;
fragment CHAR_ESCAPE: '\\' ( 'a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '\\' | '\'' | '"' );

BOOL_LIT: 'true' | 'false';

FLOAT_LIT : ( DECIMALS DOT DECIMALS? EXPONENT? | DECIMALS EXPONENT | DOT DECIMALS EXPONENT? ) | 'inf' | 'nan';
fragment EXPONENT  : ( 'e' | 'E' ) (PLUS | MINUS)? DECIMALS;
fragment DECIMALS  : DECIMAL_DIGIT+;

INT_LIT     : DECIMAL_LIT | OCTAL_LIT | HEX_LIT;
fragment DECIMAL_LIT : ( [1-9] ) DECIMAL_DIGIT*;
fragment OCTAL_LIT   : '0' OCTAL_DIGIT*;
fragment HEX_LIT     : '0' ( 'x' | 'X' ) HEX_DIGIT+ ;

IDENTIFIER: LETTER ( LETTER | DECIMAL_DIGIT )*;

fragment LETTER: [A-Za-z_];
fragment DECIMAL_DIGIT: [0-9];
fragment OCTAL_DIGIT: [0-7];
fragment HEX_DIGIT: [0-9A-Fa-f];

// comments
WS  :   [ \t\r\n\u000C]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> channel(1);
COMMENT: '/*' .*? '*/' ->  channel(1);

keywords
  : IMPORT
  | PACKAGE
  | REPEATED
  | MAP
  | BOOL
  | I8
  | U8
  | I16
  | U16
  | I32
  | U32
  | S32
  | F32
  | SF32
  | I64
  | U64
  | S64
  | F64
  | SF64
  | STRING
  | DOUBLE
  | FLOAT
  | ENUM
  | MESSAGE
  | EXTEND
  | SET
  | ARRAY
  | DEFAULT
  ;
