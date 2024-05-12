package service

import (
  "github.com/zepif/Test-service/internal/config"
  "github.com/zepif/Test-service/internal/service/handlers"
  "github.com/zepif/Test-service/internal/data/pg"
  "github.com/go-chi/chi"
  "gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
  r := chi.NewRouter()

  r.Use(
    ape.RecoverMiddleware(s.log),
    ape.LoganMiddleware(s.log),
    ape.CtxMiddleware(
      handlers.CtxLog(s.log),
      handlers.CtxDB(pg.NewStorage(cfg.DB())),
    ),
  )
  r.Route("/integrations/Test-service", func(r chi.Router) {
        r.Post("/create_short_link", handlers.CreateShortURL)
		r.Get("/get_long_link/{id}", handlers.GetShortLink) 
  })

  return r
}
