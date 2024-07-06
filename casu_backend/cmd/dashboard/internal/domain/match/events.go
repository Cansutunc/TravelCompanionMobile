package match

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	WasCreatedType = (WasCreated{}).GetType()
	WasRemovedType = (WasRemoved{}).GetType()
	WasUpdatedType = (WasUpdated{}).GetType()
)

// GetType returns event type
func (e WasCreated) GetType() string {
	return fmt.Sprintf("%T", e)
}
func (e WasRemoved) GetType() string {
	return fmt.Sprintf("%T", e)
}
func (e WasUpdated) GetType() string {
	return fmt.Sprintf("%T", e)
}

// GetID the id
func (e WasCreated) GetCompetitionID() string {
	return e.CompetitionId.String()
}

// GetName the name
func (e WasCreated) GetSubDescription() string {
	return e.SubDescription
}

// GetCode the code
func (e WasCreated) GetPlayingTime() time.Time {
	return e.PlayingTime
}

// GetCode the code
func (e WasCreated) GetAwayTeamID() string {
	return e.AwayTeamId.String()
}

// GetCode the code
func (e WasCreated) GetHomeTeamID() string {
	return e.HomeTeamId.String()
}

func (e WasCreated) GetAreaID() string {
	return e.AreaId.String()
}

func (e WasCreated) GetMatchID() string {
	//TODO implement me
	return e.MatchID.String()
}

func (e WasCreated) GetYear() string {
	return e.Year
}

// WasCreated event
type WasCreated struct {
	MatchID        uuid.UUID `json:"match_id" bson:"match_id"`
	CompetitionId  uuid.UUID `json:"competition_id" bson:"competition_id"`
	SubDescription string    `json:"sub_description" bson:"sub_description"`
	PlayingTime    time.Time `json:"playing_time" bson:"playing_time"`
	HomeTeamId     uuid.UUID `json:"home_team_id" bson:"home_team_id"`
	AwayTeamId     uuid.UUID `json:"away_team_id" bson:"away_team_id"`
	Year           string    `json:"year" bson:"year"`
	AreaId         uuid.UUID `json:"area_id" bson:"area_id"`
}

// WasRemoved event
type WasRemoved struct {
	MatchID uuid.UUID `json:"match_id" bson:"match_id"`
}

// WasUpdated event
type WasUpdated struct {
	MatchID   uuid.UUID `json:"match_id" bson:"match_id"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Code      string    `json:"code,omitempty" bson:"code,omitempty"`
	CountryID uuid.UUID `json:"country_id,omitempty" bson:"country_id,omitempty"`
}
