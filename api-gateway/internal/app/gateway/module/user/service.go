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
	var r dto.RemoteResponse[dto.UsersRemoteResponse]
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	qMap := make(map[string]string)
	qMap["limit"] = adapter.IntToString(limit)
	qMap["page"] = adapter.IntToString(page)
	u, err := httppkg.BuildURL(global.UserServiceBaseUrl+constant.UsersEndpoint, "", qMap)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	h := http.Header{}
	h.Add("tmp", "d")
	c, r, err := httppkg.Request[dto.RemoteResponse[dto.UsersRemoteResponse]](http.MethodGet, u, h, nil, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	if c != http.StatusOK {
		httpErr.Err = err
		return r.Data, httpErr
	}
	return r.Data, nil
}

func (s Service) ReadOne(id string, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	var r dto.RemoteResponse[entity.User]
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	u, err := httppkg.BuildURL(global.UserServiceBaseUrl+constant.UsersEndpoint, id, nil)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	h := http.Header{}
	h.Add("tmp", "d")
	c, r, err := httppkg.Request[dto.RemoteResponse[entity.User]](http.MethodGet, u, h, nil, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	if c != http.StatusOK {
		httpErr.Err = err
		return r.Data, httpErr
	}
	return r.Data, nil
}

func (s Service) Update(id string, d *dto.CreateUpdateUserDto, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	var r dto.RemoteResponse[entity.User]
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	u, err := httppkg.BuildURL(global.UserServiceBaseUrl+constant.UsersEndpoint, id, nil)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	h := http.Header{}
	h.Add("tmp", "d")
	c, r, err := httppkg.Request[dto.RemoteResponse[entity.User]](http.MethodGet, u, h, nil, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	if c != http.StatusOK {
		httpErr.Err = err
		return r.Data, httpErr
	}
	return r.Data, nil
}

func (s Service) Delete(id string, ctx context.Context) (entity.User, *errorpkg.HTTPError) {
	var r dto.RemoteResponse[entity.User]
	httpErr := &errorpkg.HTTPError{Code: http.StatusInternalServerError, Err: errorpkg.NewError(constant.InternalServerError)}
	u, err := httppkg.BuildURL(global.UserServiceBaseUrl+constant.UsersEndpoint, id, nil)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	h := http.Header{}
	h.Add("tmp", "d")
	c, r, err := httppkg.Request[dto.RemoteResponse[entity.User]](http.MethodGet, u, h, nil, s.HTTPClient)
	if err != nil {
		httpErr.Err = err
		return r.Data, httpErr
	}
	if c != http.StatusOK {
		httpErr.Err = err
		return r.Data, httpErr
	}
	return r.Data, nil
}
