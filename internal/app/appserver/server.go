package appserver

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/ivanmalyi/broker/internal/app/repository"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

func RegisterHooks(
	lifecycle fx.Lifecycle,
	logger *logrus.Logger,
	repo repository.Repository,
	config *TomlConfig,
	mux *mux.Router,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Listening on ", config.BindAddr)
				go http.ListenAndServe(config.BindAddr, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				repo.Close()
				return nil
			},
		},
	)
}
