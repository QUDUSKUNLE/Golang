package clients

// import (
// 	"context"

// 	"github.com/QUDUSKUNLE/microservices/shared/db"
// 	"github.com/QUDUSKUNLE/microservices/shared/dto"
// 	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
// )

// // CreateOrganization implements ports.UseCasePorts.
// func (this *organizationService) CreateOrganization(ctx context.Context, user dto.OrganizationDto) (*db.Organization, error) {
// 	req := &organization.CreateOrganizationRequest{UserId: user.UserID}
// 	resp, err := this.organizationGrpcClient.CreateOrganization(ctx, req)
// 	if err != nil {
// 		return &db.Organization{}, nil
// 	}
// 	return &db.Organization{ID: resp.Id}, nil
// }

// // GetOrganization implements ports.UseCasePorts.
// func (this *organizationService) GetOrganization(ctx context.Context, id string) (*db.Organization, error) {
// 	panic("unimplemented")
// }

// // GetOrganizationID implements ports.UseCasePorts.
// func (this *organizationService) GetOrganizationByUserID(ctx context.Context, user_id string) (*db.Organization, error) {
// 	req := &organization.GetOrganizationByUserIDRequest{UserId: user_id}
// 	res, err := this.organizationGrpcClient.GetOrganizationByUserID(ctx, req)
// 	if err != nil {
// 		return &db.Organization{}, nil
// 	}
// 	return &db.Organization{ID: res.Id, UserID: res.UserId}, nil
// }
