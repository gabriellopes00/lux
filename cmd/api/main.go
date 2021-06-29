package main

import (
	"fmt"
	"helpy/config/env"
	"helpy/infra/db"
	builds "helpy/pkg/api/builds"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	db := db.GormPG{}

	fmt.Printf("%v :: Postgres connected successfully\n", time.Now().Local().Format(time.RFC3339))
	conn, err := db.Connect()
	if err != nil {
		log.Panicln(err)
	}

	if err = db.AutoMigrate(conn); err != nil {
		log.Panicln(err)
	}

	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	CreateUserHandler := builds.NewCreateUserHandler(conn)
	r.HandleFunc("/user", CreateUserHandler.Handle).Methods(http.MethodPost)
	DeleteUserHandler := builds.NewDeleteUserHandler(conn)
	r.HandleFunc("/user/{userId}", DeleteUserHandler.Handle).Methods(http.MethodDelete)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("localhost:%d", env.PORT),
		Handler:      r,
	}

	fmt.Printf(
		"%v :: Server running on port :%d\n",
		time.Now().Local().Format(time.RFC3339),
		env.PORT,
	)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
