package theory_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
)

func TestTheoryHandler_GetScale(t *testing.T) {
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				GetScale: []interface{}{&theory.DetailedScale{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				GetScale: []interface{}{nil, theory.ErrScaleNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				GetScale: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/scales/1")
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
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListScaleKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListScaleKeys: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/scales/1/keys")
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

func TestTheoryHandler_ListScalePitches(t *testing.T) {
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListScalePitches: []interface{}{[]theory.SimplifiedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListScalePitches: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/scales/1/pitches")
			require.NoError(t, err)

			defer func() { _ = resp.Body.Close() }()

			require.Equal(t, tc.ExpectedStatus, resp.StatusCode)
			if tc.ExpectedStatus == http.StatusOK {
				var decoded []theory.SimplifiedPitch
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&decoded))
				require.NotEmpty(t, decoded)
			}
		})
	}
}

func TestTheoryHandler_ListScales(t *testing.T) {
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListScales: []interface{}{[]theory.SimplifiedScale{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListScales: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/scales")
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
