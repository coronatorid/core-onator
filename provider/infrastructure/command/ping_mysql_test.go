package command_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coronatorid/core-onator/provider/infrastructure/command"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func TestPingMYSQL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	pingDBCmd := command.NewPingMYSQL(db)

	t.Run("Use", func(t *testing.T) {
		assert.Equal(t, "ping:mysql", pingDBCmd.Use())
	})

	t.Run("Example", func(t *testing.T) {
		assert.Equal(t, "ping:mysql", pingDBCmd.Example())
	})

	t.Run("Short", func(t *testing.T) {
		assert.Equal(t, "Ping coronator mysql database", pingDBCmd.Short())
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("Given args", func(t *testing.T) {
			t.Run("When database connection is success", func(t *testing.T) {
				t.Run("It will return nil", func(t *testing.T) {
					mock.ExpectPing()
					pingDBCmd.Run([]string{})
				})
			})

			t.Run("When database connection is failed", func(t *testing.T) {
				t.Run("It will return error", func(t *testing.T) {
					mock.ExpectPing().WillReturnError(errors.New("unexpected error"))
					pingDBCmd.Run([]string{})
				})
			})
		})
	})
}
