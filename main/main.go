package main

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"yaTrivoga/controllers"
)

/**

Google OAuth:



*/

func main() {

	db, err := getDb()
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	srv := &http.Server{
		Addr:           "localhost:8080",
		Handler:        handlers.CompressHandler(router(db, rdb)),
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 0,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func getDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/raid_alert")
	if err != nil {
		log.Panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}

func router(mysql *sql.DB, redis *redis.Client) http.Handler {

	r := mux.NewRouter()
	r.StrictSlash(true)

	healthCheckController := controllers.HealthCheckController{}
	healthCheckController.SetConnections(mysql, redis)

	cityController := controllers.CityController{}
	cityController.SetConnections(mysql, redis)

	//	APIs:
	//		/city  GET - get all cities
	//		/city/status GET - get all cities' status
	//		/city/status/{id} GET - get city status
	//		/city/status/{id} POST - submit status (yes|no) - cancel option for yes

	r.HandleFunc("/healthcheck", healthCheckController.HealthCheck).Methods("GET")

	r.HandleFunc("/city", cityController.GetAll).Methods("GET")

	r.HandleFunc("/city/status", cityController.GetAll).Methods("GET")
	r.HandleFunc("/city/status/{cityId}", cityController.GetAll).Methods("GET")
	r.HandleFunc("/city/status/{cityId}", cityController.GetAll).Methods("POST")

	return r

}
