package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Board struct {
	Id                  primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name                string             `bson:"name" json:"name"`
	Description         string             `bson:"description" json:"description"`
	UserId              string             `bson:"user" json:"user"`
	Saved               bool               `bson:"saved" json:"saved"`
	Color               ColorBoard         `bson:"color" json:"color"`
	Image               string             `bson:"image" json:"image"`
	Background          string             `bson:"background" json:"background"`
	CreatedAt           string             `bson:"created_at" json:"created_at"`
	DesertRef           string             `bson:"desert_ref" json:"desert_ref"`
	DesertRefBackground string             `bson:"desert_ref_background" json:"desert_ref_background"`
}

type InsertBoard struct {
	Name                string     `bson:"name" json:"name"`
	Description         string     `bson:"description" json:"description"`
	UserId              string     `bson:"user" json:"user"`
	Saved               bool       `bson:"saved" json:"saved"`
	Color               ColorBoard `bson:"color" json:"color"`
	Image               string     `bson:"image" json:"image"`
	Background          string     `bson:"background" json:"background"`
	CreatedAt           string     `bson:"created_at" json:"created_at"`
	DesertRef           string     `bson:"desert_ref" json:"desert_ref"`
	DesertRefBackground string     `bson:"desert_ref_background" json:"desert_ref_background"`
}

type UpdateBoard struct {
	Name                string     `bson:"name" json:"name"`
	Description         string     `bson:"description" json:"description"`
	Saved               bool       `bson:"saved" json:"saved"`
	Color               ColorBoard `bson:"color" json:"color"`
	Image               string     `bson:"image" json:"image"`
	Background          string     `bson:"background" json:"background"`
	DesertRef           string     `bson:"desert_ref" json:"desert_ref"`
	DesertRefBackground string     `bson:"desert_ref_background" json:"desert_ref_background"`
}

type ColorBoard struct {
	Primary   string `bson:"primary" json:"primary"`
	Secondary string `bson:"secondary" json:"secondary"`
}
