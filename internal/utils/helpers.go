package utils

func GenerateShortURL(url string) (string, error) {
	var shortID string
	var err error

	for {
		// Generate a random 8-character ID
		shortID, err = gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 8)
		if err != nil {
			return "", err
		}

		// Check if the ID already exists in the database
		exists, err := storage.CheckIfIDExists(shortID)
		if err != nil {
			return "", err
		}

		if !exists {
			break
		}
	}

	err = storage.SaveURL(shortID, url)
	if err != nil {
		return "", err
	}

	return shortID, nil
}
