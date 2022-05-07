package test

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"github.com/samithiwat/samithiwat-backend-organization/src/service"
	"github.com/samithiwat/samithiwat-backend-organization/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneContact(t *testing.T) {
	mock.InitializeMockContact()

	var errors []string

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       &mock.Contact1,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewContactService(&mock.ContactMockClient{})
	locRes, err := locService.FindOne(1)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestFindOneErrGrpcContact(t *testing.T) {
	mock.InitializeMockContact()

	errors := []string{"Not found contact", "Grpc error"}

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewContactService(&mock.ContactMockErrClient{})
	locRes, err := locService.FindOne(1)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, locRes)
}

func TestFindMultiContact(t *testing.T) {
	mock.InitializeMockContact()

	var errors []string

	assert := assert.New(t)
	want := &proto.ContactListResponse{
		Data:       mock.Contacts,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewContactService(&mock.ContactMockClient{})
	locRes, err := locService.FindMulti([]uint32{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}
