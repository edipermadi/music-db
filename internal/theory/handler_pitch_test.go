package theory_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestTheoryHandler_GetPitch(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetPitch: []interface{}{&theory.DetailedPitch{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetPitch: []interface{}{nil, theory.ErrPitchNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetPitch: []interface{}{nil, errors.New("error")},
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

			service := &mock.TheoryService{}
			service.On("GetPitch", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.GetPitch...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/pitches/1", server.URL))
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

func TestTheoryHandler_ListPitchChords(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListPitchChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListPitchChords: []interface{}{nil, nil, errors.New("error")},
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

			service := &mock.TheoryService{}
			service.On("ListPitchChords", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.ListPitchChords...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/pitches/1/chords", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.SimplifiedChord
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListPitchKeys(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListPitchKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListPitchKeys: []interface{}{nil, nil, errors.New("error")},
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

			service := &mock.TheoryService{}
			service.On("ListPitchKeys", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.ListPitchKeys...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			resp, err := http.Get(fmt.Sprintf("%s/pitches/1/keys", server.URL))
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.SimplifiedChord
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListPitches(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("ListPitches", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.ListPitches...)

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
