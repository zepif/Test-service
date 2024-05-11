package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/go-chi/render"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"github.com/zepif/Test-service/internal/data"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	var request struct {
		OriginalURL string `json:"original_url"`
	}
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		log.WithError(err).Error("failed to decode request body")
		ape.RenderErr(w, problems.BadRequest(err))
		return
	}

	originalURL := request.OriginalURL
	hash := sha256.Sum256([]byte(originalURL))
	shortURL := base64.RawURLEncoding.EncodeToString(hash[:6])
    
    db, ok := r.Context().Value(data.ContextKey).(data.URLStorage)
	if !ok {
		log.Error("failed to get data.URLStorage from context")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	linkQ := db.Link()
	if err := linkQ.Insert(originalURL, shortURL); err != nil {
		log.WithError(err).Error("failed to save link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	log.WithField("short_url", shortURL).Info("link created")
	ape.Render(w, map[string]string{"short_url": shortURL})
}
