package test

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"github.com/samithiwat/samithiwat-backend-organization/src/service"
	"github.com/samithiwat/samithiwat-backend-organization/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneRole(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       &mock.Role1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockClient{})
	roleRes, err := roleService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}

func TestFindOneErrNotFoundRole(t *testing.T) {
	mock.InitializeMockRole()

	errors := []string{"Not found role", "Grpc error"}

	assert := assert.New(t)
	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	roleService := service.NewRoleService(&mock.RoleMockErrClient{})
	roleRes, err := roleService.FindOne(1)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, roleRes)
}

func TestFindMulti(t *testing.T) {
	mock.InitializeMockRole()

	var errors []string

	assert := assert.New(t)
	want := &proto.RoleListResponse{
		Data:       mock.Roles,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	roleService := service.NewRoleService(&mock.RoleMockClient{})
	roleRes, err := roleService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, roleRes)
}
