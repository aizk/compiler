package lexer

type Lexer interface {

}

type SimpleLexer struct {
	currentState FSAState
	tokens       []SimpleToken
	token        SimpleToken
}

type SimpleToken struct {
    text string
    _type TokenType
}

func (s *SimpleToken) reset() {
	s.text = ""
	s._type = NullToken
}

func (s *SimpleToken) appendChar(ch rune) {
	s.text += string(ch)
}

func NewSimpleLexer() *SimpleLexer {
	return &SimpleLexer{
		currentState: InitialState,
	}
}

func (l *SimpleLexer) Tokenize(code string) (err error) {
	for _, ch := range code {
		// 根据字符判断进入的状态
		switch l.currentState {
		case InitialState:
			// 根据字符获取下一个状态
			l.handleChar(ch)
			// 解析当前字符的 Token
		case IdentifierState:
			if isAlpha(ch) || isDigit(ch) { // 其他字母或数组，切换回 Id 状态
				l.token.appendChar(ch)
			} else {
				// 退出 Id 状态
				l.saveToken()
				l.handleChar(ch)
			}
		case IntKeywordState1:
			if ch == 'n' {
				l.currentState = IntKeywordState2
				l.token.appendChar(ch)
			} else if isAlpha(ch) || isDigit(ch) { // 其他字母或数组，切换回 Id 状态
				l.currentState = IdentifierState
				l.token.appendChar(ch)
			} else { // 其他字符状态结束重新开始，例如空格
				l.saveToken()
				l.handleChar(ch)
			}
		case IntKeywordState2:
			// int
			if ch == 't' {
				l.currentState = IntKeywordState3
				l.token.appendChar(ch)
			} else if isAlpha(ch) || isDigit(ch) { // 其他字母或数组，切换回 Id 状态
				l.currentState = IdentifierState
				l.token.appendChar(ch)
			} else { // 其他字符状态结束重新开始，例如空格
				l.saveToken()
				l.handleChar(ch)
			}
		case IntKeywordState3:
			// int 后面接空格字符，可以确定是 int 关键字
			if isBlank(ch) {
				l.token._type = IntToken
				l.saveToken()
				l.currentState = InitialState
			} else {
				l.currentState = IdentifierState
				l.token.appendChar(ch)
			}
		case GTState:
			if ch == '=' {
				l.currentState = GEState // 转换成 GE
				l.token._type = GEToken
				l.token.appendChar(ch)
			} else {
				l.saveToken()
				l.handleChar(ch)
			}
		case GEState:
			l.saveToken()
			l.handleChar(ch)
		case IntLiteralState:
			if isDigit(ch) {
				l.token.appendChar(ch)
			} else {
				l.saveToken()
				l.handleChar(ch)
			}
		case AssignmentState:
			if ch == '=' {
				l.currentState = EQState
				l.token._type = EQToken
				l.token.appendChar(ch)
			} else {
				l.saveToken()
				l.handleChar(ch)
			}
		}
	}

	l.saveToken()
	return
}

func (l *SimpleLexer) saveToken() {
	l.tokens = append(l.tokens, l.token)
	l.token.reset()
}

// 处理 ch 更新 currentState 和当前 Token
func (l *SimpleLexer) handleChar(ch rune) {
	switch {
	case isAlpha(ch):
		// 如果是 i 可能是 int 关键字
		if ch == 'i' {
			l.currentState = IntKeywordState1
		} else {
			l.currentState = IdentifierState
		}
		l.token._type = IdentifierToken
		l.token.appendChar(ch)
	case isDigit(ch):
		l.currentState = IntLiteralState
		l.token._type = IntLiteralToken
		l.token.appendChar(ch)
	case ch == '>':
		l.currentState = GTState
		l.token._type = GTToken
		l.token.appendChar(ch)
	case ch == '=':
		l.currentState = AssignmentState
		l.token._type = AssignmentToken
		l.token.appendChar(ch)
	default:
		l.currentState = InitialState
	}
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch > 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}