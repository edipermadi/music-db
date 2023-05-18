package theory_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	logger := mock.Logger()
	service := mock.TheoryService(mock.TheoryServiceReturnValues{})
	instance := theory.NewHandler(logger, service)
	require.Implements(t, (*theory.Handler)(nil), instance)
}

type handlerTestCase struct {
	Title               string
	GivenQueryStrings   url.Values
	ServiceReturnValues mock.TheoryServiceReturnValues
	ExpectedStatus      int
	baseURL             string
}

func (h *handlerTestCase) mockServer() *httptest.Server {
	logger := mock.Logger()
	router := mux.NewRouter()
	server := httptest.NewServer(router)
	service := mock.TheoryService(h.ServiceReturnValues)
	theory.NewHandler(logger, service).InstallEndpoints(router)
	h.baseURL = server.URL
	return server
}

func (h *handlerTestCase) rawQuery() string {
	if len(h.GivenQueryStrings) > 0 {
		return h.GivenQueryStrings.Encode()
	}

	return ""
}

func (h *handlerTestCase) httpGet(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", h.baseURL, path), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = h.rawQuery()
	return http.DefaultClient.Do(req)
}
