package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type ChatNameReq struct {
	g.Meta `path:"/chat/name" method:"post"  tags:"ChatService" summary:"Name"`
	Name   string `v:"required|max-length:21#Why not an awesome name"`
}
type ChatNameRes struct{}

type ChatWebsocketReq struct {
	g.Meta `path:"/chat/websocket" method:"get"  tags:"ChatService" summary:"Send message"`
}
type ChatWebsocketRes struct{}

type ChatCreateRoomReq struct {
	g.Meta   `path:"/chat/create_room" method:"post"  tags:"ChatService" summary:"Create room"`
	RoomID   string `v:"required#Please input a correct room id" json:"room_id"`
	RoomName string `v:"required#Please input a correct room name" json:"room_name"`
}
type ChatCreateRoomRes struct{}

type ChatJoinRoomReq struct {
	g.Meta   `path:"/chat/join_room" method:"get"  tags:"ChatService" summary:"Join room"`
	RoomID   string `v:"required#Please input a correct room id" json:"room_id"`
	UserId   string `v:"required#user_id is required" json:"user_id"`
	Username string `v:"required|max-length:21#Please input a correct user name" json:"username"`
}
type ChatJoinRoomRes struct{}

type ChatRoomListReq struct {
	g.Meta `path:"/chat/room_list" method:"get"  tags:"ChatService" summary:"Get room list"`
}
type ChatRoomListRes struct {
	Rooms []model.RoomItem `json:"rooms"`
}

type ChatClientListReq struct {
	g.Meta `path:"/chat/client_list" method:"get"  tags:"ChatService" summary:"Get client list"`
	RoomID string `v:"required#Please input a correct room id" json:"room_id"`
}
type ChatClientListRes struct {
	Clients []model.ClientItem `json:"clients"`
}
