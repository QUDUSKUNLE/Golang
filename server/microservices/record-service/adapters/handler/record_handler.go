package handler

import (
	"context"
	"io"
	"os"

	// "github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
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
	re, err := this.recordService.CreateRecord(ctx, domain.RecordDto{UserID: data.UserID, OrganizationID: organizationDetails.ID, Record: data.Record})
	if err != nil {
		return nil, status.Error(codes.Unimplemented, "Unimplemented record")
	}
	return &record.CreateRecordResponse{
		Id:             re.ID,
		OrganizationId: re.OrganizationID,
		UserId:         re.UserID,
		Record:         re.Record,
		CreatedAt:      re.CreatedAt.Time.String(),
		UpdatedAt:      re.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) GetRecord(ctx context.Context, req *record.GetRecordRequest) (*record.GetRecordResponse, error) {
	_, ok := ctx.Value("user").(*middleware.UserType)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	rec, err := this.recordService.GetRecord(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Record not found")
	}
	return &record.GetRecordResponse{
		Id:             rec.ID,
		UserId:         rec.UserID,
		OrganizationId: rec.OrganizationID,
		Record:         rec.Record,
		CreatedAt:      rec.CreatedAt.Time.String(),
		UpdatedAt:      rec.UpdatedAt.Time.String(),
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

func (this *RecordServiceStruct) ScanUpload(stream record.RecordService_ScanUploadServer) error {
	var fileName string
	var userID string
	var scanTitle string
	var file *os.File
	var fileData []byte

	// Obtain context
	ctx := stream.Context()
	organization_user, ok := ctx.Value("user").(*middleware.UserType)
	if !ok {
		return status.Error(codes.PermissionDenied, "Unauthorized to perform operation")
	}

	for {
		chunk, err := stream.Recv()
		fileName = chunk.GetFileName()
		userID = chunk.GetUserId()
		scanTitle = chunk.GetScanTitle()

		if err == io.EOF {
			return stream.SendAndClose(&record.ScanUploadResponse{
				FileName:       chunk.GetFileName(),
				UserId:         chunk.GetUserId(),
				ScanTitle:      chunk.GetScanTitle(),
				OrganizationId: organization_user.UserID,
			})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		fileData = append(fileData, chunk.GetContent()...)
		_, err = file.Write(fileData)
		if err != nil {
			return stream.SendAndClose(&record.ScanUploadResponse{
				FileName:       fileName,
				UserId:         userID,
				ScanTitle:      scanTitle,
				OrganizationId: organization_user.UserID,
			})
		}
	}
}
