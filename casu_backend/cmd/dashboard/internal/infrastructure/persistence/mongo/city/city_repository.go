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

type cityRepository struct {
	cfg        *config.Config
	collection *mongo.Collection
}

func NewCityRepository(ctx context.Context, cfg *config.Config, mongoDb *mongo.Database) (persistence.CityRepository, error) {
	collection := mongoDb.Collection("cities")

	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "city_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "country_id", Value: 1}}, Options: options.Index().SetUnique(false)},
		{Keys: bson.D{{Key: "name", Value: 1}, {Key: "code", Value: 1}}},
	}); err != nil {
		return nil, apperrors.Wrap(err)
	}

	return &cityRepository{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (r *cityRepository) Get(ctx context.Context, id string) (persistence.City, error) {
	filter := bson.M{
		"city_id": id,
	}

	var result City
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *cityRepository) Add(ctx context.Context, c persistence.City) error {
	city := City{
		CityID:          c.GetCityID(),
		Title:           c.GetTitle(),
		Description:     c.GetDescription(),
		MaxPerson:       c.GetMaxPerson(),
		CreatedUserId:   c.GetCreatedUserId(),
		CreatedUsername: c.GetCreatedUsername(),
		AllPerson:       c.GetAllPerson(),
	}

	if _, err := r.collection.InsertOne(ctx, city); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *cityRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"city_id": id,
	}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *cityRepository) Count(ctx context.Context) (int64, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, apperrors.Wrap(err)
	}

	return total, nil
}

func (r *cityRepository) GetByID(ctx context.Context, id string) (persistence.City, error) {
	c, err := r.Get(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	return c, nil
}

func (r *cityRepository) GetByName(ctx context.Context, name string) (persistence.City, error) {
	filter := bson.M{
		"name": name,
	}

	var result City
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *cityRepository) FindAll(ctx context.Context, limit, offset int64) ([]persistence.City, error) {
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	filter := bson.M{}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.City
	for cur.Next(ctx) {
		var item City
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}

func (r *cityRepository) Update(ctx context.Context, id, name string, code string, countryId string) error {
	filter := bson.M{
		"city_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"name":       name,
			"code":       code,
			"country_id": countryId,
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

func (r *cityRepository) Accept(ctx context.Context, id, username string) error {
	filter := bson.M{
		"city_id": id,
	}
	update := bson.M{
		"$push": bson.M{
			"all_person": username,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return apperrors.Wrap(err)
	}
	return nil
}

func (r *cityRepository) Remove(ctx context.Context, id, username string) error {
	filter := bson.M{
		"city_id": id,
	}
	update := bson.M{
		"$pull": bson.M{
			"all_person": username,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return apperrors.Wrap(err)
	}
	return nil
}

func (r *cityRepository) FindAllAccepted(ctx context.Context, username string) ([]persistence.City, error) {
	findOptions := options.Find()
	filter := bson.M{
		"all_person": username,
	}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.City
	for cur.Next(ctx) {
		var item City
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}
