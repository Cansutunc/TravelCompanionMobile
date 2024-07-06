package mongo

import "time"

type Match struct {
	MatchID         string    `json:"match_id" bson:"match_id"`
	CompetitionId   string    `json:"competition_id" bson:"competition_id"`
	PlayingTime     time.Time `json:"playing_time" bson:"playing_time"`
	HomeTeamId      string    `json:"home_team_id" bson:"home_team_id"`
	AwayTeamId      string    `json:"away_team_id" bson:"away_team_id"`
	AreaId          string    `json:"area_id" bson:"area_id"`
	Winner          string    `json:"winner" bson:"winner"`
	SubDescription  string    `json:"sub_description" bson:"sub_description"`
	HomeTeamPlayers []string  `json:"home_team_players" bson:"home_team_players"`
	AwayTeamPlayers []string  `json:"away_team_players" bson:"away_team_players"`
	Year            string    `json:"year" bson:"year"`
}

func (c *Match) GetMatchID() string        { return c.MatchID }
func (c *Match) GetCompetitionID() string  { return c.CompetitionId }
func (c *Match) GetSubDescription() string { return c.SubDescription }
func (c *Match) GetPlayingTime() time.Time { return c.PlayingTime }
func (c *Match) GetAwayTeamID() string     { return c.AwayTeamId }
func (c *Match) GetHomeTeamID() string     { return c.HomeTeamId }
func (c *Match) GetYear() string           { return c.Year }
func (c *Match) GetAreaID() string         { return c.AreaId }
