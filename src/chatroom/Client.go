package chatroom

import (
	"github.com/gorilla/websocket"
	"log"
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
	go client.RecvMsg()

	return
}

// 接收消息
func (client *Client) RecvMsg() {
	for {
		// Read in a new message as JSON and map it to a Message object
		var msg Message
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("recv msg error: %v", err)
			OffLine(client)
			break
		}
		if msg.Username != "" && client.Username == "" {
			client.Username = msg.Username
		}
		if msg.Message != "" {
			msg.Send()
		}

	}
}
