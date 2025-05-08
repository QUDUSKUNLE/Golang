package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
)

// ScheduleHandler is a struct that implements the ScheduleServiceServer interface.
// It contains a pointer to the ScheduleService, which is used to handle schedule-related operations.
func (h *ScheduleHandler) CreateSchedule(ctx context.Context, req *schedule.ScheduleRequest) (*schedule.ScheduleResponse, error) {
	// Call the CreateSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.CreateScheduleSession(ctx, req)
}

// GetScheduleSession retrieves a schedule session by its ID.
func (h *ScheduleHandler) GetSchedule(ctx context.Context, req *schedule.GetScheduleRequest) (*schedule.GetScheduleResponse, error) {
	// Call the GetScheduleByID method of the ScheduleService and return the response.
	return h.ScheduleService.GetScheduleSession(ctx, req)
}

// CancelScheduleSession cancels a schedule session by its ID.
func (h *ScheduleHandler) CancelSchedule(ctx context.Context, req *schedule.CancelScheduleRequest) (*schedule.CancelScheduleResponse, error) {
	// Call the CancelSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.CancelScheduleSession(ctx, *req)
}

// UpdateScheduleSession updates a schedule session by its ID.
func (h *ScheduleHandler) UpdateSchedule(ctx context.Context, req *schedule.UpdateScheduleRequest) (*schedule.UpdateScheduleResponse, error) {
	// Call the UpdateSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.UpdateScheduleSession(ctx, req)
}

// ListScheduleSessions retrieves a list of schedule sessions based on the provided parameters.
func (h *ScheduleHandler) ListSchedules(ctx context.Context, req *schedule.ListSchedulesRequest) (*schedule.ListSchedulesResponse, error) {
	// Call the ListSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.ListScheduleSessions(ctx, req)
}
