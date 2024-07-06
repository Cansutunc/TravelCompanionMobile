package main

import (
	"context"
	"fmt"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/config"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/services"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain"
	dashboardhttp "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/interfaces/http"
	"github.com/hsynrtn/dashboard-management/pkg/application"
	"github.com/hsynrtn/dashboard-management/pkg/buildinfo"
	httputils "github.com/hsynrtn/dashboard-management/pkg/http"
	"github.com/vardius/gocontainer"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	gocontainer.GlobalContainer = nil
}
func main() {
	buildinfo.PrintVersionOrContinue()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.FromEnv()
	fmt.Println("CONFIG:", cfg)

	container, err := services.NewServiceContainer(ctx, cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create service container: %w", err))
	}
	defer container.Close()

	if err := domain.RegisterCountryDomain(container); err != nil {
		panic(err)
	}
	if err := domain.RegisterCityDomain(container); err != nil {
		panic(err)
	}
	if err := domain.RegisterAreaDomain(container); err != nil {
		panic(err)
	}
	if err := domain.RegisterCompetitionDomain(container); err != nil {
		panic(err)
	}
	if err := domain.RegisterMatchDomain(container); err != nil {
		panic(err)
	}
	if err := domain.RegisterBookDomain(container); err != nil {
		panic(err)
	}

	if err := domain.CreateRegisterAllDomains(ctx, container); err != nil {
		panic(err)
	}

	router := dashboardhttp.NewRouter(cfg, container)
	app := application.New()

	app.AddAdapters(
		httputils.NewAdapter(
			&http.Server{
				Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
				ReadTimeout:  cfg.HTTP.ReadTimeout,
				WriteTimeout: cfg.HTTP.WriteTimeout,
				IdleTimeout:  cfg.HTTP.IdleTimeout, // limits server-side the amount of time a Keep-Alive connection will be kept idle before being reused
				Handler:      router,
			},
		))
	app.WithShutdownTimeout(cfg.App.ShutdownTimeout)
	app.Run(ctx)
}
