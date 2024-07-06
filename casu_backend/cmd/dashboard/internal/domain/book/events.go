package book

import (
	"fmt"
	"github.com/google/uuid"
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

func (e WasCreated) GetBookID() string {
	return e.BookID.String()
}
func (e WasCreated) GetName() string {
	return e.Name
}
func (e WasCreated) GetCode() string {
	return e.Code
}
func (e WasCreated) GetAuthor() string {
	return e.Author
}

func (e WasCreated) GetHomePage() string {
	return e.Homepage
}

// WasRemoved event
type WasRemoved struct {
	BookID uuid.UUID `json:"book_id" bson:"book_id"`
}

type WasCreated struct {
	BookID   uuid.UUID `json:"book_id" bson:"book_id"`
	Name     string    `json:"name" bson:"name"`
	Code     string    `json:"code" bson:"code"`
	Author   string    `json:"author" bson:"author"`
	Homepage string    `json:"homepage" bson:"homepage"`
}

type WasUpdated struct {
	BookID uuid.UUID `json:"book_id" bson:"book_id"`
	Name   string    `json:"name,omitempty" bson:"name,omitempty"`
	Code   string    `json:"code,omitempty" bson:"code,omitempty"`
}
