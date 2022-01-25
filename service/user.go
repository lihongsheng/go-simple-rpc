package service

import (
	"go-simple-rpc/common"
	"strconv"
)

type UserServer struct {
}

func (u UserServer) GetUser(id int) common.User {
	return common.User{
		Id:   id,
		Name: strconv.Itoa(id) + "_name",
	}
}
