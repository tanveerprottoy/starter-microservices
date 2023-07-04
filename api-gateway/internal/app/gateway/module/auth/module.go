package auth

import "github.com/tanveerprottoy/starter-microservices/gateway/pkg/httppkg"

type Module struct {
	Service *Service
}

func NewModule(c *httppkg.HTTPClient) *Module {
	m := new(Module)
	m.Service = NewService(c)
	return m
}
