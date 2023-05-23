package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Dbname   string
}

func (p *Postgres) ConnectToDB() (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", p.User, p.Password, p.Host, p.Port, p.Dbname)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	return conn, nil
}
