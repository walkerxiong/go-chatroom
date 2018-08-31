package chatroom

import "time"

type Message struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Timestamp time.Time
}

var Broadcast = make(chan Message, 100)

func (msg Message) Send() {
	Broadcast <- msg
}
