package theory_test

import (
	"net/http/httptest"
	"testing"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewHandler(t *testing.T) {
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	service := &mock.TheoryService{}
	instance := theory.NewHandler(logger, service)
	require.Implements(t, (*theory.Handler)(nil), instance)
}

func mockServer() (*httptest.Server, *mux.Router) {
	router := mux.NewRouter()
	server := httptest.NewServer(router)
	return server, router
}
