package dto

import "github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/entity"

type RemoteResponse[T any] struct {
	Data T
}

type UsersRemoteResponse struct {
	Items []entity.User `json:"items"`
	Limit int           `json:"limit"`
	Page  int           `json:"page"`
}
