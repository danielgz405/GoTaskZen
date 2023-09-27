package database

import (
	"context"

	"github.com/danielgz405/GoTaskZen/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertCategory(ctx context.Context, category *models.InsertCategory) (*models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}
	createdCategory, err := repo.GetCategoryById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (repo *MongoRepo) GetCategoryById(ctx context.Context, id string) (*models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var category models.Category
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *MongoRepo) AddTaskToCategory(ctx context.Context, data models.InsertTask, categoryId string) (*models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	oid, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return nil, err
	}
	taskOid := primitive.NewObjectID()
	task := &models.Task{
		Id:          taskOid,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		Priority:    data.Priority,
		Active:      data.Active,
	}
	update := bson.M{"$push": bson.M{"task": task}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, err
	}
	category, err := repo.GetCategoryById(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *MongoRepo) UpdateTaskToCategory(ctx context.Context, data models.UpdateTask, taskId string, categoryId string) (*models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	oid, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return nil, err
	}
	taskOid, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, err
	}
	update := bson.M{"$set": bson.M{"task.$[elem].name": data.Name, "task.$[elem].description": data.Description, "task.$[elem].priority": data.Priority, "task.$[elem].active": data.Active}}
	filter := bson.M{"_id": oid}
	_, err = collection.UpdateOne(ctx, filter, update, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem._id": taskOid},
		},
	}))
	if err != nil {
		return nil, err
	}
	category, err := repo.GetCategoryById(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *MongoRepo) RemoveTaskToCategory(ctx context.Context, taskId string, categoryId string) (*models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	oid, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return nil, err
	}
	taskOid, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return nil, err
	}
	update := bson.M{"$pull": bson.M{"task": bson.M{"_id": taskOid}}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, err
	}
	task, err := repo.GetCategoryById(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (repo *MongoRepo) ListCategorys(ctx context.Context, boardId string) ([]models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	cursor, err := collection.Find(ctx, bson.M{"table_id": boardId})
	if err != nil {
		return nil, err
	}
	var category []models.Category
	if err = cursor.All(ctx, &category); err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *MongoRepo) UpdateCategory(ctx context.Context, data *models.UpdateCategory, id string) (*models.Category, error) {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":        data.Name,
		"description": data.Description,
		"color":       data.Color,
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
	updatedCategory, err := repo.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil
}

func (repo *MongoRepo) DeleteCategory(ctx context.Context, id string) error {
	collection := repo.client.Database("GoTaskZen").Collection("categories")
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
