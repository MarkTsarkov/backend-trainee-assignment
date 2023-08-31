package storage

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sql/up.sql
var initialDB string

type SegmentStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *SegmentStorage {
	return &SegmentStorage{db: db}
}

func (storage *SegmentStorage) CreateTables(ctx context.Context) error {
	_, err := storage.db.Exec(ctx, initialDB)
	if err != nil {
		return fmt.Errorf("Error create Tables: %v", err)
	}
	return nil
}

func (storage *SegmentStorage) CreateSegment(ctx context.Context, slug string) error {
	sql := "INSERT INTO segment (slug) VALUES ($1);"
	_, err := storage.db.Exec(ctx, sql, slug)
	if err != nil {
		return fmt.Errorf("Error createSegment: %v", err)
	}
	return nil
}
