package handler

import (
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
)

func (handler *OrganizationServiceStruct) transformOrganizationRPC(req *organization.CreateOrganizationRequest) domain.OrganizationDto {
	userID := req.GetUserId()
	return domain.OrganizationDto{UserID: userID}
}
