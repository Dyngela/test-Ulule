package DTO

import "github.com/jackc/pgx/v5/pgtype"

// VisitByDateByProject API response to get the number of visits for a range of date for a given project
type VisitByDateByProject struct {
	Visits  float32     `db:"visits"`
	Date    pgtype.Date `db:"date"`
	Project string      `db:"name"`
}
