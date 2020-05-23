package model

type Message struct {
	FromUid string `json:"from_uid"`
	ToUid   string `json:"to_uid"`
	Type    string `json:"type"`
	content string `json:"content"`
}

const (
	LOGIN  string = "LOGIN"
	LOGOUT string = "LOGOUT"
	SAY    string = "SAY"
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
