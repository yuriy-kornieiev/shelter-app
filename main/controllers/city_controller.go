package controllers

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type CityController struct {
	mysql *sql.DB
	redis *redis.Client
}

func (ctrl *CityController) SetConnections(mysql *sql.DB, redis *redis.Client) {
	ctrl.mysql = mysql
	ctrl.redis = redis
}

func (ctrl *CityController) GetAll(w http.ResponseWriter, r *http.Request) {

}
