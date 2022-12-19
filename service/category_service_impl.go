package service

import (
	"belajar-goalng-rest-api/helper"
	"belajar-goalng-rest-api/model/domain"
	"belajar-goalng-rest-api/model/web"
	"belajar-goalng-rest-api/repository"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepo
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	return category
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	panic("IMplement Me")
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	panic("Implement Me")
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	panic("Impelement Me")
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	panic("Implement Me")
}
