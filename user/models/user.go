package models

import proto "github.com/robertgarayshin/driveshare/proto/users"

type User struct {
	Id int
}

func ToUserModel(user *proto.User) *User {
	return &User{
		Id: int(user.Id),
	}
}
