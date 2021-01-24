package services

import (
	"github.com/nakiscia/go-microservice-example/domain"
	"github.com/nakiscia/go-microservice-example/utils"
)

func GetUser(userId string) (*domain.User, *utils.ApplicationError) {
	return domain.GetUserById(userId)
}
