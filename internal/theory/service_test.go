package theory_test

import (
	"testing"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	logger := mock.Logger()

	repository := mock.TheoryRepository(mock.TheoryRepositoryReturnValues{})
	instance := theory.NewService(logger, repository)
	require.Implements(t, (*theory.Service)(nil), instance)
}

type serviceTestCase struct {
	Title                  string
	RepositoryReturnValues mock.TheoryRepositoryReturnValues
}

func (s serviceTestCase) mockedService() theory.Service {
	logger := mock.Logger()
	repository := mock.TheoryRepository(s.RepositoryReturnValues)
	return theory.NewService(logger, repository)
}
