package user

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/starter-microservices/service/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/data/sql/sqlxpkg"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository sqlxpkg.Repository[entity.User]
}

func NewModule(db *sql.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service, validate)
	return m
}
