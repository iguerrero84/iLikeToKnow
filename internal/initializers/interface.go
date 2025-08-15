package initializers

import (
	"iLikeToKnow.com/internal/database"
	db "iLikeToKnow.com/internal/database/db"
)

type IDatabaseService = database.Database[db.Queries]
