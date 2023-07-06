package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/dto"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/entity"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/global"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/adapter"
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

func (s Service) Create(d *dto.CreateUpdateUserDto, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	var e entity.User
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	h := http.Header{}
	h.Add("tmp", "d")
	buf, err := jsonpkg.EncodeWithEncoder(*d)
	if err != nil {
		return e, httpErr
	}
	c, e, err := httppkg.Request[entity.User](http.MethodPost, fmt.Sprintf("%s%s", global.UserServiceBaseUrl, constant.UsersEndpoint), h, &buf, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return e, httpErr
	}
	if c != http.StatusCreated {
		httpErr.Err = err
		return e, httpErr
	}
	return e, nil
}

func (s Service) ReadMany(limit, page int, ctx context.Context) (dto.UsersRemoteResponse, *errorpkg.HTTPError) {
	var d dto.RemoteResponse[dto.UsersRemoteResponse]
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	qMap := make(map[string]string)
	qMap["limit"] = adapter.IntToString(limit)
	qMap["page"] = adapter.IntToString(page)
	u, err := httppkg.BuildURL(global.UserServiceBaseUrl, "", qMap)
	if err != nil {
		httpErr.Err = err
		return d.Data, httpErr
	}
	h := http.Header{}
	h.Add("tmp", "d")
	c, d, err := httppkg.Request[dto.RemoteResponse[dto.UsersRemoteResponse]](http.MethodGet, u, h, nil, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return d.Data, httpErr
	}
	if c != http.StatusOK {
		httpErr.Err = err
		return d.Data, httpErr
	}
	return d.Data, nil
}

func (s Service) ReadOne(id string, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	var d entity.User
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	u, err := httppkg.BuildURL(global.UserServiceBaseUrl, id, nil)
	if err != nil {
		httpErr.Err = err
		return d, httpErr
	}
	h := http.Header{}
	h.Add("tmp", "d")
	c, d, err := httppkg.Request[dto.RemoteResponse[dt]](http.MethodGet, u, h, nil, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return d, httpErr
	}
	if c != http.StatusOK {
		httpErr.Err = err
		return d, httpErr
	}
	return d, nil
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
