package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const table = "url_shortener"

type PostgresClient struct {
	ctx  context.Context
	conn *pgx.Conn
}

func (p *PostgresClient) Insert(url string) (id int64, err error) {
	return 0, nil
}

func (r *PostgresClient) Retrieve(id int64) (url string, err error) {
	return "", nil
}

func (r *PostgresClient) Ping() error {
	if err := r.conn.Ping(r.ctx); err != nil {
		return err
	}
	return nil
}

func NewPostgresClient(ctx context.Context, conn *pgx.Conn) *PostgresClient {
	return &PostgresClient{
		ctx:  ctx,
		conn: conn,
	}
}
