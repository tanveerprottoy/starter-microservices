package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/httppkg"
)

type Module struct {
	Handler         *Handler
	Service         *Service
}

func NewModule(c *httppkg.HTTPClient, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Service = NewService(c)
	m.Handler = NewHandler(m.Service, validate)
	return m
}
