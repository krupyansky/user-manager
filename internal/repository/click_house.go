package repository

import (
	"fmt"
	"github.com/roistat/go-clickhouse"
)

const (
	logsTable = "logs"
)

type ConfigClickHouse struct {
	Host string
	Port string
}

func NewClickHouseDB(cfg ConfigClickHouse) (*clickhouse.Conn, error) {
	transport := clickhouse.NewHttpTransport()
	conn := clickhouse.NewConn(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), transport)
	err := conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
