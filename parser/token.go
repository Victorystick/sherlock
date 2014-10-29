package parser

type TokenType int

func (t TokenType) String() string {
	return tokMap[int(t)]
}

var tokMap = []string{
	"TokEOF",
	"TokId",
	"TokNum",
	"TokString",
	"TokLBrack",
	"TokRBrack",
	"TokNL",
	"TokPunc",
}

const (
	TokEOF TokenType = iota
	TokId
	TokNum
	TokString
	TokLBrack
	TokRBrack
	TokNL
	TokPunc
)

type Token struct {
	Type  TokenType
	Value string
}