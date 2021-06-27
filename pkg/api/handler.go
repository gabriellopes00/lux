package server

import "net/http"

type Controller interface {
	Handle(rw http.ResponseWriter, r *http.Request)
}
