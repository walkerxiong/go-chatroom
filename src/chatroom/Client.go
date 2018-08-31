package chatroom

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	Conn       *websocket.Conn
	Username   string
	status     int
	OnlineTime time.Time
}

func InitClient(ws *websocket.Conn) (client *Client) {
	client = &Client{
		Conn:       ws,
		Username:   "",
		status:     1,
		OnlineTime: time.Now(),
	}
	return
}
