package services

import (
	"context"
	"fmt"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/config"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/area"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/book"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/city"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/competition"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/country"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/match"
	dashboardpersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type containerFactory func(ctx context.Context, cfg *config.Config) (*ServiceContainer, error)

// NewServiceContainer creates new container
var NewServiceContainer containerFactory

type ServiceContainer struct {
	Mongo                            *mongo.Client
	CommandBus                       commandbus.CommandBus
	EventBus                         eventbus.EventBus
	CountryRepository                country.Repository
	CityRepository                   city.Repository
	AreaRepository                   area.Repository
	CompetitionRepository            competition.Repository
	MatchRepository                  match.Repository
	BookRepository                   book.Repository
	CountryPersistenceRepository     dashboardpersistence.CountryRepository
	CityPersistenceRepository        dashboardpersistence.CityRepository
	AreaPersistenceRepository        dashboardpersistence.AreaRepository
	CompetitionPersistenceRepository dashboardpersistence.CompetitionRepository
	MatchPersistenceRepository       dashboardpersistence.MatchRepository
	BookPersistenceRepository        dashboardpersistence.BookRepository
}

func (c *ServiceContainer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	var errs []error
	go func() {
		defer wg.Done()
		if c.Mongo != nil {
			if err := c.Mongo.Disconnect(ctx); err != nil {
				errs = append(errs, err)
			}
		}
	}()

	wg.Wait()

	var closeErr error
	for _, err := range errs {
		if closeErr == nil {
			closeErr = err
		} else {
			closeErr = fmt.Errorf("%v | %v", closeErr, err)
		}
	}

	if closeErr != nil {
		return closeErr
	}

	return ctx.Err()
}
