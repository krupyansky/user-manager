package repository

import (
	"github.com/roistat/go-clickhouse"
)

type LogsClickHouse struct {
	con *clickhouse.Conn
}

func NewLogsClickHouse(con *clickhouse.Conn) *LogsClickHouse {
	return &LogsClickHouse{con: con}
}

func (r *LogsClickHouse) SaveCreatedUserLog(logMessage string) error {
	q := clickhouse.NewQuery("INSERT INTO default.logs VALUES (toDate(now()), ?, ?)", 1, logMessage)
	err := q.Exec(r.con)
	if err != nil {
		return err
	}

	return nil
}
