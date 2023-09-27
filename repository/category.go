package repository

import (
	"context"

	"github.com/danielgz405/GoTaskZen/models"
)

func InsertCategory(ctx context.Context, category *models.InsertCategory) (*models.Category, error) {
	return implementation.InsertCategory(ctx, category)
}

func GetCategoryById(ctx context.Context, id string) (*models.Category, error) {
	return implementation.GetCategoryById(ctx, id)
}

func AddTaskToCategory(ctx context.Context, data models.InsertTask, categoryId string) (*models.Category, error) {
	return implementation.AddTaskToCategory(ctx, data, categoryId)
}

func UpdateTaskToCategory(ctx context.Context, data models.UpdateTask, taskId string, categoryId string) (*models.Category, error) {
	return implementation.UpdateTaskToCategory(ctx, data, taskId, categoryId)
}

func RemoveTaskToCategory(ctx context.Context, taskId string, categoryId string) (*models.Category, error) {
	return implementation.RemoveTaskToCategory(ctx, taskId, categoryId)
}

func ListCategorys(ctx context.Context, boardId string) ([]models.Category, error) {
	return implementation.ListCategorys(ctx, boardId)
}

func UpdateCategory(ctx context.Context, data *models.UpdateCategory, id string) (*models.Category, error) {
	return implementation.UpdateCategory(ctx, data, id)
}

func DeleteCategory(ctx context.Context, id string) error {
	return implementation.DeleteCategory(ctx, id)
}
