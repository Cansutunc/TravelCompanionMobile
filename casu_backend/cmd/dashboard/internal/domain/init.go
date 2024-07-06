package domain

import (
	"context"
	areaeventhandler "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/eventhandler/area"
	bookeventhandler "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/eventhandler/book"
	cityeventhandler "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/eventhandler/city"
	competitioneventhandler "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/eventhandler/competition"
	countryeventhandler "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/eventhandler/country"
	matcheventhandler "github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/eventhandler/match"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/services"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/area"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/book"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/city"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/competition"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/country"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/match"
	"github.com/hsynrtn/dashboard-management/pkg/commandbus"
	"github.com/hsynrtn/dashboard-management/pkg/domain"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
)

type RegisterDomain struct {
	eventTypeName  string
	eventFuncType  interface{}
	commandBusName string
	commandHandler commandbus.CommandHandler
	eventHandler   eventbus.EventHandler
}

var Domains = make(map[string]RegisterDomain)

func RegisterCountryDomain(container *services.ServiceContainer) error {
	Domains[country.CreateCountry] = RegisterDomain{eventTypeName: country.WasCreatedType, eventFuncType: &country.WasCreated{}, commandBusName: country.CreateName, commandHandler: country.OnCreate(container.CountryRepository, container.CountryPersistenceRepository), eventHandler: countryeventhandler.WhenCountryWasCreated(container.CountryPersistenceRepository)}
	Domains[country.RemoveCountry] = RegisterDomain{eventTypeName: country.WasRemovedType, eventFuncType: &country.WasRemoved{}, commandBusName: country.RemoveName,
		commandHandler: country.OnRemove(container.CountryRepository),
		eventHandler:   countryeventhandler.WhenCountryWasRemoved(container.CountryPersistenceRepository),
	}
	Domains[country.UpdateCountry] = RegisterDomain{
		eventTypeName:  country.WasUpdatedType,
		eventFuncType:  &country.WasUpdated{},
		commandBusName: country.UpdateName,
		commandHandler: country.OnUpdate(container.CountryRepository, container.CountryPersistenceRepository),
		eventHandler:   countryeventhandler.WhenCountryWasUpdated(container.CountryPersistenceRepository),
	}
	return nil
}

func RegisterBookDomain(container *services.ServiceContainer) error {
	Domains[book.CreateBook] = RegisterDomain{eventTypeName: book.WasCreatedType, eventFuncType: &book.WasCreated{}, commandBusName: book.CreateName, commandHandler: book.OnCreate(container.BookRepository, container.BookPersistenceRepository), eventHandler: bookeventhandler.WhenBookWasCreated(container.BookPersistenceRepository)}
	Domains[book.RemoveBook] = RegisterDomain{eventTypeName: book.WasRemovedType, eventFuncType: &book.WasRemoved{}, commandBusName: book.RemoveName,
		commandHandler: book.OnRemove(container.BookRepository),
		eventHandler:   bookeventhandler.WhenBookWasRemoved(container.BookPersistenceRepository),
	}
	Domains[book.UpdateBook] = RegisterDomain{
		eventTypeName:  book.WasUpdatedType,
		eventFuncType:  &book.WasUpdated{},
		commandBusName: book.UpdateName,
		commandHandler: book.OnUpdate(container.BookRepository, container.BookPersistenceRepository),
		eventHandler:   bookeventhandler.WhenBookWasUpdated(container.BookPersistenceRepository),
	}
	return nil
}

func RegisterCityDomain(container *services.ServiceContainer) error {
	Domains[city.CreateCity] = RegisterDomain{
		eventTypeName:  city.WasCreatedType,
		eventFuncType:  &city.WasCreated{},
		commandBusName: city.CreateName,
		commandHandler: city.OnCreate(container.CityRepository, container.CityPersistenceRepository),
		eventHandler:   cityeventhandler.WhenCityWasCreated(container.CityPersistenceRepository),
	}

	Domains[city.RemoveCity] = RegisterDomain{
		eventTypeName:  city.WasRemovedType,
		eventFuncType:  &city.WasRemoved{},
		commandBusName: city.RemoveName,
		commandHandler: city.OnRemove(container.CityRepository),
		eventHandler:   cityeventhandler.WhenCityWasRemoved(container.CityPersistenceRepository),
	}

	Domains[city.UpdateCity] = RegisterDomain{
		eventTypeName:  city.WasUpdatedType,
		eventFuncType:  &city.WasUpdated{},
		commandBusName: city.UpdateName,
		commandHandler: city.OnUpdate(container.CityRepository, container.CityPersistenceRepository),
		eventHandler:   cityeventhandler.WhenCityWasUpdated(container.CityPersistenceRepository),
	}

	return nil
}

func RegisterAreaDomain(container *services.ServiceContainer) error {
	Domains[area.CreateArea] = RegisterDomain{
		eventTypeName:  area.WasCreatedType,
		eventFuncType:  &area.WasCreated{},
		commandBusName: area.CreateName,
		commandHandler: area.OnCreate(container.AreaRepository, container.AreaPersistenceRepository),
		eventHandler:   areaeventhandler.WhenAreaWasCreated(container.AreaPersistenceRepository),
	}

	Domains[area.RemoveArea] = RegisterDomain{
		eventTypeName:  area.WasRemovedType,
		eventFuncType:  &area.WasRemoved{},
		commandBusName: area.RemoveName,
		commandHandler: area.OnRemove(container.AreaRepository),
		eventHandler:   areaeventhandler.WhenAreaWasRemoved(container.AreaPersistenceRepository),
	}

	Domains[area.UpdateArea] = RegisterDomain{
		eventTypeName:  area.WasUpdatedType,
		eventFuncType:  &area.WasUpdated{},
		commandBusName: area.UpdateName,
		commandHandler: area.OnUpdate(container.AreaRepository, container.AreaPersistenceRepository),
		eventHandler:   areaeventhandler.WhenAreaWasUpdated(container.AreaPersistenceRepository),
	}

	return nil
}

func RegisterCompetitionDomain(container *services.ServiceContainer) error {
	Domains[competition.CreateCompetition] = RegisterDomain{
		eventTypeName:  competition.WasCreatedType,
		eventFuncType:  &competition.WasCreated{},
		commandBusName: competition.CreateName,
		commandHandler: competition.OnCreate(container.CompetitionRepository, container.CompetitionPersistenceRepository),
		eventHandler:   competitioneventhandler.WhenCompetitionWasCreated(container.CompetitionPersistenceRepository),
	}

	Domains[competition.RemoveCompetition] = RegisterDomain{
		eventTypeName:  competition.WasRemovedType,
		eventFuncType:  &competition.WasRemoved{},
		commandBusName: competition.RemoveName,
		commandHandler: competition.OnRemove(container.CompetitionRepository),
		eventHandler:   competitioneventhandler.WhenCompetitionWasRemoved(container.CompetitionPersistenceRepository),
	}

	Domains[competition.UpdateCompetition] = RegisterDomain{
		eventTypeName:  competition.WasUpdatedType,
		eventFuncType:  &competition.WasUpdated{},
		commandBusName: competition.UpdateName,
		commandHandler: competition.OnUpdate(container.CompetitionRepository, container.CompetitionPersistenceRepository),
		eventHandler:   competitioneventhandler.WhenCompetitionWasUpdated(container.CompetitionPersistenceRepository),
	}

	return nil
}

func RegisterMatchDomain(container *services.ServiceContainer) error {
	Domains[match.CreateMatch] = RegisterDomain{
		eventTypeName:  match.WasCreatedType,
		eventFuncType:  &match.WasCreated{},
		commandBusName: match.CreateName,
		commandHandler: match.OnCreate(container.MatchRepository, container.MatchPersistenceRepository),
		eventHandler:   matcheventhandler.WhenMatchWasCreated(container.MatchPersistenceRepository),
	}

	Domains[match.RemoveMatch] = RegisterDomain{
		eventTypeName:  match.WasRemovedType,
		eventFuncType:  &match.WasRemoved{},
		commandBusName: match.RemoveName,
		commandHandler: match.OnRemove(container.MatchRepository),
		eventHandler:   matcheventhandler.WhenMatchWasRemoved(container.MatchPersistenceRepository),
	}

	Domains[match.UpdateMatch] = RegisterDomain{
		eventTypeName:  match.WasUpdatedType,
		eventFuncType:  &match.WasUpdated{},
		commandBusName: match.UpdateName,
		commandHandler: match.OnUpdate(container.MatchRepository, container.MatchPersistenceRepository),
		eventHandler:   matcheventhandler.WhenMatchWasUpdated(container.MatchPersistenceRepository),
	}

	return nil
}

func CreateRegisterAllDomains(ctx context.Context, container *services.ServiceContainer) error {
	for _, v := range Domains {
		if err := domain.RegisterEventFactory(v.eventTypeName, func() interface{} { return v.eventFuncType }); err != nil {
			return apperrors.Wrap(err)
		}
		if v.commandHandler != nil {
			if err := container.CommandBus.Subscribe(ctx, v.commandBusName, v.commandHandler); err != nil {
				return apperrors.Wrap(err)
			}
		}
		if err := container.EventBus.Subscribe(ctx, v.eventTypeName, v.eventHandler); err != nil {
			return apperrors.Wrap(err)
		}

	}
	return nil
}
