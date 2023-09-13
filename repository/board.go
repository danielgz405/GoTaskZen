package repository

import (
	"context"

	"github.com/danielgz405/GoTaskZen/models"
)

func InsertBoard(ctx context.Context, user *models.InsertBoard) (*models.Board, error) {
	return implementation.InsertBoard(ctx, user)
}

func GetBoardById(ctx context.Context, id string) (*models.Board, error) {
	return implementation.GetBoardById(ctx, id)
}

func ListBoards(ctx context.Context) ([]models.Board, error) {
	return implementation.ListBoards(ctx)
}

func DeleteBoard(ctx context.Context, id string) error {
	return implementation.DeleteBoard(ctx, id)
}

func UpdateBoard(ctx context.Context, data *models.UpdateBoard, id string) (*models.Board, error) {
	return implementation.UpdateBoard(ctx, data, id)
}
