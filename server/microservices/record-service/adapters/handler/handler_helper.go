package handler

import (
	"github.com/QUDUSKUNLE/microservices/record-service/core/domain"
	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
)

func (this *RecordServiceStruct) transformRecordRPC(req *record.CreateRecordRequest) domain.RecordDto {
	return domain.RecordDto{UserID: req.GetUserId(), Record: req.GetRecord()}
}
