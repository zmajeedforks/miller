package lib

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Miller regexes use a final 'i' to indicate case-insensitivity; Go regexes
// use an initial "(?i)".  Also (TODO) I need to find all the right things to
// backslash-escape in Go.
//
// * If the regex_string is of the form a.*b, compiles it case-sensisitively.
// * If the regex_string is of the form "a.*b", compiles a.*b case-sensisitively.
// * If the regex_string is of the form "a.*b"i, compiles a.*b case-insensitively.
func CompileMillerRegex(regexString string) (*regexp.Regexp, error) {
	if !strings.HasPrefix(regexString, "\"") {
		return regexp.Compile(regexString)
	} else {
		n := len(regexString)
		if n < 2 {
			// This means the user entered "\"" which in the parser-to-AST was
			// presented as " which we need to handle as a single-character
			// regex.
			return regexp.Compile(regexString)
		}
		if strings.HasSuffix(regexString, "\"") {
			// TODO: rethink this. This will strip out things people have entered, e.g. "\"...\"".
			// The parser-to-AST will have stripped the outer and we'll strip the inner and the
			// user's intent will be lost.
			//
			// TODO: make separate functions for calling from parser-to-AST (string
			// literals) and from verbs (like cut -r or having-fields).
			return regexp.Compile(regexString[1 : n-1])
		} else if strings.HasSuffix(regexString, "\"i") {
			return regexp.Compile("(?i)" + regexString[1:n-2])
		} else {
			// The user can enter things like "\".." which comes in as ".. which is fine.
			return regexp.Compile(regexString)
		}
	}
}

func CompileMillerRegexOrDie(regexString string) *regexp.Regexp {
	regex, err := CompileMillerRegex(regexString)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return regex
}

func RegexReplaceOnce(
	regex *regexp.Regexp,
	input string,
	replacement string,
) string {
	onFirst := true
	output := regex.ReplaceAllStringFunc(input, func(s string) string {
		if !onFirst {
			return s
		}
		onFirst = false
		return regex.ReplaceAllString(s, replacement)
	})
	return output
}

// xxx
// echo a=ab_cde | mlrgo --oxtab put '
//   $b = sub($a, "(..)_(...)", "\2-\1");
//   $c = sub($a, "(..)_(.)(..)", ":\1:\2:\3")
// '

//// ----------------------------------------------------------------
//func luminLine(regex *regexp.Regexp, input string) string {
//
//	matrix := regex.FindAllStringIndex(input, -1)
//	// fmt.Printf("%+v\n", matrix)
//	if matrix == nil || len(matrix) == 0 {
//		return input
//	}
//
//	// The key is the Go library's regex.FindAllStringIndex.  It gives us start
//	// (inclusive) and end (exclusive) indices for matches.
//	//
//	// Example: for pattern "foo" and input "abc foo def foo ghi" we'll have
//	// matrix [[4 7] [12 15]] which indicates matches from positions 4-6 and
//	// 12-14.  We simply need to print out:
//	// *  0-3  "abc "  with default color
//	// *  4-6  "foo"   with highlight color
//	// *  7-11 " def " with default color
//	// * 12-14 "foo"   with highlight color
//	// * 15-18 " ghi"  with default color.
//	//
//	// Example: with pattern "f.*o" and input "abc foo def foo ghi" we'll have
//	// matrix [[4 15]] so "foo def foo" will be highlighted.
//
//	var buffer bytes.Buffer // Faster since os.Stdout is unbuffered
//	nonMatchStartIndex := 0
//
//	for _, startEnd := range matrix {
//		buffer.WriteString(input[nonMatchStartIndex:startEnd[0]])
//		buffer.WriteString(colors.Colorize(input[startEnd[0]:startEnd[1]]))
//		nonMatchStartIndex = startEnd[1]
//	}
//
//	buffer.WriteString(input[nonMatchStartIndex:])
//
//	return buffer.String()
//}
