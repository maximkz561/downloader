package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title      string             `json:"title" bson:"title"`
	Downloaded bool               `json:"downloaded" bson:"downloaded"`
	FileId     string             `json:"file_id" bson:"file_id"`
	FormatId   string             `json:"format_id" bson:"format_id"`
	FileName   string             `json:"file_name" bson:"file_name"`
}
