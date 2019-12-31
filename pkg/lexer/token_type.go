package lexer

type TokenType int

var TokenMap = map[TokenType]string{
	NullToken: "Null Token",
	IntToken: "Int Token",
	IdentifierToken: "Identifier Token",
	IntLiteralToken: "IntLiteral Token",
	GEToken: "GE Token",
	AssignmentToken: "Assignment Token",
}

const (
	NullToken TokenType = iota
	PlusToken
	MinusToken
	StarToken
	SlashToken

	GEToken // >=
	GTToken // >
	EQToken // ==
	LEToken // <=
	LTToken // <

	SemiColonToken // ;
	LeftParenToken // (
	RightParenToken // )

	AssignmentToken // =

	IfToken
	ElseToken

	IntToken // 16

	IdentifierToken // 标识符

	IntLiteralToken // 18 整数字面量
	StringLiteralToken // 字符串字面量
)