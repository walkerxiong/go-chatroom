package gui

import "github.com/walkerxiong/go-chatroom/session"

type Gui interface {
	Run() error
	ReceiveMsg(msg session.Message)
	// SendMsg(msg session.Message)
	Online(user session.User)
	Offline(user session.User)
}
