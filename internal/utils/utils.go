package utils

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// TimestamptzToTime safely converts pgtype.Timestamptz to time.Time
func TimestamptzToTime(ts pgtype.Timestamptz) time.Time {
	if !ts.Valid {
		return time.Time{} // return zero time if null
	}
	return ts.Time
}

// Convert time.Time → pgtype.Timestamptz
func TimeToTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
