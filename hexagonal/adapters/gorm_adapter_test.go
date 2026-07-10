package adapters

import (
	"testing"

	"zachtix/hexagonal/core"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormOrderRepository_Save(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).
			AddRow("3.31.1"))

	dialector := sqlite.Dialector{Conn: sqlDB}
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}
	repo := NewGormOrderRepository(gormDB)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Save(core.Order{Total: 100})
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").
			WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		err := repo.Save(core.Order{Total: 100})
		assert.Error(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
