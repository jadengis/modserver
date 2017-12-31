package models

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/jadengis/modserver/protoapi"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	BaseModel
	Age   int    `json:"age"`
	Name  string `gorm:"type:varchar(100)" json:"name"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
}

// func (user *User) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("ID", idGenerator.NewId())
// 	return nil
// }

func (user *User) BeforeDelete(scope *gorm.Scope) error {
	scope.DB().Model(user).Update("email", fmt.Sprintf("del-%s-%s", time.Now().String(), user.Email))
	return nil
}

func NewUser(name, email string, age int) *User {
	return &User{
		BaseModel{ID: uint(idGenerator.NewId())},
		age,
		name,
		email,
	}
}

func NewUserProto(user *User) (*pb.User, error) {
	resultUser := new(pb.User)
	var err error

	resultUser.Id = &pb.Id{uint64(user.ID)}
	resultUser.Age = int64(user.Age)
	resultUser.CreatedAt, err = ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return nil, err
	}
	resultUser.UpdatedAt, err = ptypes.TimestampProto(user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if user.DeletedAt == nil {
		resultUser.DeletedAt = nil
	} else {
		resultUser.DeletedAt, err = ptypes.TimestampProto(*user.DeletedAt)
		if err != nil {
			return nil, err
		}
	}
	resultUser.Name = user.Name
	resultUser.Email = user.Email
	return resultUser, nil
}

func NewUserModel(user *pb.User) (*User, error) {
	resultUser := new(User)
	var err error

	resultUser.ID = uint(user.GetId().GetValue())
	resultUser.Age = int(user.GetAge())
	resultUser.CreatedAt, err = ptypes.Timestamp(user.GetCreatedAt())
	if err != nil {
		return nil, err
	}
	resultUser.UpdatedAt, err = ptypes.Timestamp(user.GetUpdatedAt())
	if err != nil {
		return nil, err
	}
	if user.GetDeletedAt() == nil {
		resultUser.DeletedAt = nil
	} else {
		deletedAt, err := ptypes.Timestamp(user.GetDeletedAt())
		if err != nil {
			return nil, err
		}
		resultUser.DeletedAt = &deletedAt
	}
	resultUser.Name = user.Name
	resultUser.Email = user.Email
	return resultUser, nil
}
