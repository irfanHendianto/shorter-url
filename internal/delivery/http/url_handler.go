package http

import (
	"encoding/json"
	"net/http"
	"shorter-url/internal/usecase"
)

type URLHandler struct {
	urlShortenerUseCase *usecase.ShortenURL
}

// NewURLHandler creates a new handler
func NewURLHandler(useCase *usecase.ShortenURL) *URLHandler {
	return &URLHandler{urlShortenerUseCase: useCase}
}

func (h *URLHandler) RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shortURL := r.URL.Path[len("/"):]

	// Call the use case to find the long URL
	longURL, err := h.urlShortenerUseCase.GetLongURL(ctx, shortURL)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original (long) URL
	http.Redirect(w, r, longURL, http.StatusFound)
}

// ShortenURLHandler handles the POST /shorten endpoint
func (h *URLHandler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		LongURL string `json:"long_url"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the use case to shorten the URL
	shortURL, err := h.urlShortenerUseCase.Execute(ctx, req.LongURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the shortened URL in the response
	resp := map[string]string{
		"short_url": shortURL,
	}
	json.NewEncoder(w).Encode(resp)
}
