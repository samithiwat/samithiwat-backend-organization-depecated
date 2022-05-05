package test

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
)

func RawToDtoTeam(team *model.Team) *proto.Team {
	return &proto.Team{
		Id:          uint32(team.ID),
		Name:        team.Name,
		Description: team.Description,
	}
}
