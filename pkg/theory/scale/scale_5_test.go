package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith5Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		// 5.1
		{Scale: scale.Epathitonic, Name: "Epathitonic", Cardinality: 5, Number: 2378, Perfection: 4, Imperfection: 1},
		{Scale: scale.Mynitonic, Name: "Mynitonic", Cardinality: 5, Number: 2644, Perfection: 4, Imperfection: 1},
		{Scale: scale.Rocritonic, Name: "Rocritonic", Cardinality: 5, Number: 2386, Perfection: 4, Imperfection: 1},
		{Scale: scale.Pentatonic, Name: "Pentatonic", Cardinality: 5, Number: 2708, Perfection: 4, Imperfection: 1},
		{Scale: scale.Thaptitonic, Name: "Thaptitonic", Cardinality: 5, Number: 2642, Perfection: 4, Imperfection: 1},

		// 5.2
		{Scale: scale.Magitonic, Name: "Magitonic", Cardinality: 5, Number: 2197, Perfection: 3, Imperfection: 2},
		{Scale: scale.Daditonic, Name: "Daditonic", Cardinality: 5, Number: 2392, Perfection: 3, Imperfection: 2},
		{Scale: scale.Aeolyphritonic, Name: "Aeolyphritonic", Cardinality: 5, Number: 2756, Perfection: 3, Imperfection: 2},
		{Scale: scale.Gycritonic, Name: "Gycritonic", Cardinality: 5, Number: 2834, Perfection: 3, Imperfection: 2},
		{Scale: scale.Pyritonic, Name: "Pyritonic", Cardinality: 5, Number: 3146, Perfection: 3, Imperfection: 2},

		// 5.3
		{Scale: scale.Gathitonic, Name: "Gathitonic", Cardinality: 5, Number: 2213, Perfection: 3, Imperfection: 2},
		{Scale: scale.Ionitonic, Name: "Ionitonic", Cardinality: 5, Number: 2648, Perfection: 3, Imperfection: 2},
		{Scale: scale.Phrynitonic, Name: "Phrynitonic", Cardinality: 5, Number: 2402, Perfection: 3, Imperfection: 2},
		{Scale: scale.Stathitonic, Name: "Stathitonic", Cardinality: 5, Number: 2836, Perfection: 3, Imperfection: 2},
		{Scale: scale.Thalitonic, Name: "Thalitonic", Cardinality: 5, Number: 3154, Perfection: 3, Imperfection: 2},

		// 5.4
		{Scale: scale.Zolitonic, Name: "Zolitonic", Cardinality: 5, Number: 2225, Perfection: 3, Imperfection: 2},
		{Scale: scale.Epogitonic, Name: "Epogitonic", Cardinality: 5, Number: 2840, Perfection: 3, Imperfection: 2},
		{Scale: scale.Lanitonic, Name: "Lanitonic", Cardinality: 5, Number: 3170, Perfection: 3, Imperfection: 2},
		{Scale: scale.Paptitonic, Name: "Paptitonic", Cardinality: 5, Number: 2245, Perfection: 3, Imperfection: 2},
		{Scale: scale.Ionacritonic, Name: "Ionacritonic", Cardinality: 5, Number: 3160, Perfection: 3, Imperfection: 2},

		// 5.5
		{Scale: scale.Phraditonic, Name: "Phraditonic", Cardinality: 5, Number: 2246, Perfection: 3, Imperfection: 2},
		{Scale: scale.Aeoloritonic, Name: "Aeoloritonic", Cardinality: 5, Number: 3176, Perfection: 3, Imperfection: 2},
		{Scale: scale.Gonitonic, Name: "Gonitonic", Cardinality: 5, Number: 2257, Perfection: 3, Imperfection: 2},
		{Scale: scale.Dalitonic, Name: "Dalitonic", Cardinality: 5, Number: 3352, Perfection: 3, Imperfection: 2},
		{Scale: scale.Dygitonic, Name: "Dygitonic", Cardinality: 5, Number: 2609, Perfection: 3, Imperfection: 2},

		// 5.6
		{Scale: scale.Aeracritonic, Name: "Aeracritonic", Cardinality: 5, Number: 2258, Perfection: 3, Imperfection: 2},
		{Scale: scale.Byptitonic, Name: "Byptitonic", Cardinality: 5, Number: 3368, Perfection: 3, Imperfection: 2},
		{Scale: scale.Daritonic, Name: "Daritonic", Cardinality: 5, Number: 2641, Perfection: 3, Imperfection: 2},
		{Scale: scale.Lonitonic, Name: "Lonitonic", Cardinality: 5, Number: 2374, Perfection: 3, Imperfection: 2},
		{Scale: scale.Ionycritonic, Name: "Ionycritonic", Cardinality: 5, Number: 2612, Perfection: 3, Imperfection: 2},

		// 5.7
		{Scale: scale.Lothitonic, Name: "Lothitonic", Cardinality: 5, Number: 2260, Perfection: 3, Imperfection: 2},
		{Scale: scale.Phratonic, Name: "Phratonic", Cardinality: 5, Number: 3400, Perfection: 3, Imperfection: 2},
		{Scale: scale.Aerathitonic, Name: "Aerathitonic", Cardinality: 5, Number: 2705, Perfection: 3, Imperfection: 2},
		{Scale: scale.Saritonic, Name: "Saritonic", Cardinality: 5, Number: 2630, Perfection: 3, Imperfection: 2},
		{Scale: scale.Zoptitonic, Name: "Zoptitonic", Cardinality: 5, Number: 2330, Perfection: 3, Imperfection: 2},

		// 5.8
		{Scale: scale.Dolitonic, Name: "Dolitonic", Cardinality: 5, Number: 2189, Perfection: 2, Imperfection: 3},
		{Scale: scale.Poritonic, Name: "Poritonic", Cardinality: 5, Number: 2264, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aerylitonic, Name: "Aerylitonic", Cardinality: 5, Number: 3464, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zagitonic, Name: "Zagitonic", Cardinality: 5, Number: 2833, Perfection: 2, Imperfection: 3},
		{Scale: scale.Lagitonic, Name: "Lagitonic", Cardinality: 5, Number: 3142, Perfection: 2, Imperfection: 3},

		// 5.9
		{Scale: scale.Molitonic, Name: "Molitonic", Cardinality: 5, Number: 2195, Perfection: 2, Imperfection: 3},
		{Scale: scale.Staptitonic, Name: "Staptitonic", Cardinality: 5, Number: 2360, Perfection: 2, Imperfection: 3},
		{Scale: scale.Mothitonic, Name: "Mothitonic", Cardinality: 5, Number: 2500, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeritonic, Name: "Aeritonic", Cardinality: 5, Number: 3620, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ragitonic, Name: "Ragitonic", Cardinality: 5, Number: 3145, Perfection: 2, Imperfection: 3},

		// 5.10
		{Scale: scale.Ionaditonic, Name: "Ionaditonic", Cardinality: 5, Number: 2198, Perfection: 2, Imperfection: 3},
		{Scale: scale.Bocritonic, Name: "Bocritonic", Cardinality: 5, Number: 2408, Perfection: 2, Imperfection: 3},
		{Scale: scale.Gythitonic, Name: "Gythitonic", Cardinality: 5, Number: 2884, Perfection: 2, Imperfection: 3},
		{Scale: scale.Pagitonic, Name: "Pagitonic", Cardinality: 5, Number: 3346, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeolythitonic, Name: "Aeolythitonic", Cardinality: 5, Number: 2597, Perfection: 2, Imperfection: 3},

		// 5.11
		{Scale: scale.Zacritonic, Name: "Zacritonic", Cardinality: 5, Number: 2201, Perfection: 2, Imperfection: 3},
		{Scale: scale.Laritonic, Name: "Laritonic", Cardinality: 5, Number: 2456, Perfection: 2, Imperfection: 3},
		{Scale: scale.Thacritonic, Name: "Thacritonic", Cardinality: 5, Number: 3268, Perfection: 2, Imperfection: 3},
		{Scale: scale.Styditonic, Name: "Styditonic", Cardinality: 5, Number: 2441, Perfection: 2, Imperfection: 3},
		{Scale: scale.Loritonic, Name: "Loritonic", Cardinality: 5, Number: 3148, Perfection: 2, Imperfection: 3},

		// 5.12
		{Scale: scale.Aeolyritonic, Name: "Aeolyritonic", Cardinality: 5, Number: 2204, Perfection: 2, Imperfection: 3},
		{Scale: scale.Goritonic, Name: "Goritonic", Cardinality: 5, Number: 2504, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeoloditonic, Name: "Aeoloditonic", Cardinality: 5, Number: 3652, Perfection: 2, Imperfection: 3},
		{Scale: scale.Doptitonic, Name: "Doptitonic", Cardinality: 5, Number: 3209, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeraphitonic, Name: "Aeraphitonic", Cardinality: 5, Number: 2323, Perfection: 2, Imperfection: 3},

		// 5.13
		{Scale: scale.Zathitonic, Name: "Zathitonic", Cardinality: 5, Number: 2211, Perfection: 2, Imperfection: 3},
		{Scale: scale.Raditonic, Name: "Raditonic", Cardinality: 5, Number: 2616, Perfection: 2, Imperfection: 3},
		{Scale: scale.Stonitonic, Name: "Stonitonic", Cardinality: 5, Number: 2274, Perfection: 2, Imperfection: 3},
		{Scale: scale.Syptitonic, Name: "Syptitonic", Cardinality: 5, Number: 3624, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ionythitonic, Name: "Ionythitonic", Cardinality: 5, Number: 3153, Perfection: 2, Imperfection: 3},

		// 5.14
		{Scale: scale.Aeolanitonic, Name: "Aeolanitonic", Cardinality: 5, Number: 2217, Perfection: 2, Imperfection: 3},
		{Scale: scale.Danitonic, Name: "Danitonic", Cardinality: 5, Number: 2712, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ionaritonic, Name: "Ionaritonic", Cardinality: 5, Number: 2658, Perfection: 2, Imperfection: 3},
		{Scale: scale.Dynitonic, Name: "Dynitonic", Cardinality: 5, Number: 2442, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zyditonic, Name: "Zyditonic", Cardinality: 5, Number: 3156, Perfection: 2, Imperfection: 3},

		// 5.15
		{Scale: scale.Aeolacritonic, Name: "Aeolacritonic", Cardinality: 5, Number: 2228, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zythitonic, Name: "Zythitonic", Cardinality: 5, Number: 2888, Perfection: 2, Imperfection: 3},
		{Scale: scale.Dyritonic, Name: "Dyritonic", Cardinality: 5, Number: 3362, Perfection: 2, Imperfection: 3},
		{Scale: scale.Koptitonic, Name: "Koptitonic", Cardinality: 5, Number: 2629, Perfection: 2, Imperfection: 3},
		{Scale: scale.Thocritonic, Name: "Thocritonic", Cardinality: 5, Number: 2326, Perfection: 2, Imperfection: 3},

		// 5.16
		{Scale: scale.Lycritonic, Name: "Lycritonic", Cardinality: 5, Number: 2249, Perfection: 2, Imperfection: 3},
		{Scale: scale.Daptitonic, Name: "Daptitonic", Cardinality: 5, Number: 3224, Perfection: 2, Imperfection: 3},
		{Scale: scale.Kygitonic, Name: "Kygitonic", Cardinality: 5, Number: 2353, Perfection: 2, Imperfection: 3},
		{Scale: scale.Mocritonic, Name: "Mocritonic", Cardinality: 5, Number: 2444, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zynitonic, Name: "Zynitonic", Cardinality: 5, Number: 3172, Perfection: 2, Imperfection: 3},

		// 5.17
		{Scale: scale.Epygitonic, Name: "Epygitonic", Cardinality: 5, Number: 2250, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zaptitonic, Name: "Zaptitonic", Cardinality: 5, Number: 3240, Perfection: 2, Imperfection: 3},
		{Scale: scale.Kagitonic, Name: "Kagitonic", Cardinality: 5, Number: 2385, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zogitonic, Name: "Zogitonic", Cardinality: 5, Number: 2700, Perfection: 2, Imperfection: 3},
		{Scale: scale.Epyritonic, Name: "Epyritonic", Cardinality: 5, Number: 2610, Perfection: 2, Imperfection: 3},

		// 5.18
		{Scale: scale.Zothitonic, Name: "Zothitonic", Cardinality: 5, Number: 2252, Perfection: 2, Imperfection: 3},
		{Scale: scale.Phrolitonic, Name: "Phrolitonic", Cardinality: 5, Number: 3272, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ionagitonic, Name: "Ionagitonic", Cardinality: 5, Number: 2449, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeolapritonic, Name: "Aeolapritonic", Cardinality: 5, Number: 3212, Perfection: 2, Imperfection: 3},
		{Scale: scale.Kyritonic, Name: "Kyritonic", Cardinality: 5, Number: 2329, Perfection: 2, Imperfection: 3},

		// 5.19
		{Scale: scale.Ionyptitonic, Name: "Ionyptitonic", Cardinality: 5, Number: 2276, Perfection: 2, Imperfection: 3},
		{Scale: scale.Gyritonic, Name: "Gyritonic", Cardinality: 5, Number: 3656, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zalitonic, Name: "Zalitonic", Cardinality: 5, Number: 3217, Perfection: 2, Imperfection: 3},
		{Scale: scale.Stolitonic, Name: "Stolitonic", Cardinality: 5, Number: 2339, Perfection: 2, Imperfection: 3},
		{Scale: scale.Bylitonic, Name: "Bylitonic", Cardinality: 5, Number: 2332, Perfection: 2, Imperfection: 3},

		// 5.20
		{Scale: scale.Thoditonic, Name: "Thoditonic", Cardinality: 5, Number: 2345, Perfection: 2, Imperfection: 3},
		{Scale: scale.Dogitonic, Name: "Dogitonic", Cardinality: 5, Number: 2380, Perfection: 2, Imperfection: 3},
		{Scale: scale.Phralitonic, Name: "Phralitonic", Cardinality: 5, Number: 2660, Perfection: 2, Imperfection: 3},
		{Scale: scale.Garitonic, Name: "Garitonic", Cardinality: 5, Number: 2450, Perfection: 2, Imperfection: 3},
		{Scale: scale.Soptitonic, Name: "Soptitonic", Cardinality: 5, Number: 3220, Perfection: 2, Imperfection: 3},

		// 5.21
		{Scale: scale.Kataritonic, Name: "Kataritonic", Cardinality: 5, Number: 2346, Perfection: 2, Imperfection: 3},
		{Scale: scale.Sylitonic, Name: "Sylitonic", Cardinality: 5, Number: 2388, Perfection: 2, Imperfection: 3},
		{Scale: scale.Thonitonic, Name: "Thonitonic", Cardinality: 5, Number: 2724, Perfection: 2, Imperfection: 3},
		{Scale: scale.Phropitonic, Name: "Phropitonic", Cardinality: 5, Number: 2706, Perfection: 2, Imperfection: 3},
		{Scale: scale.Staditonic, Name: "Staditonic", Cardinality: 5, Number: 2634, Perfection: 2, Imperfection: 3},

		// 5.22
		{Scale: scale.Lyditonic, Name: "Lyditonic", Cardinality: 5, Number: 2354, Perfection: 2, Imperfection: 3},
		{Scale: scale.Mythitonic, Name: "Mythitonic", Cardinality: 5, Number: 2452, Perfection: 2, Imperfection: 3},
		{Scale: scale.Sogitonic, Name: "Sogitonic", Cardinality: 5, Number: 3236, Perfection: 2, Imperfection: 3},
		{Scale: scale.Gothitonic, Name: "Gothitonic", Cardinality: 5, Number: 2377, Perfection: 2, Imperfection: 3},
		{Scale: scale.Rothitonic, Name: "Rothitonic", Cardinality: 5, Number: 2636, Perfection: 2, Imperfection: 3},

		// 5.23
		{Scale: scale.Zylitonic, Name: "Zylitonic", Cardinality: 5, Number: 2187, Perfection: 1, Imperfection: 4},
		{Scale: scale.Zoditonic, Name: "Zoditonic", Cardinality: 5, Number: 2232, Perfection: 1, Imperfection: 4},
		{Scale: scale.Zaritonic, Name: "Zaritonic", Cardinality: 5, Number: 2952, Perfection: 1, Imperfection: 4},
		{Scale: scale.Phrythitonic, Name: "Phrythitonic", Cardinality: 5, Number: 3618, Perfection: 1, Imperfection: 4},
		{Scale: scale.Rolitonic, Name: "Rolitonic", Cardinality: 5, Number: 3141, Perfection: 1, Imperfection: 4},

		// 5.24
		{Scale: scale.Ranitonic, Name: "Ranitonic", Cardinality: 5, Number: 2190, Perfection: 1, Imperfection: 4},
		{Scale: scale.Laditonic, Name: "Laditonic", Cardinality: 5, Number: 2280, Perfection: 1, Imperfection: 4},
		{Scale: scale.Poditonic, Name: "Poditonic", Cardinality: 5, Number: 3720, Perfection: 1, Imperfection: 4},
		{Scale: scale.Ionothitonic, Name: "Ionothitonic", Cardinality: 5, Number: 3345, Perfection: 1, Imperfection: 4},
		{Scale: scale.Kanitonic, Name: "Kanitonic", Cardinality: 5, Number: 2595, Perfection: 1, Imperfection: 4},

		// 5.25
		{Scale: scale.Ryphitonic, Name: "Ryphitonic", Cardinality: 5, Number: 2202, Perfection: 1, Imperfection: 4},
		{Scale: scale.Gylitonic, Name: "Gylitonic", Cardinality: 5, Number: 2472, Perfection: 1, Imperfection: 4},
		{Scale: scale.Aeolycritonic, Name: "Aeolycritonic", Cardinality: 5, Number: 3396, Perfection: 1, Imperfection: 4},
		{Scale: scale.Pynitonic, Name: "Pynitonic", Cardinality: 5, Number: 2697, Perfection: 1, Imperfection: 4},
		{Scale: scale.Zanitonic, Name: "Zanitonic", Cardinality: 5, Number: 2598, Perfection: 1, Imperfection: 4},

		// 5.26
		{Scale: scale.Phronitonic, Name: "Phronitonic", Cardinality: 5, Number: 2214, Perfection: 1, Imperfection: 4},
		{Scale: scale.Banitonic, Name: "Banitonic", Cardinality: 5, Number: 2664, Perfection: 1, Imperfection: 4},
		{Scale: scale.Aeronitonic, Name: "Aeronitonic", Cardinality: 5, Number: 2466, Perfection: 1, Imperfection: 4},
		{Scale: scale.Golitonic, Name: "Golitonic", Cardinality: 5, Number: 3348, Perfection: 1, Imperfection: 4},
		{Scale: scale.Dyptitonic, Name: "Dyptitonic", Cardinality: 5, Number: 2601, Perfection: 1, Imperfection: 4},

		// 5.27
		{Scale: scale.Aerynitonic, Name: "Aerynitonic", Cardinality: 5, Number: 2220, Perfection: 1, Imperfection: 4},
		{Scale: scale.Palitonic, Name: "Palitonic", Cardinality: 5, Number: 2760, Perfection: 1, Imperfection: 4},
		{Scale: scale.Stothitonic, Name: "Stothitonic", Cardinality: 5, Number: 2850, Perfection: 1, Imperfection: 4},
		{Scale: scale.Aerophitonic, Name: "Aerophitonic", Cardinality: 5, Number: 3210, Perfection: 1, Imperfection: 4},
		{Scale: scale.Katagitonic, Name: "Katagitonic", Cardinality: 5, Number: 2325, Perfection: 1, Imperfection: 4},

		// 5.28
		{Scale: scale.Ionoditonic, Name: "Ionoditonic", Cardinality: 5, Number: 2226, Perfection: 1, Imperfection: 4},
		{Scale: scale.Bogitonic, Name: "Bogitonic", Cardinality: 5, Number: 2856, Perfection: 1, Imperfection: 4},
		{Scale: scale.Mogitonic, Name: "Mogitonic", Cardinality: 5, Number: 3234, Perfection: 1, Imperfection: 4},
		{Scale: scale.Docritonic, Name: "Docritonic", Cardinality: 5, Number: 2373, Perfection: 1, Imperfection: 4},
		{Scale: scale.Epaditonic, Name: "Epaditonic", Cardinality: 5, Number: 2604, Perfection: 1, Imperfection: 4},

		// 5.29
		{Scale: scale.Mixitonic, Name: "Mixitonic", Cardinality: 5, Number: 2341, Perfection: 1, Imperfection: 4},
		{Scale: scale.Phrothitonic, Name: "Phrothitonic", Cardinality: 5, Number: 2348, Perfection: 1, Imperfection: 4},
		{Scale: scale.Katycritonic, Name: "Katycritonic", Cardinality: 5, Number: 2404, Perfection: 1, Imperfection: 4},
		{Scale: scale.Ionalitonic, Name: "Ionalitonic", Cardinality: 5, Number: 2852, Perfection: 1, Imperfection: 4},
		{Scale: scale.Loptitonic, Name: "Loptitonic", Cardinality: 5, Number: 3218, Perfection: 1, Imperfection: 4},

		// 5.30
		{Scale: scale.Thyritonic, Name: "Thyritonic", Cardinality: 5, Number: 2342, Perfection: 1, Imperfection: 4},
		{Scale: scale.Thoptitonic, Name: "Thoptitonic", Cardinality: 5, Number: 2356, Perfection: 1, Imperfection: 4},
		{Scale: scale.Bycritonic, Name: "Bycritonic", Cardinality: 5, Number: 2468, Perfection: 1, Imperfection: 4},
		{Scale: scale.Pathitonic, Name: "Pathitonic", Cardinality: 5, Number: 3364, Perfection: 1, Imperfection: 4},
		{Scale: scale.Myditonic, Name: "Myditonic", Cardinality: 5, Number: 2633, Perfection: 1, Imperfection: 4},

		// 5.31
		{Scale: scale.Bolitonic, Name: "Bolitonic", Cardinality: 5, Number: 2218, Perfection: 0, Imperfection: 5},
		{Scale: scale.Bothitonic, Name: "Bothitonic", Cardinality: 5, Number: 2728, Perfection: 0, Imperfection: 5},
		{Scale: scale.Kataditonic, Name: "Kataditonic", Cardinality: 5, Number: 2722, Perfection: 0, Imperfection: 5},
		{Scale: scale.Koditonic, Name: "Koditonic", Cardinality: 5, Number: 2698, Perfection: 0, Imperfection: 5},
		{Scale: scale.Tholitonic, Name: "Tholitonic", Cardinality: 5, Number: 2602, Perfection: 0, Imperfection: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Cardinality, tc.Scale.Cardinality())
			assert.Equal(t, tc.Number, tc.Scale.Number())

			result := tc.Scale.Perfection()
			assert.Equal(t, tc.Perfection, result.Perfection)
			assert.Equal(t, tc.Imperfection, result.Imperfection)
		})
	}
}
