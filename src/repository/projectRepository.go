package repository

import (
	"Ulule/src/DTO"
	"Ulule/src/utils"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

// FindProjectMilestones Database request to get project's milestone e.g first contribution, 10% completion, 20%  completion etc..
// for a range of date for a given project
func FindProjectMilestones(projectId int) ([]DTO.ProjectMilestone, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.ProjectMilestone{}, err
	}
	query := `select min(t.id) as contribution_id, slice, min(c2.created_at) as date from
					(SELECT c.id, sum(amount) OVER (ORDER BY c.id),
						case
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 10 then 0
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 20 then 10
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 30 then 20
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 40 then 30
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 50 then 40
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 60 then 50
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 70 then 60
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 80 then 70
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 90 then 80
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 < 100 then 90
						   when sum(amount) OVER (ORDER BY c.id) / p.goal * 100 >= 100 then 100
						end slice
					FROM contribution c
					inner join project p on p.id = c.project_id
					where project_id = $1) as t
                inner join contribution c2 on t.id = c2.id
                where t.slice > 0
				group by t.slice, c2.created_at
				order by t.slice, date
				limit 10;`
	var resp []DTO.ProjectMilestone
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId)
	if err != nil {
		return []DTO.ProjectMilestone{}, err
	}
	if resp == nil {
		return []DTO.ProjectMilestone{}, nil
	}
	return resp, nil
}

// FindAdvancementPercentage Database request to get the objective's completion in percentage for a range of date for a given project
func FindAdvancementPercentage(dateStart time.Time, dateRange int, projectId int) ([]DTO.ProjectAdvancement, error) {
	conn, err := utils.Pool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		return []DTO.ProjectAdvancement{}, err
	}
	query := `SELECT distinct c.created_at as date, p.name as name, 
					ROUND(sum(amount/ p.goal * 100) OVER (ORDER BY c.created_at), 2) as completion_percentage
				FROM contribution c
						 inner join project p on p.id = c.project_id
				where project_id = $1
				and c.created_at < $2
				order by c.created_at desc
				limit $3;`
	var resp []DTO.ProjectAdvancement
	err = pgxscan.Select(context.Background(), conn, &resp, query, projectId, dateStart, dateRange)
	if err != nil {
		return []DTO.ProjectAdvancement{}, err
	}
	if resp == nil {
		return []DTO.ProjectAdvancement{}, nil
	}
	return resp, nil
}
