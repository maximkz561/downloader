package storage

import (
	"context"
	"downloader/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// FileRepository ...
type FileRepository struct {
	store      *Store
	collection *mongo.Collection
}

// Create ...
func (fr *FileRepository) Create(f *File) (*File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := fr.collection.InsertOne(ctx, f)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return f, nil
}

func (fr *FileRepository) Find(id string) (*File, error) {
	file := &File{}
	filter := bson.D{{"id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := fr.collection.FindOne(ctx, filter).Decode(&file)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		utils.Logger.Error(err)
		return nil, nil
	}
	return file, nil
}
