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

type competitionRepository struct {
	cfg        *config.Config
	collection *mongo.Collection
}

func NewCompetitionRepository(ctx context.Context, cfg *config.Config, mongoDb *mongo.Database) (persistence.CompetitionRepository, error) {
	collection := mongoDb.Collection("competitons")

	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "competition_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "country_id", Value: 1}}, Options: options.Index().SetUnique(false)},
		{Keys: bson.D{{Key: "name", Value: 1}, {Key: "code", Value: 1}}},
	}); err != nil {
		return nil, apperrors.Wrap(err)
	}

	return &competitionRepository{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (r *competitionRepository) Get(ctx context.Context, id string) (persistence.Competition, error) {
	filter := bson.M{
		"competition_id": id,
	}

	var result Competition
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *competitionRepository) Add(ctx context.Context, c persistence.Competition) error {
	competition := Competition{
		CompetitionID:   c.GetCompetitionID(),
		Name:            c.GetName(),
		SporType:        c.GetSporType(),
		CompetitionType: c.GetCompetitionID(),
		Region:          c.GetRegion(),
	}

	if _, err := r.collection.InsertOne(ctx, competition); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *competitionRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"competition_id": id,
	}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *competitionRepository) Count(ctx context.Context) (int64, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, apperrors.Wrap(err)
	}

	return total, nil
}

func (r *competitionRepository) GetByID(ctx context.Context, id string) (persistence.Competition, error) {
	c, err := r.Get(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	return c, nil
}

func (r *competitionRepository) GetByName(ctx context.Context, name string) (persistence.Competition, error) {
	filter := bson.M{
		"name": name,
	}

	var result Competition
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *competitionRepository) FindAll(ctx context.Context, limit, offset int64) ([]persistence.Competition, error) {
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	filter := bson.M{}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.Competition
	for cur.Next(ctx) {
		var item Competition
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}

func (r *competitionRepository) Update(ctx context.Context, id, name string, code string, countryId string) error {
	filter := bson.M{
		"competition_id": id,
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
