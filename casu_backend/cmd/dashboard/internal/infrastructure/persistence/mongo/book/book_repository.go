package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/application/config"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/infrastructure/persistence"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bookRepository struct {
	cfg        *config.Config
	collection *mongo.Collection
}

func NewBookRepository(ctx context.Context, cfg *config.Config, mongoDb *mongo.Database) (persistence.BookRepository, error) {
	collection := mongoDb.Collection("books")

	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "book_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "name", Value: 1}, {Key: "code", Value: 1}}},
	}); err != nil {
		return nil, apperrors.Wrap(err)
	}

	return &bookRepository{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (r *bookRepository) Get(ctx context.Context, id string) (persistence.Book, error) {
	filter := bson.M{
		"book_id": id,
	}

	var result Book
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}
func (r *bookRepository) Add(ctx context.Context, c persistence.Book) error {
	book := Book{
		BookID:   c.GetBookID(),
		Name:     c.GetName(),
		Code:     c.GetCode(),
		Author:   c.GetAuthor(),
		HomePage: c.GetHomePage(),
	}

	if _, err := r.collection.InsertOne(ctx, book); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *bookRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"book_id": id,
	}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *bookRepository) GetByID(ctx context.Context, id string) (persistence.Book, error) {
	c, err := r.Get(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	return c, nil
}

func (r *bookRepository) GetByName(ctx context.Context, name string) (persistence.Book, error) {
	filter := bson.M{
		"name": name,
	}

	var result Book
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *bookRepository) FindAll(ctx context.Context, limit, offset int64) ([]persistence.Book, error) {
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	filter := bson.M{}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.Book
	for cur.Next(ctx) {
		var item Book
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}

func (r *bookRepository) Count(ctx context.Context) (int64, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, apperrors.Wrap(err)
	}

	return total, nil
}

func (r *bookRepository) Update(ctx context.Context, id, name string, code string) error {
	filter := bson.M{
		"book_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"name": name,
			"code": code,
		},
	}
	opt := options.FindOneAndUpdate().SetUpsert(false)

	if err := r.collection.FindOneAndUpdate(ctx, filter, update, opt).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return apperrors.Wrap(err)
	}

	return nil
}
