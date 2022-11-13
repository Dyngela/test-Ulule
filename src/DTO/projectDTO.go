package DTO

import "github.com/jackc/pgx/v5/pgtype"

type ProjectMilestone struct {
	ContributionId uint64      `db:"contribution_id"`
	Date           pgtype.Date `db:"date"`
	Slice          string      `db:"slice"`
}

type ProjectAdvancement struct {
	Date                 pgtype.Date `db:"date"`
	Project              string      `db:"name"`
	CompletionPercentage float32     `db:"completion_percentage"`
}
