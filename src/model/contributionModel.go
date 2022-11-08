package model

import "github.com/jackc/pgx/v5/pgtype"

type Contribution struct {
	Id        uint64
	Amount    float32
	UserId    uint64
	ProjectId uint64
	CreatedAt pgtype.Date
}
