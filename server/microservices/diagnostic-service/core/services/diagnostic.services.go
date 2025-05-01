package services

import (
		"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/repository"
)

type DiagnosticService struct {
	Repo repository.DiagnosticRepository
}

func NewDiagnosticService(repo repository.DiagnosticRepository) *DiagnosticService {
	return &DiagnosticService{
		Repo: repo,
	}
}
