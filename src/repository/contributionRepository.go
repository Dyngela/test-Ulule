package repository

import (
	"Ulule/src/DTO"
	"Ulule/src/utils"
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

func FindContributionByDateByProject(dates string, projectId int) ([]DTO.ContributionByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.ContributionByDateByProject{}, err
	}
	// I used Sprintf here instead of a query parameter because it was not working when i tried it with several dates
	// TODO Find out why it doesn't work that way. Scany limitation?
	query := fmt.Sprintf(
		`select sum(c.amount) as amount, c.created_at, p.name from project p
				inner join contribution c on c.project_id = p.id
				where p.id = $1 and c.created_at in ('%s')
				group by c.created_at, p.name
				order by created_at desc`, dates)
	var resp []DTO.ContributionByDateByProject
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId)
	if err != nil {
		return []DTO.ContributionByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.ContributionByDateByProject{}, err
	}
	return resp, nil
}

func FindNewContributorByDateByProject(dates []string, projectId int) ([]DTO.NewContributorByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.NewContributorByDateByProject{}, err
	}

	var resp []DTO.NewContributorByDateByProject
	var tempResp DTO.NewContributorByDateByProject
	for i := 0; i < len(dates); i++ {

		tempDate, parseError := time.Parse("2006-01-02", dates[i])
		if parseError != nil {
			return []DTO.NewContributorByDateByProject{}, parseError
		}
		tempDate.AddDate(0, 0, i*-1)
		date := tempDate.Format("2006-01-02")
		query := fmt.Sprintf(`select distinct count(c.user_id) as contributors, c.created_at, p.name from project p
				inner join contribution c on c.project_id = p.id
				where user_id not in
					  (select user_id from contribution
									  where created_at < '%s' and project_id = $1)
				and project_id = $2
				and amount > 0
				and created_at = '%s'
				group by c.created_at, p.name;`, date, date)
		err = pgxscan.Get(context.Background(), conn, &tempResp, query, projectId, projectId)
		resp = append(resp, tempResp)
	}
	if err != nil {
		return []DTO.NewContributorByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.NewContributorByDateByProject{}, err
	}
	return resp, nil
}

func FindNewContributionByDateByProject(dates string, projectId int) ([]DTO.NewContributionByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.NewContributionByDateByProject{}, err
	}
	query := fmt.Sprintf(
		`select distinct count(c.amount) as contributions, c.created_at, p.name from project p
				inner join contribution c on c.project_id = p.id
				where project_id = $1
				and c.amount > 0
				and c.created_at in ('%s')
				group by c.created_at, p.name
				order by c.created_at desc;`, dates)
	var resp []DTO.NewContributionByDateByProject
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId)
	if err != nil {
		return []DTO.NewContributionByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.NewContributionByDateByProject{}, err
	}
	return resp, nil
}
