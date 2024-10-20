package entity

// ShortURL represents the structure for storing shortened URLs.
type ShortURL struct {
	ID       string `bson:"_id,omitempty"` // MongoDB ID
	ShortURL string `bson:"shortURL"`
	LongURL  string `bson:"longURL"`
}
