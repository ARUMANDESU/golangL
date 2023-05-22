package store

import "github.com/jackc/pgx/v5"

type database interface {
	ConnectToDB() (pgx.Conn, error)
}
