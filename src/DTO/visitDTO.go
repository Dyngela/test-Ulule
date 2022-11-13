package DTO

import "github.com/jackc/pgx/v5/pgtype"

type VisitByDateByProject struct {
	Visits  float32     `db:"visits"`
	Date    pgtype.Date `db:"date"`
	Project string      `db:"name"`
}
