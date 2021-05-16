package models

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateRoomRequest struct {
	Name      string   `json:"name"`
	IsPrivate bool     `json:"isPrivate"`
	UsersRefs []string `json:"usersRefs"`
}
