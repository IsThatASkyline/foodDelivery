package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) getDB(ctx context.Context) Querier {
	if tx, ok := ExtractTx(ctx); ok {
		return tx
	}
	return s.db
}
