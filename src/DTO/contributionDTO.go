package DTO

import "github.com/jackc/pgx/v5/pgtype"

type ContributionByDateByProject struct {
	Amount  float32     `db:"amount"`
	Date    pgtype.Date `db:"created_at"`
	Project string      `db:"name"`
}

type NewContributorByDateByProject struct {
	Contributors uint        `db:"contributors"`
	Date         pgtype.Date `db:"created_at"`
	Project      string      `db:"name"`
}

type NewContributionByDateByProject struct {
	Contributions uint        `db:"contributions"`
	Date          pgtype.Date `db:"created_at"`
	Project       string      `db:"name"`
}

type AverageContributionByProject struct {
	Amount  float32 `db:"average"`
	Project string  `db:"name"`
}

type ContributionRateByVisitorsByDateByProject struct {
	Rate    float32     `db:"rate"`
	Date    pgtype.Date `db:"date"`
	Project string      `db:"name"`
}
