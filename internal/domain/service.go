package domain

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	db "iLikeToKnow.com/internal/database/db"
	"iLikeToKnow.com/internal/model"
	"iLikeToKnow.com/internal/utils"
	"log"
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validate input invariants for creation/update.
func (e *Event) Validate() error {
	if len(e.Title) == 0 {
		return errors.New("title is required")
	}
	if len(e.Title) > 100 {
		return errors.New("title must be <= 100 characters")
	}
	if e.StartTime.IsZero() || e.EndTime.IsZero() {
		return errors.New("start_time and end_time are required")
	}
	if !e.StartTime.Before(e.EndTime) {
		return errors.New("start_time must be before end_time")
	}
	return nil
}

type ServiceImpl struct {
	dbService Database[db.Queries]
}

func (s ServiceImpl) GetEventById(ctx context.Context, eventId uuid.UUID) (*model.Event, error) {
	var event db.Event

	if err := s.dbService.Raw(ctx, func(ctx context.Context, conn *pgxpool.Conn, q *db.Queries) error {
		var err error
		event, err = q.GetEventById(ctx, eventId)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}); err == nil && event.ID != eventId {
		fmt.Println("Do I pass by here ? ")
		return nil, errors.New(model.ErrorNoRowsFound)
	}

	// Map db.Event → model.Event
	return &model.Event{
		ID:          event.ID.String(),
		Title:       event.Title,
		Description: event.Description,
		StartTime:   utils.TimestamptzToTime(event.StartTime),
		EndTime:     utils.TimestamptzToTime(event.EndTime),
		CreatedAt:   utils.TimestamptzToTime(event.CreatedAt),
	}, nil
}

func (s ServiceImpl) CreateEvent(ctx context.Context, event model.Event) (uuid.UUID, error) {
	desc := event.Description
	newId := uuid.New()
	err := s.dbService.TX(ctx, func(ctx context.Context, conn *pgxpool.Conn, q *db.Queries) error {
		_, err := q.CreateEvent(ctx, db.CreateEventParams{
			ID:          newId,
			Title:       event.Title,
			Description: desc,
			StartTime:   utils.TimeToTimestamptz(time.Now()),
			EndTime:     utils.TimeToTimestamptz(time.Now().Add(time.Hour)),
			CreatedAt:   utils.TimeToTimestamptz(time.Now()),
		})
		return err
	})
	if err != nil {
		log.Println("TX failed:", err)
	}

	return newId, err
}

func NewService(
	dbService Database[db.Queries],
) *ServiceImpl {
	return &ServiceImpl{
		dbService: dbService,
	}
}
