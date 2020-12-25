package services

import (
	"github.com/CodingMaven1/go-microservices/mvc/domain"
	"github.com/CodingMaven1/go-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
