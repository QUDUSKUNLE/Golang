package handler

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
)

func (handler *OrganizationServiceStruct) CreateOrganization(ctx context.Context, req *organization.CreateOrganizationRequest) (*organization.CreateOrganizationResponse, error) {
	fmt.Println(req.GetUserId(), "Organization>>>>>>>>>>>>>>>>>")
	return &organization.CreateOrganizationResponse{Id: "1", CreatedAt: "Okay"}, nil
}
