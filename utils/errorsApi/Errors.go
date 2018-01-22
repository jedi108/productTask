package errorsApi

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
	ErrorsMessage []string
}

func NewErrors() *ApiErrors  {
	return &ApiErrors{}
}

func (apiErrors *ApiErrors) Add(error Error)  {
	if error.Code > apiErrors.Code {
		apiErrors.Code = error.Code
	}
	apiErrors.Errors = append(apiErrors.Errors, error)
}

func (apiErrors *ApiErrors) ResponseError(resp http.ResponseWriter) {
	errorsJson, err := json.Marshal(apiErrors)
	if err != nil {
		resp.WriteHeader(503)
		resp.Write([]byte("server error json"))
		return
	}
	resp.WriteHeader(apiErrors.Code)
	resp.Write([]byte(errorsJson))
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
