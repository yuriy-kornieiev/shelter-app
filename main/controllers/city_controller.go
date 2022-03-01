package controllers

import (
	"database/sql"
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
			Response{}.withError(w, http.StatusNotFound, "Not found")
		default:
			Response{}.withError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	var citiesDTO []dto.CityDTO
	for city := range cities {
		citiesDTO = append(citiesDTO, dto.NewCityDTO(cities[city]))
	}

	Response{}.withJSON(w, http.StatusOK, citiesDTO)
}

func (ctrl *CityController) GetOne(w http.ResponseWriter, r *http.Request) {
	Response{}.withError(w, http.StatusNotImplemented, "Not implemented")
}
