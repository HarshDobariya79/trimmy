package storage

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var urlStore = make(map[string]string)

func SaveURL(shortID, url string) error {
	urlStore[shortID] = url
	return nil
}

func GetURL(id string) (string, bool) {
	url, found := urlStore[id]
	return url, found
}

func GenerateShortURL(url string) (string, error) {
	var shortID string
	var err error
	maxTries := 15

	for maxTries > 0 {
		// Generate a random 8-character ID
		shortID, err = gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 8)

		// Check if the ID already exists in the memory store
		_, exists := urlStore[shortID]

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
