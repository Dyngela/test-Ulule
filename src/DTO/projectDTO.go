package DTO

import "github.com/jackc/pgx/v5/pgtype"

// ProjectMilestone API response to get project's milestone e.g first contribution, 10% completion, 20%  completion etc..
// for a range of date for a given project
type ProjectMilestone struct {
	ContributionId uint64      `db:"contribution_id"`
	Date           pgtype.Date `db:"date"`
	Slice          string      `db:"slice"`
}

// ProjectAdvancement API response to get the objective's completion in percentage for a range of date for a given project
type ProjectAdvancement struct {
	Date                 pgtype.Date `db:"date"`
	Project              string      `db:"name"`
	CompletionPercentage float32     `db:"completion_percentage"`
}
