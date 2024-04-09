// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
	"server/internal/ws"
)

type (
	IChat interface {
		CreateRoom(ctx context.Context, roomName string, roomId string) (err error)
		JoinRoom(ctx context.Context, client *ws.Client) (err error)
		GetRooms(ctx context.Context) (res *model.ChatRoomsOutput, err error)
		GetClients(ctx context.Context, roomId string) (res *model.ChatClientsOutput, err error)
	}
)

var (
	localChat IChat
)

func Chat() IChat {
	if localChat == nil {
		panic("implement not found for interface IChat, forgot register?")
	}
	return localChat
}

func RegisterChat(i IChat) {
	localChat = i
}
