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

type matchRepository struct {
	cfg        *config.Config
	collection *mongo.Collection
}

func NewMatchRepository(ctx context.Context, cfg *config.Config, mongoDb *mongo.Database) (persistence.MatchRepository, error) {
	collection := mongoDb.Collection("matches")

	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "match_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "country_id", Value: 1}}, Options: options.Index().SetUnique(false)},
		{Keys: bson.D{{Key: "name", Value: 1}, {Key: "code", Value: 1}}},
	}); err != nil {
		return nil, apperrors.Wrap(err)
	}

	return &matchRepository{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (r *matchRepository) Get(ctx context.Context, id string) (persistence.Match, error) {
	filter := bson.M{
		"match_id": id,
	}

	var result Match
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *matchRepository) Add(ctx context.Context, c persistence.Match) error {
	match := Match{
		MatchID:        c.GetMatchID(),
		CompetitionId:  c.GetCompetitionID(),
		SubDescription: c.GetSubDescription(),
		Year:           c.GetYear(),
		PlayingTime:    c.GetPlayingTime(),
		HomeTeamId:     c.GetHomeTeamID(),
		AwayTeamId:     c.GetAwayTeamID(),
		AreaId:         c.GetAreaID(),
	}

	if _, err := r.collection.InsertOne(ctx, match); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *matchRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"match_id": id,
	}

	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return apperrors.Wrap(err)
	}

	return nil
}

func (r *matchRepository) Count(ctx context.Context) (int64, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, apperrors.Wrap(err)
	}

	return total, nil
}

func (r *matchRepository) GetByID(ctx context.Context, id string) (persistence.Match, error) {
	c, err := r.Get(ctx, id)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}

	return c, nil
}

func (r *matchRepository) GetByName(ctx context.Context, name string) (persistence.Match, error) {
	filter := bson.M{
		"name": name,
	}

	var result Match
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperrors.Wrap(fmt.Errorf("%s: %w", err, apperrors.ErrNotFound))
		}
		return nil, apperrors.Wrap(err)
	}

	return &result, nil
}

func (r *matchRepository) FindAll(ctx context.Context, limit, offset int64) ([]persistence.Match, error) {
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	filter := bson.M{}

	cur, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	defer cur.Close(ctx)

	var result []persistence.Match
	for cur.Next(ctx) {
		var item Match
		if err := cur.Decode(&item); err != nil {
			return nil, apperrors.Wrap(err)
		}

		result = append(result, &item)
	}

	return result, nil
}

func (r *matchRepository) Update(ctx context.Context, id, name string, code string, countryId string) error {
	filter := bson.M{
		"match_id": id,
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
