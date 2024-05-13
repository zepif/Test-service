package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/go-chi/render"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	var request struct {
		OriginalURL string `json:"original_url"`
	}
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		log.WithError(err).Error("failed to decode request body")
		ape.RenderErr(w, problems.BadRequest(err)[0])
		return
	}

	originalURL := request.OriginalURL
	hash := sha256.Sum256([]byte(originalURL))
	ShortURL := base64.RawURLEncoding.EncodeToString(hash[:6])

	db := DB(r)

	linkQ := db.Link()
	if err := linkQ.Insert(originalURL, ShortURL); err != nil {
		log.WithError(err).Error("failed to save link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	log.WithField("ShortURL", ShortURL).Info("link created")
	ape.Render(w, map[string]string{"ShortURL": ShortURL})
}
