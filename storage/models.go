package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
}
