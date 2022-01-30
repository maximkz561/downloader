package storage

import (
	"context"
	"downloader/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Store struct {
	db             *mongo.Client
	FileRepository *FileRepository
}

func New() (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	config_ := config.GetConfig()
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s", config_.Mongo.Host, config_.Mongo.Port),
		),
	)
	if err != nil {
		print(1)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	store := &Store{
		db: client,
	}
	store.createFileRepository()
	return store, nil
}

func (s *Store) createFileRepository() *FileRepository {
	if s.FileRepository != nil {
		return s.FileRepository
	}
	config_ := config.GetConfig()
	collection := s.db.Database(config_.Mongo.DbName).Collection("files")

	s.FileRepository = &FileRepository{
		store:      s,
		collection: collection,
	}

	return s.FileRepository
}
