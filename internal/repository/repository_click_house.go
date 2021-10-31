package repository

import (
	"github.com/roistat/go-clickhouse"
)

type LogClickHouse interface {
	SaveCreatedUserLog(logMessage string) error
}

type ClickHouse struct {
	LogClickHouse
}

func NewRepositoryClickHouse(con *clickhouse.Conn) *ClickHouse {
	return &ClickHouse{
		LogClickHouse: NewLogsClickHouse(con),
	}
}
