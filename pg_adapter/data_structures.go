package pg_adapter

import (
	"github.com/lib/pq"
	"log"
	"net"
	"time"
)

type Queues struct {
	Id              int    `db:"id"`
	Ip              string `db:"ip"`
	ActiveJobsCount int    `db:"active_jobs_count"`

	NotActiveSince pq.NullTime `db:"not_active_since"`
	CreatedAt      time.Time   `db:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at"`
}

func (q Queues) GetIP() net.IPAddr {
	return net.IPAddr{IP: net.ParseIP(q.Ip)}
}

func (q *Queues) Insert(ip net.IPAddr) error {
	q.ActiveJobsCount = 1
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	q.Ip = ip.String()

	connection := createConnection()
	defer func() {
		err := connection.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}()

	tx, err := connection.Connection.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`set transaction isolation level serializable`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO pg_queues (ip, active_jobs_count, created_at, updated_at) VALUES($1, $2, $3, $4)",
		q.Ip, q.ActiveJobsCount, q.CreatedAt, q.UpdatedAt)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q Queues) SelectCountOfActive() (int, error) {
	connection := createConnection()
	defer func() {
		err := connection.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}()

	tx, err := connection.Connection.Begin()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`set transaction isolation level serializable`)
	if err != nil {
		return 0, err
	}

	var count int
	row := tx.QueryRow("SELECT COUNT(pg_queues.id) FROM pg_queues WHERE pg_queues.active_jobs_count > 0 AND not_active_since NOTNULL")
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q Queues) SelectNotActive() ([]net.IPAddr, error) {
	connection := createConnection()
	defer func() {
		err := connection.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}()

	tx, err := connection.Connection.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(`set transaction isolation level serializable`)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT ip FROM pg_queues WHERE pg_queues.active_jobs_count = 0 AND not_active_since NOTNULL")
	if err != nil {
		return nil, err
	}

	var ipsStr []string

	for rows.Next() {
		var tmp string
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		ipsStr = append(ipsStr, tmp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var ips []net.IPAddr

	for _, ipStr := range ipsStr {
		ips = append(ips, net.IPAddr{IP: net.ParseIP(ipStr)})
	}

	return ips, nil
}

func (q *Queues) SelectTheMostNotLoaded() error {
	connection := createConnection()
	defer func() {
		err := connection.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}()

	tx, err := connection.Connection.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`set transaction isolation level serializable`)
	if err != nil {
		return err
	}

	row := tx.QueryRow("SELECT * FROM pg_queues ORDER BY active_jobs_count ASC LIMIT 1")
	err = row.Scan(&q.Id, &q.Ip, &q.ActiveJobsCount, &q.NotActiveSince, &q.CreatedAt, &q.UpdatedAt)

	if err != nil {
		return err
	}

	rows, err := tx.Query("UPDATE pg_queues SET active_jobs_count = $1 where ip = $2", q.ActiveJobsCount+1, q.Ip)
	if err != nil {
		return err
	}

	err = rows.Close()
	if err != nil {
		return err
	}

	if q.NotActiveSince.Valid {
		rows, err = tx.Query("UPDATE pg_queues SET not_active_since = $1 where ip = $2", nil, q.Ip)
		if err != nil {
			return err
		}

		err = rows.Close()
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q Queues) UpdateActiveJobsCount(ip net.IPAddr) error {
	connection := createConnection()
	defer func() {
		err := connection.Disconnect()
		if err != nil {
			log.Println(err)
		}
	}()

	tx, err := connection.Connection.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`set transaction isolation level serializable`)
	if err != nil {
		return err
	}

	var activeJobCount int

	row := tx.QueryRow("SELECT active_jobs_count FROM pg_queues WHERE ip = $1", ip.String())
	err = row.Scan(&activeJobCount)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE pg_queues SET active_jobs_count = $1 where ip = $2", activeJobCount-1, ip.String())
	if err != nil {
		return err
	}

	if activeJobCount-1 == 0 {
		_, err = tx.Exec("UPDATE pg_queues SET not_active_since = $1 where ip = $2", time.Now(), ip.String())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
