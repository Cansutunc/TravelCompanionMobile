package mongo

type Book struct {
	BookID   string `json:"book_id" bson:"book_id"`
	Name     string `json:"name" bson:"name"`
	Code     string `json:"code" bson:"code"`
	Author   string `json:"author" bson:"author"`
	HomePage string `json:"homepage" bson:"homepage"`
}

func (c *Book) GetBookID() string {
	return c.BookID
}

func (c *Book) GetName() string {
	return c.Name
}

func (c *Book) GetCode() string {
	return c.Code
}
func (c *Book) GetAuthor() string {
	return c.Author
}

func (c *Book) GetHomePage() string {
	return c.HomePage
}
