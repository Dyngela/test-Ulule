package model

import "github.com/jackc/pgx/v5/pgtype"

// Visit Database's model
type Visit struct {
	Id        uint64
	ProjectId uint64
	Date      pgtype.Date
	PageViews uint32
	Visitors  uint32
}
