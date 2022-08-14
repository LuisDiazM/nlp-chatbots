package app

import "net/http"

// @Summary Index
// @Description This endpoint does not do nothing
// @Tags Hibot
// @Success 200 {object} object
// @Failure 401 {object} string
// @Accept  json
// @Produce  json
// @Router / [get]
func (app Application) Index(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}
