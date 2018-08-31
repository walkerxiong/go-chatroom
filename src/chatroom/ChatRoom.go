package chatroom

var RoomClient = make(map[*Client]bool) // connected clients

func Join(client *Client) {
	RoomClient[client] = true
}

func OffLine(client *Client) {
	delete(RoomClient, client)
}
