package controllers

import (
	"context"
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

	var err error

	err = ctrl.mysql.Ping()
	if err != nil {
		http.Error(w, "MySQL connection issue", http.StatusServiceUnavailable)
	}

	_, err = ctrl.redis.Ping(context.Background()).Result()
	if err != nil {
		http.Error(w, "Redis connection issue", http.StatusServiceUnavailable)
	}

	http.Error(w, "Healthy", 200)
}
