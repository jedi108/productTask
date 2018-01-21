package tests_functional

import (
	"github.com/jedi108/productTask/utils/tests"
	"testing"
)

func TestGetFailKeyForProduct(t *testing.T) {
	req := tests.SendRequest("GET", GetUrlTest()+"/v1/GetProduct?id=999", []byte(""))
	if req.Response.StatusCode != 501 {
		t.Fatal(req.Response)
	}
}

func TestGetCorrectProduct(t *testing.T) {
	req := tests.SendRequest("GET", GetUrlTest()+"/v1/GetProduct?id=1", []byte(""))
	if req.Response.StatusCode != 200 {
		t.Fatal(req.Response)
	}
}
