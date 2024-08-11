package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	DueDate     time.Time          `json:"due_date" bson:"due_date" validate:"required"`
	Status      string             `json:"status" bson:"status" validate:"required"`
}

