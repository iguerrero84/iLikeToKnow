package http

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	sqldb "iLikeToKnow.com/internal/database/db"
	"iLikeToKnow.com/internal/domain"
	"iLikeToKnow.com/internal/model"
	"net/http"
	"time"
)

// Store abstracts the subset of sqlc we need (easy to mock in tests).
type Store interface {
	CreateEvent(ctx context.Context, arg sqldb.CreateEventParams) (model.Event, error)
	ListEvents(ctx context.Context) ([]model.Event, error)
	GetEvent(ctx context.Context, id uuid.UUID) (model.Event, error)
}

// EventsHandler holds dependencies for HTTP handlers.
type EventsHandler struct {
	ds domain.Service
}

func NewEventsHandler(domainService domain.Service) *EventsHandler {
	return &EventsHandler{ds: domainService}
}

// --- Request/response DTOs ---

type createEventRequest struct {
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

//func (h *EventsHandler) listEvents(w http.ResponseWriter, r *http.Request) {
//	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
//	defer cancel()
//
//	rows, err := h.store.ListEvents(ctx)
//	if err != nil {
//		h.writeError(w, http.StatusInternalServerError, "failed to list events")
//		return
//	}
//	out := make([]model.Event, 0, len(rows))
//	for _, row := range rows {
//		out = append(out, toDomain(row))
//	}
//	w.Header().Set("Content-Type", "application/json")
//	_ = json.NewEncoder(w).Encode(out)
//}

// --- helpers ---

type ctxKey string

func getOrEmpty(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func (h *EventsHandler) writeError(w http.ResponseWriter, code int, msg string, desc string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": msg, "description": desc})
}
