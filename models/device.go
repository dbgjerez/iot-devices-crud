package models

type Device struct {
	ID   string `json:"idDevice" bson:"_id"`
	Type string `json:"type" bson:"type"`
}

type UserRepository struct{}
