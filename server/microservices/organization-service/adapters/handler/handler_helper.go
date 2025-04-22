package handler

import (
	"github.com/QUDUSKUNLE/microservices/gateway/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/domain"
)

func (handler *OrganizationServiceStruct) transformOrganizationRPC(req *organization.CreateOrganizationRequest) domain.OrganizationDto {
	userID := req.GetUserId()
	return domain.OrganizationDto{UserID: userID}
}
