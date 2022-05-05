package test

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"github.com/samithiwat/samithiwat-backend-organization/src/service"
	"github.com/samithiwat/samithiwat-backend-organization/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindAllOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)

	var result []*proto.Organization
	for _, org := range mock.Orgs {
		result = append(result, RawToDtoOrganization(org))
	}

	var errors []string

	want := &proto.OrganizationListResponse{
		Data: &proto.OrganizationPagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    4,
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockRepo{})
	orgRes, err := orgService.FindAll(mock.Context{}, &proto.FindAllOrganizationRequest{Limit: 10, Page: 1})

	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, orgRes, fmt.Sprintf("Want %v but got %v", want, orgRes))
}

func TestFindOneOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	var errors []string

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       RawToDtoOrganization(&mock.Org1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockRepo{})
	orgRes, err := orgService.FindOne(mock.Context{}, &proto.FindOneOrganizationRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, orgRes)
}

func TestFindOneErrNotFoundOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	errors := []string{"Not found organization"}

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockErrRepo{})
	orgRes, err := orgService.FindOne(mock.Context{}, &proto.FindOneOrganizationRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, orgRes)
}

func TestCreateOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	var errors []string

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       RawToDtoOrganization(&mock.Org1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockRepo{})
	orgRes, err := orgService.Create(mock.Context{}, &mock.CreateOrganizationReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, orgRes)
}

func TestUpdateOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	var errors []string

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       RawToDtoOrganization(&mock.Org1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockRepo{})
	orgRes, err := orgService.Update(mock.Context{}, &mock.UpdateOrganizationReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, orgRes)
}

func TestUpdateErrNotFoundOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	errors := []string{"Not found organization"}

	assert := assert.New(t)

	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockErrRepo{})
	orgRes, err := orgService.Update(mock.Context{}, &mock.UpdateOrganizationReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, orgRes)
}

func TestDeleteOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	var errors []string

	assert := assert.New(t)
	want := &proto.OrganizationResponse{
		Data:       RawToDtoOrganization(&mock.Org1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockRepo{})
	orgRes, err := orgService.Delete(mock.Context{}, &proto.DeleteOrganizationRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, orgRes)
}

func TestDeleteErrNotFoundOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	errors := []string{"Not found organization"}

	assert := assert.New(t)

	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	orgService := service.NewOrganizationService(&mock.OrganizationMockErrRepo{})
	orgRes, err := orgService.Delete(mock.Context{}, &proto.DeleteOrganizationRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, orgRes)
}
