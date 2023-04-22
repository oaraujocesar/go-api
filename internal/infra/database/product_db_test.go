package database

import (
	"testing"

	"github.com/oaraujocesar/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.NotEmpty(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.NotNil(t, err)
	assert.Empty(t, productFound)
}

func TestDeleteProductWhenProductIsNotFound(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Delete(product.ID)
	assert.NotNil(t, err)
}

func TestUpdateProduct(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	product.Name = "Product 2"
	product.Price = 20.0

	err = productDB.Update(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.NotEmpty(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProductWhenProductIsNotFound(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Update(product)
	assert.NotNil(t, err)
}

func TestFindProductById(t *testing.T) {
	db, err := createInMemoryDatabase(t, &entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID)

	assert.Nil(t, err)
	assert.NotEmpty(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}
