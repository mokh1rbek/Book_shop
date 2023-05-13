package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/mokh1rbek/Book_CRUD/config"
	"github.com/mokh1rbek/Book_CRUD/storage"
)

type Store struct {
	db       *pgxpool.Pool
	category *BookRepo
}

// Book implements storage.StorageI
func (*Store) Book() storage.BookRepoI {
	panic("unimplemented")
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pool,
		category: NewBookRepo(pool),
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Category() storage.BookRepoI {

	if s.category == nil {
		s.category = NewBookRepo(s.db)
	}

	return s.category
}
