package pg

import (
	"context"
	"fmt"
	"strings"
	"url-shortener/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool *pgxpool.Pool
	ctx  context.Context
	cfgp *config.Config
}

func NewPostgresStore(c *config.Config) (*PostgresStore, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName)
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	if initStore(pool) != nil {
		return nil, err
	}
	return &PostgresStore{
		pool: pool,
		cfgp: c,
		ctx:  ctx,
	}, nil
}

func initStore(pool *pgxpool.Pool) error {
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	tag, err := tx.Exec(ctx, "CREATE TABLE IF NOT EXISTS codes (url varchar(2048) NOT NULL, code varchar(8) NOT NULL, PRIMARY KEY (code))")
	if err != nil {
		return err
	}
	if !strings.HasPrefix(tag.String(), "CREATE") {
		return fmt.Errorf("tag not matching expecting 'CREATE' command: %s", tag.String())
	}
	return tx.Commit(ctx)
}

func (s *PostgresStore) Save(code, url string) error {
	ctx := context.Background()
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	tag, err := tx.Exec(ctx, "INSERT INTO codes (url, code) VALUES ($1, $2)", url, code)
	if !tag.Insert() {
		return fmt.Errorf("tag not matching expecting 'INSERT' command: %s", tag.String())
	}
	return tx.Commit(ctx)
}

func (s *PostgresStore) Get(code string) (string, bool) {
	ctx := context.Background()
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return "", false
	}
	defer tx.Rollback(ctx)
	var url string
	err = tx.QueryRow(ctx, "SELECT url FROM codes WHERE code=$1", code).Scan(&url)
	if err != nil || tx.Commit(ctx) != nil {
		return "", false
	}
	return url, true
}
