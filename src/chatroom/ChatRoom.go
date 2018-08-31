package chatroom

import "time"

var RoomClient = make(map[*Client]bool) // connected clients

func Join(client *Client) {
	RoomClient[client] = true
	Message{Username: client.Username, Message: "加入群聊", Timestamp: time.Now()}.Send()
}
