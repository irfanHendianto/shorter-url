package repository

import (
	"context"
	"shorter-url/internal/entity"
)

// URLRepository defines the repository interface for storing/retrieving URLs
type URLRepository interface {
	Save(ctx context.Context, shortURL *entity.ShortURL) error
	FindByShortURL(ctx context.Context, shortURL string) (*entity.ShortURL, error)
}
