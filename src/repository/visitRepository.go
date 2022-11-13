package repository

import (
	"Ulule/src/DTO"
	"Ulule/src/utils"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

func FindNumberOfVisitsByDateByProject(startDate time.Time, dateRange int, projectId int) ([]DTO.VisitByDateByProject, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.VisitByDateByProject{}, err
	}
	query := `select date_trunc('day', days)::date as date, n.visitors as visits, n.name as name
				from generate_series
						 ( (select min(date) from visit),
							(select max(date) from visit
							where date <= $2
						   ), '1 day'::interval) days
						 left outer join
					 ( select visitors, date as min_created_at, p.name
					   from visit v
						inner join project p on v.project_id = p.id
       					where v.project_id = $1
					   group by date, p.name, visitors
					 ) n
					 on date_trunc('day', days)::date   = n.min_created_at
				
				group by date, n.name, n.visitors
				order by date desc
				limit $3;`
	var resp []DTO.VisitByDateByProject
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId, startDate, dateRange)
	if err != nil {
		return []DTO.VisitByDateByProject{}, err
	}
	if resp == nil {
		return []DTO.VisitByDateByProject{}, err
	}
	return resp, nil
}
