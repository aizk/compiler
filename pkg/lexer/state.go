package lexer

// FSA 的各种状态
type FSAState int

const (
	InitialState    FSAState = iota // 初始状态
	If                              // 1
	IdentifierState                 // 标识符
	AssignmentState                 // 赋值 '='

	IntKeywordState1
	IntKeywordState2
	IntKeywordState3

	IntLiteralState

	GTState
	GEState
	EQState

	Plus // +
	Minus // -
	Star // *
	Slash // /
)
