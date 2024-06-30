package test

import (
	"github.com/bxcodec/faker/v4"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/isd-sgcu/rpkm67-model/model"
)

func MockCheckInModel() *model.CheckIn {
	return &model.CheckIn{
		Event: faker.Word(),
		Email: faker.Email(),
		UserID: faker.UUIDDigit(),
	}
}

func MockCheckInProto() *proto.CheckIn {
	return &proto.CheckIn{
		Id:     faker.UUIDDigit(),
		UserId: faker.UUIDDigit(),
		Email:  faker.Email(),
	}
}

func MockCheckInsProto() []*proto.CheckIn {
	var checkIns []*proto.CheckIn
	for i := 0; i < 10; i++ {
		checkIn := &proto.CheckIn{
			Id:     faker.UUIDDigit(),
			UserId: faker.UUIDDigit(),
			Email:  faker.Email(),
		}
		checkIns = append(checkIns, checkIn)
	}
	return checkIns
}
