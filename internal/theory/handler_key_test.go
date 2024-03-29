package theory_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/edipermadi/music-db/internal/platform/api"
	"github.com/edipermadi/music-db/internal/theory"
	"github.com/edipermadi/music-db/mock"
	"github.com/stretchr/testify/require"
)

func TestTheoryHandler_ListKeyChords(t *testing.T) {
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceededWithoutFilter",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithChordQualityIDFilter",
			GivenQueryStrings: url.Values{
				"chord_quality_id": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRootIDFilter",
			GivenQueryStrings: url.Values{
				"root_id": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithZeitlerNumberFilter",
			GivenQueryStrings: url.Values{
				"zeitler_number": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRingNumberFilter",
			GivenQueryStrings: url.Values{
				"ring_number": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithCardinalityFilter",
			GivenQueryStrings: url.Values{
				"cardinality": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{[]theory.SimplifiedChord{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyChords: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/keys/1/chords")
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
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceededWithoutFilter",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithScaleIDFilter",
			GivenQueryStrings: url.Values{
				"scale_id": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithTonicIDFilter",
			GivenQueryStrings: url.Values{
				"tonic_id": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithZeitlerNumberFilter",
			GivenQueryStrings: url.Values{
				"zeitler_number": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRingNumberFilter",
			GivenQueryStrings: url.Values{
				"ring_number": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithPerfectionFilter",
			GivenQueryStrings: url.Values{
				"perfection": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithImperfectionFilter",
			GivenQueryStrings: url.Values{
				"imperfection": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithBalancedFilter",
			GivenQueryStrings: url.Values{
				"balanced": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRotationalSymmetricFilter",
			GivenQueryStrings: url.Values{
				"rotational_symmetric": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRotationalSymmetryLevelFilter",
			GivenQueryStrings: url.Values{
				"rotational_symmetry_level": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithReflectionalSymmetricFilter",
			GivenQueryStrings: url.Values{
				"reflectional_symmetric": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithPalindromicFilter",
			GivenQueryStrings: url.Values{
				"palindromic": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithCardinalityFilter",
			GivenQueryStrings: url.Values{
				"cardinality": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyModes: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/keys/1/modes")
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
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyPitches: []interface{}{[]theory.SimplifiedPitch{{ID: 1, Name: "name"}}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeyPitches: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/keys/1/pitches")
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
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceededWithoutFilter",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithScaleIDFilter",
			GivenQueryStrings: url.Values{
				"scale_id": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithTonicIDFilter",
			GivenQueryStrings: url.Values{
				"tonic_id": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithZeitlerNumberFilter",
			GivenQueryStrings: url.Values{
				"zeitler_number": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRingNumberFilter",
			GivenQueryStrings: url.Values{
				"ring_number": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithPerfectionFilter",
			GivenQueryStrings: url.Values{
				"perfection": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithImperfectionFilter",
			GivenQueryStrings: url.Values{
				"imperfection": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithBalancedFilter",
			GivenQueryStrings: url.Values{
				"balanced": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRotationalSymmetricFilter",
			GivenQueryStrings: url.Values{
				"rotational_symmetric": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithRotationalSymmetryLevelFilter",
			GivenQueryStrings: url.Values{
				"rotational_symmetry_level": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithReflectionalSymmetricFilter",
			GivenQueryStrings: url.Values{
				"reflectional_symmetric": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithPalindromicFilter",
			GivenQueryStrings: url.Values{
				"palindromic": []string{"true"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns200WhenSucceededWithCardinalityFilter",
			GivenQueryStrings: url.Values{
				"cardinality": []string{"1"},
			},
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{[]theory.SimplifiedKey{{ID: 1, Name: "name"}}, &api.Pagination{}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				ListKeys: []interface{}{nil, nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/keys")
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

func TestTheoryHandler_GetKey(t *testing.T) {
	testCases := []handlerTestCase{
		{
			Title: "Returns200WhenSucceeded",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				GetKey: []interface{}{&theory.DetailedKey{ID: 1, Name: "name"}, nil},
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Title: "Returns404WhenNotFound",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				GetKey: []interface{}{nil, theory.ErrKeyNotFound},
			},
			ExpectedStatus: http.StatusNotFound,
		},
		{
			Title: "Returns500WhenFailed",
			ServiceReturnValues: mock.TheoryServiceReturnValues{
				GetKey: []interface{}{nil, errors.New("error")},
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			server := tc.mockServer()
			defer server.Close()

			resp, err := tc.httpGet("/keys/1")
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
