package test

import (
	"github.com/bxcodec/faker/v4"
	"github.com/isd-sgcu/rpkm67-model/model"
)

func MockCheckInModel() *model.CheckIn {
	return &model.CheckIn{
		Event: faker.Word(),
		Email: faker.Email(),
		UserID: faker.UUIDDigit(),
	}
}

func MockCheckInsModel() []*model.CheckIn {
	var checkIns []*model.CheckIn
	for i := 0; i < 10; i++ {
		checkIn := &model.CheckIn{
			Event: faker.Word(),
			Email: faker.Email(),
			UserID: faker.UUIDDigit(),
		}
		checkIns = append(checkIns, checkIn)
	}
	return checkIns
}
