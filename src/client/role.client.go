package client

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"net/http"
	"time"
)

type RoleClient struct {
	client proto.RoleServiceClient
}

func NewRoleClient(client proto.RoleServiceClient) *RoleClient {
	return &RoleClient{
		client: client,
	}
}

func (c *RoleClient) FindOne(role *proto.Role) (errs []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, grpcErr := c.client.FindOne(ctx, &proto.FindOneRoleRequest{Id: int32(role.Id)})
	if grpcErr != nil {
		errs = append(errs, grpcErr)
		return
	}

	if res.StatusCode != http.StatusOK {
		errs = FormatErrors(res.Errors)
		return
	}

	*role = *res.Data

	return
}

func (c *RoleClient) Create(role *proto.Role) (errs []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, grpcErr := c.client.Create(ctx, &proto.CreateRoleRequest{Role: role})
	if grpcErr != nil {
		errs = append(errs, grpcErr)
		return
	}

	if res.StatusCode != http.StatusOK {
		errs = FormatErrors(res.Errors)
		return
	}

	*role = *res.Data

	return
}

func (c *RoleClient) Update(role *proto.Role) (errs []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, grpcErr := c.client.Update(ctx, &proto.UpdateRoleRequest{Role: role})
	if grpcErr != nil {
		errs = append(errs, grpcErr)
		return
	}

	if res.StatusCode != http.StatusOK {
		errs = FormatErrors(res.Errors)
		return
	}

	*role = *res.Data

	return
}

func (c *RoleClient) Delete(role *proto.Role) (errs []error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, grpcErr := c.client.Delete(ctx, &proto.DeleteRoleRequest{Id: int32(role.Id)})
	if grpcErr != nil {
		errs = append(errs, grpcErr)
		return
	}

	if res.StatusCode != http.StatusOK {
		errs = FormatErrors(res.Errors)
		return
	}

	*role = *res.Data

	return
}
