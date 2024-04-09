package model

type ChatMsg struct {
	Type string      `json:"type" v:"required"`
	Data interface{} `json:"data" v:"required"`
	From string      `json:"name" v:""`
}

type RoomItem struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
type ChatRoomsOutput struct {
	List []RoomItem `json:"list"`
}

type ClientItem struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
type ChatClientsOutput struct {
	List []ClientItem `json:"list"`
}
