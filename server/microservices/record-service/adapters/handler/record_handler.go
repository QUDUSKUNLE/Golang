package handler

import (
	"context"

	// "fmt"

	"github.com/QUDUSKUNLE/microservices/record-service/adapters/middleware"
	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (this *RecordServiceStruct) CreateRecord(ctx context.Context, req *record.CreateRecordRequest) (*record.CreateRecordResponse, error) {
	data := this.transformRecordRPC(req)
	organization_user := ctx.Value("user").(*middleware.UserType)
	organizationDetails, err := this.organizationService.GetOrganizationByUserID(ctx, organization_user.UserID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Organization not found")
	}
	records, err := this.recordService.CreateRecord(ctx, domain.RecordDto{UserID: data.UserID, OrganizationID: organizationDetails.ID, Record: data.Record})
	if err != nil {
		return nil, status.Error(codes.Unimplemented, "Unimplemented record")
	}
	return &record.CreateRecordResponse{
		Id:             records.ID,
		OrganizationId: records.OrganizationID,
		UserId:         records.UserID,
		Record:         records.Record,
		CreatedAt:      records.CreatedAt.Time.String(),
		UpdatedAt:      records.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) GetRecord(ctx context.Context, req *record.GetRecordRequest) (*record.GetRecordResponse, error) {
	records, err := this.recordService.GetRecord(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Record not found")
	}
	return &record.GetRecordResponse{
		Id:        records.ID,
		UserId:    records.UserID,
		Record:    records.Record,
		CreatedAt: records.CreatedAt.Time.String(),
		UpdatedAt: records.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) GetRecords(ctx context.Context, req *record.GetRecordsRequest) (*record.GetRecordsResponse, error) {
	organization_user, ok := ctx.Value("user").(*middleware.UserType)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	organizationDetails, err := this.organizationService.GetOrganizationByUserID(ctx, organization_user.UserID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Organization not found")
	}
	records, err := this.recordService.GetRecords(ctx, organizationDetails.ID)
	if err != nil {
		return nil, status.Error(codes.Unimplemented, "Unimplemented record")
	}
	recordsResponse := &record.GetRecordsResponse{
		Records: []*record.Record{},
	}
	for _, re := range records {
		recordsResponse.Records = append(recordsResponse.Records, &record.Record{
			Id:             re.ID,
			OrganizationId: re.OrganizationID,
			UserId:         re.UserID,
			Record:         re.Record,
			CreatedAt:      re.CreatedAt.Time.String(),
			UpdatedAt:      re.UpdatedAt.Time.String(),
		})
	}
	return recordsResponse, nil
}
