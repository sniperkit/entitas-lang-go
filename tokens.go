package elang

// Token ...
type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE // ' '
	NEWLINE    // '\n' '\r\n'
	PERIOD     // '.'
	COLON      // ':'
	LPAREN     // '('
	RPAREN     // ')'
	COMMA      // ','
	QUOTE      // '"'
	UNDERSCORE // '_'

	WORD
	INTEGER

	KW_TARGET           // 'TARGET'
	KW_NAMESPACE        // 'NAMESPACE'
	KW_CONTEXT          // 'CONTEXT'
	KW_ALIAS            // 'ALIAS'
	KW_COMP             // 'COMP'
	KW_IN               // 'IN'
	KW_DEFAULT          // 'DEFAULT'
	KW_SYS              // 'SYS'
	KW_TRIGGER          // 'TRIGGER'
	KW_ACCESS           // 'ACCESS'
	KW_NO_FILTER        // 'NO_FILTER'
	KW_FILTER           // 'FILTER'
	KW_ADDED            // 'ADDED'
	KW_REMOVED          // 'REMOVED'
	KW_ADDED_OR_REMOVED // 'ADDED_OR_REMOVED'
	KW_ALL_OF           // 'ALL_OF'
	KW_ANY_OF           // 'ANY_OF'
	KW_NONE_OF          // 'NONE_OF'
	KW_INIT             // 'INIT'
	KW_CLEANUP          // 'CLEANUP'
	KW_TEARDOWN         // 'TEARDOWN'
)

// Represent the tokens as strings.
var TokenToString = map[Token]string{
	ILLEGAL:             "illegal",
	EOF:                 "eof",
	WHITESPACE:          "whitespace",
	NEWLINE:             "newline",
	PERIOD:              "period",
	COLON:               "colon",
	LPAREN:              "lparen",
	RPAREN:              "rparen",
	COMMA:               "comma",
	QUOTE:               "quote",
	UNDERSCORE:          "underscore",
	WORD:                "word",
	INTEGER:             "integer",
	KW_TARGET:           "target",
	KW_NAMESPACE:        "namespace",
	KW_CONTEXT:          "context",
	KW_ALIAS:            "alias",
	KW_COMP:             "comp",
	KW_IN:               "in",
	KW_DEFAULT:          "default",
	KW_SYS:              "sys",
	KW_TRIGGER:          "trigger",
	KW_ACCESS:           "access",
	KW_NO_FILTER:        "noFilter",
	KW_FILTER:           "filter",
	KW_ADDED:            "added",
	KW_REMOVED:          "removed",
	KW_ADDED_OR_REMOVED: "addedOrRemoved",
	KW_ALL_OF:           "allOf",
	KW_ANY_OF:           "anyOf",
	KW_NONE_OF:          "noneOf",
	KW_INIT:             "init",
	KW_CLEANUP:          "cleanup",
	KW_TEARDOWN:         "teardown",
}

// Represent the strings as keyword tokens.
var KeywordToToken = map[string]Token{
	"target":         KW_TARGET,
	"namespace":      KW_NAMESPACE,
	"context":        KW_CONTEXT,
	"alias":          KW_ALIAS,
	"comp":           KW_COMP,
	"in":             KW_IN,
	"default":        KW_DEFAULT,
	"sys":            KW_SYS,
	"trigger":        KW_TRIGGER,
	"acess":          KW_ACCESS,
	"noFilter":       KW_NO_FILTER,
	"filter":         KW_FILTER,
	"added":          KW_ADDED,
	"removed":        KW_REMOVED,
	"addedOrRemoved": KW_ADDED_OR_REMOVED,
	"allOf":          KW_ALL_OF,
	"anyOf":          KW_ANY_OF,
	"noneOf":         KW_NONE_OF,
	"init":           KW_INIT,
	"cleanup":        KW_CLEANUP,
	"teardown":       KW_TEARDOWN,
}

var TokenToKeyword = map[Token]string{
	KW_TARGET:           "target",
	KW_NAMESPACE:        "namespace",
	KW_CONTEXT:          "context",
	KW_ALIAS:            "alias",
	KW_COMP:             "comp",
	KW_IN:               "in",
	KW_DEFAULT:          "default",
	KW_SYS:              "sys",
	KW_TRIGGER:          "trigger",
	KW_ACCESS:           "access",
	KW_NO_FILTER:        "noFilter",
	KW_FILTER:           "filter",
	KW_ADDED:            "added",
	KW_REMOVED:          "removed",
	KW_ADDED_OR_REMOVED: "addedOrRemoved",
	KW_ALL_OF:           "allOf",
	KW_ANY_OF:           "anyOf",
	KW_NONE_OF:          "noneOf",
	KW_INIT:             "init",
	KW_CLEANUP:          "cleanup",
	KW_TEARDOWN:         "teardown",
}

// Returns the string representation of a token.
func (t Token) String() string {
	ts := TokenToString[t]
	if ts != "" {
		return ts
	}

	return "unknown"
}

// Returns true if t is a keyword token.
func containsOnly(s string, ch rune) bool {
	for _, sch := range s {
		if sch != ch {
			return false
		}
	}
	return true
}

// Returns true if t is a keyword token.
func isKeyword(t Token) bool {
	_, ok := TokenToKeyword[t]
	return ok
}

// Returns true if ch is a space or a tab.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

// Returns true if ch is a newline character.
func isNewline(ch rune) bool {
	return ch == '\r' || ch == '\n'
}

// Returns true if ch is a upper or lowercase letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// Returns true if ch is a numeric digit
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

var eof = rune(0)
