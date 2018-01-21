package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Field   string
	Message string
	Code    int
}

type ApiErrors struct {
	Code    int
	Message string
	Errors  []Error
}

func ResponseError(HttpError int, s string, resp http.ResponseWriter) {
	log.Println(s)
	resp.WriteHeader(HttpError)
	resp.Write(GetApiError(s))
}

func GetApiError(error string) []byte {
	var s []Error
	errors, _ := json.Marshal(&ApiErrors{
		Code:    400,
		Message: "test",
		Errors: append(s, Error{
			Message: error,
		}),
	})
	return errors
}
