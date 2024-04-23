package ws

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gorilla/websocket"
	"server/internal/consts"
)

type Client struct {
	WsConn   *ghttp.WebSocket
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

// WriteMessage listens to the message channel and write to the current client
func (c *Client) WriteMessage() {
	defer func() {
		c.WsConn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.WsConn.WriteJSON(message)
	}
}

// ReadMessage reads messages from current client and broadcast to all clients in the same room
func (c *Client) ReadMessage(ctx context.Context, hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.WsConn.Close()
	}()

	for {
		_, m, err := c.WsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				g.Log("error").Error(context.TODO(), err)
			}
			break
		}

		// Checks sending interval limit.
		var cacheKey = fmt.Sprintf("ChatWebSocket:%p", c.WsConn)
		if ok, _ := gcache.SetIfNotExist(ctx, cacheKey, struct{}{}, consts.ChatIntervalLimit); !ok {
			c.Message <- &Message{
				Content:  `Message sending too frequently, why not a rest first`,
				RoomID:   c.RoomID,
				Username: "System",
			}
			continue
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}
}
