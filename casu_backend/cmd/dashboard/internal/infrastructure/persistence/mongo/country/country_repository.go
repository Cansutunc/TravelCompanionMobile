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

type countryRepository struct {
	cfg        *config.Config
	collection *mongo.Collection
}

func NewCountryRepository(ctx context.Context, cfg *config.Config, mongoDb *mongo.Database) (persistence.CountryRepository, error) {
	collection := mongoDb.Collection("countries")

	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "country_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "name", Value: 1}, {Key: "code", Value: 1}}},
	}); err != nil {
		return nil, apperrors.Wrap(err)
	}

	return &countryRepository{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (r *countryRepository) Get(ctx context.Context, id string) (persistence.Country, error) {
	filter := bson.M{
		"country_id": id,
	}

	var result Country
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}
func (r *countryRepository) Add(ctx context.Context, c persistence.Country) error {
	country := Country{
		CountryID: c.GetCountryID(),
		Province:  c.GetProvince(),
		Country:   c.GetCountry(),
		UserId:    c.GetUserId(),
		UserName:  c.GetUserName(),
	}

	if _, err := r.collection.InsertOne(ctx, country); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *countryRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"country_id": id,
	}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *countryRepository) GetByID(ctx context.Context, id string) (persistence.Country, error) {
	c, err := r.Get(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	return c, nil
}

func (r *countryRepository) GetByName(ctx context.Context, name string) (persistence.Country, error) {
	filter := bson.M{
		"name": name,
	}

	var result Country
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *countryRepository) GetByUserIdAndProvince(ctx context.Context, userId string, province string) (persistence.Country, error) {
	filter := bson.M{
		"province": province,
		"user_id":  userId,
	}

	var result Country
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *countryRepository) FindAll(ctx context.Context, limit, offset int64) ([]persistence.Country, error) {
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	filter := bson.M{}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.Country
	for cur.Next(ctx) {
		var item Country
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}

func (r *countryRepository) Count(ctx context.Context) (int64, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, apperrors.Wrap(err)
	}

	return total, nil
}

func (r *countryRepository) Update(ctx context.Context, id, country string, province string, userId string) error {
	filter := bson.M{
		"country_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"country":  country,
			"user_id":  userId,
			"province": province,
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
