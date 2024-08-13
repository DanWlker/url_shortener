package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const table = "url_shortener"

type PostgresClient struct {
	ctx  context.Context
	conn *pgx.Conn
}

func (p *PostgresClient) Insert(url string) (id int64, err error) {
	rows, err := p.conn.Query(
		p.ctx,
		fmt.Sprintf("insert into %v(url) values ('%v') returning id", table, url),
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("Well you shouldn't be here, something def went wrong")
	}

	if err := rows.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PostgresClient) Retrieve(id int64) (url string, err error) {
	rows, err := p.conn.Query(
		p.ctx,
		fmt.Sprintf("select * from %v where id=%d", table, id),
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", IdNotExistError
	}

	if err := rows.Scan(nil, &url); err != nil {
		return "", err
	}

	return url, nil
}

func (p *PostgresClient) Ping() error {
	if err := p.conn.Ping(p.ctx); err != nil {
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
