package parser

func Tokenize(source string) ([]Token, error) {
	lexer := NewLexer()

	lexer.Stream([]byte(source))

	lexer.Lex()

	return lexer.Tokens(), nil
}

type Lexer struct {
	stream               []byte
	start, curr, nesting int
	fn                   LexFunc
	tokens               []Token
}

type LexFunc func(*Lexer) LexFunc

func NewLexer() Lexer {
	return Lexer{
		fn: LexUnknown,
	}
}

func (l *Lexer) Stream(bytes []byte) {
	l.stream = bytes
}

func (l Lexer) Tokens() []Token {
	return l.tokens
}

func (l *Lexer) Lex() {
	for l.fn != nil {
		l.fn = l.fn(l)
	}
}

func (l *Lexer) next() (byte, bool) {
	curr := l.curr

	if l.curr >= len(l.stream) {
		return 0, false
	}

	l.curr++
	return l.stream[curr], l.curr <= len(l.stream)
}

func (l *Lexer) currentString() string {
	start := l.start
	l.start = l.curr
	str := string(l.stream[start:l.curr])
	return str
}

func (l *Lexer) addToken(t Token) {
	l.tokens = append(l.tokens, t)
}

func (l *Lexer) addCurrentToken(t TokenType) {
	l.addToken(Token{
		Type:  t,
		Value: l.currentString(),
	})
}

func (l *Lexer) addConstStrToken(t TokenType, str string) {
	l.addToken(Token{
		Type:  t,
		Value: str,
	})
}

func (l *Lexer) back(num int) *Lexer {
	l.curr -= num
	return l
}

func LexUnknown(l *Lexer) LexFunc {
	ch, ok := l.next()

	switch {
	case 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z':
		return LexId
	case '0' <= ch && ch <= '9' || ch == '-':
		return LexNum
	case '\n' == ch:
		l.addConstStrToken(TokNL, "<NL>")
	case ' ' == ch:
		// ignore
		l.start++
	case '{' == ch:
		l.addCurrentToken(TokLBrack)
	case '}' == ch:
		l.addCurrentToken(TokRBrack)
	default:
		if ch != 0 {
			l.addCurrentToken(TokPunc)
		}
	}

	if !ok {
		l.addConstStrToken(TokEOF, "<EOF>")
		return nil
	}

	return LexUnknown
}

func LexId(l *Lexer) LexFunc {
	// fmt.Println("parse id")

	for {
		char, ok := l.next()

		if !ok || char == ' ' || 'a' > char || char > 'z' {
			break
		}
	}

	l.back(1)

	l.addCurrentToken(TokId)

	return LexUnknown
}

func LexNum(l *Lexer) LexFunc {
	// fmt.Println("parse num")

	for {
		char, ok := l.next()

		if !ok {
			break
		}

		if '0' > char || char > '9' {
			l.back(1)
			break
		}
	}

	l.addCurrentToken(TokNum)

	return LexUnknown
}
