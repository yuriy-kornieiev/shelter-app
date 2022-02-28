package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"yaTrivoga/business"
	"yaTrivoga/dto"
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
	cityBO := business.NewCityBO()
	cities, err := cityBO.FindAll(ctrl.mysql)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	var citiesDTO []dto.CityDTO
	for city := range cities {
		citiesDTO = append(citiesDTO, dto.NewCityDTO(cities[city]))
	}

	respondWithJSON(w, http.StatusOK, citiesDTO)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
