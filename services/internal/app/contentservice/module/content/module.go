package content

import "github.com/jmoiron/sqlx"

type Module struct {
	RPC        *RPC
	Service    *Service
	Repository *Repository
}

func NewModule(db *sqlx.DB) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.RPC = NewRPC(m.Service)
	return m
}
