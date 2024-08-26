package storage

import (
	"context"
	"trimmy/pkg/config"

	gonanoid "github.com/matoous/go-nanoid/v2"
	redis "github.com/redis/go-redis/v9"
)

var urlCache = redis.NewClient(&redis.Options{
	Addr:     config.Env["REDIS_HOST"],
	Password: config.Env["REDIS_PASSWORD"],
	DB:       0,
})

func SaveURL(shortID, url string) error {
	urlCache.Set(context.Background(), shortID, url, 0)
	return nil
}

func GetURL(id string) (string, bool) {
	url, err := urlCache.Get(context.Background(), id).Result()
	return url, err == nil
}

func GenerateShortURL(url string) (string, error) {
	var shortID string
	var err error
	maxTries := 15

	for maxTries > 0 {
		// Generate a random 8-character ID
		shortID, err = gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 8)

		if err != nil {
			return "", err
		}

		// Check if the ID already exists in the memory store
		_, exists := GetURL(shortID)

		if !exists {
			break
		}
		maxTries--
	}

	err = SaveURL(shortID, url)

	if err != nil {
		return "", err
	}

	return shortID, nil
}
