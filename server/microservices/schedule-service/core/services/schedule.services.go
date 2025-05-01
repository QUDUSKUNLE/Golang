package services

import (
	"github.com/QUDUSKUNLE/microservices/schedule-service/adapters/repository"
)

type ScheduleService struct {
	Repo repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{
		Repo: repo,
	}
}
