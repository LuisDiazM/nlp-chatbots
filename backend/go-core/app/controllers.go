package app

import "net/http"

func (app Application) Index(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}
