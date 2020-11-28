package adapter

import (
	"database/sql"

	"github.com/coronatorid/core-onator/provider"
)

// SQLRow wrap single sql row
type SQLRow struct {
	row provider.Row
}

// AdaptSQLRow wrap provider row
func AdaptSQLRow(row provider.Row) *SQLRow {
	return &SQLRow{row: row}
}

// Scan warp default row scan function
func (s *SQLRow) Scan(dest ...interface{}) error {
	err := s.row.Scan(dest...)
	if err == sql.ErrNoRows {
		return provider.ErrDBNotFound
	} else if err != nil {
		return err
	}

	return nil
}
