package dto

type UserData struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
}
