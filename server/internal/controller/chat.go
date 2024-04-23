package controller

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "server/api/v1"
	"server/internal/service"
	"server/internal/ws"
)

var Chat = cChat{
	Users: gmap.New(true),
	Names: gset.NewStrSet(true),
}

type cChat struct {
	Users *gmap.Map    // All users in chat.
	Names *gset.StrSet // All names in chat for unique name validation.
}

func (c *cChat) CreateRoom(ctx context.Context, req *v1.ChatCreateRoomReq) (res *v1.ChatCreateRoomRes, err error) {
	err = service.Chat().CreateRoom(ctx, req.RoomName, req.RoomID)
	return
}

func (c *cChat) JoinRoom(ctx context.Context, req *v1.ChatJoinRoomReq) (res *v1.ChatJoinRoomRes, err error) {
	var (
		r      = g.RequestFromCtx(ctx)
		wsConn *ghttp.WebSocket
	)
	// WebSocket connect.
	if wsConn, err = r.WebSocket(); err != nil {
		return
	}

	// create a client for current websocket.
	err = service.Chat().JoinRoom(ctx, &ws.Client{
		WsConn:   wsConn,
		Message:  make(chan *ws.Message, 10),
		ID:       req.UserId,
		RoomID:   req.RoomID,
		Username: req.Username,
	})

	return
}

func (c *cChat) RoomList(ctx context.Context, req *v1.ChatRoomListReq) (res *v1.ChatRoomListRes, err error) {
	rooms, err := service.Chat().GetRooms(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.ChatRoomListRes{
		Rooms: rooms.List,
	}
	return
}

func (c *cChat) ClientList(ctx context.Context, req *v1.ChatClientListReq) (res *v1.ChatClientListRes, err error) {
	clients, err := service.Chat().GetClients(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}

	res = &v1.ChatClientListRes{
		Clients: clients.List,
	}
	return
}
