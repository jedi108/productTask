package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/jedi108/productTask/models"
	"github.com/jedi108/productTask/utils/errorsApi"
)

func PostProducts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	contextProduct := models.NewContextProduct()
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		contextProduct.ApiErrors.Add(errorsApi.Error{Code: 400, Message: "error request"})
		contextProduct.ApiErrors.ResponseError(resp)
		return
	}
	products := &models.Products{}

	err = json.Unmarshal(buf[:len(buf)], &products)
	if err != nil {
		contextProduct.ApiErrors.Add(errorsApi.Error{Code: 401, Message: "error data"})
		contextProduct.ApiErrors.ResponseError(resp)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		fmt.Sprintf("sakura response: %q", buf)
		return
	}

	ids := contextProduct.Insert(products)

	if contextProduct.HasError() {
		contextProduct.ApiErrors.ResponseError(resp)
		return
	}
	jsonIds, err := json.Marshal(ids)
	if err != nil {
		contextProduct.ApiErrors.Add(errorsApi.Error{Code: 401, Message: "error data"})
		contextProduct.ApiErrors.ResponseError(resp)
		return
	}
	resp.Write([]byte(jsonIds))
}

func GetProductById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	apiErrors := errorsApi.ApiErrors{}
	ids, ok := req.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		apiErrors.Add(errorsApi.Error{Code: 400, Message: "not found key id for product"})
		apiErrors.ResponseError(resp)
		return
	}

	id := ids[0]
	product, err := models.GetProductById(id)
	if err != nil {
		apiErrors.Add(errorsApi.Error{Code: 501, Message: err.Error()})
		apiErrors.ResponseError(resp)
		return
	}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		apiErrors.Add(errorsApi.Error{Code: 501, Message: err.Error()})
		apiErrors.ResponseError(resp)
		return
	}
	resp.Write([]byte(jsonProduct))
}
