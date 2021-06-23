package main

import (
	"fmt"
	"helpy/config/env"
	router "helpy/pkg/server/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	r := router.GetRouter()
	r.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("localhost:%d", env.PORT),
		Handler:      r,
	}

	fmt.Printf(
		"%v :: Server running at http://localhost:%d\n",
		time.Now().Local().Format(time.RFC3339),
		env.PORT,
	)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
