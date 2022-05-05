package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"net/http"
)

type TeamRepository interface {
	FindAll(*proto.PaginationMetadata, *[]*model.Team) error
	FindOne(uint, *model.Team) error
	Create(*model.Team) error
	Update(uint, *model.Team) error
	Delete(uint, *model.Team) error
}

type TeamService struct {
	repository TeamRepository
}

func NewTeamService(repository TeamRepository) *TeamService {
	return &TeamService{
		repository: repository,
	}
}

func (s *TeamService) FindAll(_ context.Context, req *proto.FindAllTeamRequest) (res *proto.TeamListResponse, err error) {

	meta := proto.PaginationMetadata{
		ItemsPerPage: req.Limit,
		CurrentPage:  req.Page,
	}

	var ts []*model.Team
	var errors []string

	res = &proto.TeamListResponse{
		Data: &proto.TeamPagination{
			Items: nil,
			Meta:  &meta,
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindAll(&meta, &ts)
	if err != nil {
		errors = append(errors, err.Error())
		res.StatusCode = http.StatusBadRequest
		return
	}

	var result []*proto.Team

	for _, t := range ts {
		result = append(result, RawToDtoTeam(t))
	}

	res.Data.Items = result

	return
}

func (s *TeamService) FindOne(_ context.Context, req *proto.FindOneTeamRequest) (res *proto.TeamResponse, err error) {
	t := model.Team{}
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(uint(req.Id), &t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoTeam(&t)
	res.Data = result

	return
}

func (s *TeamService) Create(_ context.Context, req *proto.CreateTeamRequest) (res *proto.TeamResponse, err error) {
	t := DtoToRawTeam(req.Team)
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return
	}

	result := RawToDtoTeam(t)
	res.Data = result

	return
}

func (s *TeamService) Update(_ context.Context, req *proto.UpdateTeamRequest) (res *proto.TeamResponse, err error) {
	t := DtoToRawTeam(req.Team)
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(uint(t.ID), t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoTeam(t)
	res.Data = result

	return
}

func (s *TeamService) Delete(_ context.Context, req *proto.DeleteTeamRequest) (res *proto.TeamResponse, err error) {
	t := model.Team{}
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(uint(req.Id), &t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoTeam(&t)
	res.Data = result

	return
}
