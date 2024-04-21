package models

import pb "github.com/robertgarayshin/driveshare/proto"

type User struct {
	Id int
}

func ToUserModel(user *pb.User) *User {
	return &User{
		Id: user.Id,
	}
}
