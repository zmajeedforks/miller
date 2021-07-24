// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

// CharLiteralValue returns the string value of the char literal.
func (t *Token) CharLiteralValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

// Float32Value returns the float32 value of the token or an error if the token literal does not
// denote a valid float32.
func (t *Token) Float32Value() (float32, error) {
	if v, err := strconv.ParseFloat(string(t.Lit), 32); err != nil {
		return 0, err
	} else {
		return float32(v), nil
	}
}

// Float64Value returns the float64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Float64Value() (float64, error) {
	return strconv.ParseFloat(string(t.Lit), 64)
}

// IDValue returns the string representation of an identifier token.
func (t *Token) IDValue() string {
	return string(t.Lit)
}

// Int32Value returns the int32 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int32Value() (int32, error) {
	if v, err := strconv.ParseInt(string(t.Lit), 10, 64); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

// Int64Value returns the int64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int64Value() (int64, error) {
	return strconv.ParseInt(string(t.Lit), 10, 64)
}

// UTF8Rune decodes the UTF8 rune in the token literal. It returns utf8.RuneError if
// the token literal contains an invalid rune.
func (t *Token) UTF8Rune() (rune, error) {
	r, _ := utf8.DecodeRune(t.Lit)
	if r == utf8.RuneError {
		err := fmt.Errorf("Invalid rune")
		return r, err
	}
	return r, nil
}

// StringValue returns the string value of the token literal.
func (t *Token) StringValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"empty",
		";",
		"{",
		"}",
		"=",
		"unset",
		"filter",
		">",
		">>",
		"|",
		"stdout",
		"stderr",
		"print",
		",",
		"printn",
		"eprint",
		"eprintn",
		"dump",
		"edump",
		"tee",
		"emitf",
		"emit",
		"(",
		")",
		"emitp",
		"field_name",
		"$[",
		"]",
		"braced_field_name",
		"$[[",
		"$[[[",
		"full_srec",
		"oosvar_name",
		"@[",
		"braced_oosvar_name",
		"full_oosvar",
		"all",
		"non_sigil_name",
		"arr",
		"bool",
		"float",
		"int",
		"map",
		"num",
		"str",
		"var",
		"||=",
		"^^=",
		"&&=",
		"??=",
		"???=",
		"|=",
		"&=",
		"^=",
		"<<=",
		">>=",
		">>>=",
		"+=",
		".=",
		"-=",
		"*=",
		"/=",
		"//=",
		"%=",
		"**=",
		"?",
		":",
		"||",
		"^^",
		"&&",
		"??",
		"???",
		"=~",
		"!=~",
		"==",
		"!=",
		">=",
		"<",
		"<=",
		"^",
		"&",
		"<<",
		">>>",
		"+",
		"-",
		".+",
		".-",
		".",
		"*",
		"/",
		"//",
		"%",
		".*",
		"./",
		".//",
		"!",
		"~",
		"**",
		"string_literal",
		"regex_case_insensitive",
		"int_literal",
		"float_literal",
		"boolean_literal",
		"null_literal",
		"const_M_PI",
		"const_M_E",
		"panic",
		"[",
		"ctx_IPS",
		"ctx_IFS",
		"ctx_IRS",
		"ctx_OPS",
		"ctx_OFS",
		"ctx_ORS",
		"ctx_FLATSEP",
		"ctx_NF",
		"ctx_NR",
		"ctx_FNR",
		"ctx_FILENAME",
		"ctx_FILENUM",
		"env",
		"[[",
		"[[[",
		"call",
		"begin",
		"end",
		"if",
		"elif",
		"else",
		"while",
		"do",
		"for",
		"in",
		"break",
		"continue",
		"func",
		"subr",
		"return",
	},

	idMap: map[string]Type{
		"INVALID":                0,
		"$":                      1,
		"empty":                  2,
		";":                      3,
		"{":                      4,
		"}":                      5,
		"=":                      6,
		"unset":                  7,
		"filter":                 8,
		">":                      9,
		">>":                     10,
		"|":                      11,
		"stdout":                 12,
		"stderr":                 13,
		"print":                  14,
		",":                      15,
		"printn":                 16,
		"eprint":                 17,
		"eprintn":                18,
		"dump":                   19,
		"edump":                  20,
		"tee":                    21,
		"emitf":                  22,
		"emit":                   23,
		"(":                      24,
		")":                      25,
		"emitp":                  26,
		"field_name":             27,
		"$[":                     28,
		"]":                      29,
		"braced_field_name":      30,
		"$[[":                    31,
		"$[[[":                   32,
		"full_srec":              33,
		"oosvar_name":            34,
		"@[":                     35,
		"braced_oosvar_name":     36,
		"full_oosvar":            37,
		"all":                    38,
		"non_sigil_name":         39,
		"arr":                    40,
		"bool":                   41,
		"float":                  42,
		"int":                    43,
		"map":                    44,
		"num":                    45,
		"str":                    46,
		"var":                    47,
		"||=":                    48,
		"^^=":                    49,
		"&&=":                    50,
		"??=":                    51,
		"???=":                   52,
		"|=":                     53,
		"&=":                     54,
		"^=":                     55,
		"<<=":                    56,
		">>=":                    57,
		">>>=":                   58,
		"+=":                     59,
		".=":                     60,
		"-=":                     61,
		"*=":                     62,
		"/=":                     63,
		"//=":                    64,
		"%=":                     65,
		"**=":                    66,
		"?":                      67,
		":":                      68,
		"||":                     69,
		"^^":                     70,
		"&&":                     71,
		"??":                     72,
		"???":                    73,
		"=~":                     74,
		"!=~":                    75,
		"==":                     76,
		"!=":                     77,
		">=":                     78,
		"<":                      79,
		"<=":                     80,
		"^":                      81,
		"&":                      82,
		"<<":                     83,
		">>>":                    84,
		"+":                      85,
		"-":                      86,
		".+":                     87,
		".-":                     88,
		".":                      89,
		"*":                      90,
		"/":                      91,
		"//":                     92,
		"%":                      93,
		".*":                     94,
		"./":                     95,
		".//":                    96,
		"!":                      97,
		"~":                      98,
		"**":                     99,
		"string_literal":         100,
		"regex_case_insensitive": 101,
		"int_literal":            102,
		"float_literal":          103,
		"boolean_literal":        104,
		"null_literal":           105,
		"const_M_PI":             106,
		"const_M_E":              107,
		"panic":                  108,
		"[":                      109,
		"ctx_IPS":                110,
		"ctx_IFS":                111,
		"ctx_IRS":                112,
		"ctx_OPS":                113,
		"ctx_OFS":                114,
		"ctx_ORS":                115,
		"ctx_FLATSEP":            116,
		"ctx_NF":                 117,
		"ctx_NR":                 118,
		"ctx_FNR":                119,
		"ctx_FILENAME":           120,
		"ctx_FILENUM":            121,
		"env":                    122,
		"[[":                     123,
		"[[[":                    124,
		"call":                   125,
		"begin":                  126,
		"end":                    127,
		"if":                     128,
		"elif":                   129,
		"else":                   130,
		"while":                  131,
		"do":                     132,
		"for":                    133,
		"in":                     134,
		"break":                  135,
		"continue":               136,
		"func":                   137,
		"subr":                   138,
		"return":                 139,
	},
}
