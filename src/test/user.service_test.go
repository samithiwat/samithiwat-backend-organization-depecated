package test

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"github.com/samithiwat/samithiwat-backend-organization/src/service"
	"github.com/samithiwat/samithiwat-backend-organization/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneUser(t *testing.T) {
	mock.InitializeMockUser()

	var errors []string

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       &mock.User1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewUserService(&mock.UserMockClient{})
	roleRes, err := roleService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestFindOneErrGrpcUser(t *testing.T) {
	mock.InitializeMockUser()

	errors := []string{"Not found user", "Grpc error"}

	assert := assert.New(t)
	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	roleService := service.NewUserService(&mock.UserMockErrClient{})
	roleRes, err := roleService.FindOne(1)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, roleRes)
}

func TestFindMultiUser(t *testing.T) {
	mock.InitializeMockUser()

	var errors []string

	assert := assert.New(t)
	want := &proto.UserListResponse{
		Data:       mock.Users,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewUserService(&mock.UserMockClient{})
	roleRes, err := roleService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}
