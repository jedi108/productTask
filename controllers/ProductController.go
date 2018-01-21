package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/jedi108/productTask/models"
)

func PostProducts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ResponseError(500, "error request", resp)
		return
	}
	products := &models.Products{}

	err = json.Unmarshal(buf[:len(buf)], &products)
	if err != nil {
		ResponseError(501, err.Error(), resp)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		fmt.Sprintf("sakura response: %q", buf)
		return
	}

	ids, err := models.NewProducts(products)
	if err != nil {
		ResponseError(501, err.Error(), resp)
		return
	}
	jsonIds, err := json.Marshal(ids)
	if err != nil {
		ResponseError(501, err.Error(), resp)
		return
	}
	resp.Write([]byte(jsonIds))
}

func GetProductById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	ids, ok := req.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		ResponseError(400, "not found key id for product", resp)
		return
	}

	id := ids[0]
	product, err := models.GetProductById(id)
	if err != nil {
		ResponseError(501, err.Error(), resp)
		return
	}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		ResponseError(500, err.Error(), resp)
		return
	}
	resp.Write([]byte(jsonProduct))
}
