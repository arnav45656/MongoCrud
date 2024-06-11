package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID    primitive.ObjectID `bson:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Age   int                `json:"age"`
}
