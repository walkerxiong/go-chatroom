package chatroom

import "time"

type Message struct {
	Username  string
	Message   string
	Timestamp time.Time
}

var Broadcast = make(chan Message)

func (msg Message) Send() {
	Broadcast <- msg
}
