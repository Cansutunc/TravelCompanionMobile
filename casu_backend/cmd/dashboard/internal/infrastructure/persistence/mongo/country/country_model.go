package mongo

type Country struct {
	CountryID string `json:"country_id" bson:"country_id"`
	Province  string `json:"province" bson:"province"`
	UserId    string `json:"user_id" bson:"user_id"`
	Country   string `json:"country" bson:"country"`
	UserName  string `json:"user_name" bson:"user_name"`
}

func (c *Country) GetCountryID() string {
	return c.CountryID
}

func (c *Country) GetProvince() string {
	return c.Province
}
func (c *Country) GetUserName() string {
	return c.UserName
}

func (c *Country) GetUserId() string {
	return c.UserId
}
func (c *Country) GetCountry() string {
	return c.Country
}
