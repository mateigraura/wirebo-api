package data

type User struct {
	Id       string `json:"id"`
	Name     string `pg:",unique" json:"name"`
	Username string `json:"userName"`
	Email    string `json:"email"`
}

type Message struct {
}
