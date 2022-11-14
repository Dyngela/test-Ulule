package repository

import (
	"Ulule/src/DTO"
	"Ulule/src/utils"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

// FindContributionByDateByProject Database request to get the sum of the contributions each day on a range of date for a given project
func FindContributionByDateByProject(dateStart time.Time, dateRange int, projectId int) ([]DTO.ContributionByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.ContributionByDateByProject{}, err
	}
	query := `select date_trunc('day', days)::date as created_at, n.amount as amount, n.name
				from generate_series
						 ( (select min(created_at) from contribution)
						 , (select max(created_at) from contribution
							where  project_id = $1
							  and created_at <= $2
						   ), '1 day'::interval) days
						 left outer join
					 ( select sum(amount) as amount,created_at as min_created_at, p.name
					   from contribution
								inner join project p on contribution.project_id = p.id
					   group by p.name, created_at
					 ) n
					 on date_trunc('day', days)::date = n.min_created_at
				order by created_at desc
				limit $3;`
	var resp []DTO.ContributionByDateByProject
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId, dateStart, dateRange)
	if err != nil {
		return []DTO.ContributionByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.ContributionByDateByProject{}, err
	}
	return resp, nil
}

// FindNewContributorByDateByProject Database request to get the count of new contributors each day on a range of date for a given project
func FindNewContributorByDateByProject(startDate time.Time, dateRange int, projectId int) ([]DTO.NewContributorByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.NewContributorByDateByProject{}, err
	}
	var resp []DTO.NewContributorByDateByProject
	query := `select date_trunc('day', days)::date as created_at, count(n.user_id) as contributors, n.name
				from generate_series
						 ( (select min(created_at) from contribution)
						 , (select max(created_at) from contribution
							where  project_id = $1
							  and created_at <= $2
						   ), '1 day'::interval) days
						 left outer join
					 ( select user_id, min(created_at) as min_created_at, p.name
					   from contribution
						inner join project p on contribution.project_id = p.id
					   group by user_id, p.name
					 ) n
					 on date_trunc('day', days)::date = n.min_created_at
				
				group by created_at, n.name
				order by created_at desc
				limit $3;`

	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId, startDate, dateRange)
	if err != nil {
		return []DTO.NewContributorByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.NewContributorByDateByProject{}, err
	}
	return resp, nil
}

// FindNewContributionByDateByProject Database request to get the count of new contributions each day on a range of date for a given project
func FindNewContributionByDateByProject(startDate time.Time, dateRange int, projectId int) ([]DTO.NewContributionByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.NewContributionByDateByProject{}, err
	}
	query := `select date_trunc('day', days)::date as created_at, n.counter as contributions, n.name
				from generate_series
						 ( (select min(created_at) from contribution),
							(select max(created_at) from contribution
							where  project_id = $1
							  and created_at <= $2
						   ), '1 day'::interval) days
						 left outer join
					 ( select count(amount) as counter, created_at as min_created_at, p.name
					   from contribution
						inner join project p on contribution.project_id = p.id
					   group by created_at, p.name
					 ) n
					 on date_trunc('day', days)::date = n.min_created_at
				
				group by created_at, n.name, n.counter
				order by created_at desc
				limit $3;`
	var resp []DTO.NewContributionByDateByProject
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId, startDate, dateRange)
	if err != nil {
		return []DTO.NewContributionByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.NewContributionByDateByProject{}, err
	}
	return resp, nil
}

// FindAverageContributionAmount Database request to get the average of the contribution for a given project
func FindAverageContributionAmount(projectId int) (DTO.AverageContributionByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return DTO.AverageContributionByProject{}, err
	}
	query := `select ROUND(avg(c.amount),2) as average, p.name as name from contribution c
				inner join project p on p.id = c.project_id
				where project_id = $1
				group by p.name;`
	var resp DTO.AverageContributionByProject
	err = pgxscan.Get(context.Background(), conn, &resp, query, projectId)
	if err != nil {
		return DTO.AverageContributionByProject{}, err
	}
	return resp, nil
}

// FindContributionRateByVisitorsByDateByProject Database request to get the rate of contribution according to the number of visitors
// on a range of date for a given project
func FindContributionRateByVisitorsByDateByProject(startDate time.Time, dateRange int, projectId int) ([]DTO.ContributionRateByVisitorsByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.ContributionRateByVisitorsByDateByProject{}, err
	}
	query := `select date_trunc('day', days)::date as date, n.name, ROUND((cast(n.counter as decimal)/ n.visitors * 100),2) as rate
				from generate_series
						 ( (select min(date) from visit),
						   (select max(date) from visit
							where date <= $1
						   ), '1 day'::interval) days
						 left outer join
					 ( select distinct count(c.amount) as counter, created_at as min_created_at, p.name, v.visitors
					   from visit v
					   inner join contribution c on v.project_id = c.project_id and v.date = c.created_at
					   inner join project p on c.project_id = p.id
					   where p.id = $2 and c.amount > 0
					   group by created_at, p.name, v.visitors
					 ) n
					 on date_trunc('day', days)::date = n.min_created_at
				
				group by date, n.name, n.counter, n.visitors
				order by date desc
				limit $3;`
	var resp []DTO.ContributionRateByVisitorsByDateByProject
	err = pgxscan.Select(context.Background(), conn, &resp, query, startDate, projectId, dateRange)
	if err != nil {
		return []DTO.ContributionRateByVisitorsByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.ContributionRateByVisitorsByDateByProject{}, err
	}
	return resp, nil
}
