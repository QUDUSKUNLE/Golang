package handler

import (
	"context"
	"fmt"

	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/shared/constants"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (this *RecordServiceStruct) ScanUpload(ctx context.Context, req *record.ScanUploadRequest) (*record.ScanUploadResponse, error) {
	log := logger.GetLogger()
	// Retrieve the user from the context
	diagnostic_centre, ok := ctx.Value("user").(*constants.UserType)
	if !ok || diagnostic_centre.Type != string(db.UserEnumDIAGNOSTICCENTRE) {
		log.Warn("Unauthorized access to perform operation.", zap.String("userType", diagnostic_centre.Type))
		return nil, status.Error(codes.Unauthenticated, constants.ErrUnauthorized)
	}
	// Check if user is registered
	_, err := this.userService.GetUser(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.NotFound, constants.ErrUserNotFound)
	}
	// Write file to upload path
	filePath := fmt.Sprintf("uploads/%s", req.GetFileName())
	if err := this.fileService.SaveFile(filePath, req.GetContent()); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}
	// Upload to Cloudinary
	uploadedFile, err := this.fileService.UploadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload file: %v", err)
	}
	// Remove uploaded file to free memory
	_ = this.fileService.DeleteFile(filePath)

	// Create record
	scanRecord, err := this.recordService.CreateRecord(ctx, domain.RecordDto{
		UserID: req.GetUserId(),
		// OrganizationID: organizationDetails.ID,
		Record:    uploadedFile,
		ScanTitle: req.GetScanTitle(),
	})
	if err != nil {
		return nil, status.Error(codes.Unimplemented, err.Error())
	}
	log.Info("Record created successfully", zap.String("recordId", scanRecord.ID))
	return &record.ScanUploadResponse{
		RecordId:     scanRecord.ID,
		UserId:       req.GetUserId(),
		ScanTitle:    req.GetScanTitle(),
		DiagnosticId: scanRecord.DiagnosticID,
		CreatedAt:    scanRecord.CreatedAt.Time.String(),
		UpdatedAt:    scanRecord.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) GetRecord(ctx context.Context, req *record.GetRecordRequest) (*record.GetRecordResponse, error) {
	// Retrieve the user from the context
	diagnostic_centre, ok := ctx.Value("user").(*constants.UserType)
	if !ok || diagnostic_centre.Type != string(db.UserEnumDIAGNOSTICCENTRE) {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}

	rec, err := this.recordService.GetRecord(ctx, req.GetRecordId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Record not found")
	}
	return &record.GetRecordResponse{
		RecordId:     rec.ID,
		UserId:       rec.UserID,
		Record:       rec.Record,
		ScanTitle:    rec.ScanTitle,
		DiagnosticId: rec.DiagnosticID,
		CreatedAt:    rec.CreatedAt.Time.String(),
		UpdatedAt:    rec.UpdatedAt.Time.String(),
	}, nil
}

func (this *RecordServiceStruct) GetRecords(ctx context.Context, req *record.GetRecordsRequest) (*record.GetRecordsResponse, error) {
	// Retrieve the user from the context
	diagnostic_centre, ok := ctx.Value("user").(*constants.UserType)
	if !ok || diagnostic_centre.Type != string(db.UserEnumDIAGNOSTICCENTRE) {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}

	recordsResponse := &record.GetRecordsResponse{
		Records: []*record.Record{},
	}
	return recordsResponse, nil
}

func (this *RecordServiceStruct) SearchRecord(ctx context.Context, req *record.SearchRecordRequest) (*record.SearchRecordResponse, error) {
	// Retrieve the user from the context
	diagnostic_centre, ok := ctx.Value("user").(*constants.UserType)
	if !ok || diagnostic_centre.Type != string(db.UserEnumDIAGNOSTICCENTRE) {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized to perform operation.")
	}
	records, err := this.recordService.SearchRecord(ctx, domain.GetRecordDto{UserID: req.GetUserId(), ScanTitle: req.GetScanTitle()})
	recordsResponse := &record.SearchRecordResponse{
		Records: []*record.Record{},
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, re := range records {
		recordsResponse.Records = append(recordsResponse.Records, &record.Record{
			RecordId:     re.ID,
			UserId:       re.UserID,
			Record:       re.Record,
			ScanTitle:    re.ScanTitle,
			DiagnosticId: re.DiagnosticID,
			CreatedAt:    re.CreatedAt.Time.String(),
			UpdatedAt:    re.UpdatedAt.Time.String(),
		})
	}
	return recordsResponse, nil
}

func (this *RecordServiceStruct) SearchByNin(ctx context.Context, req *record.SearchByNinRequest) (*record.SearchRecordResponse, error) {
	// Retrieve the user from the context
	diagnostic_centre, ok := ctx.Value("user").(*constants.UserType)
	if !ok || diagnostic_centre.Type != string(db.UserEnumDIAGNOSTICCENTRE) {
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
			RecordId:     re.ID,
			UserId:       re.UserID,
			Record:       re.Record,
			ScanTitle:    re.ScanTitle,
			DiagnosticId: re.DiagnosticID,
			CreatedAt:    re.CreatedAt.Time.String(),
			UpdatedAt:    re.UpdatedAt.Time.String(),
		})
	}
	return recordsResponse, nil
}
