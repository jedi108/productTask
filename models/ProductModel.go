package models

import (
	"github.com/jedi108/skylib/app"
	"database/sql"
	"gopkg.in/guregu/null.v3/zero"
	"github.com/jedi108/productTask/utils/errorsApi"
	"errors"
)

type ContextProducts struct {
	Products  *Products
	ApiErrors errorsApi.ApiErrors
	tx        *sql.Tx
}

func NewContextProduct() *ContextProducts {
	return &ContextProducts{}
}

type Products []Product

type Product struct {
	Id   int64
	Url  string `json:"url"`
	Meta Meta   `json:"meta"`
}

type Meta struct {
	Title   string    `json:"title"`
	Price   string    `json:"price"`
	Image   string    `json:"image"`
	Aviable zero.Bool `json:"aviable"`
}

func GetProductById(id string) (*Product, error) {
	product := &Product{Meta: Meta{}}
	err := app.GetDB().QueryRow(`
		SELECT 
			id, url, title, price, image, IFNULL (aviable, FALSE)
		FROM 
			Product 
		WHERE id = ?`, id).Scan(&product.Id, &product.Url, &product.Meta.Title, &product.Meta.Price, &product.Meta.Image, &product.Meta.Aviable)
	return product, err
}

func (contxProd *ContextProducts) Insert(products *Products) []int64 {
	contxProd.Products = products
	var Ids []int64
	tx, err := app.GetDB().Begin()
	if err != nil {
		contxProd.ApiErrors.Add(errorsApi.Error{Code: 500, Message: "Error connection"})
		app.Log(err.Error())
		return []int64{}
	}
	contxProd.tx = tx
	for _, product := range *contxProd.Products {
		Id, err := contxProd.InsertProduct(&product)
		Ids = append(Ids, Id)
		if err != nil {
			tx.Rollback()
			app.Log(err.Error())
			contxProd.ApiErrors.Add(errorsApi.Error{Code: 500, Message: "Error server data"})
			return Ids
		}
	}
	err = tx.Commit()
	return Ids
}

func (contxProd *ContextProducts) ProductValidate(product *Product) {
	if product.Url == "" {
		contxProd.ApiErrors.Add(errorsApi.Error{Code: 422, Field: "Url", Message: "is empty"})
	}
	if product.Meta.Title == "" {
		contxProd.ApiErrors.Add(errorsApi.Error{Code: 422, Field: "Meta.Title", Message: "is empty"})
	}
	if product.Meta.Image == "" {
		contxProd.ApiErrors.Add(errorsApi.Error{Code: 422, Field: "Meta.Image", Message: "is empty"})
	}
	if product.Meta.Price == "" {
		contxProd.ApiErrors.Add(errorsApi.Error{Code: 422, Field: "Meta.Price", Message: "is empty"})
	}
}

func (contxProd *ContextProducts) HasError() bool {
	if len(contxProd.ApiErrors.Errors) > 0 {
		return true
	}
	return false
}

func (contxProd *ContextProducts) InsertProduct(product *Product) (int64, error) {
	contxProd.ProductValidate(product)
	if contxProd.HasError() {
		return 0, errors.New("Validate error")
	}
	result, err := app.GetDB().Exec(`
		INSERT INTO
			Product (
				Title,
				Url,
				Price,
				Image
			)
			VALUES
				(?,?,?,?)`,
		product.Meta.Title,
		product.Url,
		product.Meta.Price,
		product.Meta.Image,
	)

	if err != nil {
		return 0, err
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return LastInsertId, err
}
