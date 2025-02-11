package handlers

import (
	"net/http"
)

type GameHandler interface {
	Ws(w http.ResponseWriter, r *http.Request)
}
