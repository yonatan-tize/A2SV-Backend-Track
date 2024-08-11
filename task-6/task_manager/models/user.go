package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Username      string             `bson:"username" json:"username" validate:"required"`
	Password      string             `bson:"password" json:"password" validate:"required"`
	Role          string             `bson:"role" json:"role" validate:"required,eq=ADMIN|eq=USER"`
	Token         string             `json:"token"`
	Refresh_token string             `json:"refresh_token"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
