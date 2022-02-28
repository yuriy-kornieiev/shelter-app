package controllers

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type HealthCheckController struct {
	mysql *sql.DB
	redis *redis.Client
}

func (ctrl *HealthCheckController) SetConnections(mysql *sql.DB, redis *redis.Client) {
	ctrl.mysql = mysql
	ctrl.redis = redis
}

func (ctrl *HealthCheckController) HealthCheck(w http.ResponseWriter, r *http.Request) {

	//var err error
	//
	//err = ctrl.connections["bac"].GetConnection().Ping()
	//if err != nil {
	//	http.Error(w, "BAC MySQL connection issue", http.StatusServiceUnavailable)
	//}
	//
	//err = ctrl.connections["config"].GetConnection().Ping()
	//if err != nil {
	//	http.Error(w, "Config MySQL connection issue", http.StatusServiceUnavailable)
	//}
	//
	//err = ctrl.connections["opxus"].GetConnection().Ping()
	//if err != nil {
	//	http.Error(w, "OPX US MySQL connection issue", http.StatusServiceUnavailable)
	//}
	//
	//err = ctrl.connections["opxin"].GetConnection().Ping()
	//if err != nil {
	//	http.Error(w, "OPX IN MySQL connection issue", http.StatusServiceUnavailable)
	//}
	//
	//err = ctrl.connections["oss"].GetConnection().Ping()
	//if err != nil {
	//	http.Error(w, "OSS MySQL connection issue", http.StatusServiceUnavailable)
	//}

	http.Error(w, "Healthy", 200)
}
