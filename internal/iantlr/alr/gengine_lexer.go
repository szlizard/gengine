// Code generated from gengine.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GengineLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func genginelexerLexerInit() {
	staticData := &GengineLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'@name'", "'@id'", "'@desc'", "'@sal'", "", "", "'&&'",
		"'||'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "'+'", "'-'", "'/'", "'*'", "'=='", "'>'", "'<'", "'>='", "'<='",
		"'!='", "'!'", "':='", "'='", "'+='", "'-='", "'*='", "'/='", "'['",
		"']'", "';'", "'{'", "'}'", "'('", "')'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE",
		"RETURN", "FOR", "BREAK", "FORRANGE", "CONTINUE", "TRUE", "FALSE", "NULL_LITERAL",
		"SALIENCE", "BEGIN", "END", "IN", "SIMPLENAME", "INT", "PLUS", "MINUS",
		"DIV", "MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT",
		"ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL",
		"LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET",
		"RR_BRACKET", "DOT", "DQUOTA_STRING", "BACKQUOTA_STRING", "DOTTEDNAME",
		"DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "DEC_DIGIT", "A", "B", "C",
		"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q",
		"R", "S", "T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "HEX_DIGIT",
		"OCTAL_DIGIT", "ESCAPED_VALUE", "NIL", "RULE", "AND", "OR", "CONC",
		"IF", "ELSE", "RETURN", "FOR", "BREAK", "FORRANGE", "CONTINUE", "TRUE",
		"FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", "IN", "SIMPLENAME",
		"INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", "GTE", "LTE",
		"NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL",
		"DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE",
		"LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "BACKQUOTA_STRING",
		"DOTTEDNAME", "DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 58, 567, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		32, 1, 32, 3, 32, 259, 8, 32, 1, 32, 4, 32, 262, 8, 32, 11, 32, 12, 32,
		263, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3,
		35, 296, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 4,
		55, 400, 8, 55, 11, 55, 12, 55, 401, 1, 55, 5, 55, 405, 8, 55, 10, 55,
		12, 55, 408, 9, 55, 1, 56, 4, 56, 411, 8, 56, 11, 56, 12, 56, 412, 1, 57,
		1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1,
		62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 66,
		1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 1,
		70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73,
		1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77, 1, 78, 1, 78, 1,
		79, 1, 79, 1, 80, 1, 80, 1, 81, 1, 81, 1, 82, 1, 82, 1, 82, 5, 82, 477,
		8, 82, 10, 82, 12, 82, 480, 9, 82, 1, 82, 1, 82, 1, 83, 1, 83, 5, 83, 486,
		8, 83, 10, 83, 12, 83, 489, 9, 83, 1, 83, 1, 83, 1, 84, 1, 84, 1, 84, 1,
		84, 1, 85, 1, 85, 1, 85, 1, 85, 1, 85, 1, 85, 1, 86, 4, 86, 504, 8, 86,
		11, 86, 12, 86, 505, 3, 86, 508, 8, 86, 1, 86, 1, 86, 4, 86, 512, 8, 86,
		11, 86, 12, 86, 513, 1, 86, 4, 86, 517, 8, 86, 11, 86, 12, 86, 518, 1,
		86, 1, 86, 1, 86, 1, 86, 4, 86, 525, 8, 86, 11, 86, 12, 86, 526, 3, 86,
		529, 8, 86, 1, 86, 1, 86, 4, 86, 533, 8, 86, 11, 86, 12, 86, 534, 1, 86,
		1, 86, 1, 86, 4, 86, 540, 8, 86, 11, 86, 12, 86, 541, 1, 86, 1, 86, 3,
		86, 546, 8, 86, 1, 87, 1, 87, 1, 87, 1, 87, 5, 87, 552, 8, 87, 10, 87,
		12, 87, 555, 9, 87, 1, 87, 1, 87, 1, 87, 1, 87, 1, 88, 4, 88, 562, 8, 88,
		11, 88, 12, 88, 563, 1, 88, 1, 88, 2, 487, 553, 0, 89, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27,
		0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0,
		49, 0, 51, 0, 53, 0, 55, 0, 57, 0, 59, 0, 61, 0, 63, 0, 65, 0, 67, 0, 69,
		0, 71, 0, 73, 6, 75, 7, 77, 8, 79, 9, 81, 10, 83, 11, 85, 12, 87, 13, 89,
		14, 91, 15, 93, 16, 95, 17, 97, 18, 99, 19, 101, 20, 103, 21, 105, 22,
		107, 23, 109, 24, 111, 25, 113, 26, 115, 27, 117, 28, 119, 29, 121, 30,
		123, 31, 125, 32, 127, 33, 129, 34, 131, 35, 133, 36, 135, 37, 137, 38,
		139, 39, 141, 40, 143, 41, 145, 42, 147, 43, 149, 44, 151, 45, 153, 46,
		155, 47, 157, 48, 159, 49, 161, 50, 163, 51, 165, 52, 167, 53, 169, 54,
		171, 55, 173, 56, 175, 57, 177, 58, 1, 0, 34, 1, 0, 48, 57, 2, 0, 65, 65,
		97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100,
		100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103,
		103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106,
		106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109,
		109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112,
		112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115,
		115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118,
		118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121,
		121, 2, 0, 90, 90, 122, 122, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 55,
		9, 0, 34, 34, 39, 39, 92, 92, 97, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 118, 118, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13,
		32, 32, 560, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0,
		0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107,
		1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0,
		0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1,
		0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0,
		129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0,
		0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143,
		1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0,
		0, 151, 1, 0, 0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1,
		0, 0, 0, 0, 159, 1, 0, 0, 0, 0, 161, 1, 0, 0, 0, 0, 163, 1, 0, 0, 0, 0,
		165, 1, 0, 0, 0, 0, 167, 1, 0, 0, 0, 0, 169, 1, 0, 0, 0, 0, 171, 1, 0,
		0, 0, 0, 173, 1, 0, 0, 0, 0, 175, 1, 0, 0, 0, 0, 177, 1, 0, 0, 0, 1, 179,
		1, 0, 0, 0, 3, 181, 1, 0, 0, 0, 5, 187, 1, 0, 0, 0, 7, 191, 1, 0, 0, 0,
		9, 197, 1, 0, 0, 0, 11, 202, 1, 0, 0, 0, 13, 204, 1, 0, 0, 0, 15, 206,
		1, 0, 0, 0, 17, 208, 1, 0, 0, 0, 19, 210, 1, 0, 0, 0, 21, 212, 1, 0, 0,
		0, 23, 214, 1, 0, 0, 0, 25, 216, 1, 0, 0, 0, 27, 218, 1, 0, 0, 0, 29, 220,
		1, 0, 0, 0, 31, 222, 1, 0, 0, 0, 33, 224, 1, 0, 0, 0, 35, 226, 1, 0, 0,
		0, 37, 228, 1, 0, 0, 0, 39, 230, 1, 0, 0, 0, 41, 232, 1, 0, 0, 0, 43, 234,
		1, 0, 0, 0, 45, 236, 1, 0, 0, 0, 47, 238, 1, 0, 0, 0, 49, 240, 1, 0, 0,
		0, 51, 242, 1, 0, 0, 0, 53, 244, 1, 0, 0, 0, 55, 246, 1, 0, 0, 0, 57, 248,
		1, 0, 0, 0, 59, 250, 1, 0, 0, 0, 61, 252, 1, 0, 0, 0, 63, 254, 1, 0, 0,
		0, 65, 256, 1, 0, 0, 0, 67, 265, 1, 0, 0, 0, 69, 267, 1, 0, 0, 0, 71, 269,
		1, 0, 0, 0, 73, 297, 1, 0, 0, 0, 75, 301, 1, 0, 0, 0, 77, 306, 1, 0, 0,
		0, 79, 309, 1, 0, 0, 0, 81, 312, 1, 0, 0, 0, 83, 317, 1, 0, 0, 0, 85, 320,
		1, 0, 0, 0, 87, 325, 1, 0, 0, 0, 89, 332, 1, 0, 0, 0, 91, 336, 1, 0, 0,
		0, 93, 342, 1, 0, 0, 0, 95, 351, 1, 0, 0, 0, 97, 360, 1, 0, 0, 0, 99, 365,
		1, 0, 0, 0, 101, 371, 1, 0, 0, 0, 103, 376, 1, 0, 0, 0, 105, 385, 1, 0,
		0, 0, 107, 391, 1, 0, 0, 0, 109, 395, 1, 0, 0, 0, 111, 399, 1, 0, 0, 0,
		113, 410, 1, 0, 0, 0, 115, 414, 1, 0, 0, 0, 117, 416, 1, 0, 0, 0, 119,
		418, 1, 0, 0, 0, 121, 420, 1, 0, 0, 0, 123, 422, 1, 0, 0, 0, 125, 425,
		1, 0, 0, 0, 127, 427, 1, 0, 0, 0, 129, 429, 1, 0, 0, 0, 131, 432, 1, 0,
		0, 0, 133, 435, 1, 0, 0, 0, 135, 438, 1, 0, 0, 0, 137, 440, 1, 0, 0, 0,
		139, 443, 1, 0, 0, 0, 141, 445, 1, 0, 0, 0, 143, 448, 1, 0, 0, 0, 145,
		451, 1, 0, 0, 0, 147, 454, 1, 0, 0, 0, 149, 457, 1, 0, 0, 0, 151, 459,
		1, 0, 0, 0, 153, 461, 1, 0, 0, 0, 155, 463, 1, 0, 0, 0, 157, 465, 1, 0,
		0, 0, 159, 467, 1, 0, 0, 0, 161, 469, 1, 0, 0, 0, 163, 471, 1, 0, 0, 0,
		165, 473, 1, 0, 0, 0, 167, 483, 1, 0, 0, 0, 169, 492, 1, 0, 0, 0, 171,
		496, 1, 0, 0, 0, 173, 545, 1, 0, 0, 0, 175, 547, 1, 0, 0, 0, 177, 561,
		1, 0, 0, 0, 179, 180, 5, 44, 0, 0, 180, 2, 1, 0, 0, 0, 181, 182, 5, 64,
		0, 0, 182, 183, 5, 110, 0, 0, 183, 184, 5, 97, 0, 0, 184, 185, 5, 109,
		0, 0, 185, 186, 5, 101, 0, 0, 186, 4, 1, 0, 0, 0, 187, 188, 5, 64, 0, 0,
		188, 189, 5, 105, 0, 0, 189, 190, 5, 100, 0, 0, 190, 6, 1, 0, 0, 0, 191,
		192, 5, 64, 0, 0, 192, 193, 5, 100, 0, 0, 193, 194, 5, 101, 0, 0, 194,
		195, 5, 115, 0, 0, 195, 196, 5, 99, 0, 0, 196, 8, 1, 0, 0, 0, 197, 198,
		5, 64, 0, 0, 198, 199, 5, 115, 0, 0, 199, 200, 5, 97, 0, 0, 200, 201, 5,
		108, 0, 0, 201, 10, 1, 0, 0, 0, 202, 203, 7, 0, 0, 0, 203, 12, 1, 0, 0,
		0, 204, 205, 7, 1, 0, 0, 205, 14, 1, 0, 0, 0, 206, 207, 7, 2, 0, 0, 207,
		16, 1, 0, 0, 0, 208, 209, 7, 3, 0, 0, 209, 18, 1, 0, 0, 0, 210, 211, 7,
		4, 0, 0, 211, 20, 1, 0, 0, 0, 212, 213, 7, 5, 0, 0, 213, 22, 1, 0, 0, 0,
		214, 215, 7, 6, 0, 0, 215, 24, 1, 0, 0, 0, 216, 217, 7, 7, 0, 0, 217, 26,
		1, 0, 0, 0, 218, 219, 7, 8, 0, 0, 219, 28, 1, 0, 0, 0, 220, 221, 7, 9,
		0, 0, 221, 30, 1, 0, 0, 0, 222, 223, 7, 10, 0, 0, 223, 32, 1, 0, 0, 0,
		224, 225, 7, 11, 0, 0, 225, 34, 1, 0, 0, 0, 226, 227, 7, 12, 0, 0, 227,
		36, 1, 0, 0, 0, 228, 229, 7, 13, 0, 0, 229, 38, 1, 0, 0, 0, 230, 231, 7,
		14, 0, 0, 231, 40, 1, 0, 0, 0, 232, 233, 7, 15, 0, 0, 233, 42, 1, 0, 0,
		0, 234, 235, 7, 16, 0, 0, 235, 44, 1, 0, 0, 0, 236, 237, 7, 17, 0, 0, 237,
		46, 1, 0, 0, 0, 238, 239, 7, 18, 0, 0, 239, 48, 1, 0, 0, 0, 240, 241, 7,
		19, 0, 0, 241, 50, 1, 0, 0, 0, 242, 243, 7, 20, 0, 0, 243, 52, 1, 0, 0,
		0, 244, 245, 7, 21, 0, 0, 245, 54, 1, 0, 0, 0, 246, 247, 7, 22, 0, 0, 247,
		56, 1, 0, 0, 0, 248, 249, 7, 23, 0, 0, 249, 58, 1, 0, 0, 0, 250, 251, 7,
		24, 0, 0, 251, 60, 1, 0, 0, 0, 252, 253, 7, 25, 0, 0, 253, 62, 1, 0, 0,
		0, 254, 255, 7, 26, 0, 0, 255, 64, 1, 0, 0, 0, 256, 258, 7, 5, 0, 0, 257,
		259, 5, 45, 0, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 261,
		1, 0, 0, 0, 260, 262, 3, 11, 5, 0, 261, 260, 1, 0, 0, 0, 262, 263, 1, 0,
		0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 66, 1, 0, 0, 0,
		265, 266, 7, 27, 0, 0, 266, 68, 1, 0, 0, 0, 267, 268, 7, 28, 0, 0, 268,
		70, 1, 0, 0, 0, 269, 295, 5, 92, 0, 0, 270, 271, 5, 117, 0, 0, 271, 272,
		3, 67, 33, 0, 272, 273, 3, 67, 33, 0, 273, 274, 3, 67, 33, 0, 274, 275,
		3, 67, 33, 0, 275, 296, 1, 0, 0, 0, 276, 277, 5, 85, 0, 0, 277, 278, 3,
		67, 33, 0, 278, 279, 3, 67, 33, 0, 279, 280, 3, 67, 33, 0, 280, 281, 3,
		67, 33, 0, 281, 282, 3, 67, 33, 0, 282, 283, 3, 67, 33, 0, 283, 284, 3,
		67, 33, 0, 284, 285, 3, 67, 33, 0, 285, 296, 1, 0, 0, 0, 286, 296, 7, 29,
		0, 0, 287, 288, 3, 69, 34, 0, 288, 289, 3, 69, 34, 0, 289, 290, 3, 69,
		34, 0, 290, 296, 1, 0, 0, 0, 291, 292, 5, 120, 0, 0, 292, 293, 3, 67, 33,
		0, 293, 294, 3, 67, 33, 0, 294, 296, 1, 0, 0, 0, 295, 270, 1, 0, 0, 0,
		295, 276, 1, 0, 0, 0, 295, 286, 1, 0, 0, 0, 295, 287, 1, 0, 0, 0, 295,
		291, 1, 0, 0, 0, 296, 72, 1, 0, 0, 0, 297, 298, 3, 39, 19, 0, 298, 299,
		3, 29, 14, 0, 299, 300, 3, 35, 17, 0, 300, 74, 1, 0, 0, 0, 301, 302, 3,
		47, 23, 0, 302, 303, 3, 53, 26, 0, 303, 304, 3, 35, 17, 0, 304, 305, 3,
		21, 10, 0, 305, 76, 1, 0, 0, 0, 306, 307, 5, 38, 0, 0, 307, 308, 5, 38,
		0, 0, 308, 78, 1, 0, 0, 0, 309, 310, 5, 124, 0, 0, 310, 311, 5, 124, 0,
		0, 311, 80, 1, 0, 0, 0, 312, 313, 3, 17, 8, 0, 313, 314, 3, 41, 20, 0,
		314, 315, 3, 39, 19, 0, 315, 316, 3, 17, 8, 0, 316, 82, 1, 0, 0, 0, 317,
		318, 3, 29, 14, 0, 318, 319, 3, 23, 11, 0, 319, 84, 1, 0, 0, 0, 320, 321,
		3, 21, 10, 0, 321, 322, 3, 35, 17, 0, 322, 323, 3, 49, 24, 0, 323, 324,
		3, 21, 10, 0, 324, 86, 1, 0, 0, 0, 325, 326, 3, 47, 23, 0, 326, 327, 3,
		21, 10, 0, 327, 328, 3, 51, 25, 0, 328, 329, 3, 53, 26, 0, 329, 330, 3,
		47, 23, 0, 330, 331, 3, 39, 19, 0, 331, 88, 1, 0, 0, 0, 332, 333, 3, 23,
		11, 0, 333, 334, 3, 41, 20, 0, 334, 335, 3, 47, 23, 0, 335, 90, 1, 0, 0,
		0, 336, 337, 3, 15, 7, 0, 337, 338, 3, 47, 23, 0, 338, 339, 3, 21, 10,
		0, 339, 340, 3, 13, 6, 0, 340, 341, 3, 33, 16, 0, 341, 92, 1, 0, 0, 0,
		342, 343, 3, 23, 11, 0, 343, 344, 3, 41, 20, 0, 344, 345, 3, 47, 23, 0,
		345, 346, 3, 47, 23, 0, 346, 347, 3, 13, 6, 0, 347, 348, 3, 39, 19, 0,
		348, 349, 3, 25, 12, 0, 349, 350, 3, 21, 10, 0, 350, 94, 1, 0, 0, 0, 351,
		352, 3, 17, 8, 0, 352, 353, 3, 41, 20, 0, 353, 354, 3, 39, 19, 0, 354,
		355, 3, 51, 25, 0, 355, 356, 3, 29, 14, 0, 356, 357, 3, 39, 19, 0, 357,
		358, 3, 53, 26, 0, 358, 359, 3, 21, 10, 0, 359, 96, 1, 0, 0, 0, 360, 361,
		3, 51, 25, 0, 361, 362, 3, 47, 23, 0, 362, 363, 3, 53, 26, 0, 363, 364,
		3, 21, 10, 0, 364, 98, 1, 0, 0, 0, 365, 366, 3, 23, 11, 0, 366, 367, 3,
		13, 6, 0, 367, 368, 3, 35, 17, 0, 368, 369, 3, 49, 24, 0, 369, 370, 3,
		21, 10, 0, 370, 100, 1, 0, 0, 0, 371, 372, 3, 39, 19, 0, 372, 373, 3, 53,
		26, 0, 373, 374, 3, 35, 17, 0, 374, 375, 3, 35, 17, 0, 375, 102, 1, 0,
		0, 0, 376, 377, 3, 49, 24, 0, 377, 378, 3, 13, 6, 0, 378, 379, 3, 35, 17,
		0, 379, 380, 3, 29, 14, 0, 380, 381, 3, 21, 10, 0, 381, 382, 3, 39, 19,
		0, 382, 383, 3, 17, 8, 0, 383, 384, 3, 21, 10, 0, 384, 104, 1, 0, 0, 0,
		385, 386, 3, 15, 7, 0, 386, 387, 3, 21, 10, 0, 387, 388, 3, 25, 12, 0,
		388, 389, 3, 29, 14, 0, 389, 390, 3, 39, 19, 0, 390, 106, 1, 0, 0, 0, 391,
		392, 3, 21, 10, 0, 392, 393, 3, 39, 19, 0, 393, 394, 3, 19, 9, 0, 394,
		108, 1, 0, 0, 0, 395, 396, 3, 29, 14, 0, 396, 397, 3, 39, 19, 0, 397, 110,
		1, 0, 0, 0, 398, 400, 7, 30, 0, 0, 399, 398, 1, 0, 0, 0, 400, 401, 1, 0,
		0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 406, 1, 0, 0, 0,
		403, 405, 7, 31, 0, 0, 404, 403, 1, 0, 0, 0, 405, 408, 1, 0, 0, 0, 406,
		404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 112, 1, 0, 0, 0, 408, 406,
		1, 0, 0, 0, 409, 411, 2, 48, 57, 0, 410, 409, 1, 0, 0, 0, 411, 412, 1,
		0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 114, 1, 0, 0,
		0, 414, 415, 5, 43, 0, 0, 415, 116, 1, 0, 0, 0, 416, 417, 5, 45, 0, 0,
		417, 118, 1, 0, 0, 0, 418, 419, 5, 47, 0, 0, 419, 120, 1, 0, 0, 0, 420,
		421, 5, 42, 0, 0, 421, 122, 1, 0, 0, 0, 422, 423, 5, 61, 0, 0, 423, 424,
		5, 61, 0, 0, 424, 124, 1, 0, 0, 0, 425, 426, 5, 62, 0, 0, 426, 126, 1,
		0, 0, 0, 427, 428, 5, 60, 0, 0, 428, 128, 1, 0, 0, 0, 429, 430, 5, 62,
		0, 0, 430, 431, 5, 61, 0, 0, 431, 130, 1, 0, 0, 0, 432, 433, 5, 60, 0,
		0, 433, 434, 5, 61, 0, 0, 434, 132, 1, 0, 0, 0, 435, 436, 5, 33, 0, 0,
		436, 437, 5, 61, 0, 0, 437, 134, 1, 0, 0, 0, 438, 439, 5, 33, 0, 0, 439,
		136, 1, 0, 0, 0, 440, 441, 5, 58, 0, 0, 441, 442, 5, 61, 0, 0, 442, 138,
		1, 0, 0, 0, 443, 444, 5, 61, 0, 0, 444, 140, 1, 0, 0, 0, 445, 446, 5, 43,
		0, 0, 446, 447, 5, 61, 0, 0, 447, 142, 1, 0, 0, 0, 448, 449, 5, 45, 0,
		0, 449, 450, 5, 61, 0, 0, 450, 144, 1, 0, 0, 0, 451, 452, 5, 42, 0, 0,
		452, 453, 5, 61, 0, 0, 453, 146, 1, 0, 0, 0, 454, 455, 5, 47, 0, 0, 455,
		456, 5, 61, 0, 0, 456, 148, 1, 0, 0, 0, 457, 458, 5, 91, 0, 0, 458, 150,
		1, 0, 0, 0, 459, 460, 5, 93, 0, 0, 460, 152, 1, 0, 0, 0, 461, 462, 5, 59,
		0, 0, 462, 154, 1, 0, 0, 0, 463, 464, 5, 123, 0, 0, 464, 156, 1, 0, 0,
		0, 465, 466, 5, 125, 0, 0, 466, 158, 1, 0, 0, 0, 467, 468, 5, 40, 0, 0,
		468, 160, 1, 0, 0, 0, 469, 470, 5, 41, 0, 0, 470, 162, 1, 0, 0, 0, 471,
		472, 5, 46, 0, 0, 472, 164, 1, 0, 0, 0, 473, 478, 5, 34, 0, 0, 474, 477,
		8, 32, 0, 0, 475, 477, 3, 71, 35, 0, 476, 474, 1, 0, 0, 0, 476, 475, 1,
		0, 0, 0, 477, 480, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 478, 479, 1, 0, 0,
		0, 479, 481, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 481, 482, 5, 34, 0, 0, 482,
		166, 1, 0, 0, 0, 483, 487, 5, 96, 0, 0, 484, 486, 9, 0, 0, 0, 485, 484,
		1, 0, 0, 0, 486, 489, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 487, 485, 1, 0,
		0, 0, 488, 490, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 491, 5, 96, 0, 0,
		491, 168, 1, 0, 0, 0, 492, 493, 3, 111, 55, 0, 493, 494, 3, 163, 81, 0,
		494, 495, 3, 111, 55, 0, 495, 170, 1, 0, 0, 0, 496, 497, 3, 111, 55, 0,
		497, 498, 3, 163, 81, 0, 498, 499, 3, 111, 55, 0, 499, 500, 3, 163, 81,
		0, 500, 501, 3, 111, 55, 0, 501, 172, 1, 0, 0, 0, 502, 504, 3, 11, 5, 0,
		503, 502, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505,
		506, 1, 0, 0, 0, 506, 508, 1, 0, 0, 0, 507, 503, 1, 0, 0, 0, 507, 508,
		1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 511, 5, 46, 0, 0, 510, 512, 3, 11,
		5, 0, 511, 510, 1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 511, 1, 0, 0, 0,
		513, 514, 1, 0, 0, 0, 514, 546, 1, 0, 0, 0, 515, 517, 3, 11, 5, 0, 516,
		515, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 516, 1, 0, 0, 0, 518, 519,
		1, 0, 0, 0, 519, 520, 1, 0, 0, 0, 520, 521, 5, 46, 0, 0, 521, 522, 3, 65,
		32, 0, 522, 546, 1, 0, 0, 0, 523, 525, 3, 11, 5, 0, 524, 523, 1, 0, 0,
		0, 525, 526, 1, 0, 0, 0, 526, 524, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527,
		529, 1, 0, 0, 0, 528, 524, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 530,
		1, 0, 0, 0, 530, 532, 5, 46, 0, 0, 531, 533, 3, 11, 5, 0, 532, 531, 1,
		0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 534, 535, 1, 0, 0,
		0, 535, 536, 1, 0, 0, 0, 536, 537, 3, 65, 32, 0, 537, 546, 1, 0, 0, 0,
		538, 540, 3, 11, 5, 0, 539, 538, 1, 0, 0, 0, 540, 541, 1, 0, 0, 0, 541,
		539, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 544,
		3, 65, 32, 0, 544, 546, 1, 0, 0, 0, 545, 507, 1, 0, 0, 0, 545, 516, 1,
		0, 0, 0, 545, 528, 1, 0, 0, 0, 545, 539, 1, 0, 0, 0, 546, 174, 1, 0, 0,
		0, 547, 548, 5, 47, 0, 0, 548, 549, 5, 47, 0, 0, 549, 553, 1, 0, 0, 0,
		550, 552, 9, 0, 0, 0, 551, 550, 1, 0, 0, 0, 552, 555, 1, 0, 0, 0, 553,
		554, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0, 554, 556, 1, 0, 0, 0, 555, 553,
		1, 0, 0, 0, 556, 557, 5, 10, 0, 0, 557, 558, 1, 0, 0, 0, 558, 559, 6, 87,
		0, 0, 559, 176, 1, 0, 0, 0, 560, 562, 7, 33, 0, 0, 561, 560, 1, 0, 0, 0,
		562, 563, 1, 0, 0, 0, 563, 561, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564,
		565, 1, 0, 0, 0, 565, 566, 6, 88, 0, 0, 566, 178, 1, 0, 0, 0, 22, 0, 258,
		263, 295, 401, 404, 406, 412, 476, 478, 487, 505, 507, 513, 518, 526, 528,
		534, 541, 545, 553, 563, 1, 6, 0, 0,
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

// gengineLexerInit initializes any static state used to implement gengineLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewgengineLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GengineLexerInit() {
	staticData := &GengineLexerLexerStaticData
	staticData.once.Do(genginelexerLexerInit)
}

// NewgengineLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewgengineLexer(input antlr.CharStream) *gengineLexer {
	GengineLexerInit()
	l := new(gengineLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GengineLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0             = 1
	gengineLexerT__1             = 2
	gengineLexerT__2             = 3
	gengineLexerT__3             = 4
	gengineLexerT__4             = 5
	gengineLexerNIL              = 6
	gengineLexerRULE             = 7
	gengineLexerAND              = 8
	gengineLexerOR               = 9
	gengineLexerCONC             = 10
	gengineLexerIF               = 11
	gengineLexerELSE             = 12
	gengineLexerRETURN           = 13
	gengineLexerFOR              = 14
	gengineLexerBREAK            = 15
	gengineLexerFORRANGE         = 16
	gengineLexerCONTINUE         = 17
	gengineLexerTRUE             = 18
	gengineLexerFALSE            = 19
	gengineLexerNULL_LITERAL     = 20
	gengineLexerSALIENCE         = 21
	gengineLexerBEGIN            = 22
	gengineLexerEND              = 23
	gengineLexerIN               = 24
	gengineLexerSIMPLENAME       = 25
	gengineLexerINT              = 26
	gengineLexerPLUS             = 27
	gengineLexerMINUS            = 28
	gengineLexerDIV              = 29
	gengineLexerMUL              = 30
	gengineLexerEQUALS           = 31
	gengineLexerGT               = 32
	gengineLexerLT               = 33
	gengineLexerGTE              = 34
	gengineLexerLTE              = 35
	gengineLexerNOTEQUALS        = 36
	gengineLexerNOT              = 37
	gengineLexerASSIGN           = 38
	gengineLexerSET              = 39
	gengineLexerPLUSEQUAL        = 40
	gengineLexerMINUSEQUAL       = 41
	gengineLexerMULTIEQUAL       = 42
	gengineLexerDIVEQUAL         = 43
	gengineLexerLSQARE           = 44
	gengineLexerRSQARE           = 45
	gengineLexerSEMICOLON        = 46
	gengineLexerLR_BRACE         = 47
	gengineLexerRR_BRACE         = 48
	gengineLexerLR_BRACKET       = 49
	gengineLexerRR_BRACKET       = 50
	gengineLexerDOT              = 51
	gengineLexerDQUOTA_STRING    = 52
	gengineLexerBACKQUOTA_STRING = 53
	gengineLexerDOTTEDNAME       = 54
	gengineLexerDOUBLEDOTTEDNAME = 55
	gengineLexerREAL_LITERAL     = 56
	gengineLexerSL_COMMENT       = 57
	gengineLexerWS               = 58
)
