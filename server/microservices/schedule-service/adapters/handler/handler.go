package handler

import (
	"context"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
)

// ScheduleHandler is a struct that implements the ScheduleServiceServer interface.
// It contains a pointer to the ScheduleService, which is used to handle schedule-related operations.
func (h *ScheduleHandler) CreateScheduleSession(ctx context.Context, req *schedule.ScheduleRequest) (*schedule.ScheduleResponse, error) {
	// Call the CreateSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.CreateScheduleSession(ctx, req)
}

// GetScheduleSession retrieves a schedule session by its ID.
func (h *ScheduleHandler) GetScheduleSession(ctx context.Context, req *schedule.GetScheduledSessionRequest) (*schedule.GetScheduledSessionResponse, error) {
	// Call the GetScheduleByID method of the ScheduleService and return the response.
	return h.ScheduleService.GetScheduleSession(ctx, req)
}

// CancelScheduleSession cancels a schedule session by its ID.
func (h *ScheduleHandler) CancelScheduleSession(ctx context.Context, req *schedule.CancelScheduledSessionRequest) (*schedule.CancelScheduledSessionResponse, error) {
	// Call the CancelSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.CancelScheduleSession(ctx, *req)
}

// UpdateScheduleSession updates a schedule session by its ID.
func (h *ScheduleHandler) UpdateScheduleSession(ctx context.Context, req *schedule.UpdateScheduledSessionRequest) (*schedule.UpdateScheduledSessionResponse, error) {
	// Call the UpdateSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.UpdateScheduleSession(ctx, req)
}

// ListScheduleSessions retrieves a list of schedule sessions based on the provided parameters.
func (h *ScheduleHandler) ListScheduleSessions(ctx context.Context, req *schedule.ListScheduledSessionsRequest) (*schedule.ListScheduledSessionsResponse, error) {
	// Call the ListSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.ListScheduleSessions(ctx, req)
}

// Get Diagnostic centre schedule by schedule ID
func (h *ScheduleHandler) GetDiagnosticCentreSchedule(ctx context.Context, req *schedule.GetDiagnosticCentreScheduleRequest) (*schedule.GetDiagnosticCentreScheduleResponse, error) {
	// Call the GetDiagnosticCentreSchedule method of the ScheduleService and return the response.
	return h.ScheduleService.GetScheduleByDiagnosticCentre(ctx, req)
}

// Get Diagnostic centre schedule by status and date
func (h *ScheduleHandler) ListDiagnosticCentreSchedules(ctx context.Context, req *schedule.ListDiagnosticCentreSchedulesRequest) (*schedule.ListDiagnosticCentreSchedulesResponse, error) {
	// Call the GetDiagnosticCentreScheduleByStatusAndDate method of the ScheduleService and return the response.
	return h.ScheduleService.ListDiagnosticCentreSchedules(ctx, req)
}
