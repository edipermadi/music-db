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

func TestTheoryHandler_GetScale(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetScale: []interface{}{&theory.DetailedScale{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetScale: []interface{}{nil, theory.ErrScaleNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("GetScale", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.GetScale...)

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

func TestTheoryHandler_ListScaleKeys(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListScaleKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("ListScaleKeys", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.ListScaleKeys...)

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
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListScales: []interface{}{[]theory.DetailedScale{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("ListScales", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.ListScales...)

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
