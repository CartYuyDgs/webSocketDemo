package model

type Message struct {
	FromUid string `json:"fromUid"`
	ToUid   string `json:"toUid"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

const (
	LOGIN  string = "LOGIN"
	LOGOUT string = "LOGOUT"
	SAY    string = "SAY"
	PONG   string = "PONG"
)

func (m *Message) isLogin() bool {
	return m.Type == LOGIN
}

func (m *Message) isLogout() bool {
	return m.Type == LOGOUT
}

func (m *Message) isSay() bool {
	return m.Type == SAY
}

func (m *Message) isPong() bool {
	return m.Type == PONG
}
