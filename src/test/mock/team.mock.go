package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"gorm.io/gorm"
)

var Team1 model.Team
var Team2 model.Team
var Team3 model.Team
var Team4 model.Team
var Teams []*model.Team
var CreateTeamReqMock proto.CreateTeamRequest
var UpdateTeamReqMock proto.UpdateTeamRequest

type TeamMockRepo struct {
}

func (r *TeamMockRepo) FindAll(meta *proto.PaginationMetadata, teams *[]*model.Team) error {
	meta.CurrentPage = 1
	meta.TotalPage = 1
	meta.ItemCount = 4
	meta.TotalItem = 4
	meta.ItemsPerPage = 10

	*teams = Teams
	return nil
}

func (r *TeamMockRepo) FindOne(_ uint, team *model.Team) error {
	*team = Team1
	return nil
}

func (r *TeamMockRepo) Create(team *model.Team) error {
	*team = Team1
	return nil
}

func (r *TeamMockRepo) Update(_ uint, team *model.Team) error {
	*team = Team1
	return nil
}

func (r *TeamMockRepo) Delete(_ uint, team *model.Team) error {
	*team = Team1
	return nil
}

type TeamMockErrRepo struct {
}

func (r *TeamMockErrRepo) FindAll(*proto.PaginationMetadata, *[]*model.Team) error {
	return nil
}

func (r *TeamMockErrRepo) FindOne(uint, *model.Team) error {
	return errors.New("Not found team")
}

func (r *TeamMockErrRepo) Create(*model.Team) error {
	return nil
}

func (r *TeamMockErrRepo) Update(uint, *model.Team) error {
	return errors.New("Not found team")
}

func (r *TeamMockErrRepo) Delete(uint, *model.Team) error {
	return errors.New("Not found team")
}

func InitializeMockTeam() {
	Team1 = model.Team{
		Model:       gorm.Model{ID: 1},
		Name:        faker.Word(),
		Description: faker.Word(),
	}

	Team2 = model.Team{
		Model:       gorm.Model{ID: 2},
		Name:        faker.Word(),
		Description: faker.Word(),
	}

	Team3 = model.Team{
		Model:       gorm.Model{ID: 3},
		Name:        faker.Word(),
		Description: faker.Word(),
	}

	Team4 = model.Team{
		Model:       gorm.Model{ID: 4},
		Name:        faker.Word(),
		Description: faker.Word(),
	}

	CreateTeamReqMock = proto.CreateTeamRequest{
		Team: &proto.Team{
			Name:        Team1.Name,
			Description: Team1.Description,
		},
	}

	UpdateTeamReqMock = proto.UpdateTeamRequest{
		Team: &proto.Team{
			Id:          uint32(Team1.ID),
			Name:        Team1.Name,
			Description: Team1.Description,
		},
	}

	Teams = append(Teams, &Team1, &Team2, &Team3, &Team4)
}
