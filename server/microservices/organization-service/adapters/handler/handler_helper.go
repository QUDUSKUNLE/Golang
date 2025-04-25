package handler

import (
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
)

func (handler *OrganizationServiceStruct) transformOrganizationRPC(req *organization.CreateOrganizationRequest) dto.OrganizationDto {
	userID := req.GetUserId()
	return dto.OrganizationDto{UserID: userID}
}
