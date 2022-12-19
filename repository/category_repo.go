package repository

import (
	"belajar-goalng-rest-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepo interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindAll(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
}
