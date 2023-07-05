package content

import (
	"context"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RPC struct {
	proto.UnimplementedContentServiceServer
	service *Service
}

func NewRPC(service *Service) *RPC {
	s := new(RPC)
	s.service = service
	return s
}

func (h *RPC) CreateContent(ctx context.Context,u *proto.Content) (*proto.Content, error) {
	e, httpErr := h.service.Create(&d, ctx)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
		return
	}
	response.Respond(http.StatusCreated, e, w)
	
	return h.service.Create(
		ctx,
		u,
	)
}

func (h *RPC) ReadContents(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Contents, error) {
	return h.service.ReadMany(
		ctx,
		v,
	)
}

/* func (h *ContentRPC) ReadContentStream(
	v *proto.VoidParam,
	serv proto.ContentService_ReadContentStreamServer,
) (*proto.Contents, error) {
	return nil, nil
	h.service.ReadMany(
		ctx,
		v,
	)
} */

func (h *RPC) ReadContent(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*proto.Content, error) {
	return h.service.ReadOne(
		ctx,
		strVal,
	)
}

func (h *RPC) UpdateContent(
	ctx context.Context,
	p *proto.UpdateContentParam,
) (*proto.Content, error) {
	return h.service.Update(
		ctx,
		p,
	)
}

func (h *RPC) DeleteContent(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*wrapperspb.BoolValue, error) {
	return h.service.Delete(
		ctx,
		strVal,
	)
}
