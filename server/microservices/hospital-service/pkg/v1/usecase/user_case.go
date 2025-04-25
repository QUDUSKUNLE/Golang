package usecase

import (
	v1 "github.com/QUDUSKUNLE/microservices/hospital-service/pkg/v1"
)

type UseCase struct {
	repo v1.RepositoryInterface
}
