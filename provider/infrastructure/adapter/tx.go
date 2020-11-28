package adapter

import (
	"context"
	"database/sql"

	"github.com/coronatorid/core-onator/provider"
)

// A Tx adapater for golang sql
type Tx struct {
	tx *sql.Tx
}

// AdaptTx do adapting mysql transaction
func AdaptTx(tx *sql.Tx) *Tx {
	return &Tx{tx: tx}
}

// ExecContext wrap sql ExecContext function
func (t *Tx) ExecContext(ctx context.Context, queryKey, query string, args ...interface{}) (provider.Result, error) {
	var result provider.Result
	var err error

	_ = runWithSQLAnalyzer(ctx, "tx", func() error {
		result, err = t.tx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// QueryContext wrap sql QueryContext function
func (t *Tx) QueryContext(ctx context.Context, queryKey, query string, args ...interface{}) (provider.Rows, error) {
	var rows provider.Rows
	var err error

	_ = runWithSQLAnalyzer(ctx, "tx", func() error {
		rows, err = t.tx.QueryContext(ctx, query, args...)
		if err == sql.ErrNoRows {
			err = provider.ErrDBNotFound
			return provider.ErrDBNotFound
		} else if err != nil {
			return err
		}

		return nil
	})

	return rows, err
}

// QueryRowContext wrap sql QueryRowContext function
func (t *Tx) QueryRowContext(ctx context.Context, queryKey, query string, args ...interface{}) provider.Row {
	var row provider.Row

	_ = runWithSQLAnalyzer(ctx, "tx", func() error {
		row = t.tx.QueryRowContext(ctx, query, args...)
		return nil
	})

	return AdaptSQLRow(row)
}
