package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"shorter-url/internal/entity"
	"shorter-url/internal/repository"
)

var (
	ShortURLOptions string
	ShortURLLength  int
)

type ShortenURL struct {
	repo repository.URLRepository
}

// NewShortenURL creates a new use case instance
func NewShortenURL(repo repository.URLRepository) *ShortenURL {
	return &ShortenURL{repo: repo}
}

// Execute shortens a URL and saves it in the repository
func (u *ShortenURL) Execute(ctx context.Context, longURL string) (string, error) {
	if longURL == "" {
		return "", errors.New("long URL cannot be empty")
	}

	// Generate the short URL (this is just a simple example)
	shortURL := generateShortURL(longURL)

	urlEntity := &entity.ShortURL{
		ShortURL: shortURL,
		LongURL:  longURL,
	}

	err := u.repo.Save(ctx, urlEntity)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (u *ShortenURL) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	urlEntity, err := u.repo.FindByShortURL(ctx, shortURL)

	if err != nil {
		return "", errors.New("short URL not found")
	}

	return urlEntity.LongURL, nil
}

// generateShortURL generates a random short URL with the length of `shortURLLength`
func generateShortURL(longURL string) string {
	result := make([]byte, ShortURLLength)
	// Assuming result is filled with characters at this point
	fmt.Printf("Length of result: %d\n", len(result))

	for i := range result {
		// Get a random index from the allowed characters (letters and digits)
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(ShortURLOptions))))
		if err != nil {
			return "" // handle error appropriately in a real-world app
		}
		result[i] = ShortURLOptions[num.Int64()]
	}

	return string(result) // Return the generated random string as the short URL
}
