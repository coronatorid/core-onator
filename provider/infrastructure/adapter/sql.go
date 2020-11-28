package adapter

import (
	"context"
	"database/sql"

	"github.com/coronatorid/core-onator/provider"
)

// A SQL adapter for golang sql
type SQL struct {
	db *sql.DB
}

// AdaptSQL adapting golang sql.DB
func AdaptSQL(db *sql.DB) *SQL {
	return &SQL{db: db}
}

// Transaction wrap mysql transaction into a bit of simpler way
func (s *SQL) Transaction(ctx context.Context, transactionKey string, f func(tx provider.TX) error) error {
	return runWithSQLAnalyzer(ctx, "db", func() error {
		tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return err
		}

		adaptedTx := &Tx{tx: tx}
		if err := f(adaptedTx); err != nil {
			_ = tx.Rollback()
			return err
		}

		if err := tx.Commit(); err != nil {
			_ = tx.Rollback()
			return err
		}

		return nil
	})

}

// ExecContext wrap sql ExecContext function
func (s *SQL) ExecContext(ctx context.Context, queryKey, query string, args ...interface{}) (provider.Result, error) {
	var result provider.Result
	var err error

	_ = runWithSQLAnalyzer(ctx, "db", func() error {
		result, err = s.db.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// QueryContext wrap sql QueryContext function
func (s *SQL) QueryContext(ctx context.Context, queryKey, query string, args ...interface{}) (provider.Rows, error) {
	var rows provider.Rows
	var err error

	_ = runWithSQLAnalyzer(ctx, "db", func() error {
		rows, err = s.db.QueryContext(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})

	return rows, err
}

// QueryRowContext wrap sql QueryRowContext function
func (s *SQL) QueryRowContext(ctx context.Context, queryKey, query string, args ...interface{}) provider.Row {
	var row provider.Row

	_ = runWithSQLAnalyzer(ctx, "db", func() error {
		row = s.db.QueryRowContext(ctx, query, args...)
		return nil
	})

	return row
}

func runWithSQLAnalyzer(ctx context.Context, executionLevel string, f func() error) error {
	// TODO add log here
	return f()
}
