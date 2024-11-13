package usecase

import (
	interfaces "github.com/QUDUSKUNLE/microservices/services/auth-service/pkg/v1"
)

type UseCase struct {
	repo interfaces.RepositoryInterface
}

func New(repo interfaces.RepositoryInterface) interfaces.RepositoryInterface {
	return &UseCase{repo: repo}
}
