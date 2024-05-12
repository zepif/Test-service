package service

import (
    "net"
    "net/http"

    "gitlab.com/distributed_lab/logan/v3/errors"
    "gitlab.com/distributed_lab/kit/copus/types"
    "github.com/zepif/Test-service/internal/config"
    "gitlab.com/distributed_lab/logan/v3"
)

type service struct {
    log      *logan.Entry
    copus    types.Copus
    listener net.Listener
}

func (s *service) run(cfg config.Config) error {
    s.log.Info("Service started")
    r := s.router(cfg)

    if err := s.copus.RegisterChi(r); err != nil {
        return errors.Wrap(err, "cop failed")
    }


    return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
    return &service{
        log:        cfg.Log(),
        copus:      cfg.Copus(),
        listener:   cfg.Listener(),
    }
}

func Run(cfg config.Config) {
    if err := newService(cfg).run(cfg); err != nil {
        panic(err)
    }
}
