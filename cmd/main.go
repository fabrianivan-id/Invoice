package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	logger "esb-test/library/logger"
	"esb-test/src/app"
	"esb-test/src/middleware/request"
	v1 "esb-test/src/v1"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

//	@title			CMS Service
//	@version		1.0
//	@description	CMS Administrator API

// @in							header
// @name						Authorization
func main() {
	initCtx := context.Background()
	if err := app.Init(initCtx); err != nil {
		panic(err)
	}

	startService(initCtx)
}

func startService(ctx context.Context) {
	address := fmt.Sprintf(":%d", app.Config().BindAddress)
	//logger.GetLogger(ctx).Infof("Starting cms-service service on %s", address)

	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)
	r.Use(request.RequestIDContext(request.DefaultGenerator))
	r.Use(request.RequestAttributesContext)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	deps := v1.Dependencies(ctx)
	v1.Router(r, deps)

	err := http.ListenAndServe(address, r)
	if err != nil {
		logger.GetLogger(ctx).Errorf("ListenAndServe err %s", err)
	}
}
