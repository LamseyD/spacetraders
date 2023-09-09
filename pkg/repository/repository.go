package repository

import (
	"context"
	"log"
	"space-traders-playground/pkg/dom"
	"space-traders-playground/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Repository interface {
	AddUser(ctx context.Context, user dom.User) *dom.Error
}

type repository struct {
	client *mongo.Client
	db     *mongo.Database
	logger *zap.Logger
}

func NewRepository(conf Config, logger *zap.Logger) Repository {

	uri := conf.GetConnectionString()
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check if the connection is successful
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(conf.Database)

	return &repository{
		client: client,
		db:     db,
		logger: logger,
	}
}

// AddUser implements Repository.
func (r *repository) AddUser(ctx context.Context, user dom.User) *dom.Error {
	collection := r.db.Collection("users")
	// Insert the person document into the collection
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		r.logger.Error(errors.ErrFailedInsertMongo.Error(), zap.String("collection", "users"), zap.Error(err))
		return &errors.ErrFailedInsertMongo
	}

	return nil
}
