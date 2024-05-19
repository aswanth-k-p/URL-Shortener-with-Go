package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Original string             `bson:"original" json:"original"`
	Short    string             `bson:"short" json:"short"`
}
