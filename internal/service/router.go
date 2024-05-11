package service

import (
  "github.com/zepif/Test-service/internal/service/handlers"
  "github.com/go-chi/chi"
  "gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
  r := chi.NewRouter()

  r.Use(
    ape.RecoverMiddleware(s.log),
    ape.LoganMiddleware(s.log),
    ape.CtxMiddleware(
      handlers.CtxLog(s.log),
    ),
  )
  r.Route("/integrations/Test-service", func(r chi.Router) {
        r.Post("/create_short_link", handlers.GetNonce)
		r.Get("/get_long_link", handlers.RefreshToken)
  })

  return r
}
