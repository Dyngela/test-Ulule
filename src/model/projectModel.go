package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/text/currency"
)

type Project struct {
	Id               uint64
	Name             string
	Slug             string
	Currency         currency.Unit
	DateStart        pgtype.Date
	DateEnd          pgtype.Date
	DateEndExtraTime pgtype.Date
}
