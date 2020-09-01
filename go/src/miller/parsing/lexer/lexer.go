// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"miller/parsing/token"
)

const (
	NoState    = -1
	NumStates  = 125
	NumSymbols = 211
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '"'
1: '"'
2: '0'
3: 'x'
4: '.'
5: '-'
6: '.'
7: '.'
8: '-'
9: '.'
10: '.'
11: '-'
12: 'I'
13: 'P'
14: 'S'
15: 'I'
16: 'F'
17: 'S'
18: 'I'
19: 'R'
20: 'S'
21: 'O'
22: 'P'
23: 'S'
24: 'O'
25: 'F'
26: 'S'
27: 'O'
28: 'R'
29: 'S'
30: 'N'
31: 'F'
32: 'N'
33: 'R'
34: 'F'
35: 'N'
36: 'R'
37: 'F'
38: 'I'
39: 'L'
40: 'E'
41: 'N'
42: 'A'
43: 'M'
44: 'E'
45: 'F'
46: 'I'
47: 'L'
48: 'E'
49: 'N'
50: 'U'
51: 'M'
52: '$'
53: ';'
54: '='
55: '|'
56: '|'
57: '='
58: '^'
59: '^'
60: '='
61: '&'
62: '&'
63: '='
64: '|'
65: '='
66: '^'
67: '='
68: '&'
69: '='
70: '<'
71: '<'
72: '='
73: '>'
74: '>'
75: '='
76: '+'
77: '='
78: '.'
79: '='
80: '-'
81: '='
82: '*'
83: '='
84: '/'
85: '='
86: '/'
87: '/'
88: '='
89: '%'
90: '='
91: '*'
92: '*'
93: '='
94: '?'
95: ':'
96: '|'
97: '|'
98: '^'
99: '^'
100: '&'
101: '&'
102: '='
103: '~'
104: '!'
105: '='
106: '~'
107: '='
108: '='
109: '!'
110: '='
111: '>'
112: '>'
113: '='
114: '<'
115: '<'
116: '='
117: '|'
118: '^'
119: '&'
120: '<'
121: '<'
122: '>'
123: '>'
124: '+'
125: '-'
126: '.'
127: '+'
128: '.'
129: '-'
130: '.'
131: '*'
132: '/'
133: '/'
134: '/'
135: '%'
136: '.'
137: '*'
138: '.'
139: '/'
140: '.'
141: '/'
142: '/'
143: '!'
144: '~'
145: '*'
146: '*'
147: '('
148: ')'
149: '$'
150: '['
151: ']'
152: '_'
153: ' '
154: '!'
155: '#'
156: '$'
157: '%'
158: '&'
159: '''
160: '\'
161: '('
162: ')'
163: '*'
164: '+'
165: ','
166: '-'
167: '.'
168: '/'
169: ':'
170: ';'
171: '<'
172: '='
173: '>'
174: '?'
175: '@'
176: '['
177: ']'
178: '^'
179: '_'
180: '`'
181: '{'
182: '|'
183: '}'
184: '~'
185: 'e'
186: 'E'
187: 't'
188: 'r'
189: 'u'
190: 'e'
191: 'f'
192: 'a'
193: 'l'
194: 's'
195: 'e'
196: ' '
197: '\t'
198: '\n'
199: '\r'
200: 'a'-'z'
201: 'A'-'Z'
202: '0'-'9'
203: '0'-'9'
204: 'a'-'f'
205: 'A'-'F'
206: 'A'-'Z'
207: 'a'-'z'
208: '0'-'9'
209: \u0100-\U0010ffff
210: .
*/