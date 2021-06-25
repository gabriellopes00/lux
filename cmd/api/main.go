package main

import (
	"fmt"
	"helpy/config/env"
	"helpy/infra/db"
	"helpy/infra/db/repositories"
	"helpy/infra/utils"
	"helpy/infra/validation"
	"helpy/pkg/server/handlers"
	usecase "helpy/pkg/user/usecases"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func main() {
	db := db.GormPG{}

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

	var CreateUserHandler = handlers.CreateUserHandler{
		Usecase: usecase.CreateUser{
			Uuid:      utils.UUIDGenerator{},
			Hasher:    utils.Argon2Hasher{},
			Validator: validation.UserGoValidator{},
			Repository: repositories.PgUserRepository{
				Db: conn,
			},
		},
	}

	r.HandleFunc("/user", CreateUserHandler.Handle).Methods(http.MethodPost)

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

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
