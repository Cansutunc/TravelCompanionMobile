package services

import (
	"context"
	"fmt"

	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/config"
	areapersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence/mongo/area"
	bookpersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence/mongo/book"
	citypersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence/mongo/city"
	competitionpersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence/mongo/competition"
	countrypersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence/mongo/country"
	matchpersistence "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence/mongo/match"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/repository"
	memorycommandbus "github.com/hsynrtn/dashboard-management/pkg/commandbus/memory"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	memoryeventbus "github.com/hsynrtn/dashboard-management/pkg/eventbus/memory"
	mongoeventstore "github.com/hsynrtn/dashboard-management/pkg/eventstore/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	NewServiceContainer = newMongoServiceContainer
}

func newMongoServiceContainer(ctx context.Context, cfg *config.Config) (*ServiceContainer, error) {
	commandBus := memorycommandbus.New(cfg.CommandBus.QueueSize)
	mongoConnection, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoDB.User, cfg.MongoDB.Pass, cfg.MongoDB.Host, cfg.MongoDB.Port),
	))
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	mongoDB := mongoConnection.Database(cfg.MongoDB.Database)
	eventStore, err := mongoeventstore.New(ctx, "events", mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	eventBus := memoryeventbus.New(cfg.EventBus.QueueSize)
	cityPersistenceRepository, err := citypersistence.NewCityRepository(ctx, cfg, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	cityRepository := repository.NewCityRepository(eventStore, eventBus)
	countryRepository := repository.NewCountryRepository(eventStore, eventBus)
	areaRepository := repository.NewAreaRepository(eventStore, eventBus)
	competitionRepository := repository.NewCompetitionRepository(eventStore, eventBus)
	matchRepository := repository.NewMatchRepository(eventStore, eventBus)
	bookRepository := repository.NewBookRepository(eventStore, eventBus)

	countryPersistenceRepository, err := countrypersistence.NewCountryRepository(ctx, cfg, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	areaPersistenceRepository, err := areapersistence.NewAreaRepository(ctx, cfg, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	competitionPersistenceRepository, err := competitionpersistence.NewCompetitionRepository(ctx, cfg, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	matchPersistenceRepository, err := matchpersistence.NewMatchRepository(ctx, cfg, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	bookPersistenceRepository, err := bookpersistence.NewBookRepository(ctx, cfg, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	//authenticator := auth.NewSecretAuthenticator([]byte(cfg.AuthClient.Secret))
	//claimsProvider := auth.NewClaimsProvider(authenticator)
	//tokenAuthorizer := auth.NewJWTTokenAuthorizer(grpAuthClient, claimsProvider, authenticator)

	return &ServiceContainer{
		Mongo:                            mongoConnection,
		CommandBus:                       commandBus,
		EventBus:                         eventBus,
		CityPersistenceRepository:        cityPersistenceRepository,
		CityRepository:                   cityRepository,
		CountryRepository:                countryRepository,
		CountryPersistenceRepository:     countryPersistenceRepository,
		AreaRepository:                   areaRepository,
		AreaPersistenceRepository:        areaPersistenceRepository,
		CompetitionRepository:            competitionRepository,
		CompetitionPersistenceRepository: competitionPersistenceRepository,
		MatchRepository:                  matchRepository,
		MatchPersistenceRepository:       matchPersistenceRepository,
		BookPersistenceRepository:        bookPersistenceRepository,
		BookRepository:                   bookRepository,
	}, nil
}
