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

func TestFindAllTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)

	var result []*proto.Team
	for _, team := range mock.Teams {
		result = append(result, RawToDtoTeam(team))
	}

	var errors []string

	want := &proto.TeamListResponse{
		Data: &proto.TeamPagination{
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

	teamService := service.NewTeamService(&mock.TeamMockRepo{})
	teamRes, err := teamService.FindAll(mock.Context{}, &proto.FindAllTeamRequest{Limit: 10, Page: 1})

	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, teamRes, fmt.Sprintf("Want %v but got %v", want, teamRes))
}

func TestFindOneTeam(t *testing.T) {
	mock.InitializeMockTeam()

	var errors []string

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       RawToDtoTeam(&mock.Team1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	teamService := service.NewTeamService(&mock.TeamMockRepo{})
	teamRes, err := teamService.FindOne(mock.Context{}, &proto.FindOneTeamRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, teamRes)
}

func TestFindOneErrNotFoundTeam(t *testing.T) {
	mock.InitializeMockTeam()

	errors := []string{"Not found team"}

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	teamService := service.NewTeamService(&mock.TeamMockErrRepo{})
	teamRes, err := teamService.FindOne(mock.Context{}, &proto.FindOneTeamRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, teamRes)
}

func TestCreateTeam(t *testing.T) {
	mock.InitializeMockTeam()

	var errors []string

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       RawToDtoTeam(&mock.Team1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	teamService := service.NewTeamService(&mock.TeamMockRepo{})
	teamRes, err := teamService.Create(mock.Context{}, &mock.CreateTeamReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, teamRes)
}

func TestUpdateTeam(t *testing.T) {
	mock.InitializeMockTeam()

	var errors []string

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       RawToDtoTeam(&mock.Team1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	teamService := service.NewTeamService(&mock.TeamMockRepo{})
	teamRes, err := teamService.Update(mock.Context{}, &mock.UpdateTeamReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, teamRes)
}

func TestUpdateErrNotFoundTeam(t *testing.T) {
	mock.InitializeMockTeam()

	errors := []string{"Not found team"}

	assert := assert.New(t)

	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	teamService := service.NewTeamService(&mock.TeamMockErrRepo{})
	teamRes, err := teamService.Update(mock.Context{}, &mock.UpdateTeamReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, teamRes)
}

func TestDeleteTeam(t *testing.T) {
	mock.InitializeMockTeam()

	var errors []string

	assert := assert.New(t)
	want := &proto.TeamResponse{
		Data:       RawToDtoTeam(&mock.Team1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	teamService := service.NewTeamService(&mock.TeamMockRepo{})
	teamRes, err := teamService.Delete(mock.Context{}, &proto.DeleteTeamRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, teamRes)
}

func TestDeleteErrNotFoundTeam(t *testing.T) {
	mock.InitializeMockTeam()

	errors := []string{"Not found team"}

	assert := assert.New(t)

	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	teamService := service.NewTeamService(&mock.TeamMockErrRepo{})
	teamRes, err := teamService.Delete(mock.Context{}, &proto.DeleteTeamRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, teamRes)
}
