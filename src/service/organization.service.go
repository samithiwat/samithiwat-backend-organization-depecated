package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"net/http"
)

type OrganizationService struct {
	repository OrganizationRepository
}

type OrganizationRepository interface {
	FindAll(*proto.PaginationMetadata, *[]*model.Organization) error
	FindOne(uint, *model.Organization) error
	Create(*model.Organization) error
	Update(uint, *model.Organization) error
	Delete(uint, *model.Organization) error
}

func NewOrganizationService(repository OrganizationRepository) *OrganizationService {
	return &OrganizationService{repository: repository}
}

func (s *OrganizationService) FindAll(_ context.Context, req *proto.FindAllOrganizationRequest) (res *proto.OrganizationListResponse, err error) {

	meta := proto.PaginationMetadata{
		ItemsPerPage: req.Limit,
		CurrentPage:  req.Page,
	}

	var orgs []*model.Organization
	var errors []string

	res = &proto.OrganizationListResponse{
		Data: &proto.OrganizationPagination{
			Items: nil,
			Meta:  &meta,
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindAll(&meta, &orgs)
	if err != nil {
		errors = append(errors, err.Error())
		res.StatusCode = http.StatusBadRequest
		return
	}

	var result []*proto.Organization

	for _, org := range orgs {
		result = append(result, RawToDtoOrganization(org))
	}

	res.Data.Items = result

	return
}

func (s *OrganizationService) FindOne(_ context.Context, req *proto.FindOneOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := model.Organization{}
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(uint(req.Id), &org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoOrganization(&org)
	res.Data = result
	return
}

func (s *OrganizationService) Create(_ context.Context, req *proto.CreateOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := DtoToRawOrganization(req.Organization)
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return
	}

	result := RawToDtoOrganization(org)
	res.Data = result

	return
}

func (s *OrganizationService) Update(_ context.Context, req *proto.UpdateOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := DtoToRawOrganization(req.Organization)
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(uint(org.ID), org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoOrganization(org)
	res.Data = result

	return
}

func (s *OrganizationService) Delete(_ context.Context, req *proto.DeleteOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := model.Organization{}
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(uint(req.Id), &org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoOrganization(&org)
	res.Data = result

	return
}
