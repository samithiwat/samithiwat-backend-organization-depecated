package service

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"gorm.io/gorm"
)

func RawToDtoSubTeams(team []*model.Team) []*proto.Team {
	var subTeams []*proto.Team
	for _, t := range team {
		subTeam := proto.Team{
			Id:          uint32(t.ID),
			Name:        t.Name,
			Description: t.Description,
			SubTeams:    RawToDtoSubTeams(t.SubTeams),
		}
		subTeams = append(subTeams, &subTeam)
	}
	return subTeams
}

func RawToDtoTeam(team *model.Team) *proto.Team {

	subTeams := RawToDtoSubTeams(team.SubTeams)

	return &proto.Team{
		Id:          uint32(team.ID),
		Name:        team.Name,
		Description: team.Description,
		SubTeams:    subTeams,
	}
}

func DtoToRawSubTeams(team []*proto.Team) []*model.Team {
	var subTeams []*model.Team
	for _, t := range team {
		subTeam := model.Team{
			Model:       gorm.Model{ID: uint(t.Id)},
			Name:        t.Name,
			Description: t.Description,
			SubTeams:    DtoToRawSubTeams(t.SubTeams),
		}
		subTeams = append(subTeams, &subTeam)
	}
	return subTeams
}

func DtoToRawTeam(team *proto.Team) *model.Team {
	return &model.Team{
		Model:       gorm.Model{ID: uint(team.Id)},
		Name:        team.Name,
		Description: team.Description,
		SubTeams:    DtoToRawSubTeams(team.SubTeams),
	}
}
