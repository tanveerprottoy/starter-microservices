package user

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/service/internal/app/userservice/module/user/dto"
	"github.com/tanveerprottoy/starter-microservices/service/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-microservices/service/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/data/sql/sqlxpkg"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/errorpkg"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/timepkg"
)

type Service struct {
	repository sqlxpkg.Repository[entity.User]
}

func NewService(r sqlxpkg.Repository[entity.User]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s Service) readOneInternal(id string) (entity.User, error) {
	return s.repository.ReadOne(id)
}

func (s Service) Create(d *dto.CreateUpdateUserDto, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	// convert dto to entity
	b := entity.User{}
	b.Name = d.Name
	n := timepkg.NowUnixMilli()
	b.CreatedAt = n
	b.UpdatedAt = n
	err := s.repository.Create(&b)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	return b, nil
}

func (s Service) ReadMany(limit, page int, ctx context.Context) (map[string]any, *errorpkg.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.User, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return m, errorpkg.HandleDBError(err)
	}
	m["items"] = d
	return m, nil
}

func (s Service) ReadOne(id string, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	return b, nil
}

func (s Service) Update(id string, d *dto.CreateUpdateUserDto, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	b.Name = d.Name
	b.UpdatedAt = timepkg.NowUnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorpkg.HTTPError{Code: http.StatusBadRequest, Err: errorpkg.NewError(constant.OperationNotSuccess)}
}

func (s Service) Delete(id string, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorpkg.HTTPError{Code: http.StatusBadRequest, Err: errorpkg.NewError(constant.OperationNotSuccess)}
}
