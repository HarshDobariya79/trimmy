package services

import (
	"errors"
	"trimmy/internal/storage"
)

func GenerateShortURL(url string) (string, error) {
	shortID, err := storage.GenerateShortURL(url)
	if err != nil {
		return "", err
	}
	return shortID, nil
}

func GetOriginalURL(shortID string) (string, error) {
	url, found := storage.GetURL(shortID)
	if !found {
		return "", errors.New("URL not found")
	}
	return url, nil
}
