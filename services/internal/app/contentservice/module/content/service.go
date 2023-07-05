package content

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tanveerprottoy/starter-microservices/service/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/errorpkg"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/grpcpkg"
	"google.golang.org/grpc/codes"
)

type Service struct {
	repository *Repository
}

func NewService(r *Repository) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(u *proto.Content, ctx context.Context) (*proto.User, errorpkg.GRPCError) {
	res, err := s.repository.Create(
		ctx,
		u,
	)
	if err != nil {
		return nil,  grpcpkg.RespondError(codes.Unknown, constant.UnknownError)
	}
	fmt.Println("create.res: ", res)
	return u, nil
}

func (s *Service) ReadMany(ctx context.Context, v *proto.VoidParam) (*proto.Users, error) {
	opts := mongodb.BuildPaginatedOpts(limit, skip)
	c, err := s.repository.ReadMany(
		r.Context(),
		bson.D{},
		&opts,
	)
	log.Print("ReadMany rpc")
	d := &proto.Users{}
	rows, err := s.repository.ReadMany()
	if err != nil {
		return nil, grpcpkg.RespondError(codes.Unknown, constant.UnknownError)
	}
	var (
		users      []*proto.User
		id         string
		name       string
		created_at time.Time
		updated_at time.Time
	)
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&id, &name, &created_at, &updated_at); err != nil {
			return nil, fmt.Errorf("ReadMany %v", err)
		}
		users = append(users, &proto.User{
			Id:   id,
			Name: name,
			CreatedAt: timestamppb.New(
				created_at,
			),
			UpdatedAt: timestamppb.New(
				updated_at,
			),
		})
	}
	d.Users = users
	return d, err
}

/* func (s *Service) ReadUserStream(
	v *proto.VoidParam,
	serv proto.ServiceRPC_ReadUserStreamServer,
) (*proto.Users, error) {
	return &proto.Users{}, nil
} */

func (s *Service) ReadOne(ctx context.Context, strVal *wrapperspb.StringValue) (*proto.User, error) {
	row := s.repository.ReadOne(
		strVal.Value,
	)
	if row == nil {
		return nil, grpcpkg.RespondError(codes.NotFound, constantglobal.NotFound)
	}
	var (
		id         string
		name       string
		created_at time.Time
		updated_at time.Time
	)
	if err := row.Scan(&id, &name, &created_at, &updated_at); err != nil {
		return nil, fmt.Errorf("ReadOne %v", err)
	}
	u := &proto.User{
		Id:   id,
		Name: name,
		CreatedAt: timestamppb.New(
			created_at,
		),
		UpdatedAt: timestamppb.New(
			updated_at,
		),
	}
	return u, nil
}

func (s *Service) Update(ctx context.Context, p *proto.UpdateUserParam) (*proto.User, error) {
	r, err := s.repository.Update(
		p.Id,
		p.User,
	)
	if err != nil || r <= 0 {
		return nil, grpcpkg.RespondError(codes.Unknown, constant.UnknownError)
	}
	return p.User, nil
}

func (s *Service) Delete(ctx context.Context, strVal *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	r, err := s.repository.Delete(
		strVal.Value,
	)
	if err != nil || r <= 0 {
		return nil, grpcpkg.RespondError(codes.Unknown, constant.UnknownError)
	}
	return &wrapperspb.BoolValue{Value: true}, nil
}
