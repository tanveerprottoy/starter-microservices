package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/dto"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/entity"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/adapter"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/config"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/errorpkg"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/httppkg"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/jsonpkg"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/timepkg"
)

type Service struct {
	HTTPClient *httppkg.HTTPClient
}

func NewService(c *httppkg.HTTPClient) *Service {
	s := new(Service)
	s.HTTPClient = c
	return s
}

func (s Service) Create(d *dto.CreateUpdateUserDto, ctx context.Context) (*entity.User, *errorpkg.HTTPError) {
	hErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	h := http.Header{}
	h.Add("tmp", "d")
	buf, err := jsonpkg.EncodeWithEncoder(*d)
	if err != nil {
		return nil, hErr
	}
	c, e, err := httppkg.Request[entity.User](
		http.MethodPost,
		fmt.Sprintf("%s%s", config.GetEnvValue("USER_SERVICE_BASE_URL"), constant.UserServiceAuthEndpoint),
		h,
		&buf,
		s.HTTPClient,
	)
	if err != nil {
		hErr.Err = err
		return e, hErr
	}
	if c != http.StatusCreated {
		hErr.Err = err
		return e, hErr
	}
	return e, nil
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
