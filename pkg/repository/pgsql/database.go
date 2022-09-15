package pgsql

import (
	"context"
	"database/sql"
)

type txKey struct{}

func injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}

type Database struct {
	conn *sql.DB
}

func NewDatabase(c *sql.DB) *Database {
	return &Database{c}
}

func (d *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}
	return d.conn.ExecContext(ctx, query, args...)
}

func (d *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}
	return d.conn.QueryContext(ctx, query, args...)
}

func (d *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return d.conn.QueryRowContext(ctx, query, args...)
}

func (d *Database) WithinTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	ctx = injectTx(ctx, tx)
	err = fn(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
