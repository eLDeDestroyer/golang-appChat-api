package model

type Friend struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Friend uint   `json:"friend"`
	UserId uint   `json:"user_id"`
	RoomId uint   `json:"room_id"`
}
