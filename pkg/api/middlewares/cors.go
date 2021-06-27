package middleware

import "net/http"

func Cors(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		return
	}

	next(rw, r)
}
