package models

import (
	"github.com/jadengis/modserver/duid"
	"github.com/jinzhu/gorm"
)

var idGenerator duid.IdGenerator

type User struct {
	gorm.Model
	Age   int
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique_index"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", idGenerator.NewId())
	return nil
}

func NewUser(name, email string, age int) *User {
	return &User{
		gorm.Model{ID: uint(idGenerator.NewId())},
		age,
		name,
		email,
	}
}

func SetIdGenerator(generator duid.IdGenerator) {
	idGenerator = generator
}
