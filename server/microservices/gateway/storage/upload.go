package storage

import (
	"io"

	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorageServer struct {
	storage_ Manager
}

func NewStorageServer(storage Manager) *StorageServer {
	return &StorageServer{
		storage_: storage,
	}
}

func (s StorageServer) Upload(stream record.RecordService_ScanUploadServer) error {
	file := NewFile("scan.pdf")
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			if err := s.storage_.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
			return stream.SendAndClose(&record.ScanUploadResponse{CreatedAt: ""})
		}
		if err := file.Write(chunk.GetContent()); err != nil {
			return err
		}
		if err := s.storage_.Store(file); err != nil {
			return err
		}
		return stream.SendAndClose(&record.ScanUploadResponse{CreatedAt: ""})
	}
}
