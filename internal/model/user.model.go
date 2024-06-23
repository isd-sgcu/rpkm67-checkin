package model

import constants "github.com/isd-sgcu/rpkm67-checkin/constant"

type User struct {
	Base
	Email     string         `json:"email" gorm:"tinytext;unique"`
	Password  string         `json:"password" gorm:"tinytext"`
	Firstname string         `json:"firstname" gorm:"tinytext"`
	Lastname  string         `json:"lastname" gorm:"tinytext"`
	Role      constants.Role `json:"role" gorm:"tinytext"`
	Checkins  []Checkin      `json:"checkins" gorm:"foreignKey:UserID"`
}
