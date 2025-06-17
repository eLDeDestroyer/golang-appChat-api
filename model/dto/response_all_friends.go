package dto

type FriendsUser struct {
	Id       uint   `json:"id"`
	RoomId int    `json:"room_id"`
	Name     string `json:"name"`
	Number   int    `gorm:"column:unique_number" json:"unique_number"`
}
