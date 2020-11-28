package adapter_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/infrastructure/adapter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTx(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	t.Run("QueryRowContext", func(t *testing.T) {
		t.Run("When querying done it will return provider.Row", func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockDB.ExpectBegin()
			tx, _ := db.Begin()

			mockDB.ExpectQuery(`select id from test_table where id = \?`).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			sql := adapter.AdaptTx(tx)
			_ = sql.QueryRowContext(ctx, "test-query-1", "select id from test_table where id = ?", 1)
			assert.Nil(t, mockDB.ExpectationsWereMet())
		})
	})

	t.Run("QueryContext", func(t *testing.T) {
		t.Run("When querying done it will return provider.Rows", func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockDB.ExpectBegin()
			tx, _ := db.Begin()

			mockDB.ExpectQuery(`select id from test_table`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(1))

			sql := adapter.AdaptTx(tx)
			_, err = sql.QueryContext(ctx, "test-query-2", "select id from test_table")
			assert.Nil(t, err)
			assert.Nil(t, mockDB.ExpectationsWereMet())
		})

		t.Run("When there are error it will return it", func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockDB.ExpectBegin()
			tx, _ := db.Begin()

			mockDB.ExpectQuery(`select id from test_table`).WillReturnError(errors.New("unexpected error"))

			sql := adapter.AdaptTx(tx)
			_, err = sql.QueryContext(ctx, "test-query-2", "select id from test_table")
			assert.NotNil(t, err)
			assert.Nil(t, mockDB.ExpectationsWereMet())
		})

		t.Run("When there are sql.ErrNoRows error it will return provider.ErrDBNotFound", func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockDB.ExpectBegin()
			tx, _ := db.Begin()

			mockDB.ExpectQuery(`select id from test_table`).WillReturnError(sql.ErrNoRows)

			sql := adapter.AdaptTx(tx)
			_, err = sql.QueryContext(ctx, "test-query-2", "select id from test_table")
			assert.Equal(t, provider.ErrDBNotFound, err)
			assert.Nil(t, mockDB.ExpectationsWereMet())
		})
	})

	t.Run("ExecuteContext", func(t *testing.T) {
		t.Run("When execution done it will return provider.Result", func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockDB.ExpectBegin()
			tx, _ := db.Begin()

			mockDB.ExpectExec(`insert into users \(name\) values\('insomnius'\)`).WillReturnResult(sqlmock.NewResult(1, 0))

			sql := adapter.AdaptTx(tx)
			_, err = sql.ExecContext(ctx, "test-query-2", "insert into users (name) values('insomnius')")
			assert.Nil(t, err)
			assert.Nil(t, mockDB.ExpectationsWereMet())
		})

		t.Run("When there are error it will return it", func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mockDB.ExpectBegin()
			tx, _ := db.Begin()

			mockDB.ExpectExec(`insert into users \(name\) values\('insomnius'\)`).WillReturnError(errors.New("unexpected error"))

			sql := adapter.AdaptTx(tx)
			_, err = sql.ExecContext(ctx, "test-query-2", "insert into users (name) values('insomnius')")
			assert.NotNil(t, err)
			assert.Nil(t, mockDB.ExpectationsWereMet())
		})
	})
}
