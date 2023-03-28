package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/config"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/data/pg"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/handlers"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/service/helpers"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxDB(pg.NewMasterQ(cfg.DB())),
			helpers.CtxServiceConfig(cfg.ServiceConfig()),
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
			helpers.CtxNetworkConnector(*cfg.NetworkConnector()),
		),
	)

	r.Route("/integrations/nonce-auth-svc", func(r chi.Router) {
		r.Post("/nonce", handlers.GetNonce)
		r.Post("/register", handlers.Register)
		r.Get("/refresh-token", handlers.RefreshToken)
		r.Get("/created-at", handlers.CreatedAt)

		r.Get("/validate", handlers.Validate)

		r.Route("/login", func(r chi.Router) {
			//r.Post("/", handlers.Login)
			r.Post("/admin", handlers.AdminLogin)
		})
	})

	return r
}
