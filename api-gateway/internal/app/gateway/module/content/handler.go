package content

import (
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/dto"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/service"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/adapter"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/httppkg"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/response"
)

type Handler struct {
	service *service.ServiceRPC
}

func NewHandler(s *service.ServiceRPC) *Handler {
	h := new(Handler)
	h.service = s
	return h
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(d, w, r)
}

func (h *Handler) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httppkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httppkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	h.service.ReadMany(limit, page, w, r)
}

func (h *Handler) ReadManyWithNestedDocQuery(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httppkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httppkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	k0 := httppkg.GetQueryParam(r, "key0")
	k1 := httppkg.GetQueryParam(r, "key1")
	h.service.ReadManyWithNestedDocQuery(limit, page, k0, k1, w, r)
}

func (h *Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, d, w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
