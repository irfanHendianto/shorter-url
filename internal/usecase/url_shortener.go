package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"shorter-url/internal/entity"
	"shorter-url/internal/repository"
	"strconv"
)

var (
	ShortURLOptions string
	ShortURLLength  int
)

type ShortenURL struct {
	repo repository.URLRepository
}

func InitializeShortener() {
	// Parse SHORT_URL_OPTIONS as a string
	ShortURLOptions = os.Getenv("SHORT_URL_OPTIONS")

	// Parse SHORT_URL_LENGTH as an integer
	lengthStr := os.Getenv("SHORT_URL_LENGTH")
	if lengthStr != "" {
		var err error
		ShortURLLength, err = strconv.Atoi(lengthStr)
		if err != nil {
			fmt.Printf("Invalid SHORT_URL_LENGTH: %v\n", err)
			ShortURLLength = 0 // Set to default value if there's an error
		}
	} else {
		ShortURLLength = 6 // Default value if not set
	}
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
	shortURL := generateShortURL()

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
func generateShortURL() string {
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
