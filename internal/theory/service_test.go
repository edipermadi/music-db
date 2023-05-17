package theory_test

import (
	"testing"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewService(t *testing.T) {
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	repository := &mock.TheoryRepository{}
	instance := theory.NewService(logger, repository)
	require.Implements(t, (*theory.Service)(nil), instance)
}
