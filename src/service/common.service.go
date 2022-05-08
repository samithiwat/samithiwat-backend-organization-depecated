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

func RawToDtoOrganization(org *model.Organization) *proto.Organization {
	var users []*proto.User
	for _, member := range org.Members {
		users = append(users, &proto.User{Id: uint32(member.ID)})
	}

	var roles []*proto.Role
	for _, role := range org.Roles {
		roles = append(roles, &proto.Role{Id: uint32(role.ID)})
	}

	return &proto.Organization{
		Id:          uint32(org.ID),
		Name:        org.Name,
		Description: org.Description,
		Teams:       RawToDtoSubTeams(org.Teams),
		Members:     users,
		Roles:       roles,
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
	}
}

func DtoToRawOrganization(org *proto.Organization) *model.Organization {
	var users []*model.User
	for _, member := range org.Members {
		users = append(users, &model.User{Model: gorm.Model{ID: uint(member.Id)}})
	}

	var roles []*model.Role
	for _, role := range org.Roles {
		roles = append(roles, &model.Role{
			Model:          gorm.Model{ID: uint(role.Id)},
			OrganizationID: uint(org.Id),
		})
	}

	return &model.Organization{
		Model:       gorm.Model{ID: uint(org.Id)},
		Name:        org.Name,
		Description: org.Description,
		Teams:       DtoToRawSubTeams(org.Teams),
		Members:     users,
		Roles:       roles,
	}
}
