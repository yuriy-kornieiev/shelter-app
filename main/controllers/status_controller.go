package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type StatusController struct {
	mysql *sql.DB
	redis *redis.Client
}

func (ctrl *StatusController) SetConnections(mysql *sql.DB, redis *redis.Client) {
	ctrl.mysql = mysql
	ctrl.redis = redis
}

func (ctrl *StatusController) GetAllStatus(w http.ResponseWriter, r *http.Request) {
	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")

}

func (ctrl *StatusController) GetStatus(w http.ResponseWriter, r *http.Request) {
	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")

}

func (ctrl *StatusController) SetStatus(w http.ResponseWriter, r *http.Request) {

	cityId := mux.Vars(r)["cityId"]
	submitType := mux.Vars(r)["type"]

	if submitType != "yes" && submitType != "no" {
		Response{}.withError(w, http.StatusBadRequest, "Not implemented")
		return
	}

	// @todo validate cityId

	// 10 last minutes
	dt := time.Now()
	for i := 1; i <= 10; i++ {
		dt = dt.AddDate(0, 0, -1)
		key := dt.Format("2006_01_02_15_04") + ":" + cityId + ":" + submitType
		ctrl.redis.Incr(context.Background(), key)
		fmt.Println(key)
	}

	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")

}

func (ctrl *StatusController) GetAllStatuses(w http.ResponseWriter, r *http.Request) {
	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")

}
