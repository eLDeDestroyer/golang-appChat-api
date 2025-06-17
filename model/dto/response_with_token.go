package dto

type ResponseWithToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
