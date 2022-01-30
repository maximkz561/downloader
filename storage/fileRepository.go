package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	res, err := fr.collection.InsertOne(ctx, f)
	if err != nil {
		return nil, err
	}
	mongoFileId := res.InsertedID.(primitive.ObjectID)
	f.Id = mongoFileId
	fmt.Println(res)
	return f, nil
}

func (fr *FileRepository) Find(id primitive.ObjectID) (*File, error) {
	file := &File{}
	filter := bson.D{{"_id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := fr.collection.FindOne(ctx, filter).Decode(&file)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, nil
	}
	return file, nil
}

func (fr *FileRepository) Update(id primitive.ObjectID, fields bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := fr.collection.UpdateByID(ctx, id, bson.D{{"$set", fields}})
	if err != nil {
		return err
	}
	return nil
}
