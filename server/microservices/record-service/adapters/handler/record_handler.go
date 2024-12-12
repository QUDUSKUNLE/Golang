package handler

import (
	"context"
	"fmt"

	// "github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (this *RecordServiceStruct) CreateRecord(ctx context.Context, req *record.CreateRecordRequest) (*record.CreateRecordResponse, error) {
	data := this.transformRecordRPC(req)
	if data.Record == "" || data.UserID == "" {
		return nil, status.Error(codes.InvalidArgument, "Record or UserID cannot be empty")
	}
	userRecord, err := this.authService.GetUser(ctx, data.UserID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "User record not found")
	}
	fmt.Println(userRecord, ">>>>>>>>>>>>>>>>>>")
	// organization, err :=
	// record, err := this.useCase.CreateRecord(ctx, domain.RecordDto{})
	return &record.CreateRecordResponse{Record: "Record created successfully"}, nil
}
