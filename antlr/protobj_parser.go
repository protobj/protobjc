// Code generated from D:/code/protobj/protobj-java/src/main/resources\Protobj.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr // Protobj
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ProtobjParser struct {
	*antlr.BaseParser
}

var protobjParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protobjParserInit() {
	staticData := &protobjParserStaticData
	staticData.literalNames = []string{
		"", "'import'", "'package'", "'lst'", "'map'", "'bool'", "'i8'", "'u8'",
		"'i16'", "'u16'", "'i32'", "'u32'", "'s32'", "'f32'", "'sf32'", "'i64'",
		"'u64'", "'s64'", "'f64'", "'sf64'", "'string'", "'double'", "'float'",
		"'enum'", "'message'", "'ext'", "'set'", "'arr'", "'dft'", "';'", "'='",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'", "'.'", "','",
		"':'", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "IMPORT", "PACKAGE", "REPEATED", "MAP", "BOOL", "I8", "U8", "I16",
		"U16", "I32", "U32", "S32", "F32", "SF32", "I64", "U64", "S64", "F64",
		"SF64", "STRING", "DOUBLE", "FLOAT", "ENUM", "MESSAGE", "EXTEND", "SET",
		"ARRAY", "DEFAULT", "SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC",
		"LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "STR_LIT", "BOOL_LIT",
		"FLOAT_LIT", "INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.ruleNames = []string{
		"protobj", "packageStatement", "importStatement", "optionName", "field",
		"modifier", "fieldOptions", "fieldOption", "fieldNumber", "extendsField",
		"mapField", "mapType", "keyType", "type_", "topLevelDef", "enumDef",
		"enumBody", "enumElement", "enumField", "messageDef", "messageIndex",
		"messageBody", "messageElement", "constant", "blockLit", "emptyStatement_",
		"ident", "fullIdent", "messageName", "enumName", "fieldName", "mapName",
		"messageType", "enumType", "intLit", "strLit", "boolLit", "floatLit",
		"keywords",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 309, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 83, 8, 0, 10,
		0, 12, 0, 86, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 109, 8, 4, 1, 4, 1, 4, 1, 5, 3, 5, 114, 8, 5, 1, 6, 1, 6, 1, 6,
		5, 6, 119, 8, 6, 10, 6, 12, 6, 122, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 139,
		8, 9, 1, 9, 1, 9, 1, 10, 3, 10, 144, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 154, 8, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		3, 13, 170, 8, 13, 1, 14, 1, 14, 3, 14, 174, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 5, 16, 182, 8, 16, 10, 16, 12, 16, 185, 9, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 3, 17, 191, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 204, 8, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 212, 8, 21, 10, 21, 12, 21,
		215, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 223, 8, 22,
		1, 23, 1, 23, 3, 23, 227, 8, 23, 1, 23, 1, 23, 3, 23, 231, 8, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 237, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 5, 24, 244, 8, 24, 10, 24, 12, 24, 247, 9, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 26, 1, 26, 3, 26, 255, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 260,
		8, 27, 10, 27, 12, 27, 263, 9, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 3, 32, 274, 8, 32, 1, 32, 1, 32, 1, 32, 5, 32,
		279, 8, 32, 10, 32, 12, 32, 282, 9, 32, 1, 32, 1, 32, 1, 33, 3, 33, 287,
		8, 33, 1, 33, 1, 33, 1, 33, 5, 33, 292, 8, 33, 10, 33, 12, 33, 295, 9,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 0, 0, 39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 70, 72, 74, 76, 0, 4, 2, 0, 3, 3, 26, 28, 1, 0, 6,
		22, 1, 0, 42, 43, 1, 0, 1, 28, 302, 0, 78, 1, 0, 0, 0, 2, 89, 1, 0, 0,
		0, 4, 93, 1, 0, 0, 0, 6, 97, 1, 0, 0, 0, 8, 99, 1, 0, 0, 0, 10, 113, 1,
		0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 123, 1, 0, 0, 0, 16, 127, 1, 0, 0, 0,
		18, 129, 1, 0, 0, 0, 20, 143, 1, 0, 0, 0, 22, 157, 1, 0, 0, 0, 24, 164,
		1, 0, 0, 0, 26, 169, 1, 0, 0, 0, 28, 173, 1, 0, 0, 0, 30, 175, 1, 0, 0,
		0, 32, 179, 1, 0, 0, 0, 34, 190, 1, 0, 0, 0, 36, 192, 1, 0, 0, 0, 38, 197,
		1, 0, 0, 0, 40, 207, 1, 0, 0, 0, 42, 209, 1, 0, 0, 0, 44, 222, 1, 0, 0,
		0, 46, 236, 1, 0, 0, 0, 48, 238, 1, 0, 0, 0, 50, 250, 1, 0, 0, 0, 52, 254,
		1, 0, 0, 0, 54, 256, 1, 0, 0, 0, 56, 264, 1, 0, 0, 0, 58, 266, 1, 0, 0,
		0, 60, 268, 1, 0, 0, 0, 62, 270, 1, 0, 0, 0, 64, 273, 1, 0, 0, 0, 66, 286,
		1, 0, 0, 0, 68, 298, 1, 0, 0, 0, 70, 300, 1, 0, 0, 0, 72, 302, 1, 0, 0,
		0, 74, 304, 1, 0, 0, 0, 76, 306, 1, 0, 0, 0, 78, 84, 3, 2, 1, 0, 79, 83,
		3, 4, 2, 0, 80, 83, 3, 28, 14, 0, 81, 83, 3, 50, 25, 0, 82, 79, 1, 0, 0,
		0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		87, 88, 5, 0, 0, 1, 88, 1, 1, 0, 0, 0, 89, 90, 5, 2, 0, 0, 90, 91, 3, 54,
		27, 0, 91, 92, 5, 29, 0, 0, 92, 3, 1, 0, 0, 0, 93, 94, 5, 1, 0, 0, 94,
		95, 3, 54, 27, 0, 95, 96, 5, 29, 0, 0, 96, 5, 1, 0, 0, 0, 97, 98, 3, 54,
		27, 0, 98, 7, 1, 0, 0, 0, 99, 100, 3, 10, 5, 0, 100, 101, 3, 26, 13, 0,
		101, 102, 3, 60, 30, 0, 102, 103, 5, 30, 0, 0, 103, 108, 3, 16, 8, 0, 104,
		105, 5, 33, 0, 0, 105, 106, 3, 12, 6, 0, 106, 107, 5, 34, 0, 0, 107, 109,
		1, 0, 0, 0, 108, 104, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 111, 5, 29, 0, 0, 111, 9, 1, 0, 0, 0, 112, 114, 7, 0, 0, 0,
		113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 11, 1, 0, 0, 0, 115, 120,
		3, 14, 7, 0, 116, 117, 5, 40, 0, 0, 117, 119, 3, 14, 7, 0, 118, 116, 1,
		0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0,
		0, 121, 13, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 3, 6, 3, 0, 124,
		125, 5, 30, 0, 0, 125, 126, 3, 46, 23, 0, 126, 15, 1, 0, 0, 0, 127, 128,
		3, 68, 34, 0, 128, 17, 1, 0, 0, 0, 129, 130, 5, 25, 0, 0, 130, 131, 3,
		64, 32, 0, 131, 132, 3, 60, 30, 0, 132, 133, 5, 30, 0, 0, 133, 138, 3,
		16, 8, 0, 134, 135, 5, 33, 0, 0, 135, 136, 3, 12, 6, 0, 136, 137, 5, 34,
		0, 0, 137, 139, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 141, 5, 29, 0, 0, 141, 19, 1, 0, 0, 0, 142,
		144, 5, 28, 0, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145,
		1, 0, 0, 0, 145, 146, 3, 22, 11, 0, 146, 147, 3, 62, 31, 0, 147, 148, 5,
		30, 0, 0, 148, 153, 3, 16, 8, 0, 149, 150, 5, 33, 0, 0, 150, 151, 3, 12,
		6, 0, 151, 152, 5, 34, 0, 0, 152, 154, 1, 0, 0, 0, 153, 149, 1, 0, 0, 0,
		153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 5, 29, 0, 0, 156,
		21, 1, 0, 0, 0, 157, 158, 5, 4, 0, 0, 158, 159, 5, 37, 0, 0, 159, 160,
		3, 24, 12, 0, 160, 161, 5, 40, 0, 0, 161, 162, 3, 26, 13, 0, 162, 163,
		5, 38, 0, 0, 163, 23, 1, 0, 0, 0, 164, 165, 7, 1, 0, 0, 165, 25, 1, 0,
		0, 0, 166, 170, 3, 24, 12, 0, 167, 170, 3, 64, 32, 0, 168, 170, 3, 66,
		33, 0, 169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0,
		170, 27, 1, 0, 0, 0, 171, 174, 3, 38, 19, 0, 172, 174, 3, 30, 15, 0, 173,
		171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 29, 1, 0, 0, 0, 175, 176, 5,
		23, 0, 0, 176, 177, 3, 58, 29, 0, 177, 178, 3, 32, 16, 0, 178, 31, 1, 0,
		0, 0, 179, 183, 5, 35, 0, 0, 180, 182, 3, 34, 17, 0, 181, 180, 1, 0, 0,
		0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 36, 0, 0, 187, 33,
		1, 0, 0, 0, 188, 191, 3, 36, 18, 0, 189, 191, 3, 50, 25, 0, 190, 188, 1,
		0, 0, 0, 190, 189, 1, 0, 0, 0, 191, 35, 1, 0, 0, 0, 192, 193, 3, 52, 26,
		0, 193, 194, 5, 30, 0, 0, 194, 195, 3, 68, 34, 0, 195, 196, 5, 29, 0, 0,
		196, 37, 1, 0, 0, 0, 197, 198, 5, 24, 0, 0, 198, 203, 3, 56, 28, 0, 199,
		200, 5, 33, 0, 0, 200, 201, 3, 40, 20, 0, 201, 202, 5, 34, 0, 0, 202, 204,
		1, 0, 0, 0, 203, 199, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 1, 0,
		0, 0, 205, 206, 3, 42, 21, 0, 206, 39, 1, 0, 0, 0, 207, 208, 5, 47, 0,
		0, 208, 41, 1, 0, 0, 0, 209, 213, 5, 35, 0, 0, 210, 212, 3, 44, 22, 0,
		211, 210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213,
		214, 1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217,
		5, 36, 0, 0, 217, 43, 1, 0, 0, 0, 218, 223, 3, 8, 4, 0, 219, 223, 3, 18,
		9, 0, 220, 223, 3, 20, 10, 0, 221, 223, 3, 50, 25, 0, 222, 218, 1, 0, 0,
		0, 222, 219, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223,
		45, 1, 0, 0, 0, 224, 237, 3, 54, 27, 0, 225, 227, 7, 2, 0, 0, 226, 225,
		1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 237, 3, 68,
		34, 0, 229, 231, 7, 2, 0, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0,
		231, 232, 1, 0, 0, 0, 232, 237, 3, 74, 37, 0, 233, 237, 3, 70, 35, 0, 234,
		237, 3, 72, 36, 0, 235, 237, 3, 48, 24, 0, 236, 224, 1, 0, 0, 0, 236, 226,
		1, 0, 0, 0, 236, 230, 1, 0, 0, 0, 236, 233, 1, 0, 0, 0, 236, 234, 1, 0,
		0, 0, 236, 235, 1, 0, 0, 0, 237, 47, 1, 0, 0, 0, 238, 245, 5, 35, 0, 0,
		239, 240, 3, 52, 26, 0, 240, 241, 5, 41, 0, 0, 241, 242, 3, 46, 23, 0,
		242, 244, 1, 0, 0, 0, 243, 239, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245,
		243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 248, 249, 5, 36, 0, 0, 249, 49, 1, 0, 0, 0, 250, 251, 5, 29,
		0, 0, 251, 51, 1, 0, 0, 0, 252, 255, 5, 48, 0, 0, 253, 255, 3, 76, 38,
		0, 254, 252, 1, 0, 0, 0, 254, 253, 1, 0, 0, 0, 255, 53, 1, 0, 0, 0, 256,
		261, 3, 52, 26, 0, 257, 258, 5, 39, 0, 0, 258, 260, 3, 52, 26, 0, 259,
		257, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262,
		1, 0, 0, 0, 262, 55, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 265, 3, 52,
		26, 0, 265, 57, 1, 0, 0, 0, 266, 267, 3, 52, 26, 0, 267, 59, 1, 0, 0, 0,
		268, 269, 3, 52, 26, 0, 269, 61, 1, 0, 0, 0, 270, 271, 3, 52, 26, 0, 271,
		63, 1, 0, 0, 0, 272, 274, 5, 39, 0, 0, 273, 272, 1, 0, 0, 0, 273, 274,
		1, 0, 0, 0, 274, 280, 1, 0, 0, 0, 275, 276, 3, 52, 26, 0, 276, 277, 5,
		39, 0, 0, 277, 279, 1, 0, 0, 0, 278, 275, 1, 0, 0, 0, 279, 282, 1, 0, 0,
		0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282,
		280, 1, 0, 0, 0, 283, 284, 3, 56, 28, 0, 284, 65, 1, 0, 0, 0, 285, 287,
		5, 39, 0, 0, 286, 285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 293, 1, 0,
		0, 0, 288, 289, 3, 52, 26, 0, 289, 290, 5, 39, 0, 0, 290, 292, 1, 0, 0,
		0, 291, 288, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 297,
		3, 58, 29, 0, 297, 67, 1, 0, 0, 0, 298, 299, 5, 47, 0, 0, 299, 69, 1, 0,
		0, 0, 300, 301, 5, 44, 0, 0, 301, 71, 1, 0, 0, 0, 302, 303, 5, 45, 0, 0,
		303, 73, 1, 0, 0, 0, 304, 305, 5, 46, 0, 0, 305, 75, 1, 0, 0, 0, 306, 307,
		7, 3, 0, 0, 307, 77, 1, 0, 0, 0, 25, 82, 84, 108, 113, 120, 138, 143, 153,
		169, 173, 183, 190, 203, 213, 222, 226, 230, 236, 245, 254, 261, 273, 280,
		286, 293,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ProtobjParserInit initializes any static state used to implement ProtobjParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProtobjParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProtobjParserInit() {
	staticData := &protobjParserStaticData
	staticData.once.Do(protobjParserInit)
}

// NewProtobjParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProtobjParser(input antlr.TokenStream) *ProtobjParser {
	ProtobjParserInit()
	this := new(ProtobjParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &protobjParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Protobj.g4"

	return this
}

// ProtobjParser tokens.
const (
	ProtobjParserEOF          = antlr.TokenEOF
	ProtobjParserIMPORT       = 1
	ProtobjParserPACKAGE      = 2
	ProtobjParserREPEATED     = 3
	ProtobjParserMAP          = 4
	ProtobjParserBOOL         = 5
	ProtobjParserI8           = 6
	ProtobjParserU8           = 7
	ProtobjParserI16          = 8
	ProtobjParserU16          = 9
	ProtobjParserI32          = 10
	ProtobjParserU32          = 11
	ProtobjParserS32          = 12
	ProtobjParserF32          = 13
	ProtobjParserSF32         = 14
	ProtobjParserI64          = 15
	ProtobjParserU64          = 16
	ProtobjParserS64          = 17
	ProtobjParserF64          = 18
	ProtobjParserSF64         = 19
	ProtobjParserSTRING       = 20
	ProtobjParserDOUBLE       = 21
	ProtobjParserFLOAT        = 22
	ProtobjParserENUM         = 23
	ProtobjParserMESSAGE      = 24
	ProtobjParserEXTEND       = 25
	ProtobjParserSET          = 26
	ProtobjParserARRAY        = 27
	ProtobjParserDEFAULT      = 28
	ProtobjParserSEMI         = 29
	ProtobjParserEQ           = 30
	ProtobjParserLP           = 31
	ProtobjParserRP           = 32
	ProtobjParserLB           = 33
	ProtobjParserRB           = 34
	ProtobjParserLC           = 35
	ProtobjParserRC           = 36
	ProtobjParserLT           = 37
	ProtobjParserGT           = 38
	ProtobjParserDOT          = 39
	ProtobjParserCOMMA        = 40
	ProtobjParserCOLON        = 41
	ProtobjParserPLUS         = 42
	ProtobjParserMINUS        = 43
	ProtobjParserSTR_LIT      = 44
	ProtobjParserBOOL_LIT     = 45
	ProtobjParserFLOAT_LIT    = 46
	ProtobjParserINT_LIT      = 47
	ProtobjParserIDENTIFIER   = 48
	ProtobjParserWS           = 49
	ProtobjParserLINE_COMMENT = 50
	ProtobjParserCOMMENT      = 51
)

// ProtobjParser rules.
const (
	ProtobjParserRULE_protobj          = 0
	ProtobjParserRULE_packageStatement = 1
	ProtobjParserRULE_importStatement  = 2
	ProtobjParserRULE_optionName       = 3
	ProtobjParserRULE_field            = 4
	ProtobjParserRULE_modifier         = 5
	ProtobjParserRULE_fieldOptions     = 6
	ProtobjParserRULE_fieldOption      = 7
	ProtobjParserRULE_fieldNumber      = 8
	ProtobjParserRULE_extendsField     = 9
	ProtobjParserRULE_mapField         = 10
	ProtobjParserRULE_mapType          = 11
	ProtobjParserRULE_keyType          = 12
	ProtobjParserRULE_type_            = 13
	ProtobjParserRULE_topLevelDef      = 14
	ProtobjParserRULE_enumDef          = 15
	ProtobjParserRULE_enumBody         = 16
	ProtobjParserRULE_enumElement      = 17
	ProtobjParserRULE_enumField        = 18
	ProtobjParserRULE_messageDef       = 19
	ProtobjParserRULE_messageIndex     = 20
	ProtobjParserRULE_messageBody      = 21
	ProtobjParserRULE_messageElement   = 22
	ProtobjParserRULE_constant         = 23
	ProtobjParserRULE_blockLit         = 24
	ProtobjParserRULE_emptyStatement_  = 25
	ProtobjParserRULE_ident            = 26
	ProtobjParserRULE_fullIdent        = 27
	ProtobjParserRULE_messageName      = 28
	ProtobjParserRULE_enumName         = 29
	ProtobjParserRULE_fieldName        = 30
	ProtobjParserRULE_mapName          = 31
	ProtobjParserRULE_messageType      = 32
	ProtobjParserRULE_enumType         = 33
	ProtobjParserRULE_intLit           = 34
	ProtobjParserRULE_strLit           = 35
	ProtobjParserRULE_boolLit          = 36
	ProtobjParserRULE_floatLit         = 37
	ProtobjParserRULE_keywords         = 38
)

// IProtobjContext is an interface to support dynamic dispatch.
type IProtobjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtobjContext differentiates from other interfaces.
	IsProtobjContext()
}

type ProtobjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtobjContext() *ProtobjContext {
	var p = new(ProtobjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_protobj
	return p
}

func (*ProtobjContext) IsProtobjContext() {}

func NewProtobjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtobjContext {
	var p = new(ProtobjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_protobj

	return p
}

func (s *ProtobjContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtobjContext) PackageStatement() IPackageStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *ProtobjContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEOF, 0)
}

func (s *ProtobjContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtobjContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProtobjContext) AllTopLevelDef() []ITopLevelDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDefContext); ok {
			tst[i] = t.(ITopLevelDefContext)
			i++
		}
	}

	return tst
}

func (s *ProtobjContext) TopLevelDef(i int) ITopLevelDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *ProtobjContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *ProtobjContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ProtobjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtobjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtobjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterProtobj(s)
	}
}

func (s *ProtobjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitProtobj(s)
	}
}

func (p *ProtobjParser) Protobj() (localctx IProtobjContext) {
	this := p
	_ = this

	localctx = NewProtobjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProtobjParserRULE_protobj)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.PackageStatement()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserIMPORT)|(1<<ProtobjParserENUM)|(1<<ProtobjParserMESSAGE)|(1<<ProtobjParserSEMI))) != 0 {
		p.SetState(82)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ProtobjParserIMPORT:
			{
				p.SetState(79)
				p.ImportStatement()
			}

		case ProtobjParserENUM, ProtobjParserMESSAGE:
			{
				p.SetState(80)
				p.TopLevelDef()
			}

		case ProtobjParserSEMI:
			{
				p.SetState(81)
				p.EmptyStatement_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(ProtobjParserEOF)
	}

	return localctx
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(ProtobjParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *ProtobjParser) PackageStatement() (localctx IPackageStatementContext) {
	this := p
	_ = this

	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProtobjParserRULE_packageStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(ProtobjParserPACKAGE)
	}
	{
		p.SetState(90)
		p.FullIdent()
	}
	{
		p.SetState(91)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserIMPORT, 0)
}

func (s *ImportStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *ProtobjParser) ImportStatement() (localctx IImportStatementContext) {
	this := p
	_ = this

	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProtobjParserRULE_importStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(ProtobjParserIMPORT)
	}
	{
		p.SetState(94)
		p.FullIdent()
	}
	{
		p.SetState(95)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *ProtobjParser) OptionName() (localctx IOptionNameContext) {
	this := p
	_ = this

	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProtobjParserRULE_optionName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.FullIdent()
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Modifier() IModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *FieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEQ, 0)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *FieldContext) LB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLB, 0)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) RB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRB, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ProtobjParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProtobjParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Modifier()
	}
	{
		p.SetState(100)
		p.Type_()
	}
	{
		p.SetState(101)
		p.FieldName()
	}
	{
		p.SetState(102)
		p.Match(ProtobjParserEQ)
	}
	{
		p.SetState(103)
		p.FieldNumber()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserLB {
		{
			p.SetState(104)
			p.Match(ProtobjParserLB)
		}
		{
			p.SetState(105)
			p.FieldOptions()
		}
		{
			p.SetState(106)
			p.Match(ProtobjParserRB)
		}

	}
	{
		p.SetState(110)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_modifier
	return p
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(ProtobjParserREPEATED, 0)
}

func (s *ModifierContext) SET() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSET, 0)
}

func (s *ModifierContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(ProtobjParserARRAY, 0)
}

func (s *ModifierContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserDEFAULT, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (p *ProtobjParser) Modifier() (localctx IModifierContext) {
	this := p
	_ = this

	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProtobjParserRULE_modifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(112)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserREPEATED)|(1<<ProtobjParserSET)|(1<<ProtobjParserARRAY)|(1<<ProtobjParserDEFAULT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldOptionContext); ok {
			len++
		}
	}

	tst := make([]IFieldOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldOptionContext); ok {
			tst[i] = t.(IFieldOptionContext)
			i++
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProtobjParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProtobjParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (p *ProtobjParser) FieldOptions() (localctx IFieldOptionsContext) {
	this := p
	_ = this

	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProtobjParserRULE_fieldOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.FieldOption()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtobjParserCOMMA {
		{
			p.SetState(116)
			p.Match(ProtobjParserCOMMA)
		}
		{
			p.SetState(117)
			p.FieldOption()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEQ, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (p *ProtobjParser) FieldOption() (localctx IFieldOptionContext) {
	this := p
	_ = this

	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProtobjParserRULE_fieldOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.OptionName()
	}
	{
		p.SetState(124)
		p.Match(ProtobjParserEQ)
	}
	{
		p.SetState(125)
		p.Constant()
	}

	return localctx
}

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_fieldNumber
	return p
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (p *ProtobjParser) FieldNumber() (localctx IFieldNumberContext) {
	this := p
	_ = this

	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProtobjParserRULE_fieldNumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.IntLit()
	}

	return localctx
}

// IExtendsFieldContext is an interface to support dynamic dispatch.
type IExtendsFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendsFieldContext differentiates from other interfaces.
	IsExtendsFieldContext()
}

type ExtendsFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendsFieldContext() *ExtendsFieldContext {
	var p = new(ExtendsFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_extendsField
	return p
}

func (*ExtendsFieldContext) IsExtendsFieldContext() {}

func NewExtendsFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendsFieldContext {
	var p = new(ExtendsFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_extendsField

	return p
}

func (s *ExtendsFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendsFieldContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEXTEND, 0)
}

func (s *ExtendsFieldContext) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *ExtendsFieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *ExtendsFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEQ, 0)
}

func (s *ExtendsFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *ExtendsFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *ExtendsFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLB, 0)
}

func (s *ExtendsFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *ExtendsFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRB, 0)
}

func (s *ExtendsFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendsFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendsFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterExtendsField(s)
	}
}

func (s *ExtendsFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitExtendsField(s)
	}
}

func (p *ProtobjParser) ExtendsField() (localctx IExtendsFieldContext) {
	this := p
	_ = this

	localctx = NewExtendsFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProtobjParserRULE_extendsField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(ProtobjParserEXTEND)
	}
	{
		p.SetState(130)
		p.MessageType()
	}
	{
		p.SetState(131)
		p.FieldName()
	}
	{
		p.SetState(132)
		p.Match(ProtobjParserEQ)
	}
	{
		p.SetState(133)
		p.FieldNumber()
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserLB {
		{
			p.SetState(134)
			p.Match(ProtobjParserLB)
		}
		{
			p.SetState(135)
			p.FieldOptions()
		}
		{
			p.SetState(136)
			p.Match(ProtobjParserRB)
		}

	}
	{
		p.SetState(140)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEQ, 0)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *MapFieldContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserDEFAULT, 0)
}

func (s *MapFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLB, 0)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRB, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *ProtobjParser) MapField() (localctx IMapFieldContext) {
	this := p
	_ = this

	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProtobjParserRULE_mapField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserDEFAULT {
		{
			p.SetState(142)
			p.Match(ProtobjParserDEFAULT)
		}

	}
	{
		p.SetState(145)
		p.MapType()
	}
	{
		p.SetState(146)
		p.MapName()
	}
	{
		p.SetState(147)
		p.Match(ProtobjParserEQ)
	}
	{
		p.SetState(148)
		p.FieldNumber()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserLB {
		{
			p.SetState(149)
			p.Match(ProtobjParserLB)
		}
		{
			p.SetState(150)
			p.FieldOptions()
		}
		{
			p.SetState(151)
			p.Match(ProtobjParserRB)
		}

	}
	{
		p.SetState(155)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(ProtobjParserMAP, 0)
}

func (s *MapTypeContext) LT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLT, 0)
}

func (s *MapTypeContext) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProtobjParserCOMMA, 0)
}

func (s *MapTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapTypeContext) GT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserGT, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (p *ProtobjParser) MapType() (localctx IMapTypeContext) {
	this := p
	_ = this

	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ProtobjParserRULE_mapType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ProtobjParserMAP)
	}
	{
		p.SetState(158)
		p.Match(ProtobjParserLT)
	}
	{
		p.SetState(159)
		p.KeyType()
	}
	{
		p.SetState(160)
		p.Match(ProtobjParserCOMMA)
	}
	{
		p.SetState(161)
		p.Type_()
	}
	{
		p.SetState(162)
		p.Match(ProtobjParserGT)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) I8() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI8, 0)
}

func (s *KeyTypeContext) U8() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU8, 0)
}

func (s *KeyTypeContext) I16() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI16, 0)
}

func (s *KeyTypeContext) U16() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU16, 0)
}

func (s *KeyTypeContext) I32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI32, 0)
}

func (s *KeyTypeContext) U32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU32, 0)
}

func (s *KeyTypeContext) S32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserS32, 0)
}

func (s *KeyTypeContext) F32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserF32, 0)
}

func (s *KeyTypeContext) SF32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSF32, 0)
}

func (s *KeyTypeContext) I64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI64, 0)
}

func (s *KeyTypeContext) U64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU64, 0)
}

func (s *KeyTypeContext) S64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserS64, 0)
}

func (s *KeyTypeContext) F64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserF64, 0)
}

func (s *KeyTypeContext) SF64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSF64, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSTRING, 0)
}

func (s *KeyTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ProtobjParserDOUBLE, 0)
}

func (s *KeyTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserFLOAT, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *ProtobjParser) KeyType() (localctx IKeyTypeContext) {
	this := p
	_ = this

	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ProtobjParserRULE_keyType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserI8)|(1<<ProtobjParserU8)|(1<<ProtobjParserI16)|(1<<ProtobjParserU16)|(1<<ProtobjParserI32)|(1<<ProtobjParserU32)|(1<<ProtobjParserS32)|(1<<ProtobjParserF32)|(1<<ProtobjParserSF32)|(1<<ProtobjParserI64)|(1<<ProtobjParserU64)|(1<<ProtobjParserS64)|(1<<ProtobjParserF64)|(1<<ProtobjParserSF64)|(1<<ProtobjParserSTRING)|(1<<ProtobjParserDOUBLE)|(1<<ProtobjParserFLOAT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *Type_Context) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *ProtobjParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ProtobjParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.KeyType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.MessageType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.EnumType()
		}

	}

	return localctx
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_topLevelDef
	return p
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) MessageDef() IMessageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (p *ProtobjParser) TopLevelDef() (localctx ITopLevelDefContext) {
	this := p
	_ = this

	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ProtobjParserRULE_topLevelDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProtobjParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.MessageDef()
		}

	case ProtobjParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.EnumDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_enumDef
	return p
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ProtobjParserENUM, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *ProtobjParser) EnumDef() (localctx IEnumDefContext) {
	this := p
	_ = this

	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ProtobjParserRULE_enumDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(ProtobjParserENUM)
	}
	{
		p.SetState(176)
		p.EnumName()
	}
	{
		p.SetState(177)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumElementContext); ok {
			len++
		}
	}

	tst := make([]IEnumElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumElementContext); ok {
			tst[i] = t.(IEnumElementContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *ProtobjParser) EnumBody() (localctx IEnumBodyContext) {
	this := p
	_ = this

	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ProtobjParserRULE_enumBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(ProtobjParserLC)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserIMPORT)|(1<<ProtobjParserPACKAGE)|(1<<ProtobjParserREPEATED)|(1<<ProtobjParserMAP)|(1<<ProtobjParserBOOL)|(1<<ProtobjParserI8)|(1<<ProtobjParserU8)|(1<<ProtobjParserI16)|(1<<ProtobjParserU16)|(1<<ProtobjParserI32)|(1<<ProtobjParserU32)|(1<<ProtobjParserS32)|(1<<ProtobjParserF32)|(1<<ProtobjParserSF32)|(1<<ProtobjParserI64)|(1<<ProtobjParserU64)|(1<<ProtobjParserS64)|(1<<ProtobjParserF64)|(1<<ProtobjParserSF64)|(1<<ProtobjParserSTRING)|(1<<ProtobjParserDOUBLE)|(1<<ProtobjParserFLOAT)|(1<<ProtobjParserENUM)|(1<<ProtobjParserMESSAGE)|(1<<ProtobjParserEXTEND)|(1<<ProtobjParserSET)|(1<<ProtobjParserARRAY)|(1<<ProtobjParserDEFAULT)|(1<<ProtobjParserSEMI))) != 0) || _la == ProtobjParserIDENTIFIER {
		{
			p.SetState(180)
			p.EnumElement()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(186)
		p.Match(ProtobjParserRC)
	}

	return localctx
}

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_enumElement
	return p
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEnumElement(s)
	}
}

func (s *EnumElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEnumElement(s)
	}
}

func (p *ProtobjParser) EnumElement() (localctx IEnumElementContext) {
	this := p
	_ = this

	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ProtobjParserRULE_enumElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProtobjParserIMPORT, ProtobjParserPACKAGE, ProtobjParserREPEATED, ProtobjParserMAP, ProtobjParserBOOL, ProtobjParserI8, ProtobjParserU8, ProtobjParserI16, ProtobjParserU16, ProtobjParserI32, ProtobjParserU32, ProtobjParserS32, ProtobjParserF32, ProtobjParserSF32, ProtobjParserI64, ProtobjParserU64, ProtobjParserS64, ProtobjParserF64, ProtobjParserSF64, ProtobjParserSTRING, ProtobjParserDOUBLE, ProtobjParserFLOAT, ProtobjParserENUM, ProtobjParserMESSAGE, ProtobjParserEXTEND, ProtobjParserSET, ProtobjParserARRAY, ProtobjParserDEFAULT, ProtobjParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.EnumField()
		}

	case ProtobjParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.EmptyStatement_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *ProtobjParser) EnumField() (localctx IEnumFieldContext) {
	this := p
	_ = this

	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ProtobjParserRULE_enumField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Ident()
	}
	{
		p.SetState(193)
		p.Match(ProtobjParserEQ)
	}
	{
		p.SetState(194)
		p.IntLit()
	}
	{
		p.SetState(195)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IMessageDefContext is an interface to support dynamic dispatch.
type IMessageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageDefContext differentiates from other interfaces.
	IsMessageDefContext()
}

type MessageDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefContext() *MessageDefContext {
	var p = new(MessageDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_messageDef
	return p
}

func (*MessageDefContext) IsMessageDefContext() {}

func NewMessageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefContext {
	var p = new(MessageDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_messageDef

	return p
}

func (s *MessageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(ProtobjParserMESSAGE, 0)
}

func (s *MessageDefContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageDefContext) MessageBody() IMessageBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageDefContext) LB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLB, 0)
}

func (s *MessageDefContext) MessageIndex() IMessageIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageIndexContext)
}

func (s *MessageDefContext) RB() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRB, 0)
}

func (s *MessageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMessageDef(s)
	}
}

func (s *MessageDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMessageDef(s)
	}
}

func (p *ProtobjParser) MessageDef() (localctx IMessageDefContext) {
	this := p
	_ = this

	localctx = NewMessageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ProtobjParserRULE_messageDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(ProtobjParserMESSAGE)
	}
	{
		p.SetState(198)
		p.MessageName()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserLB {
		{
			p.SetState(199)
			p.Match(ProtobjParserLB)
		}
		{
			p.SetState(200)
			p.MessageIndex()
		}
		{
			p.SetState(201)
			p.Match(ProtobjParserRB)
		}

	}
	{
		p.SetState(205)
		p.MessageBody()
	}

	return localctx
}

// IMessageIndexContext is an interface to support dynamic dispatch.
type IMessageIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageIndexContext differentiates from other interfaces.
	IsMessageIndexContext()
}

type MessageIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageIndexContext() *MessageIndexContext {
	var p = new(MessageIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_messageIndex
	return p
}

func (*MessageIndexContext) IsMessageIndexContext() {}

func NewMessageIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageIndexContext {
	var p = new(MessageIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_messageIndex

	return p
}

func (s *MessageIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageIndexContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserINT_LIT, 0)
}

func (s *MessageIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMessageIndex(s)
	}
}

func (s *MessageIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMessageIndex(s)
	}
}

func (p *ProtobjParser) MessageIndex() (localctx IMessageIndexContext) {
	this := p
	_ = this

	localctx = NewMessageIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ProtobjParserRULE_messageIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(ProtobjParserINT_LIT)
	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLC, 0)
}

func (s *MessageBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRC, 0)
}

func (s *MessageBodyContext) AllMessageElement() []IMessageElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageElementContext); ok {
			len++
		}
	}

	tst := make([]IMessageElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageElementContext); ok {
			tst[i] = t.(IMessageElementContext)
			i++
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageElement(i int) IMessageElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageElementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (p *ProtobjParser) MessageBody() (localctx IMessageBodyContext) {
	this := p
	_ = this

	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ProtobjParserRULE_messageBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(ProtobjParserLC)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserIMPORT)|(1<<ProtobjParserPACKAGE)|(1<<ProtobjParserREPEATED)|(1<<ProtobjParserMAP)|(1<<ProtobjParserBOOL)|(1<<ProtobjParserI8)|(1<<ProtobjParserU8)|(1<<ProtobjParserI16)|(1<<ProtobjParserU16)|(1<<ProtobjParserI32)|(1<<ProtobjParserU32)|(1<<ProtobjParserS32)|(1<<ProtobjParserF32)|(1<<ProtobjParserSF32)|(1<<ProtobjParserI64)|(1<<ProtobjParserU64)|(1<<ProtobjParserS64)|(1<<ProtobjParserF64)|(1<<ProtobjParserSF64)|(1<<ProtobjParserSTRING)|(1<<ProtobjParserDOUBLE)|(1<<ProtobjParserFLOAT)|(1<<ProtobjParserENUM)|(1<<ProtobjParserMESSAGE)|(1<<ProtobjParserEXTEND)|(1<<ProtobjParserSET)|(1<<ProtobjParserARRAY)|(1<<ProtobjParserDEFAULT)|(1<<ProtobjParserSEMI))) != 0) || _la == ProtobjParserDOT || _la == ProtobjParserIDENTIFIER {
		{
			p.SetState(210)
			p.MessageElement()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(ProtobjParserRC)
	}

	return localctx
}

// IMessageElementContext is an interface to support dynamic dispatch.
type IMessageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageElementContext differentiates from other interfaces.
	IsMessageElementContext()
}

type MessageElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageElementContext() *MessageElementContext {
	var p = new(MessageElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_messageElement
	return p
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_messageElement

	return p
}

func (s *MessageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageElementContext) ExtendsField() IExtendsFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendsFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendsFieldContext)
}

func (s *MessageElementContext) MapField() IMapFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *MessageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMessageElement(s)
	}
}

func (s *MessageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMessageElement(s)
	}
}

func (p *ProtobjParser) MessageElement() (localctx IMessageElementContext) {
	this := p
	_ = this

	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ProtobjParserRULE_messageElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.ExtendsField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.MapField()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)
			p.EmptyStatement_()
		}

	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ProtobjParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ProtobjParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) BlockLit() IBlockLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *ProtobjParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ProtobjParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProtobjParserPLUS || _la == ProtobjParserMINUS {
			{
				p.SetState(225)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProtobjParserPLUS || _la == ProtobjParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(228)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProtobjParserPLUS || _la == ProtobjParserMINUS {
			{
				p.SetState(229)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProtobjParserPLUS || _la == ProtobjParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(232)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(233)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(234)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(235)
			p.BlockLit()
		}

	}

	return localctx
}

// IBlockLitContext is an interface to support dynamic dispatch.
type IBlockLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockLitContext differentiates from other interfaces.
	IsBlockLitContext()
}

type BlockLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockLitContext() *BlockLitContext {
	var p = new(BlockLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_blockLit
	return p
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(ProtobjParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(ProtobjParserRC, 0)
}

func (s *BlockLitContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *BlockLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ProtobjParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ProtobjParserCOLON, i)
}

func (s *BlockLitContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *BlockLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterBlockLit(s)
	}
}

func (s *BlockLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitBlockLit(s)
	}
}

func (p *ProtobjParser) BlockLit() (localctx IBlockLitContext) {
	this := p
	_ = this

	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ProtobjParserRULE_blockLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(ProtobjParserLC)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserIMPORT)|(1<<ProtobjParserPACKAGE)|(1<<ProtobjParserREPEATED)|(1<<ProtobjParserMAP)|(1<<ProtobjParserBOOL)|(1<<ProtobjParserI8)|(1<<ProtobjParserU8)|(1<<ProtobjParserI16)|(1<<ProtobjParserU16)|(1<<ProtobjParserI32)|(1<<ProtobjParserU32)|(1<<ProtobjParserS32)|(1<<ProtobjParserF32)|(1<<ProtobjParserSF32)|(1<<ProtobjParserI64)|(1<<ProtobjParserU64)|(1<<ProtobjParserS64)|(1<<ProtobjParserF64)|(1<<ProtobjParserSF64)|(1<<ProtobjParserSTRING)|(1<<ProtobjParserDOUBLE)|(1<<ProtobjParserFLOAT)|(1<<ProtobjParserENUM)|(1<<ProtobjParserMESSAGE)|(1<<ProtobjParserEXTEND)|(1<<ProtobjParserSET)|(1<<ProtobjParserARRAY)|(1<<ProtobjParserDEFAULT))) != 0) || _la == ProtobjParserIDENTIFIER {
		{
			p.SetState(239)
			p.Ident()
		}
		{
			p.SetState(240)
			p.Match(ProtobjParserCOLON)
		}
		{
			p.SetState(241)
			p.Constant()
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(248)
		p.Match(ProtobjParserRC)
	}

	return localctx
}

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_emptyStatement_
	return p
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (p *ProtobjParser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	this := p
	_ = this

	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ProtobjParserRULE_emptyStatement_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(ProtobjParserSEMI)
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ProtobjParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (p *ProtobjParser) Ident() (localctx IIdentContext) {
	this := p
	_ = this

	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ProtobjParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProtobjParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(ProtobjParserIDENTIFIER)
		}

	case ProtobjParserIMPORT, ProtobjParserPACKAGE, ProtobjParserREPEATED, ProtobjParserMAP, ProtobjParserBOOL, ProtobjParserI8, ProtobjParserU8, ProtobjParserI16, ProtobjParserU16, ProtobjParserI32, ProtobjParserU32, ProtobjParserS32, ProtobjParserF32, ProtobjParserSF32, ProtobjParserI64, ProtobjParserU64, ProtobjParserS64, ProtobjParserF64, ProtobjParserSF64, ProtobjParserSTRING, ProtobjParserDOUBLE, ProtobjParserFLOAT, ProtobjParserENUM, ProtobjParserMESSAGE, ProtobjParserEXTEND, ProtobjParserSET, ProtobjParserARRAY, ProtobjParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Keywords()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ProtobjParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ProtobjParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *ProtobjParser) FullIdent() (localctx IFullIdentContext) {
	this := p
	_ = this

	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ProtobjParserRULE_fullIdent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Ident()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProtobjParserDOT {
		{
			p.SetState(257)
			p.Match(ProtobjParserDOT)
		}
		{
			p.SetState(258)
			p.Ident()
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_messageName
	return p
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (p *ProtobjParser) MessageName() (localctx IMessageNameContext) {
	this := p
	_ = this

	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ProtobjParserRULE_messageName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Ident()
	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (p *ProtobjParser) EnumName() (localctx IEnumNameContext) {
	this := p
	_ = this

	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ProtobjParserRULE_enumName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Ident()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *ProtobjParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ProtobjParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Ident()
	}

	return localctx
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_mapName
	return p
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMapName(s)
	}
}

func (p *ProtobjParser) MapName() (localctx IMapNameContext) {
	this := p
	_ = this

	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ProtobjParserRULE_mapName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Ident()
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ProtobjParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ProtobjParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *ProtobjParser) MessageType() (localctx IMessageTypeContext) {
	this := p
	_ = this

	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ProtobjParserRULE_messageType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserDOT {
		{
			p.SetState(272)
			p.Match(ProtobjParserDOT)
		}

	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(275)
				p.Ident()
			}
			{
				p.SetState(276)
				p.Match(ProtobjParserDOT)
			}

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	{
		p.SetState(283)
		p.MessageName()
	}

	return localctx
}

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_enumType
	return p
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ProtobjParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ProtobjParserDOT, i)
}

func (s *EnumTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *EnumTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterEnumType(s)
	}
}

func (s *EnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitEnumType(s)
	}
}

func (p *ProtobjParser) EnumType() (localctx IEnumTypeContext) {
	this := p
	_ = this

	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ProtobjParserRULE_enumType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProtobjParserDOT {
		{
			p.SetState(285)
			p.Match(ProtobjParserDOT)
		}

	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(288)
				p.Ident()
			}
			{
				p.SetState(289)
				p.Match(ProtobjParserDOT)
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}
	{
		p.SetState(296)
		p.EnumName()
	}

	return localctx
}

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_intLit
	return p
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitIntLit(s)
	}
}

func (p *ProtobjParser) IntLit() (localctx IIntLitContext) {
	this := p
	_ = this

	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ProtobjParserRULE_intLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(ProtobjParserINT_LIT)
	}

	return localctx
}

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_strLit
	return p
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSTR_LIT, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterStrLit(s)
	}
}

func (s *StrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitStrLit(s)
	}
}

func (p *ProtobjParser) StrLit() (localctx IStrLitContext) {
	this := p
	_ = this

	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ProtobjParserRULE_strLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(ProtobjParserSTR_LIT)
	}

	return localctx
}

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_boolLit
	return p
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

func (p *ProtobjParser) BoolLit() (localctx IBoolLitContext) {
	this := p
	_ = this

	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ProtobjParserRULE_boolLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(ProtobjParserBOOL_LIT)
	}

	return localctx
}

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_floatLit
	return p
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitFloatLit(s)
	}
}

func (p *ProtobjParser) FloatLit() (localctx IFloatLitContext) {
	this := p
	_ = this

	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ProtobjParserRULE_floatLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(ProtobjParserFLOAT_LIT)
	}

	return localctx
}

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProtobjParserRULE_keywords
	return p
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProtobjParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserIMPORT, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(ProtobjParserPACKAGE, 0)
}

func (s *KeywordsContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(ProtobjParserREPEATED, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(ProtobjParserMAP, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ProtobjParserBOOL, 0)
}

func (s *KeywordsContext) I8() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI8, 0)
}

func (s *KeywordsContext) U8() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU8, 0)
}

func (s *KeywordsContext) I16() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI16, 0)
}

func (s *KeywordsContext) U16() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU16, 0)
}

func (s *KeywordsContext) I32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI32, 0)
}

func (s *KeywordsContext) U32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU32, 0)
}

func (s *KeywordsContext) S32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserS32, 0)
}

func (s *KeywordsContext) F32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserF32, 0)
}

func (s *KeywordsContext) SF32() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSF32, 0)
}

func (s *KeywordsContext) I64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserI64, 0)
}

func (s *KeywordsContext) U64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserU64, 0)
}

func (s *KeywordsContext) S64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserS64, 0)
}

func (s *KeywordsContext) F64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserF64, 0)
}

func (s *KeywordsContext) SF64() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSF64, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ProtobjParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserFLOAT, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ProtobjParserENUM, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(ProtobjParserMESSAGE, 0)
}

func (s *KeywordsContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(ProtobjParserEXTEND, 0)
}

func (s *KeywordsContext) SET() antlr.TerminalNode {
	return s.GetToken(ProtobjParserSET, 0)
}

func (s *KeywordsContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(ProtobjParserARRAY, 0)
}

func (s *KeywordsContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ProtobjParserDEFAULT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProtobjListener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (p *ProtobjParser) Keywords() (localctx IKeywordsContext) {
	this := p
	_ = this

	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ProtobjParserRULE_keywords)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProtobjParserIMPORT)|(1<<ProtobjParserPACKAGE)|(1<<ProtobjParserREPEATED)|(1<<ProtobjParserMAP)|(1<<ProtobjParserBOOL)|(1<<ProtobjParserI8)|(1<<ProtobjParserU8)|(1<<ProtobjParserI16)|(1<<ProtobjParserU16)|(1<<ProtobjParserI32)|(1<<ProtobjParserU32)|(1<<ProtobjParserS32)|(1<<ProtobjParserF32)|(1<<ProtobjParserSF32)|(1<<ProtobjParserI64)|(1<<ProtobjParserU64)|(1<<ProtobjParserS64)|(1<<ProtobjParserF64)|(1<<ProtobjParserSF64)|(1<<ProtobjParserSTRING)|(1<<ProtobjParserDOUBLE)|(1<<ProtobjParserFLOAT)|(1<<ProtobjParserENUM)|(1<<ProtobjParserMESSAGE)|(1<<ProtobjParserEXTEND)|(1<<ProtobjParserSET)|(1<<ProtobjParserARRAY)|(1<<ProtobjParserDEFAULT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
