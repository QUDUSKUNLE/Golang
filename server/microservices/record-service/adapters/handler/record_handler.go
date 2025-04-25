package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/QUDUSKUNLE/microservices/record-service/adapters/thirdparty"
	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (this *RecordServiceStruct) GetRecord(ctx context.Context, req *record.GetRecordRequest) (*record.GetRecordResponse, error) {
	_, ok := ctx.Value("user").(*constants.UserType)
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
		Record:         rec.Record,
		ScanTitle:      rec.ScanTitle,
		OrganizationId: rec.OrganizationID,
		CreatedAt:      rec.CreatedAt.Time.String(),
		UpdatedAt:      rec.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) GetRecords(ctx context.Context, req *record.GetRecordsRequest) (*record.GetRecordsResponse, error) {
	organization_user, ok := ctx.Value("user").(*constants.UserType)
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
			UserId:         re.UserID,
			Record:         re.Record,
			ScanTitle:      re.ScanTitle,
			OrganizationId: re.OrganizationID,
			CreatedAt:      re.CreatedAt.Time.String(),
			UpdatedAt:      re.UpdatedAt.Time.String(),
		})
	}
	return recordsResponse, nil
}

func (this *RecordServiceStruct) ScanUpload(ctx context.Context, req *record.ScanUploadRequest) (*record.ScanUploadResponse, error) {
	organization_user, ok := ctx.Value("user").(*constants.UserType)
	// Check authorization right
	if !ok || organization_user.Type != "ORGANIZATION" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	// Get organization details
	organizationDetails, err := this.organizationService.GetOrganizationByUserID(ctx, organization_user.UserID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Organization not found")
	}
	// Get present directory
	directory, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error gettting current directory %v", err)
	}
	// Write file to upload path
	filePath := fmt.Sprintf("%s/uploads/%s", directory, req.GetFileName())
	if err := os.WriteFile(filePath, req.GetContent(), 0644); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}
	// Upload to Cloudinary
	uploadedFile, err := thirdparty.CloudinaryUploader(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload file: %v", err)
	}
	// Remove uploaded file to free memory
	_ = os.Remove(filePath)
	scanRecord, err := this.recordService.CreateRecord(ctx, domain.RecordDto{
		UserID:         req.GetUserId(),
		OrganizationID: organizationDetails.ID,
		Record:         uploadedFile,
		ScanTitle:      req.GetScanTitle(),
	})
	if err != nil {
		return nil, status.Error(codes.Unimplemented, err.Error())
	}
	return &record.ScanUploadResponse{
		Id:             scanRecord.ID,
		UserId:         req.GetUserId(),
		ScanTitle:      req.GetScanTitle(),
		OrganizationId: scanRecord.OrganizationID,
		CreatedAt:      scanRecord.CreatedAt.Time.String(),
		UpdatedAt:      scanRecord.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) SearchRecord(ctx context.Context, req *record.SearchRecordRequest) (*record.SearchRecordResponse, error) {
	organization_user, ok := ctx.Value("user").(*constants.UserType)
	if !ok || organization_user.Type != "ORGANIZATION" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	records, err := this.recordService.SearchRecord(ctx, domain.GetRecordDto{UserID: req.GetUserId(), ScanTitle: req.GetScanTitle()})
	recordsResponse := &record.SearchRecordResponse{
		Records: []*record.Record{},
	}
	if err != nil {
		return nil, status.Error(codes.Unimplemented, err.Error())
	}
	for _, re := range records {
		recordsResponse.Records = append(recordsResponse.Records, &record.Record{
			Id:             re.ID,
			UserId:         re.UserID,
			Record:         re.Record,
			ScanTitle:      re.ScanTitle,
			OrganizationId: re.OrganizationID,
			CreatedAt:      re.CreatedAt.Time.String(),
			UpdatedAt:      re.UpdatedAt.Time.String(),
		})
	}
	return recordsResponse, nil
}

func (this *RecordServiceStruct) SearchByNin(ctx context.Context, req *record.SearchByNinRequest) (*record.SearchRecordResponse, error) {
	organization_user, ok := ctx.Value("user").(*constants.UserType)
	if !ok || organization_user.Type != "ORGANIZATION" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	records, err := this.recordService.SearchRecordByNin(ctx, domain.GetRecordByNinDto{Nin: req.GetNin(), ScanTitle: req.GetScanTitle()})
	if err != nil {
		return nil, status.Error(codes.Unimplemented, err.Error())
	}
	recordsResponse := &record.SearchRecordResponse{
		Records: []*record.Record{},
	}
	for _, re := range records {
		recordsResponse.Records = append(recordsResponse.Records, &record.Record{
			Id:             re.ID,
			UserId:         re.UserID,
			Record:         re.Record,
			ScanTitle:      re.ScanTitle,
			OrganizationId: re.OrganizationID,
			CreatedAt:      re.CreatedAt.Time.String(),
			UpdatedAt:      re.UpdatedAt.Time.String(),
		})
	}
	return recordsResponse, nil
}
