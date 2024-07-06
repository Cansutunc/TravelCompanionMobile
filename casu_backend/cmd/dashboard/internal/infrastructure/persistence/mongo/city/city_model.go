package mongo

type City struct {
	CityID          string   `json:"city_id" bson:"city_id"`
	Title           string   `json:"title" bson:"title"`
	Description     string   `json:"description" bson:"description"`
	MaxPerson       int      `json:"max_person" bson:"max_person"`
	CreatedUserId   string   `json:"created_user_id" bson:"created_user_id"`
	CreatedUsername string   `json:"created_username" bson:"created_username"`
	AllPerson       []string `json:"all_person,omitempty" bson:"all_person,omitempty"`
}

func (c *City) GetAllPerson() []string {
	return c.AllPerson
}

func (c *City) GetCityID() string { return c.CityID }
func (c *City) GetTitle() string {
	return c.Title
}
func (c *City) GetDescription() string {
	return c.Description
}
func (c *City) GetMaxPerson() int {
	return c.MaxPerson
}
func (c *City) GetCreatedUserId() string {
	return c.CreatedUserId
}
func (c *City) GetCreatedUsername() string {
	return c.CreatedUsername
}
