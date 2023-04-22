package database

import (
	"github.com/oaraujocesar/go-api/internal/entity"
	uuid "github.com/oaraujocesar/go-api/pkg/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) Delete(id uuid.ID) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(&entity.Product{}, id).Error
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID)
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) FindByID(id uuid.ID) (*entity.Product, error) {
	var product entity.Product

	err := p.DB.First(&product, "id = ?", id).Error

	return &product, err
}
