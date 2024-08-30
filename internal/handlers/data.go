package handlers

import (
	"net/http"
	"time"

	"my-redis-app/internal/models"
	"my-redis-app/internal/storage"

	"github.com/gin-gonic/gin"
)

func GetDataHandler(redisClient *storage.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var data models.Data
		err := redisClient.Get(c.Request.Context(), id, &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data from cache"})
			return
		}

		if data.ID != "" {
			c.JSON(http.StatusOK, data)
			return
		}

		dataPtr, err := storage.FetchDataFromDB(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data from database"})
			return
		}

		data = *dataPtr

		err = redisClient.Set(c.Request.Context(), id, data, 5*time.Minute)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set data in cache"})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
