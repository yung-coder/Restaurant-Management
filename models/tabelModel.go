package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tabel struct {
	ID              primitive.ObjectID `bson:"_id"`
	Number_of_guets *int               `json:"number_of_guets" validate:"required"`
	Tabel_number    *int               `json:"tabel_number" validate:"required"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	Tabel_id        string             `json:"table_id"`
}
