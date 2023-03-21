package service

import (
	"context"

	"github.com/devfullcycle/goexpert/tree/main/14-gRPC/internal/database"
	"github.com/devfullcycle/goexpert/tree/main/14-gRPC/internal/pb"
)

type Category struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CatcategoryResponse{
		Category: catecategoryResponse,
	}, nil
}
