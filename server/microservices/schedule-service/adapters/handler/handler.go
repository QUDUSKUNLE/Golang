package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
)

// ScheduleHandler is a struct that implements the ScheduleServiceServer interface.
// It contains a pointer to the ScheduleService, which is used to handle schedule-related operations.
func (h *ScheduleHandler) CreateScheduleSession(ctx context.Context, req *schedule.ScheduleRequest) (*schedule.ScheduleResponse, error) {
	// Call the CreateSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.CreateSchedule(ctx, req)
}

// GetScheduleSession retrieves a schedule session by its ID.
func (h *ScheduleHandler) GetScheduleSession(ctx context.Context, req *schedule.GetScheduledSessionRequest) (*schedule.GetScheduledSessionResponse, error) {
	// Call the GetScheduleByID method of the ScheduleService and return the response.
	return h.ScheduleService.GetScheduleSession(ctx, req)
}
