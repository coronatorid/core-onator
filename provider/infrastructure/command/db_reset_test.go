package command_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/coronatorid/core-onator/provider/infrastructure/command"
	"github.com/coronatorid/core-onator/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDBRese(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dbResetCMD := command.NewDBReset(db, "coronator_development", "file://migration/")

	t.Run("Use", func(t *testing.T) {
		assert.Equal(t, "db:reset", dbResetCMD.Use())
	})

	t.Run("Example", func(t *testing.T) {
		assert.Equal(t, "db:reset", dbResetCMD.Example())
	})

	t.Run("Short", func(t *testing.T) {
		assert.Equal(t, "Migrate coronator database all way down", dbResetCMD.Short())
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("Given args", func(t *testing.T) {
			dbResetCMD := command.NewDBReset(db, "coronator_development", "file://migration_down/")

			migrationContent := "DROP TABLE `exposed_users`;"

			t.Run("When migration process is success", func(t *testing.T) {
				testutil.GenerateTempTestFiles("./migration_down/", migrationContent, "01_migration.up.sql", 0666)

				mockDB.ExpectPing()

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WillReturnRows(sqlmock.NewRows([]string{
					"GET_LOCK(1, 10)",
				}).AddRow(1))

				mockDB.ExpectQuery(`SHOW TABLES LIKE "db_versions"`).WillReturnRows(sqlmock.NewRows([]string{
					"TABLES IN DATABASES",
				}).AddRow("db_versions"))

				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{
					"GET_LOCK(1, 10)",
				}).AddRow(1))

				mockDB.ExpectQuery("SELECT version, dirty FROM `db_versions` LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{"version", "dirty"}).AddRow(1, 0))
				mockDB.ExpectBegin()
				mockDB.ExpectExec("TRUNCATE `db_versions`").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectExec("INSERT INTO `db_versions` \\(version, dirty\\) VALUES \\(\\?, \\?\\)").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectCommit()

				mockDB.ExpectBegin()
				mockDB.ExpectExec("TRUNCATE `db_versions`").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectCommit()
				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				dbResetCMD.Run([]string{})

				testutil.RemoveTempTestFiles("./migration_down/")
			})

			t.Run("When there is no change it will not panic", func(t *testing.T) {
				testutil.GenerateTempTestFiles("./migration_down/", migrationContent, "01_migration.up.sql", 0666)

				mockDB.ExpectPing()

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WillReturnRows(sqlmock.NewRows([]string{
					"GET_LOCK(1, 10)",
				}).AddRow(1))

				mockDB.ExpectQuery(`SHOW TABLES LIKE "db_versions"`).WillReturnRows(sqlmock.NewRows([]string{
					"TABLES IN DATABASES",
				}).AddRow("db_versions"))

				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{
					"GET_LOCK(1, 10)",
				}).AddRow(1))

				mockDB.ExpectQuery("SELECT version, dirty FROM `db_versions` LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{"version", "dirty"}))
				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				dbResetCMD.Run([]string{})

				testutil.RemoveTempTestFiles("./migration_down/")
			})

			t.Run("When database instantiation failed it will not panic", func(t *testing.T) {
				testutil.GenerateTempTestFiles("./migration_down/", migrationContent, "01_migration.up.sql", 0666)

				mockDB.ExpectPing().WillReturnError(errors.New("unexpected error"))

				dbResetCMD.Run([]string{})

				testutil.RemoveTempTestFiles("./migration_down/")
			})

		})
	})
}
