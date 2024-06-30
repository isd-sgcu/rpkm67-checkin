package checkin

import (
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/isd-sgcu/rpkm67-model/model"
)

func ModelToProto(in *model.CheckIn) *proto.CheckIn {
	return &proto.CheckIn{
		Id:     in.ID.String(),
		Email:  in.Email,
		Event:  in.Event,
		UserId: in.UserID,
	}
}

func ModelToProtoList(in []*model.CheckIn) []*proto.CheckIn {
	var out []*proto.CheckIn
	for _, v := range in {
		out = append(out, ModelToProto(v))
	}
	return out
}
