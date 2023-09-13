package database

import (
	"context"

	"github.com/danielgz405/GoTaskZen/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertBoard(ctx context.Context, board *models.InsertBoard) (*models.Board, error) {
	collection := repo.client.Database("GoTaskZen").Collection("boards")
	result, err := collection.InsertOne(ctx, board)
	if err != nil {
		return nil, err
	}
	createdBoard, err := repo.GetBoardById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdBoard, nil
}

func (repo *MongoRepo) GetBoardById(ctx context.Context, id string) (*models.Board, error) {
	collection := repo.client.Database("GoTaskZen").Collection("boards")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var board models.Board
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&board)
	if err != nil {
		return nil, err
	}
	return &board, nil
}

func (repo *MongoRepo) ListBoards(ctx context.Context) ([]models.Board, error) {
	collection := repo.client.Database("GoTaskZen").Collection("boards")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var boards []models.Board
	if err = cursor.All(ctx, &boards); err != nil {
		return nil, err
	}
	return boards, nil
}

func (repo *MongoRepo) UpdateBoard(ctx context.Context, data *models.UpdateBoard, id string) (*models.Board, error) {
	collection := repo.client.Database("GoTaskZen").Collection("boards")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":                  data.Name,
		"description":           data.Description,
		"saved":                 data.Saved,
		"color":                 data.Color,
		"image":                 data.Image,
		"background":            data.Background,
		"desert_ref":            data.DesertRef,
		"desert_ref_background": data.DesertRefBackground,
	}
	for key, value := range iterableData {
		if value != nil && value != "" {
			update["$set"].(bson.M)[key] = value
		}
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, err
	}
	updatedBoard, err := repo.GetBoardById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedBoard, nil
}

func (repo *MongoRepo) DeleteBoard(ctx context.Context, id string) error {
	collection := repo.client.Database("GoTaskZen").Collection("boards")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
