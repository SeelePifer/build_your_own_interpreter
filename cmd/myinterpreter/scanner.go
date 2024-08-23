package main

type TokenType string

const (
	LEFT_PAREN    TokenType = "LEFT_PAREN"    // '('
	RIGHT_PAREN   TokenType = "RIGHT_PAREN"   // ')'
	LEFT_BRACE    TokenType = "LEFT_BRACE"    // '{'
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"   // '}'
	DOT           TokenType = "DOT"           // '.'
	COMMA         TokenType = "COMMA"         // ','
	MINUS         TokenType = "MINUS"         // '-'
	PLUS          TokenType = "PLUS"          // '+'
	SEMICOLON     TokenType = "SEMICOLON"     // ';'
	STAR          TokenType = "STAR"          // '*'
	EQUAL         TokenType = "EQUAL"         // '='
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"   // '=='
	BANG          TokenType = "BANG"          // '!'
	BANG_EQUAL    TokenType = "BANG_EQUAL"    // '!='
	LESS          TokenType = "LESS"          // '<'
	LESS_EQUAL    TokenType = "LESS_EQUAL"    // '<='
	GREATER       TokenType = "GREATER"       // '>'
	GREATER_EQUAL TokenType = "GREATER_EQUAL" // '<='
	SLASH         TokenType = "SLASH"         // '/'
	COMMENT       TokenType = "Comment"       // '//'
	STRING        TokenType = "STRING"        // "string"
	NUMBER        TokenType = "NUMBER"        // 1234.1234
	IDENTIFIER    TokenType = "IDENTIFIER"    // foo bar _hello
	AND           TokenType = "AND"           // and
	CLASS         TokenType = "CLASS"         // class
	ELSE          TokenType = "ELSE"          // else
	FALSE         TokenType = "FALSE"         // false
	FOR           TokenType = "FOR"           // for
	FUN           TokenType = "FUN"           // fun
	IF            TokenType = "IF"            // if
	NIL           TokenType = "NIL"           // nil
	OR            TokenType = "OR"            // or
	PRINT         TokenType = "PRINT"         // print
	RETURN        TokenType = "RETURN"        // return
	SUPER         TokenType = "SUPER"         // super
	THIS          TokenType = "THIS"          // this
	TRUE          TokenType = "TRUE"          // true
	VAR           TokenType = "VAR"           // var
	WHILE         TokenType = "WHILE"         // while
	EOF           TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal string
}
type Tokens []Token

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
func (tokens *Tokens) addToken(tokenType TokenType, tokenLexame string, tokenLiteral ...string) {
	var literalValue string
	if len(tokenLiteral) == 1 {
		literalValue = tokenLiteral[0]
	} else if len(tokenLiteral) == 0 {
		literalValue = "null"
	} else {
		panic("tokenLiteral must be 'string'!")
	}
	*tokens = append(*tokens, Token{Type: tokenType, Lexeme: tokenLexame, Literal: literalValue})
}
