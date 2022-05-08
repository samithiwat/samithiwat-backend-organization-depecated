package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
)

var Org1 model.Organization
var Org2 model.Organization
var Org3 model.Organization
var Org4 model.Organization
var Orgs []*model.Organization
var CreateOrganizationReqMock proto.CreateOrganizationRequest
var UpdateOrganizationReqMock proto.UpdateOrganizationRequest

type OrganizationMockRepo struct {
}

func (r *OrganizationMockRepo) FindAll(pagination *model.OrganizationPagination) error {
	pagination.Meta.CurrentPage = 1
	pagination.Meta.TotalPage = 1
	pagination.Meta.ItemCount = 4
	pagination.Meta.TotalItem = 4
	pagination.Meta.ItemsPerPage = 10

	*pagination.Items = Orgs
	return nil
}

func (r *OrganizationMockRepo) FindOne(_ uint, org *model.Organization) error {
	*org = Org1
	return nil
}

func (r *OrganizationMockRepo) FindMulti(_ []uint32, orgs *[]*model.Organization) error {
	*orgs = Orgs
	return nil
}

func (r *OrganizationMockRepo) Create(org *model.Organization) error {
	*org = Org1
	return nil
}

func (r *OrganizationMockRepo) Update(_ uint, org *model.Organization) error {
	*org = Org1
	return nil
}

func (r *OrganizationMockRepo) Delete(_ uint, org *model.Organization) error {
	*org = Org1
	return nil
}

type OrganizationMockErrRepo struct {
}

func (r *OrganizationMockErrRepo) FindAll(*model.OrganizationPagination) error {
	return nil
}

func (r *OrganizationMockErrRepo) FindOne(uint, *model.Organization) error {
	return errors.New("Not found organization")
}

func (r *OrganizationMockErrRepo) FindMulti([]uint32, *[]*model.Organization) error {
	return nil
}

func (r *OrganizationMockErrRepo) Create(*model.Organization) error {
	return nil
}

func (r *OrganizationMockErrRepo) Update(uint, *model.Organization) error {
	return errors.New("Not found organization")
}

func (r *OrganizationMockErrRepo) Delete(uint, *model.Organization) error {
	return errors.New("Not found organization")
}

func InitializeMockOrganization() (err error) {
	Org1 = model.Organization{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	Org2 = model.Organization{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	Org3 = model.Organization{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	Org4 = model.Organization{
		Name:        faker.Name(),
		Description: faker.Paragraph(),
	}

	CreateOrganizationReqMock = proto.CreateOrganizationRequest{
		Organization: &proto.Organization{
			Name:        Org1.Name,
			Description: Org1.Description,
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	UpdateOrganizationReqMock = proto.UpdateOrganizationRequest{
		Organization: &proto.Organization{
			Id:          uint32(Org1.ID),
			Name:        Org1.Name,
			Description: Org1.Description,
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	Orgs = append(Orgs, &Org1, &Org2, &Org3, &Org4)

	return nil
}
