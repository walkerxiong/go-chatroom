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
	UUID       string
}

func InitClient(ws *websocket.Conn, username string) (client *Client) {
	client = &Client{
		Conn:       ws,
		Username:   username,
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
		if msg.Message != "" {
			client.SendMessage(msg)
		}

	}
}

func (client *Client) SendMsgBody(msgBody string) {
	Broadcast <- Message{Username: client.Username, Timestamp: time.Now(), Message: msgBody}
}

func (client *Client) SendMessage(msg Message) {
	if msg.Username == "" {
		msg.Username = client.Username
	}
	msg.Timestamp = time.Now()
	Broadcast <- msg
}
