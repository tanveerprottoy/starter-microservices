package adapter

import (
	"database/sql"

	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/apigateway/module/user/entity"
	sqlpkg "github.com/tanveerprottoy/starter-microservices/gateway/pkg/data/sql"
)

func RowToUserEntity(row *sql.Row) (*entity.User, error) {
	e := new(entity.User)
	return sqlpkg.GetEntity(
		row,
		e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
}
