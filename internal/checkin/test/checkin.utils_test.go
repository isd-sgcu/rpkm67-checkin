package test

import (
	"github.com/bxcodec/faker/v4"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
)

func MockCheckInProto() *proto.CheckIn {
	return &proto.CheckIn{
		Id:     faker.UUIDDigit(),
		UserId: faker.UUIDDigit(),
		Email:  faker.Email(),
		Event: faker.Word(),
	}
}

func MockCheckInsProto() []*proto.CheckIn {
	var checkIns []*proto.CheckIn
	for i := 0; i < 10; i++ {
		checkIn := &proto.CheckIn{
			Id:     faker.UUIDDigit(),
			UserId: faker.UUIDDigit(),
			Email:  faker.Email(),
			Event: faker.Word(),
		}
		checkIns = append(checkIns, checkIn)
	}
	return checkIns
}
