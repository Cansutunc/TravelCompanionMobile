package mongo

type Area struct {
	AreaID string `json:"area_id" bson:"area_id"`
	CityID string `json:"city_id" bson:"city_id"`
	Name   string `json:"name" bson:"name"`
}

func (c *Area) GetAreaID() string { return c.AreaID }
func (c *Area) GetCityID() string { return c.CityID }
func (c *Area) GetName() string {
	return c.Name
}
