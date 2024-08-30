package storage

import (
	"context"
	"time"

	"my-redis-app/internal/models"
)

func FetchDataFromDB(ctx context.Context, id string) (*models.Data, error) {
	time.Sleep(1 * time.Second)

	data := &models.Data{
		ID:   id,
		Info: "Some information",
	}

	return data, nil
}
