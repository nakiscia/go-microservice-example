package domain

import (
	"net/http"

	"github.com/nakiscia/go-microservice-example/utils"
)

var (
	users = map[string]*User{
		"1234": &User{FirstName: "Ahmet", LastName: "Nakisci", UserId: 1234},
	}
)

func GetUserById(userId string) (*User, *utils.ApplicationError) {

	user := users[userId]
	if user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: "User not found", StatusCode: http.StatusNotFound, Code: "not_found"}
}
