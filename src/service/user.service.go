package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"time"
)

type UserService struct {
	client proto.UserServiceClient
}

func NewUserService(client proto.UserServiceClient) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) FindOne(id uint) (res *proto.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneUserRequest{Id: int32(id)})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}

func (s *UserService) FindMulti(ids []uint32) (res *proto.UserListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindMulti(ctx, &proto.FindMultiUserRequest{Ids: ids})
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		return res, nil
	}

	return
}
