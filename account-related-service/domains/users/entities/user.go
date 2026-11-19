package entities

import (

)

type User struct {
	Id         string `gorm:"primaryKey;type:varchar(255)"`
	Email      string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Name       string
	Password   string
	Enabled    int
}
