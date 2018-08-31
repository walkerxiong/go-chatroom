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

func InitClinet(ws *websocket.Conn, username string) (client *Client) {
	client = &Client{
		conn:       ws,
		Username:   username,
		status:     1,
		OnlineTime: time.Now(),
	}
	return
}
