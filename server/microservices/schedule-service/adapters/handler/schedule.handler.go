package handler

import (
	"github.com/QUDUSKUNLE/microservices/schedule-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
)

type ScheduleHandler struct {
	ScheduleService *services.ScheduleService
	schedule.UnimplementedScheduleServiceServer
}

// NewScheduleHandler creates a new ScheduleHandler with the given ScheduleService.
// This function initializes the ScheduleHandler with the provided ScheduleService.
// It returns a pointer to the newly created ScheduleHandler instance.
func NewScheduleHandler(scheduleService *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		ScheduleService: scheduleService,
	}
}
