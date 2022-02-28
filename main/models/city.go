package models

import (
	"database/sql"
)

type City struct {
	Id   sql.NullInt64  `json:"ID"`
	Name sql.NullString `json:"Name"`
}
