package auth

import "github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/service"

type Module struct {
	Service *Service
}

func NewModule(s *Service) *Module {
	m := new(Module)
	m.Service = NewService(s)
	return m
}
