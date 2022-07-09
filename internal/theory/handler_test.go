package theory_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewHandler(t *testing.T) {
	logger, err := zap.NewProduction()
	require.NoError(t, err)

	service := &MockService{}
	instance := theory.NewHandler(logger, service)
	require.Implements(t, (*theory.Handler)(nil), instance)
}

func TestTheoryHandler_GetChord(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChord: []interface{}{&theory.DetailedChord{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChord: []interface{}{nil, theory.ErrChordNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChord: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("GetChord", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetChord...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/chords/1", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded theory.DetailedChord
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_GetChordQuality(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChordQuality: []interface{}{&theory.DetailedChordQuality{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChordQuality: []interface{}{nil, theory.ErrChordQualityNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetChordQuality: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("GetChordQuality", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetChordQuality...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/chords/1/quality", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded theory.DetailedChordQuality
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_GetKey(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetKey: []interface{}{&theory.DetailedKey{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetKey: []interface{}{nil, theory.ErrKeyNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetKey: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("GetKey", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetKey...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/keys/1", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded theory.DetailedKey
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_GetScale(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetScale: []interface{}{&theory.DetailedScale{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetScale: []interface{}{nil, theory.ErrScaleNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				GetScale: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("GetScale", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.GetScale...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/scales/1", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded theory.DetailedScale
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListChordKeys(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordKeys: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListChordKeys", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListChordKeys...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/chords/1/keys", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedKey
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListChordPitches(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChordPitches: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListChordPitches", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListChordPitches...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/chords/1/pitches", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedPitch
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListChords(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListChords: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListChords", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListChords...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/chords", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedChord
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListKeyChords(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyChords: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListKeyChords", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeyChords...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/keys/1/chords", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedChord
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListKeyModes(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyModes: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyModes: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListKeyModes", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeyModes...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/keys/1/modes", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedKey
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListKeyPitches(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeyPitches: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListKeyPitches", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeyPitches...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/keys/1/pitches", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedPitch
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListKeys(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListKeys: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListKeys", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListKeys...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/keys", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedKey
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListPitches(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListPitches: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListPitches", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListPitches...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/pitches", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedPitch
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListScaleKeys(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScaleKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScaleKeys: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListScaleKeys", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListScaleKeys...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/scales/1/keys", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedKey
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListScales(t *testing.T) {
	type testCase struct {
		Title                      string
		MockRepositoryReturnValues MockRepositoryReturnValues
		ExpectedStatus             int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScales: []interface{}{[]theory.DetailedScale{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			MockRepositoryReturnValues: MockRepositoryReturnValues{
				ListScales: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			logger, err := zap.NewProduction()
			require.NoError(t, err)

			server, router := mockServer()
			defer server.Close()

			service := &MockService{}
			service.On("ListScales", mock.Anything, mock.Anything).
				Return(tc.MockRepositoryReturnValues.ListScales...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/scales", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.DetailedScale
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func mockServer() (*httptest.Server, *mux.Router) {
	router := mux.NewRouter()
	server := httptest.NewServer(router)
	return server, router
}
