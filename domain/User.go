package domain

type User struct {
	UserId    int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
