package adapters

import (
	"testing"

	"zachtix/hexagonal/core"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockGormDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	return gormDB, mock, func() { db.Close() }
}

func TestGormOrderRepository_Save(t *testing.T) {
	gormDB, mock, cleanup := newMockGormDB(t)
	defer cleanup()

	repo := NewGormOrderRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectQuery(`^INSERT INTO "orders" (.+)$`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err := repo.Save(&core.Order{Total: 100.0})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
