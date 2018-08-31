package chatroom

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	conn       *websocket.Conn
	Username   string
	status     int
	OnlineTime time.Time
}

func InitClient(ws *websocket.Conn) (client *Client) {
	client = &Client{
		conn:       ws,
		Username:   "",
		status:     1,
		OnlineTime: time.Now(),
	}
	return
}
