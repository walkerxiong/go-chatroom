package chatroom

import "log"

var RoomClient = make(map[*Client]bool) // connected clients

func Join(client *Client) {
	RoomClient[client] = true
}

func OffLine(client *Client) {
	client.Conn.Close()
	delete(RoomClient, client)
}

// 发送消息
func BroadcastMsg() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-Broadcast
		// Send it out to every client that is currently connected
		for client, _ := range RoomClient {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("send msg error: %v", err)
				OffLine(client)
			}
		}

	}
}
