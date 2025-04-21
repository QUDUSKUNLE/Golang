package usecase

import (
	interfaces "github.com/QUDUSKUNLE/microservices/hospital-service/pkg/v1"
)

type UseCase struct {
	repo interfaces.RepositoryInterface
}

