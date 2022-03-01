package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"yaTrivoga/controllers"
)

type App struct {
	Router *mux.Router
	MySQL  *sql.DB
	Redis  *redis.Client
}

func (a *App) Initialize(
	dbUser, dbPwd, dbHost, dbPort, dbName,
	rdHost, rdPort, rdPwd string) {

	a.MySQL = a.getDbConnection(dbUser, dbPwd, dbHost, dbPort, dbName)
	a.Redis = a.getRedisConnection(rdHost, rdPort, rdPwd)

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {

	srv := &http.Server{
		Addr:           "localhost:8080",
		Handler:        handlers.CompressHandler(a.Router),
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 0,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) initializeRoutes() {

	a.Router.StrictSlash(true)

	healthCheckController := controllers.HealthCheckController{}
	healthCheckController.SetConnections(a.MySQL, a.Redis)

	cityController := controllers.CityController{}
	cityController.SetConnections(a.MySQL, a.Redis)

	statusController := controllers.StatusController{}
	statusController.SetConnections(a.MySQL, a.Redis)

	a.Router.HandleFunc("/healthcheck", healthCheckController.HealthCheck).Methods("GET")

	a.Router.HandleFunc("/city", cityController.GetAll).Methods("GET")
	a.Router.HandleFunc("/city/{cityId}", cityController.GetOne).Methods("GET")

	a.Router.HandleFunc("/city/{cityId}/status", statusController.GetStatus).Methods("GET")        // get status
	a.Router.HandleFunc("/city/{cityId}/status/{type}", statusController.SetStatus).Methods("GET") // POST // submit status
	a.Router.HandleFunc("/city/{cityId}/status", statusController.SetStatus).Methods("DELETE")     // Cancel

}

func (a *App) getRedisConnection(rdHost, rdPort, rdPwd string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdHost + ":" + rdPort, // "localhost:6379"
		Password: rdPwd,
		DB:       0,
	})
	return rdb
}

func (a *App) getDbConnection(dbUser, dbPwd, dbHost, dbPort, dbName string) *sql.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
