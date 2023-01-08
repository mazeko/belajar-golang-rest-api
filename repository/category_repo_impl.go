package repository

import (
	"belajar-goalng-rest-api/helper"
	"belajar-goalng-rest-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepoImpl struct {
}

func NewCategoryRepository() CategoryRepo {
	return &CategoryRepoImpl{}
}

func (repository *CategoryRepoImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicError(err)

	id, _ := result.LastInsertId()
	category.Id = int(id)
	return category
}

func (repository *CategoryRepoImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicError(err)

	return category
}

func (repository *CategoryRepoImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicError(err)
}

func (repository *CategoryRepoImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT * FROM category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		categories = append(categories, category)
	}

	return categories
}
