package pg_adapter

import "time"

type Queues struct {
	Id              int    `db:"id"`
	Ip              string `db:"ip"`
	ActiveJobsCount int    `db:"active_jobs_count"`

	NotActiveSince time.Time `db:"not_active_since"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
