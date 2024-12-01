package pkg

import (
	"net/http"

	"github.com/go-chi/render"
)

func SuccessResp(w http.ResponseWriter, r *http.Request, d interface{}) {
	render.Status(r, http.StatusAccepted)
	render.JSON(w, r, d)
}

func ErrorResp(w http.ResponseWriter, r *http.Request, status int, err error) {
	render.Status(r, status)
	render.JSON(w, r, map[string]string{"message": err.Error()})
}
