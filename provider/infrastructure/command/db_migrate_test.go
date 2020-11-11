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

func TestDBMigrate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, mockDB, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dbMigrateCMD := command.NewDBMigrate(db, "coronator_development", "file://migration/")

	t.Run("Use", func(t *testing.T) {
		assert.Equal(t, "db:migrate", dbMigrateCMD.Use())
	})

	t.Run("Example", func(t *testing.T) {
		assert.Equal(t, "db:migrate", dbMigrateCMD.Example())
	})

	t.Run("Short", func(t *testing.T) {
		assert.Equal(t, "Migrate coronator database to a newer version", dbMigrateCMD.Short())
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("Given args", func(t *testing.T) {
			migrationContent := "CREATE TABLE `exposed_users` (\n`id` int PRIMARY KEY AUTO_INCREMENT,\n`user_id` int NOT NULL,\n`confirmed_cases_id` int NOT NULL,\n`lat` double NOT NULL,\n`long` double NOT NULL,\n`created_at` datetime NOT NULL,\n`updated_at` datetime NOT NULL,\nKEY `user_id` (`user_id`)\n) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"

			t.Run("When migration process is success", func(t *testing.T) {

				testutil.GenerateTempTestFiles("./migration/", migrationContent, "01_migration.up.sql", 0666)

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
				mockDB.ExpectBegin()
				mockDB.ExpectExec("TRUNCATE `db_versions`").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectExec("INSERT INTO `db_versions` \\(version, dirty\\) VALUES \\(\\?, \\?\\)").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectCommit()
				mockDB.ExpectExec("CREATE TABLE `exposed_users`*").WillReturnResult(sqlmock.NewResult(0, 0))

				mockDB.ExpectBegin()
				mockDB.ExpectExec("TRUNCATE `db_versions`").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectExec("INSERT INTO `db_versions` \\(version, dirty\\) VALUES \\(\\?, \\?\\)").WillReturnResult(sqlmock.NewResult(0, 0))
				mockDB.ExpectCommit()
				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				dbMigrateCMD.Run([]string{})

				testutil.RemoveTempTestFiles("./migration/")
			})

			t.Run("When database migration ping failed it will not panic", func(t *testing.T) {
				dbMigrateCMD := command.NewDBMigrate(db, "coronator_development", "file://migration_ping/")
				testutil.GenerateTempTestFiles("./migration_ping/", migrationContent, "01_migration.up.sql", 0666)

				mockDB.ExpectPing().WillReturnError(errors.New("unexpected error"))

				dbMigrateCMD.Run([]string{})

				testutil.RemoveTempTestFiles("./migration_ping/")
			})

			t.Run("When migration instantiation fail it will not panic", func(t *testing.T) {
				dbMigrateCMD := command.NewDBMigrate(db, "coronator_development", "")
				mockDB.ExpectPing()

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WillReturnRows(sqlmock.NewRows([]string{
					"GET_LOCK(1, 10)",
				}).AddRow(1))

				mockDB.ExpectQuery(`SHOW TABLES LIKE "db_versions"`).WillReturnRows(sqlmock.NewRows([]string{
					"TABLES IN DATABASES",
				}).AddRow("db_versions"))

				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				dbMigrateCMD.Run([]string{})
			})

			t.Run("When migration up process is failed it will not panic", func(t *testing.T) {

				testutil.GenerateTempTestFiles("./migration/", migrationContent, "01_migration.up.sql", 0666)

				mockDB.ExpectPing()

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WillReturnRows(sqlmock.NewRows([]string{
					"GET_LOCK(1, 10)",
				}).AddRow(1))

				mockDB.ExpectQuery(`SHOW TABLES LIKE "db_versions"`).WillReturnRows(sqlmock.NewRows([]string{
					"TABLES IN DATABASES",
				}).AddRow("db_versions"))

				mockDB.ExpectExec(`SELECT RELEASE_LOCK\(\?\)`).WillReturnResult(sqlmock.NewResult(0, 1))

				mockDB.ExpectQuery(`SELECT GET_LOCK\(\?, 10\)`).WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("unexpected error"))

				dbMigrateCMD.Run([]string{})

				testutil.RemoveTempTestFiles("./migration/")
			})
		})
	})
}
