package DTO

import "github.com/jackc/pgx/v5/pgtype"

// ContributionByDateByProject API response to get the sum of the contributions each day on a range of date for a given project
type ContributionByDateByProject struct {
	Amount  float32     `db:"amount"`
	Date    pgtype.Date `db:"created_at"`
	Project string      `db:"name"`
}

// NewContributorByDateByProject API response to get the count of new contributors each day on a range of date for a given project
type NewContributorByDateByProject struct {
	Contributors uint        `db:"contributors"`
	Date         pgtype.Date `db:"created_at"`
	Project      string      `db:"name"`
}

// NewContributionByDateByProject API response to get the count of new contributions each day on a range of date for a given project
type NewContributionByDateByProject struct {
	Contributions uint        `db:"contributions"`
	Date          pgtype.Date `db:"created_at"`
	Project       string      `db:"name"`
}

// AverageContributionByProject API response to get the average of the contribution for a given project
type AverageContributionByProject struct {
	Amount  float32 `db:"average"`
	Project string  `db:"name"`
}

// ContributionRateByVisitorsByDateByProject API response to get the rate of contribution according to the number of visitors
// on a range of date for a given project
type ContributionRateByVisitorsByDateByProject struct {
	Rate    float32     `db:"rate"`
	Date    pgtype.Date `db:"date"`
	Project string      `db:"name"`
}
