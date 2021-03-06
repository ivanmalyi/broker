package main

import (
	"github.com/gorilla/mux"
	"github.com/ivanmalyi/broker/internal/app/appserver"
	"github.com/ivanmalyi/broker/internal/app/handler"
	"github.com/ivanmalyi/broker/internal/app/repository/pgrepository"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(appserver.NewConfig),
		fx.Provide(appserver.NewLogger),
		fx.Provide(mux.NewRouter),
		fx.Provide(pgrepository.New),
		fx.Invoke(handler.New),
		fx.Invoke(appserver.RegisterHooks),
	).Run()
}
