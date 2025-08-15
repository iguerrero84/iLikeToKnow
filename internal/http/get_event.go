package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"iLikeToKnow.com/internal/model"
	"net/http"
	"strings"
	"time"
)

func (h *EventsHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	idStr := strings.TrimPrefix(r.URL.Path, "/events/")

	// Parse string to uuid.UUID
	eventId, err := uuid.Parse(idStr)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "ErrInvalidEventId", "invalid event id")
		fmt.Println("Invalid UUID:", err)
		return
	}

	row, err := h.ds.GetEventById(ctx, eventId)
	fmt.Printf("GetEventById error type: %T, value: %v\n", err, err)
	if err != nil {
		if err.Error() == model.ErrorNoRowsFound {
			fmt.Println("No rows found, I got it!")
			h.writeError(w, http.StatusNotFound, "ResourceNotFound", "event not found")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(row)
}
