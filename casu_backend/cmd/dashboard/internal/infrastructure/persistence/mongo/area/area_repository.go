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

type areaRepository struct {
	cfg        *config.Config
	collection *mongo.Collection
}

func NewAreaRepository(ctx context.Context, cfg *config.Config, mongoDb *mongo.Database) (persistence.AreaRepository, error) {
	collection := mongoDb.Collection("areas")

	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "area_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "name", Value: 1}, {Key: "code", Value: 1}}},
	}); err != nil {
		return nil, apperrors.Wrap(err)
	}

	return &areaRepository{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (r *areaRepository) Get(ctx context.Context, id string) (persistence.Area, error) {
	filter := bson.M{
		"area_id": id,
	}

	var result Area
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *areaRepository) Add(ctx context.Context, c persistence.Area) error {
	area := Area{

		AreaID: c.GetAreaID(),
		Name:   c.GetName(),
		CityID: c.GetCityID(),
	}

	if _, err := r.collection.InsertOne(ctx, area); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *areaRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"area_id": id,
	}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *areaRepository) Count(ctx context.Context) (int64, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, apperrors.Wrap(err)
	}

	return total, nil
}

func (r *areaRepository) GetByID(ctx context.Context, id string) (persistence.Area, error) {
	c, err := r.Get(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	return c, nil
}

func (r *areaRepository) GetByName(ctx context.Context, name string) (persistence.Area, error) {
	filter := bson.M{
		"name": name,
	}

	var result Area
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *areaRepository) FindAll(ctx context.Context, limit, offset int64) ([]persistence.Area, error) {
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	filter := bson.M{}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.Area
	for cur.Next(ctx) {
		var item Area
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}

func (r *areaRepository) Update(ctx context.Context, id, name string, cityID string) error {
	filter := bson.M{
		"area_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"name":    name,
			"city_id": cityID,
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
