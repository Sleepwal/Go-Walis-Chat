package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

var HubObj = NewHub()

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register: // register client to room
			// get room by client.RoomID
			if _, ok := h.Rooms[client.RoomID]; ok {
				room := h.Rooms[client.RoomID]

				// get client by client.ID, if not exist, add it to room.Clients
				if _, ok := room.Clients[client.ID]; !ok {
					room.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister: // unregister client from room
			// get room by client.RoomID
			if _, ok := h.Rooms[client.RoomID]; ok {
				room := h.Rooms[client.RoomID]

				if _, ok := room.Clients[client.ID]; ok {
					// if room has other clients, broadcast user left the room
					if len(room.Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user has left the room",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}

					// delete client from room.Clients
					delete(room.Clients, client.ID)
					close(client.Message)
				}
			}
		case message := <-h.Broadcast:
			// if room not exist, broadcast message to all clients in this room
			if room, ok := h.Rooms[message.RoomID]; ok {
				for _, client := range room.Clients {
					client.Message <- message
				}
			}
		}
	}
}
