package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}
	command := os.Args[1]
	commands := []string{"tokenize", "parse"}
	if !contains(commands, command) {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	var tokens Tokens
	exitCode := 0
	line := 1
	strFileContents := string(fileContents)
	for i := 0; i < len(strFileContents); i++ {
		token := rune(strFileContents[i])
		switch token {
		case '\n':
			line++
		// Start of ignored cases
		case '\r':
		case '\t':
		case ' ':
		// End of ignored cases
		case '(':
			tokens.addToken(LEFT_PAREN, string(token))
		case ')':
			tokens.addToken(RIGHT_PAREN, string(token))
		case '{':
			tokens.addToken(LEFT_BRACE, string(token))
		case '}':
			tokens.addToken(RIGHT_BRACE, string(token))
		case ',':
			tokens.addToken(COMMA, string(token))
		case '.':
			tokens.addToken(DOT, string(token))
		case '-':
			tokens.addToken(MINUS, string(token))
		case '+':
			tokens.addToken(PLUS, string(token))
		case '*':
			tokens.addToken(STAR, string(token))
		case ';':
			tokens.addToken(SEMICOLON, string(token))
		case '=':
			if i+1 < len(strFileContents) && strFileContents[i+1] == '=' {
				tokens.addToken(EQUAL_EQUAL, "==")
				i++
			} else {
				tokens.addToken(EQUAL, string(token))
			}
		case '!':
			if i+1 < len(strFileContents) && strFileContents[i+1] == '=' {
				tokens.addToken(BANG_EQUAL, "!=")
				i++
			} else {
				tokens.addToken(BANG, string(token))
			}
		case '<':
			if i+1 < len(strFileContents) && strFileContents[i+1] == '=' {
				tokens.addToken(LESS_EQUAL, "<=")
				i++
			} else {
				tokens.addToken(LESS, string(token))
			}
		case '>':
			if i+1 < len(strFileContents) && strFileContents[i+1] == '=' {
				tokens.addToken(GREATER_EQUAL, ">=")
				i++
			} else {
				tokens.addToken(GREATER, string(token))
			}
		case '/':
			if i+1 < len(strFileContents) && strFileContents[i+1] == '/' {
				i++
				for {
					if strFileContents[i+1] == '\n' {
						break
					} else {
						i++
						if i+1 == len(strFileContents) {
							break
						}
					}
				}
			} else {
				tokens.addToken(SLASH, string(token))
			}
		case '"':
			start := i + 1
			isUnterminatedString := false
			for {
				i++
				if strFileContents[i] == '"' {
					break
				} else if strFileContents[i] == '\n' {
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
					isUnterminatedString = true
					line++
					break
				} else if i+1 == len(strFileContents) {
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
					isUnterminatedString = true
					break
				}
			}
			if isUnterminatedString {
				exitCode = 65
				break
			}
			str := strFileContents[start:i]
			tokens.addToken(STRING, fmt.Sprintf("%c%s%c", '"', str, '"'), str)
		default:
			if unicode.IsDigit(token) {
				start := i
				isDot := false
				for {
					if i+1 == len(strFileContents) {
						break
					}
					if unicode.IsDigit(rune(strFileContents[i+1])) {
						i++
					} else if rune(strFileContents[i+1]) == '.' {
						if isDot || len(strFileContents) == i+2 || !unicode.IsDigit(rune(strFileContents[i+2])) {
							break
						}
						isDot = true
						i++
					} else {
						break
					}
				}
				number := strFileContents[start : i+1]
				numberLiteral := number
				numberLiteralParts := strings.Split(numberLiteral, ".")
				if len(numberLiteralParts) < 2 {
					numberLiteral += ".0"
				} else {
					for {
						if numberLiteral[len(numberLiteral)-1] == '0' && numberLiteral[len(numberLiteral)-2] != '.' {
							numberLiteral = numberLiteral[:len(numberLiteral)-1]
						} else {
							break
						}
					}
				}
				tokens.addToken(NUMBER, number, numberLiteral)
				continue
			}
			if unicode.IsLetter(token) || token == '_' {
				switch token {
				case 'a':
					if i+2 < len(strFileContents) && strFileContents[i+1] == 'n' && strFileContents[i+2] == 'd' {
						tokens.addToken(AND, "and")
						i += 2
						continue
					}
				case 'c':
					if i+4 < len(strFileContents) && strFileContents[i+1] == 'l' && strFileContents[i+2] == 'a' && strFileContents[i+3] == 's' && strFileContents[i+4] == 's' {
						tokens.addToken(CLASS, "class")
						i += 4
						continue
					}
				case 'e':
					if i+3 < len(strFileContents) && strFileContents[i+1] == 'l' && strFileContents[i+2] == 's' && strFileContents[i+3] == 'e' {
						tokens.addToken(ELSE, "else")
						i += 3
						continue
					}
				case 'f':
					if i+4 < len(strFileContents) && strFileContents[i+1] == 'a' && strFileContents[i+2] == 'l' && strFileContents[i+3] == 's' && strFileContents[i+4] == 'e' {
						tokens.addToken(FALSE, "false")
						i += 4
						continue
					}
					if i+2 < len(strFileContents) && strFileContents[i+1] == 'o' && strFileContents[i+2] == 'r' {
						tokens.addToken(FOR, "for")
						i += 2
						continue
					}
					if i+2 < len(strFileContents) && strFileContents[i+1] == 'u' && strFileContents[i+2] == 'n' {
						tokens.addToken(FUN, "fun")
						i += 2
						continue
					}
				case 'i':
					if i+1 < len(strFileContents) && strFileContents[i+1] == 'f' {
						tokens.addToken(IF, "if")
						i += 1
						continue
					}
				case 'n':
					if i+2 < len(strFileContents) && strFileContents[i+1] == 'i' && strFileContents[i+2] == 'l' {
						tokens.addToken(NIL, "nil")
						i += 2
						continue
					}
				case 'o':
					if i+1 < len(strFileContents) && strFileContents[i+1] == 'r' {
						tokens.addToken(OR, "or")
						i += 1
						continue
					}
				case 'p':
					if i+4 < len(strFileContents) && strFileContents[i+1] == 'r' && strFileContents[i+2] == 'i' && strFileContents[i+3] == 'n' && strFileContents[i+4] == 't' {
						tokens.addToken(PRINT, "print")
						i += 4
						continue
					}
				case 'r':
					if i+5 < len(strFileContents) && strFileContents[i+1] == 'e' && strFileContents[i+2] == 't' && strFileContents[i+3] == 'u' && strFileContents[i+4] == 'r' && strFileContents[i+5] == 'n' {
						tokens.addToken(RETURN, "return")
						i += 5
						continue
					}
				case 's':
					if i+4 < len(strFileContents) && strFileContents[i+1] == 'u' && strFileContents[i+2] == 'p' && strFileContents[i+3] == 'e' && strFileContents[i+4] == 'r' {
						tokens.addToken(SUPER, "super")
						i += 4
						continue
					}
				case 't':
					if i+3 < len(strFileContents) && strFileContents[i+1] == 'h' && strFileContents[i+2] == 'i' && strFileContents[i+3] == 's' {
						tokens.addToken(THIS, "this")
						i += 3
						continue
					}
					if i+3 < len(strFileContents) && strFileContents[i+1] == 'r' && strFileContents[i+2] == 'u' && strFileContents[i+3] == 'e' {
						tokens.addToken(TRUE, "true")
						i += 3
						continue
					}
				case 'v':
					if i+2 < len(strFileContents) && strFileContents[i+1] == 'a' && strFileContents[i+2] == 'r' {
						tokens.addToken(VAR, "var")
						i += 2
						continue
					}
				case 'w':
					if i+4 < len(strFileContents) && strFileContents[i+1] == 'h' && strFileContents[i+2] == 'i' && strFileContents[i+3] == 'l' && strFileContents[i+4] == 'e' {
						tokens.addToken(WHILE, "while")
						i += 4
						continue
					}
				}
				start := i
				for {
					if i+1 == len(strFileContents) || strFileContents[i+1] == ' ' {
						break
					}
					if !unicode.IsLetter(rune(strFileContents[i+1])) && rune(strFileContents[i+1]) != '_' && !unicode.IsNumber(rune(strFileContents[i+1])) {
						break
					}
					i++
				}
				tokens.addToken(IDENTIFIER, strFileContents[start:i+1])
				continue
			}
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, token)
			exitCode = 65
			continue
		}
	}
	tokens.addToken(EOF, "")
	if command == "tokenize" {
		for _, token := range tokens {
			fmt.Printf("%s %s %s\n", token.Type, token.Lexeme, token.Literal)
		}
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
	if command == "parse" {
		for i := 0; i < len(tokens); i++ {
			if tokens[i].Type == "NUMBER" &&
				contains([]string{"PLUS", "MINUS"}, string(tokens[i+1].Type)) &&
				tokens[i+2].Type == "NUMBER" {
				fmt.Println([]any{tokens[i+1].Lexeme, tokens[i].Literal, tokens[i+2].Literal})
			}
			if contains([]string{"TRUE", "FALSE", "NIL"}, string(tokens[i].Type)) {
				fmt.Println(tokens[i].Lexeme)
			}
		}
	}
}
