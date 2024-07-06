package mongo

type Competition struct {
	CompetitionID   string `json:"competition_id" bson:"competition_id"`
	Name            string `json:"name" bson:"name"`
	CompetitionType string `json:"competition_type" bson:"competition_type"`
	Region          string `json:"region" bson:"region"`
	SporType        string `json:"spor_type" bson:"spor_type"`
}

func (c *Competition) GetCompetitionID() string { return c.CompetitionID }
func (c *Competition) GetName() string          { return c.Name }
func (c *Competition) GetRegion() string {
	return c.Region
}
func (c *Competition) GetCompetitionType() string { return c.CompetitionType }
func (c *Competition) GetSporType() string        { return c.SporType }
