package models

import (
	"github.com/jedi108/skylib/app"
	"database/sql"
	"errors"
)

type Products []Product


type Product struct {
	Id   int64
	Url  string `json:"url"`
	Meta Meta   `json:"meta"`
}

type Meta struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Image string `json:"image"`
}

func GetProductById(id string) (*Product, error) {
	product := &Product{Meta: Meta{}}
	err := app.GetDB().QueryRow(`
		SELECT 
			id, url, title, price, image 
		FROM 
			Product 
		WHERE id = ?`, id).Scan(&product.Id, &product.Url, &product.Meta.Title, &product.Meta.Price, &product.Meta.Image)
	return product, err
}

func NewProducts(products *Products) ([]int64, error) {
	var Ids []int64
	tx, err := app.GetDB().Begin()
	if err != nil {
		return []int64{}, errors.New("Error connection")
	}
	for _, product := range *products {
		Id, err := product.Insert(tx)
		Ids = append(Ids, Id)
		if err != nil {
			tx.Rollback()
			return Ids, err
		}
	}
	err = tx.Commit()
	return Ids, err
}

func (product *Product) Insert(tx *sql.Tx) (int64, error) {
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
