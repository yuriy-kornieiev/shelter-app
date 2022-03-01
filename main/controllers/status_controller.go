package controllers

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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
	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")

}

func (ctrl *StatusController) GetAllStatuses(w http.ResponseWriter, r *http.Request) {
	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")

}
