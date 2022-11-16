package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Project Database's model
type Project struct {
	Id               uint64
	Name             string
	Slug             string
	Currency         string
	DateStart        pgtype.Date `db:"date_start"`
	DateEnd          pgtype.Date `db:"date_end"`
	DateEndExtraTime pgtype.Date `db:"date_end_extra_time"`
	Type             string
	Goal             float64
}
