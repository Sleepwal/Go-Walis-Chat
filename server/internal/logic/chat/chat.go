package chat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/service"
	"server/internal/ws"
)

type sChat struct {
	hub *ws.Hub
}

func New() *sChat {
	return &sChat{
		hub: ws.HubObj,
	}
}

func init() {
	service.RegisterChat(New())
}

func (s *sChat) CreateRoom(ctx context.Context, roomName string, roomId string) (err error) {
	s.hub.Rooms[roomId] = &ws.Room{
		ID:      roomId,
		Name:    roomName,
		Clients: make(map[string]*ws.Client),
	}

	g.Log().Infof(ctx, "Create room: %#v", s.hub.Rooms[roomId])
	return
}

func (s *sChat) JoinRoom(ctx context.Context, client *ws.Client) (err error) {
	msg := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   client.RoomID,
		Username: client.Username,
	}

	g.Log().Infof(ctx, "Join room: %#v", client)

	// Register a new client in the room.
	s.hub.Register <- client
	// Broadcast the message to all clients in the room.
	s.hub.Broadcast <- msg

	// Start a new goroutine to write message.
	go client.WriteMessage()
	// Read messages.
	client.ReadMessage(ctx, s.hub)

	return
}

func (s *sChat) GetRooms(ctx context.Context) (out *model.ChatRoomsOutput, err error) {
	out = &model.ChatRoomsOutput{
		List: make([]model.RoomItem, 0),
	}

	for _, room := range s.hub.Rooms {
		out.List = append(out.List, model.RoomItem{
			ID:   room.ID,
			Name: room.Name,
		})
	}
	return
}

func (s *sChat) GetClients(ctx context.Context, roomId string) (out *model.ChatClientsOutput, err error) {
	room, ok := s.hub.Rooms[roomId]
	if !ok {
		return
	}

	out = &model.ChatClientsOutput{
		List: make([]model.ClientItem, 0),
	}
	for _, client := range room.Clients {
		out.List = append(out.List, model.ClientItem{
			ID:       client.ID,
			Username: client.Username,
		})
	}
	return
}
