package model

type Chat struct {
	Id     uint   `json:"id"`
	Text   string `json:"chat"`
	RoomId uint   `json:"room_id"`
	UserId uint   `json:"user_id"`
}
