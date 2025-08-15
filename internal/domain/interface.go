package domain

import (
	"context"
	"github.com/google/uuid"
	"iLikeToKnow.com/internal/database"
	db "iLikeToKnow.com/internal/database/db"
	"iLikeToKnow.com/internal/model"
)

type Service interface {
	GetEventById(ctx context.Context, eventId uuid.UUID) (*model.Event, error)
	CreateEvent(ctx context.Context, event model.Event) (uuid.UUID, error)
}

type Database[Q db.Queries] interface {
	TX(ctx context.Context, f database.QueryFunc[Q]) error
	Raw(ctx context.Context, f database.QueryFunc[Q]) error
}
