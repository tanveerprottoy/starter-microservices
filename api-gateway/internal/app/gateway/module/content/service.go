package content

import (
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user/proto"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/module/user/dto"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/module/user/entity"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/adapter"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/response"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/timepkg"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Service struct {
}

func (s *Service) Create(d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	v.CreatedAt = timepkg.Now()
	v.UpdatedAt = timepkg.Now()
	// ctx := context.Background()
	// send to service
	e, err := grpc.ServiceRPCClient.CreateUser(
		r.Context(),
		&proto.User{
			Name: v.Name,
		},
	)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(v), w)
}

func (s *Service) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	u, err := grpc.ServiceRPCClient.ReadUsers(
		r.Context(),
		&proto.VoidParam{},
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(u), w)
}

func (s *Service) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	u, err := grpc.ServiceRPCClient.ReadUser(
		r.Context(),
		&wrapperspb.StringValue{Value: id},
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(u), w)
}

func (s *Service) Update(id string, d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	u, err := grpc.ServiceRPCClient.UpdateUser(
		r.Context(),
		&proto.UpdateUserParam{
			Id: id,
			User: &proto.User{
				Name: d.Name,
			},
		},
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(u), w)
}

func (s *Service) Delete(id string, w http.ResponseWriter, r *http.Request) {
	u, err := grpc.ServiceRPCClient.DeleteUser(
		r.Context(),
		&wrapperspb.StringValue{Value: id},
		nil,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(u), w)
}
