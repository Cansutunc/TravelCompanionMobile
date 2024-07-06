package persistence

import "context"

// Book the client persistence model interface
type Book interface {
	GetBookID() string
	GetName() string
	GetCode() string
	GetAuthor() string
	GetHomePage() string
}

// BookRepository allows to get/save current state of user to memory storage
type BookRepository interface {
	Get(ctx context.Context, id string) (Book, error)
	Add(ctx context.Context, client Book) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id string) (Book, error)
	GetByName(ctx context.Context, name string) (Book, error)
	FindAll(ctx context.Context, limit, offset int64) ([]Book, error)
	Update(ctx context.Context, id, name string, code string) error
}
