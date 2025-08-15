package http

import (
	"context"
	"encoding/json"
	"iLikeToKnow.com/internal/model"
	"net/http"
	"time"
)

func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var in createEventRequest
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		h.writeError(w, http.StatusBadRequest, "ErrInvalidInputBody", "invalid JSON body")
		return
	}

	evt := model.Event{
		Title:       in.Title,
		Description: in.Description,
		StartTime:   in.StartTime,
		EndTime:     in.EndTime,
		CreatedAt:   time.Now(),
	}

	row, err := h.ds.CreateEvent(ctx, evt)

	if err != nil {
		h.writeError(w, http.StatusInternalServerError, "ErrEventCreation", "failed to create event")
		return
	}

	//out := toDomain(row)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(row)
}
