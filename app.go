package main

import (
	"ShopItemsService/controllers"
	"fmt"
	"net/http"
	"github.com/jedi108/skylib/app"
	"time"
)

func main() {
	app.GetConnection()
	app.GetDB().SetMaxOpenConns(200)
	app.GetDB().SetMaxIdleConns(100)
	app.GetDB().SetConnMaxLifetime(time.Second * 5)
	app.InitLog()
	app.Log("Service started")

	fmt.Println("start service")
	http.HandleFunc("/v1/GetProduct", controllers.GetProductById)
	http.HandleFunc("/v1/PostProduct", controllers.PostProducts)
	if err := http.ListenAndServe(":8001", nil); err != nil {
		panic(err)
	}
}
