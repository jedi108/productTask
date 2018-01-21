package tests_functional

import (
	"github.com/jedi108/productTask/utils/tests"
	"encoding/json"
	"fmt"
	"testing"
	"github.com/jedi108/productTask/models"
)

func Test_InsertNullUrl(t *testing.T) {
	TestItem := &models.Product{
		Meta: models.Meta{
			Title: "Product1",
			Image: "/src/image.jpg",
			Price: "80",
		},
	}
	var TestItems models.Products
	TestItems = append(TestItems, *TestItem)
	jsonItem, err := json.Marshal(TestItems)
	if err != nil {
		panic(err)
	}
	req := tests.SendRequest("GET", GetUrlTest()+"/v1/PostProduct", []byte(jsonItem))
	fmt.Println(req.Body)
}

func Test_InsertItemCorrect(t *testing.T) {
	TestItem := &models.Product{
		Url: "sd",
		Meta: models.Meta{
			Title: "Product1",
			Image: "/src/image.jpg",
			Price: "80",
		},
	}

	var TestItems models.Products
	TestItems = append(TestItems, *TestItem)
	TestItems = append(TestItems, models.Product{
		Url: "ssd",
		Meta: models.Meta{
			Title: "Product2",
			Image: "/src/image2.jpg",
			Price: "800",
		},
	})

	TestItems = append(TestItems, models.Product{
		Url: "monitor",
		Meta: models.Meta{
			Title: "Product3",
			Image: "/src/image4.jpg",
			Price: "100",
		},
	})

	jsonItem, err := json.Marshal(TestItems)
	if err != nil {
		panic(err)
	}
	req := tests.SendRequest("GET", GetUrlTest()+"/v1/PostProduct", []byte(jsonItem))
	fmt.Println(req.Body)
}

/**
[
  {
    "url": "https://www.amazon.co.uk/gp/product/1509836071",
    "meta": {
        "title": "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
        "price": "8.49",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51kB2nKZ47L._SX382_BO1,204,203,200_.jpg",
    }
  },
  // ...
]
*/
