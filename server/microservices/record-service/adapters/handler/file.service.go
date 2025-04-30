package handler

import (
	"os"
	"github.com/QUDUSKUNLE/microservices/record-service/adapters/thirdparty"
)

type FileService interface {
	SaveFile(filePath string, content []byte) error
	UploadFile(filePath string) (string, error)
	DeleteFile(filePath string) error
}

type LocalFileService struct {}

func (l *LocalFileService) SaveFile(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}

func (l *LocalFileService) UploadFile(filePath string) (string, error) {
	// Simulate file upload and return the file path
	return thirdparty.CloudinaryUploader(filePath)
}

func (l *LocalFileService) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}
