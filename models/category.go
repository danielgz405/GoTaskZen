package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   string             `bson:"created_at" json:"created_at"`
	Task        []Task             `bson:"task" json:"task"`
	Color       ColorBoard         `bson:"color" json:"color"`
	UserId      string             `bson:"user_id" json:"user_id"`
	TableId     string             `bson:"table_id" json:"table_id"`
}

type InsertCategory struct {
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	CreatedAt   string     `bson:"created_at" json:"created_at"`
	Color       ColorBoard `bson:"color" json:"color"`
	UserId      string     `bson:"user_id" json:"user_id"`
	TableId     string     `bson:"table_id" json:"table_id"`
}

type UpdateCategory struct {
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	Color       ColorBoard `bson:"color" json:"color"`
}

type Task struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   string             `bson:"created_at" json:"created_at"`
	Priority    string             `bson:"priority" json:"priority"`
	Active      bool               `bson:"active" json:"active"`
}

type InsertTask struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	CreatedAt   string `bson:"created_at" json:"created_at"`
	Priority    string `bson:"priority" json:"priority"`
	Active      bool   `bson:"active" json:"active"`
}

type UpdateTask struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Priority    string `bson:"priority" json:"priority"`
	Active      bool   `bson:"active" json:"active"`
}
