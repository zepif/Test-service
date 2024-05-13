package handlers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

type GetShortLinkResponse struct {
	FullURL string `json:"full_url"`
}

func GetShortLink(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	ShortURL := chi.URLParam(r, "ShortURL")

	db := DB(r)
	linkQ := db.Link()

	FullURL, err := linkQ.Get(ShortURL)
	if err != nil {
		if err == sql.ErrNoRows {
			log.WithError(err).Error("link not found")
			ape.RenderErr(w, problems.NotFound())
			return
		}
		log.WithError(err).Error("failed to get link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	log.WithFields(map[string]interface{}{
		"FullURL":  FullURL,
		"ShortURL": ShortURL,
	}).Info("link retrieved")

	resp := GetShortLinkResponse{FullURL: FullURL}
	ape.Render(w, resp)
}
