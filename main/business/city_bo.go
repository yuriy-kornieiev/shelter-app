package business

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"yaTrivoga/models"
)

type CityBO struct {
	mysql *sql.DB
	redis *redis.Client
}

func NewCityBO() *CityBO {
	bo := CityBO{}
	return &bo
}

func (bo CityBO) FindAll(conn *sql.DB) ([]models.City, error) {

	query := "SELECT id, name FROM cities"
	rows, err := conn.Query(query)
	if err != nil {
		return []models.City{}, err
	}
	defer rows.Close()

	var row models.City
	var cities []models.City
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.Name)
		if err != nil {
			return []models.City{}, err
		}
		cities = append(cities, row)
	}

	return cities, nil
}

func (bo CityBO) FindById(conn *sql.DB, id int64) (models.City, error) {
	var row models.City

	query := "SELECT id, name FROM cities WHERE id = ?"

	err := conn.QueryRow(query, id).Scan(&row.Id, &row.Name)

	if err != nil {
		return models.City{}, err
	}

	return row, nil
}

func (bo CityBO) FindAllAsIdMap(cities []models.City) map[int64][]models.City {
	citiesMap := map[int64][]models.City{}
	for _, value := range cities {
		citiesMap[value.Id.Int64] = append(citiesMap[value.Id.Int64], value)
	}
	return citiesMap
}
