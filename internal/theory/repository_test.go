package theory_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewRepository(t *testing.T) {
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	db, _, err := mockDatabase()
	require.NoError(t, err)

	repository := theory.NewRepository(logger, db)
	require.Implements(t, (*theory.Repository)(nil), repository)
}

func mockDatabase() (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, sqlMock, err := sqlmock.New(
		sqlmock.MonitorPingsOption(true),
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	if err != nil {
		return nil, nil, err
	}

	dbx := sqlx.NewDb(db, "postgres")
	return dbx, sqlMock, err
}
