package service

import (
	"context"
	"nishant/internal/data"
	"nishant/internal/models"
)

type ProductService struct {
	productStore *data.ProductStore
}

func NewProductService(productStore *data.ProductStore) *ProductService {
	return &ProductService{
		productStore: productStore,
	}
}

func (s *ProductService) GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	products, err := s.productStore.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	return s.productStore.GetProductById(ctx, id)
}
