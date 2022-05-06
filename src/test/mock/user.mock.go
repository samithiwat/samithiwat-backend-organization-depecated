package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var User1 proto.User
var User2 proto.User
var User3 proto.User
var User4 proto.User
var Users []*proto.User

type UserMockClient struct {
}

func (*UserMockClient) FindAll(_ context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	return nil, nil
}

func (*UserMockClient) FindOne(_ context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &User1,
	}, nil
}

func (*UserMockClient) FindMulti(_ context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return &proto.UserListResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       Users,
	}, nil
}

func (*UserMockClient) Create(_ context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, nil
}

func (*UserMockClient) Update(_ context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, nil
}

func (*UserMockClient) Delete(_ context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, nil
}

type UserMockErrClient struct {
}

func (*UserMockErrClient) FindAll(_ context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	return nil, nil
}

func (*UserMockErrClient) FindOne(_ context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, errors.New("Grpc error")
}

func (*UserMockErrClient) FindMulti(_ context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (*UserMockErrClient) Create(_ context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, nil
}

func (*UserMockErrClient) Update(_ context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, nil
}

func (*UserMockErrClient) Delete(_ context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, nil
}

func InitializeMockUser() {
	User1 = proto.User{
		Id:        1,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User2 = proto.User{
		Id:        2,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User3 = proto.User{
		Id:        3,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User4 = proto.User{
		Id:        4,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	Users = append(Users, &User1, &User2, &User3, &User4)
}
