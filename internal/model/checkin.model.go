package model

type Checkin struct {
	Base
	Email  string `json:"email" gorm:"tinytext"`
	Event  string `json:"event" gorm:"tinytext"`
	UserID uint   `json:"user_id"`
}
