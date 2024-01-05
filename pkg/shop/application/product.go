package application

import (
	"github.com/pkg/errors"
	"github.com/ssr0016/monolithToMicro/pkg/common/price"
	"github.com/ssr0016/monolithToMicro/pkg/shop/domain"
)

type productReadModell interface {
	AllProdcuts() ([]products.Product, error)
}

type ProductsService struct {
	repo      domain.Repository
	readModel productReadModell
}

func NewProductsService(repo products.Repository, readModel productReadModell) productReadModell {
	return &ProductsService{
		repo,
		readModel,
	}
}

// Struct Method
func (s ProductsService) AllProducts() ([]products.Product, error) { // []products.Product {
	return s.readModel.AllProdcuts()
}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {

	price, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
	if err != nil {
		return errors.Wrap(err, "invalid product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.Wrap(err, "product could not be saved")
	}

	return nil

}
