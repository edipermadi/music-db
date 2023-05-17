package theory_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestTheoryHandler_ListChordKeys(t *testing.T) {
	type testCase struct {
		Title               string
		GivenQueryString    url.Values
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedFilter      theory.KeyFilter
		ExpectedStatus      int
	}

	intValue := 1
	boolValue := true
	testCases := []testCase{
		{
			Title: "Returns200WhenSucceededWithoutFilter",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithScaleIDFilter",
			GivenQueryString: url.Values{
				"scale_id": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{ScaleID: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithTonicIDFilter",
			GivenQueryString: url.Values{
				"tonic_id": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{TonicID: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithNumberFilter",
			GivenQueryString: url.Values{
				"number": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{Number: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithPerfectionFilter",
			GivenQueryString: url.Values{
				"perfection": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{Perfection: &intValue},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithImperfectionFilter",
			GivenQueryString: url.Values{
				"imperfection": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{Imperfection: &intValue},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithBalancedFilter",
			GivenQueryString: url.Values{
				"balanced": []string{"true"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{Balanced: &boolValue},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRotationalSymmetricFilter",
			GivenQueryString: url.Values{
				"rotational_symmetric": []string{"true"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{RotationalSymmetric: &boolValue},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRotationalSymmetryLevelFilter",
			GivenQueryString: url.Values{
				"rotational_symmetry_level": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{RotationalSymmetryLevel: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithReflectionalSymmetricFilter",
			GivenQueryString: url.Values{
				"reflectional_symmetric": []string{"true"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{ReflectionalSymmetric: &boolValue},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithPalindromicFilter",
			GivenQueryString: url.Values{
				"palindromic": []string{"true"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{Palindromic: &boolValue},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithCardinalityFilter",
			GivenQueryString: url.Values{
				"cardinality": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordKeys: []interface{}{[]theory.DetailedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.KeyFilter{Cardinality: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("ListChordKeys", mock2.Anything, mock2.Anything, tc.ExpectedFilter, mock2.Anything).
				Return(tc.ServiceReturnValues.ListChordKeys...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/chords/1/keys", server.URL), nil)
			require.NoError(t, err)

			if len(tc.GivenQueryString) > 0 {
				req.URL.RawQuery = tc.GivenQueryString.Encode()
			}

			resp, err := http.DefaultClient.Do(req)
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
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChordPitches: []interface{}{[]theory.DetailedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("ListChordPitches", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.ListChordPitches...)

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
		Title               string
		GivenQueryStrings   url.Values
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedFilter      theory.ChordFilter
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceededWithoutFilter",
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithChordQualityIDFilter",
			GivenQueryStrings: url.Values{
				"chord_quality_id": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.ChordFilter{ChordQualityID: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRootIDFilter",
			GivenQueryStrings: url.Values{
				"root_id": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.ChordFilter{RootID: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithNumberFilter",
			GivenQueryStrings: url.Values{
				"number": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.ChordFilter{Number: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithCardinalityFilter",
			GivenQueryStrings: url.Values{
				"cardinality": []string{"1"},
			},
			ServiceReturnValues: mock.ServiceReturnValues{
				ListChords: []interface{}{[]theory.DetailedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedFilter: theory.ChordFilter{Cardinality: 1},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("ListChords", mock2.Anything, tc.ExpectedFilter, mock2.Anything).
				Return(tc.ServiceReturnValues.ListChords...)

			theory.NewHandler(logger, service).InstallEndpoints(router)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/chords", server.URL), nil)
			require.NoError(t, err)

			if len(tc.GivenQueryStrings) > 0 {
				req.URL.RawQuery = tc.GivenQueryStrings.Encode()
			}

			resp, err := http.DefaultClient.Do(req)
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

func TestTheoryHandler_GetChord(t *testing.T) {
	type testCase struct {
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetChord: []interface{}{&theory.DetailedChord{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetChord: []interface{}{nil, theory.ErrChordNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("GetChord", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.GetChord...)

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
		Title               string
		ServiceReturnValues mock.ServiceReturnValues
		ExpectedStatus      int
	}

	testCases := []testCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetChordQuality: []interface{}{&theory.DetailedChordQuality{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			ServiceReturnValues: mock.ServiceReturnValues{
				GetChordQuality: []interface{}{nil, theory.ErrChordQualityNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.ServiceReturnValues{
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

			service := &mock.TheoryService{}
			service.On("GetChordQuality", mock2.Anything, mock2.Anything).
				Return(tc.ServiceReturnValues.GetChordQuality...)

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
