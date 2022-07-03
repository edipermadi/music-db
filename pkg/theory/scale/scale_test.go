package scale_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/pitch"
	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScale_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		{Scale: scale.Invalid, Name: "InvalidScale"},

		// 3 notes
		{Scale: scale.Minoric, Name: "Minoric", Number: 2184, Imperfection: 3},

		// 4 notes

		// 4.1
		{Scale: scale.Thaptic, Name: "Thaptic", Number: 2193, Perfection: 2, Imperfection: 2},
		{Scale: scale.Lothic, Name: "Lothic", Number: 2328, Perfection: 2, Imperfection: 2},
		{Scale: scale.Phratic, Name: "Phratic", Number: 2244, Perfection: 2, Imperfection: 2},
		{Scale: scale.Aerathic, Name: "Aerathic", Number: 3144, Perfection: 2, Imperfection: 2},

		// 4.2
		{Scale: scale.Epathic, Name: "Epathic", Number: 2196, Perfection: 2, Imperfection: 2},
		{Scale: scale.Mynic, Name: "Mynic", Number: 2376, Perfection: 2, Imperfection: 2},
		{Scale: scale.Rothic, Name: "Rothic", Number: 2628, Perfection: 2, Imperfection: 2},
		{Scale: scale.Eporic, Name: "Eporic", Number: 2322, Perfection: 2, Imperfection: 2},

		// 4.3
		{Scale: scale.Zyphic, Name: "Zyphic", Number: 2185, Perfection: 1, Imperfection: 3},
		{Scale: scale.Epogic, Name: "Epogic", Number: 2200, Perfection: 1, Imperfection: 3},
		{Scale: scale.Lanic, Name: "Lanic", Number: 2440, Perfection: 1, Imperfection: 3},
		{Scale: scale.Pyrric, Name: "Pyrric", Number: 3140, Perfection: 1, Imperfection: 3},

		// 4.4
		{Scale: scale.Aeoloric, Name: "Aeoloric", Number: 2188, Perfection: 1, Imperfection: 3},
		{Scale: scale.Gonic, Name: "Gonic", Number: 2248, Perfection: 1, Imperfection: 3},
		{Scale: scale.Dalic, Name: "Dalic", Number: 3208, Perfection: 1, Imperfection: 3},
		{Scale: scale.Dygic, Name: "Dygic", Number: 2321, Perfection: 1, Imperfection: 3},

		// 4.5
		{Scale: scale.Daric, Name: "Daric", Number: 2194, Perfection: 1, Imperfection: 3},
		{Scale: scale.Lonic, Name: "Lonic", Number: 2344, Perfection: 1, Imperfection: 3},
		{Scale: scale.Phradic, Name: "Phradic", Number: 2372, Perfection: 1, Imperfection: 3},
		{Scale: scale.Bolic, Name: "Bolic", Number: 2596, Perfection: 1, Imperfection: 3},

		// 4.6
		{Scale: scale.Saric, Name: "Saric", Number: 2212, Perfection: 1, Imperfection: 3},
		{Scale: scale.Zoptic, Name: "Zoptic", Number: 2632, Perfection: 1, Imperfection: 3},
		{Scale: scale.Aeraphic, Name: "Aeraphic", Number: 2338, Perfection: 1, Imperfection: 3},
		{Scale: scale.Byptic, Name: "Byptic", Number: 2324, Perfection: 1, Imperfection: 3},

		// 4.7
		{Scale: scale.Aeolic, Name: "Aeolic", Number: 2186, Perfection: 0, Imperfection: 4},
		{Scale: scale.Koptic, Name: "Koptic", Number: 2216, Perfection: 0, Imperfection: 4},
		{Scale: scale.Mixolyric, Name: "Mixolyric", Number: 2696, Perfection: 0, Imperfection: 4},
		{Scale: scale.Lydic, Name: "Lydic", Number: 2594, Perfection: 0, Imperfection: 4},

		// 4.8
		{Scale: scale.Stathic, Name: "Stathic", Number: 2210, Perfection: 0, Imperfection: 4},
		{Scale: scale.Dadic, Name: "Dadic", Number: 2600, Perfection: 0, Imperfection: 4},

		// 4.9
		{Scale: scale.Phrynic, Name: "Phrynic", Number: 2340, Perfection: 0, Imperfection: 4},

		// 5 notes

		// 5.1
		{Scale: scale.Epathitonic, Name: "Epathitonic", Number: 2378, Perfection: 4, Imperfection: 1},
		{Scale: scale.Mynitonic, Name: "Mynitonic", Number: 2644, Perfection: 4, Imperfection: 1},
		{Scale: scale.Rocritonic, Name: "Rocritonic", Number: 2386, Perfection: 4, Imperfection: 1},
		{Scale: scale.Pentatonic, Name: "Pentatonic", Number: 2708, Perfection: 4, Imperfection: 1},
		{Scale: scale.Thaptitonic, Name: "Thaptitonic", Number: 2642, Perfection: 4, Imperfection: 1},

		// 5.2
		{Scale: scale.Magitonic, Name: "Magitonic", Number: 2197, Perfection: 3, Imperfection: 2},
		{Scale: scale.Daditonic, Name: "Daditonic", Number: 2392, Perfection: 3, Imperfection: 2},
		{Scale: scale.Aeolyphritonic, Name: "Aeolyphritonic", Number: 2756, Perfection: 3, Imperfection: 2},
		{Scale: scale.Gycritonic, Name: "Gycritonic", Number: 2834, Perfection: 3, Imperfection: 2},
		{Scale: scale.Pyritonic, Name: "Pyritonic", Number: 3146, Perfection: 3, Imperfection: 2},

		// 5.3
		{Scale: scale.Gathitonic, Name: "Gathitonic", Number: 2213, Perfection: 3, Imperfection: 2},
		{Scale: scale.Ionitonic, Name: "Ionitonic", Number: 2648, Perfection: 3, Imperfection: 2},
		{Scale: scale.Phrynitonic, Name: "Phrynitonic", Number: 2402, Perfection: 3, Imperfection: 2},
		{Scale: scale.Stathitonic, Name: "Stathitonic", Number: 2836, Perfection: 3, Imperfection: 2},
		{Scale: scale.Thalitonic, Name: "Thalitonic", Number: 3154, Perfection: 3, Imperfection: 2},

		// 5.4
		{Scale: scale.Zolitonic, Name: "Zolitonic", Number: 2225, Perfection: 3, Imperfection: 2},
		{Scale: scale.Epogitonic, Name: "Epogitonic", Number: 2840, Perfection: 3, Imperfection: 2},
		{Scale: scale.Lanitonic, Name: "Lanitonic", Number: 3170, Perfection: 3, Imperfection: 2},
		{Scale: scale.Paptitonic, Name: "Paptitonic", Number: 2245, Perfection: 3, Imperfection: 2},
		{Scale: scale.Ionacritonic, Name: "Ionacritonic", Number: 3160, Perfection: 3, Imperfection: 2},

		// 5.5
		{Scale: scale.Phraditonic, Name: "Phraditonic", Number: 2246, Perfection: 3, Imperfection: 2},
		{Scale: scale.Aeoloritonic, Name: "Aeoloritonic", Number: 3176, Perfection: 3, Imperfection: 2},
		{Scale: scale.Gonitonic, Name: "Gonitonic", Number: 2257, Perfection: 3, Imperfection: 2},
		{Scale: scale.Dalitonic, Name: "Dalitonic", Number: 3352, Perfection: 3, Imperfection: 2},
		{Scale: scale.Dygitonic, Name: "Dygitonic", Number: 2609, Perfection: 3, Imperfection: 2},

		// 5.6
		{Scale: scale.Aeracritonic, Name: "Aeracritonic", Number: 2258, Perfection: 3, Imperfection: 2},
		{Scale: scale.Byptitonic, Name: "Byptitonic", Number: 3368, Perfection: 3, Imperfection: 2},
		{Scale: scale.Daritonic, Name: "Daritonic", Number: 2641, Perfection: 3, Imperfection: 2},
		{Scale: scale.Lonitonic, Name: "Lonitonic", Number: 2374, Perfection: 3, Imperfection: 2},
		{Scale: scale.Ionycritonic, Name: "Ionycritonic", Number: 2612, Perfection: 3, Imperfection: 2},

		// 5.7
		{Scale: scale.Lothitonic, Name: "Lothitonic", Number: 2260, Perfection: 3, Imperfection: 2},
		{Scale: scale.Phratonic, Name: "Phratonic", Number: 3400, Perfection: 3, Imperfection: 2},
		{Scale: scale.Aerathitonic, Name: "Aerathitonic", Number: 2705, Perfection: 3, Imperfection: 2},
		{Scale: scale.Saritonic, Name: "Saritonic", Number: 2630, Perfection: 3, Imperfection: 2},
		{Scale: scale.Zoptitonic, Name: "Zoptitonic", Number: 2330, Perfection: 3, Imperfection: 2},

		// 5.8
		{Scale: scale.Dolitonic, Name: "Dolitonic", Number: 2189, Perfection: 2, Imperfection: 3},
		{Scale: scale.Poritonic, Name: "Poritonic", Number: 2264, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aerylitonic, Name: "Aerylitonic", Number: 3464, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zagitonic, Name: "Zagitonic", Number: 2833, Perfection: 2, Imperfection: 3},
		{Scale: scale.Lagitonic, Name: "Lagitonic", Number: 3142, Perfection: 2, Imperfection: 3},

		// 5.9
		{Scale: scale.Molitonic, Name: "Molitonic", Number: 2195, Perfection: 2, Imperfection: 3},
		{Scale: scale.Staptitonic, Name: "Staptitonic", Number: 2360, Perfection: 2, Imperfection: 3},
		{Scale: scale.Mothitonic, Name: "Mothitonic", Number: 2500, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeritonic, Name: "Aeritonic", Number: 3620, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ragitonic, Name: "Ragitonic", Number: 3145, Perfection: 2, Imperfection: 3},

		// 5.10
		{Scale: scale.Ionaditonic, Name: "Ionaditonic", Number: 2198, Perfection: 2, Imperfection: 3},
		{Scale: scale.Bocritonic, Name: "Bocritonic", Number: 2408, Perfection: 2, Imperfection: 3},
		{Scale: scale.Gythitonic, Name: "Gythitonic", Number: 2884, Perfection: 2, Imperfection: 3},
		{Scale: scale.Pagitonic, Name: "Pagitonic", Number: 3346, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeolythitonic, Name: "Aeolythitonic", Number: 2597, Perfection: 2, Imperfection: 3},

		// 5.11
		{Scale: scale.Zacritonic, Name: "Zacritonic", Number: 2201, Perfection: 2, Imperfection: 3},
		{Scale: scale.Laritonic, Name: "Laritonic", Number: 2456, Perfection: 2, Imperfection: 3},
		{Scale: scale.Thacritonic, Name: "Thacritonic", Number: 3268, Perfection: 2, Imperfection: 3},
		{Scale: scale.Styditonic, Name: "Styditonic", Number: 2441, Perfection: 2, Imperfection: 3},
		{Scale: scale.Loritonic, Name: "Loritonic", Number: 3148, Perfection: 2, Imperfection: 3},

		// 5.12
		{Scale: scale.Aeolyritonic, Name: "Aeolyritonic", Number: 2204, Perfection: 2, Imperfection: 3},
		{Scale: scale.Goritonic, Name: "Goritonic", Number: 2504, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeoloditonic, Name: "Aeoloditonic", Number: 3652, Perfection: 2, Imperfection: 3},
		{Scale: scale.Doptitonic, Name: "Doptitonic", Number: 3209, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeraphitonic, Name: "Aeraphitonic", Number: 2323, Perfection: 2, Imperfection: 3},

		// 5.13
		{Scale: scale.Zathitonic, Name: "Zathitonic", Number: 2211, Perfection: 2, Imperfection: 3},
		{Scale: scale.Raditonic, Name: "Raditonic", Number: 2616, Perfection: 2, Imperfection: 3},
		{Scale: scale.Stonitonic, Name: "Stonitonic", Number: 2274, Perfection: 2, Imperfection: 3},
		{Scale: scale.Syptitonic, Name: "Syptitonic", Number: 3624, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ionythitonic, Name: "Ionythitonic", Number: 3153, Perfection: 2, Imperfection: 3},

		// 5.14
		{Scale: scale.Aeolanitonic, Name: "Aeolanitonic", Number: 2217, Perfection: 2, Imperfection: 3},
		{Scale: scale.Danitonic, Name: "Danitonic", Number: 2712, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ionaritonic, Name: "Ionaritonic", Number: 2658, Perfection: 2, Imperfection: 3},
		{Scale: scale.Dynitonic, Name: "Dynitonic", Number: 2442, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zyditonic, Name: "Zyditonic", Number: 3156, Perfection: 2, Imperfection: 3},

		// 5.15
		{Scale: scale.Aeolacritonic, Name: "Aeolacritonic", Number: 2228, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zythitonic, Name: "Zythitonic", Number: 2888, Perfection: 2, Imperfection: 3},
		{Scale: scale.Dyritonic, Name: "Dyritonic", Number: 3362, Perfection: 2, Imperfection: 3},
		{Scale: scale.Koptitonic, Name: "Koptitonic", Number: 2629, Perfection: 2, Imperfection: 3},
		{Scale: scale.Thocritonic, Name: "Thocritonic", Number: 2326, Perfection: 2, Imperfection: 3},

		// 5.16
		{Scale: scale.Lycritonic, Name: "Lycritonic", Number: 2249, Perfection: 2, Imperfection: 3},
		{Scale: scale.Daptitonic, Name: "Daptitonic", Number: 3224, Perfection: 2, Imperfection: 3},
		{Scale: scale.Kygitonic, Name: "Kygitonic", Number: 2353, Perfection: 2, Imperfection: 3},
		{Scale: scale.Mocritonic, Name: "Mocritonic", Number: 2444, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zynitonic, Name: "Zynitonic", Number: 3172, Perfection: 2, Imperfection: 3},

		// 5.17
		{Scale: scale.Epygitonic, Name: "Epygitonic", Number: 2250, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zaptitonic, Name: "Zaptitonic", Number: 3240, Perfection: 2, Imperfection: 3},
		{Scale: scale.Kagitonic, Name: "Kagitonic", Number: 2385, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zogitonic, Name: "Zogitonic", Number: 2700, Perfection: 2, Imperfection: 3},
		{Scale: scale.Epyritonic, Name: "Epyritonic", Number: 2610, Perfection: 2, Imperfection: 3},

		// 5.18
		{Scale: scale.Zothitonic, Name: "Zothitonic", Number: 2252, Perfection: 2, Imperfection: 3},
		{Scale: scale.Phrolitonic, Name: "Phrolitonic", Number: 3272, Perfection: 2, Imperfection: 3},
		{Scale: scale.Ionagitonic, Name: "Ionagitonic", Number: 2449, Perfection: 2, Imperfection: 3},
		{Scale: scale.Aeolapritonic, Name: "Aeolapritonic", Number: 3212, Perfection: 2, Imperfection: 3},
		{Scale: scale.Kyritonic, Name: "Kyritonic", Number: 2329, Perfection: 2, Imperfection: 3},

		// 5.19
		{Scale: scale.Ionyptitonic, Name: "Ionyptitonic", Number: 2276, Perfection: 2, Imperfection: 3},
		{Scale: scale.Gyritonic, Name: "Gyritonic", Number: 3656, Perfection: 2, Imperfection: 3},
		{Scale: scale.Zalitonic, Name: "Zalitonic", Number: 3217, Perfection: 2, Imperfection: 3},
		{Scale: scale.Stolitonic, Name: "Stolitonic", Number: 2339, Perfection: 2, Imperfection: 3},
		{Scale: scale.Bylitonic, Name: "Bylitonic", Number: 2332, Perfection: 2, Imperfection: 3},

		// 5.20
		{Scale: scale.Thoditonic, Name: "Thoditonic", Number: 2345, Perfection: 2, Imperfection: 3},
		{Scale: scale.Dogitonic, Name: "Dogitonic", Number: 2380, Perfection: 2, Imperfection: 3},
		{Scale: scale.Phralitonic, Name: "Phralitonic", Number: 2660, Perfection: 2, Imperfection: 3},
		{Scale: scale.Garitonic, Name: "Garitonic", Number: 2450, Perfection: 2, Imperfection: 3},
		{Scale: scale.Soptitonic, Name: "Soptitonic", Number: 3220, Perfection: 2, Imperfection: 3},

		// 5.21
		{Scale: scale.Kataritonic, Name: "Kataritonic", Number: 2346, Perfection: 2, Imperfection: 3},
		{Scale: scale.Sylitonic, Name: "Sylitonic", Number: 2388, Perfection: 2, Imperfection: 3},
		{Scale: scale.Thonitonic, Name: "Thonitonic", Number: 2724, Perfection: 2, Imperfection: 3},
		{Scale: scale.Phropitonic, Name: "Phropitonic", Number: 2706, Perfection: 2, Imperfection: 3},
		{Scale: scale.Staditonic, Name: "Staditonic", Number: 2634, Perfection: 2, Imperfection: 3},

		// 5.22
		{Scale: scale.Lyditonic, Name: "Lyditonic", Number: 2354, Perfection: 2, Imperfection: 3},
		{Scale: scale.Mythitonic, Name: "Mythitonic", Number: 2452, Perfection: 2, Imperfection: 3},
		{Scale: scale.Sogitonic, Name: "Sogitonic", Number: 3236, Perfection: 2, Imperfection: 3},
		{Scale: scale.Gothitonic, Name: "Gothitonic", Number: 2377, Perfection: 2, Imperfection: 3},
		{Scale: scale.Rothitonic, Name: "Rothitonic", Number: 2636, Perfection: 2, Imperfection: 3},

		// 5.23
		{Scale: scale.Zylitonic, Name: "Zylitonic", Number: 2187, Perfection: 1, Imperfection: 4},
		{Scale: scale.Zoditonic, Name: "Zoditonic", Number: 2232, Perfection: 1, Imperfection: 4},
		{Scale: scale.Zaritonic, Name: "Zaritonic", Number: 2952, Perfection: 1, Imperfection: 4},
		{Scale: scale.Phrythitonic, Name: "Phrythitonic", Number: 3618, Perfection: 1, Imperfection: 4},
		{Scale: scale.Rolitonic, Name: "Rolitonic", Number: 3141, Perfection: 1, Imperfection: 4},

		// 5.24
		{Scale: scale.Ranitonic, Name: "Ranitonic", Number: 2190, Perfection: 1, Imperfection: 4},
		{Scale: scale.Laditonic, Name: "Laditonic", Number: 2280, Perfection: 1, Imperfection: 4},
		{Scale: scale.Poditonic, Name: "Poditonic", Number: 3720, Perfection: 1, Imperfection: 4},
		{Scale: scale.Ionothitonic, Name: "Ionothitonic", Number: 3345, Perfection: 1, Imperfection: 4},
		{Scale: scale.Kanitonic, Name: "Kanitonic", Number: 2595, Perfection: 1, Imperfection: 4},

		// 5.25
		{Scale: scale.Ryphitonic, Name: "Ryphitonic", Number: 2202, Perfection: 1, Imperfection: 4},
		{Scale: scale.Gylitonic, Name: "Gylitonic", Number: 2472, Perfection: 1, Imperfection: 4},
		{Scale: scale.Aeolycritonic, Name: "Aeolycritonic", Number: 3396, Perfection: 1, Imperfection: 4},
		{Scale: scale.Pynitonic, Name: "Pynitonic", Number: 2697, Perfection: 1, Imperfection: 4},
		{Scale: scale.Zanitonic, Name: "Zanitonic", Number: 2598, Perfection: 1, Imperfection: 4},

		// 5.26
		{Scale: scale.Phronitonic, Name: "Phronitonic", Number: 2214, Perfection: 1, Imperfection: 4},
		{Scale: scale.Banitonic, Name: "Banitonic", Number: 2664, Perfection: 1, Imperfection: 4},
		{Scale: scale.Aeronitonic, Name: "Aeronitonic", Number: 2466, Perfection: 1, Imperfection: 4},
		{Scale: scale.Golitonic, Name: "Golitonic", Number: 3348, Perfection: 1, Imperfection: 4},
		{Scale: scale.Dyptitonic, Name: "Dyptitonic", Number: 2601, Perfection: 1, Imperfection: 4},

		// 5.27
		{Scale: scale.Aerynitonic, Name: "Aerynitonic", Number: 2220, Perfection: 1, Imperfection: 4},
		{Scale: scale.Palitonic, Name: "Palitonic", Number: 2760, Perfection: 1, Imperfection: 4},
		{Scale: scale.Stothitonic, Name: "Stothitonic", Number: 2850, Perfection: 1, Imperfection: 4},
		{Scale: scale.Aerophitonic, Name: "Aerophitonic", Number: 3210, Perfection: 1, Imperfection: 4},
		{Scale: scale.Katagitonic, Name: "Katagitonic", Number: 2325, Perfection: 1, Imperfection: 4},

		// 5.28
		{Scale: scale.Ionoditonic, Name: "Ionoditonic", Number: 2226, Perfection: 1, Imperfection: 4},
		{Scale: scale.Bogitonic, Name: "Bogitonic", Number: 2856, Perfection: 1, Imperfection: 4},
		{Scale: scale.Mogitonic, Name: "Mogitonic", Number: 3234, Perfection: 1, Imperfection: 4},
		{Scale: scale.Docritonic, Name: "Docritonic", Number: 2373, Perfection: 1, Imperfection: 4},
		{Scale: scale.Epaditonic, Name: "Epaditonic", Number: 2604, Perfection: 1, Imperfection: 4},

		// 5.29
		{Scale: scale.Mixitonic, Name: "Mixitonic", Number: 2341, Perfection: 1, Imperfection: 4},
		{Scale: scale.Phrothitonic, Name: "Phrothitonic", Number: 2348, Perfection: 1, Imperfection: 4},
		{Scale: scale.Katycritonic, Name: "Katycritonic", Number: 2404, Perfection: 1, Imperfection: 4},
		{Scale: scale.Ionalitonic, Name: "Ionalitonic", Number: 2852, Perfection: 1, Imperfection: 4},
		{Scale: scale.Loptitonic, Name: "Loptitonic", Number: 3218, Perfection: 1, Imperfection: 4},

		// 5.30
		{Scale: scale.Thyritonic, Name: "Thyritonic", Number: 2342, Perfection: 1, Imperfection: 4},
		{Scale: scale.Thoptitonic, Name: "Thoptitonic", Number: 2356, Perfection: 1, Imperfection: 4},
		{Scale: scale.Bycritonic, Name: "Bycritonic", Number: 2468, Perfection: 1, Imperfection: 4},
		{Scale: scale.Pathitonic, Name: "Pathitonic", Number: 3364, Perfection: 1, Imperfection: 4},
		{Scale: scale.Myditonic, Name: "Myditonic", Number: 2633, Perfection: 1, Imperfection: 4},

		// 5.31
		{Scale: scale.Bolitonic, Name: "Bolitonic", Number: 2218, Perfection: 0, Imperfection: 5},
		{Scale: scale.Bothitonic, Name: "Bothitonic", Number: 2728, Perfection: 0, Imperfection: 5},
		{Scale: scale.Kataditonic, Name: "Kataditonic", Number: 2722, Perfection: 0, Imperfection: 5},
		{Scale: scale.Koditonic, Name: "Koditonic", Number: 2698, Perfection: 0, Imperfection: 5},
		{Scale: scale.Tholitonic, Name: "Tholitonic", Number: 2602, Perfection: 0, Imperfection: 5},

		// 6.1
		{Scale: scale.Epathimic, Name: "Epathimic", Number: 2394, Perfection: 5, Imperfection: 1},
		{Scale: scale.Mynimic, Name: "Mynimic", Number: 2772, Perfection: 5, Imperfection: 1},
		{Scale: scale.Rocrimic, Name: "Rocrimic", Number: 2898, Perfection: 5, Imperfection: 1},
		{Scale: scale.Eporimic, Name: "Eporimic", Number: 3402, Perfection: 5, Imperfection: 1},
		{Scale: scale.Thaptimic, Name: "Thaptimic", Number: 2709, Perfection: 5, Imperfection: 1},
		{Scale: scale.Lothimic, Name: "Lothimic", Number: 2646, Perfection: 5, Imperfection: 1},

		// 6.2
		{Scale: scale.Dyrimic, Name: "Dyrimic", Number: 2229, Perfection: 4, Imperfection: 2},
		{Scale: scale.Koptimic, Name: "Koptimic", Number: 2904, Perfection: 4, Imperfection: 2},
		{Scale: scale.Thocrimic, Name: "Thocrimic", Number: 3426, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolanimic, Name: "Aeolanimic", Number: 2757, Perfection: 4, Imperfection: 2},
		{Scale: scale.Danimic, Name: "Danimic", Number: 2838, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionarimic, Name: "Ionarimic", Number: 3162, Perfection: 4, Imperfection: 2},

		// 6.3
		{Scale: scale.Daptimic, Name: "Daptimic", Number: 2247, Perfection: 4, Imperfection: 2},
		{Scale: scale.Kygimic, Name: "Kygimic", Number: 3192, Perfection: 4, Imperfection: 2},
		{Scale: scale.Mocrimic, Name: "Mocrimic", Number: 2289, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zynimic, Name: "Zynimic", Number: 3864, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolimic, Name: "Aeolimic", Number: 3633, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zythimic, Name: "Zythimic", Number: 3171, Perfection: 4, Imperfection: 2},

		// 6.4
		{Scale: scale.Epygimic, Name: "Epygimic", Number: 2259, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zaptimic, Name: "Zaptimic", Number: 3384, Perfection: 4, Imperfection: 2},
		{Scale: scale.Kagimic, Name: "Kagimic", Number: 2673, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zogimic, Name: "Zogimic", Number: 2502, Perfection: 4, Imperfection: 2},
		{Scale: scale.Epyrimic, Name: "Epyrimic", Number: 3636, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lycrimic, Name: "Lycrimic", Number: 3177, Perfection: 4, Imperfection: 2},

		// 6.5
		{Scale: scale.Bylimic, Name: "Bylimic", Number: 2261, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zothimic, Name: "Zothimic", Number: 3416, Perfection: 4, Imperfection: 2},
		{Scale: scale.Phrolimic, Name: "Phrolimic", Number: 2737, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionagimic, Name: "Ionagimic", Number: 2758, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolaphimic, Name: "Aeolaphimic", Number: 2842, Perfection: 4, Imperfection: 2},
		{Scale: scale.Kycrimic, Name: "Kycrimic", Number: 3178, Perfection: 4, Imperfection: 2},

		// 6.6
		{Scale: scale.Garimic, Name: "Garimic", Number: 2262, Perfection: 4, Imperfection: 2},
		{Scale: scale.Soptimic, Name: "Soptimic", Number: 3432, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionyptimic, Name: "Ionyptimic", Number: 2769, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gyrimic, Name: "Gyrimic", Number: 2886, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zalimic, Name: "Zalimic", Number: 3354, Perfection: 4, Imperfection: 2},
		{Scale: scale.Stolimic, Name: "Stolimic", Number: 2613, Perfection: 4, Imperfection: 2},

		// 6.7
		{Scale: scale.Thonimic, Name: "Thonimic", Number: 2275, Perfection: 4, Imperfection: 2},
		{Scale: scale.Stadimic, Name: "Stadimic", Number: 3640, Perfection: 4, Imperfection: 2},
		{Scale: scale.Thodimic, Name: "Thodimic", Number: 3185, Perfection: 4, Imperfection: 2},

		// 6.8
		{Scale: scale.Mythimic, Name: "Mythimic", Number: 2277, Perfection: 4, Imperfection: 2},
		{Scale: scale.Sogimic, Name: "Sogimic", Number: 3672, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gogimic, Name: "Gogimic", Number: 3249, Perfection: 4, Imperfection: 2},
		{Scale: scale.Rothimic, Name: "Rothimic", Number: 2403, Perfection: 4, Imperfection: 2},
		{Scale: scale.Katarimic, Name: "Katarimic", Number: 2844, Perfection: 4, Imperfection: 2},
		{Scale: scale.Sylimic, Name: "Sylimic", Number: 3186, Perfection: 4, Imperfection: 2},

		// 6.9
		{Scale: scale.Mixolimic, Name: "Mixolimic", Number: 2379, Perfection: 4, Imperfection: 2},
		{Scale: scale.Dadimic, Name: "Dadimic", Number: 2652, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolyphimic, Name: "Aeolyphimic", Number: 2418, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gycrimic, Name: "Gycrimic", Number: 2964, Perfection: 4, Imperfection: 2},
		{Scale: scale.Pyrimic, Name: "Pyrimic", Number: 3666, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lydimic, Name: "Lydimic", Number: 3237, Perfection: 4, Imperfection: 2},

		// 6.10
		{Scale: scale.Ionacrimic, Name: "Ionacrimic", Number: 2382, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gathimic, Name: "Gathimic", Number: 2676, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionynimic, Name: "Ionynimic", Number: 2514, Perfection: 4, Imperfection: 2},
		{Scale: scale.Phrynimic, Name: "Phrynimic", Number: 3732, Perfection: 4, Imperfection: 2},
		{Scale: scale.Stathimic, Name: "Stathimic", Number: 3369, Perfection: 4, Imperfection: 2},
		{Scale: scale.Thatimic, Name: "Thatimic", Number: 2643, Perfection: 4, Imperfection: 2},

		// 6.11
		{Scale: scale.Dalimic, Name: "Dalimic", Number: 2387, Perfection: 4, Imperfection: 2},
		{Scale: scale.Dygimic, Name: "Dygimic", Number: 2716, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zolimic, Name: "Zolimic", Number: 2674, Perfection: 4, Imperfection: 2},
		{Scale: scale.Epogimic, Name: "Epogimic", Number: 2506, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lanimic, Name: "Lanimic", Number: 3668, Perfection: 4, Imperfection: 2},
		{Scale: scale.Paptimic, Name: "Paptimic", Number: 3241, Perfection: 4, Imperfection: 2},

		// 6.12
		{Scale: scale.Darmic, Name: "Darmic", Number: 2390, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lonimic, Name: "Lonimic", Number: 2740, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionycrimic, Name: "Ionycrimic", Number: 2770, Perfection: 4, Imperfection: 2},
		{Scale: scale.Phradimic, Name: "Phradimic", Number: 2890, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolorimic, Name: "Aeolorimic", Number: 3370, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gonimic, Name: "Gonimic", Number: 2645, Perfection: 4, Imperfection: 2},

		// 6.13
		{Scale: scale.Phracrimic, Name: "Phracrimic", Number: 2410, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aerathimic, Name: "Aerathimic", Number: 2900, Perfection: 4, Imperfection: 2},
		{Scale: scale.Sarimic, Name: "Sarimic", Number: 3410, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zoptimic, Name: "Zoptimic", Number: 2725, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zeracrimic, Name: "Zeracrimic", Number: 2710, Perfection: 4, Imperfection: 2},
		{Scale: scale.Byptimic, Name: "Byptimic", Number: 2650, Perfection: 4, Imperfection: 2},

		// 6.14
		{Scale: scale.Starimic, Name: "Starimic", Number: 2199, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrathimic, Name: "Phrathimic", Number: 2424, Perfection: 3, Imperfection: 3},
		{Scale: scale.Saptimic, Name: "Saptimic", Number: 3012, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerodimic, Name: "Aerodimic", Number: 3858, Perfection: 3, Imperfection: 3},
		{Scale: scale.Macrimic, Name: "Macrimic", Number: 3621, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rogimic, Name: "Rogimic", Number: 3147, Perfection: 3, Imperfection: 3},

		// 6.15
		{Scale: scale.Bygimic, Name: "Bygimic", Number: 2205, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thycrimic, Name: "Thycrimic", Number: 2520, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeoladimic, Name: "Aeoladimic", Number: 3780, Perfection: 3, Imperfection: 3},
		{Scale: scale.Dylimic, Name: "Dylimic", Number: 3465, Perfection: 3, Imperfection: 3},
		{Scale: scale.Eponimic, Name: "Eponimic", Number: 2835, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katygimic, Name: "Katygimic", Number: 3150, Perfection: 3, Imperfection: 3},

		// 6.16
		{Scale: scale.Stalimic, Name: "Stalimic", Number: 2215, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stoptimic, Name: "Stoptimic", Number: 2680, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zygimic, Name: "Zygimic", Number: 2530, Perfection: 3, Imperfection: 3},
		{Scale: scale.Kataptimic, Name: "Kataptimic", Number: 3860, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeolaptimic, Name: "Aeolaptimic", Number: 3625, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pothimic, Name: "Pothimic", Number: 3155, Perfection: 3, Imperfection: 3},

		// 6.17
		{Scale: scale.Rycrimic, Name: "Rycrimic", Number: 2221, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ronimic, Name: "Ronimic", Number: 2776, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stycrimic, Name: "Stycrimic", Number: 2914, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katorimic, Name: "Katorimic", Number: 3466, Perfection: 3, Imperfection: 3},
		{Scale: scale.Epythimic, Name: "Epythimic", Number: 2837, Perfection: 3, Imperfection: 3},
		{Scale: scale.Kaptimic, Name: "Kaptimic", Number: 3158, Perfection: 3, Imperfection: 3},

		// 6.18
		{Scale: scale.Katythimic, Name: "Katythimic", Number: 2227, Perfection: 3, Imperfection: 3},
		{Scale: scale.Madimic, Name: "Madimic", Number: 2872, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerygimic, Name: "Aerygimic", Number: 3298, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pylimic, Name: "Pylimic", Number: 2501, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionathimic, Name: "Ionathimic", Number: 3628, Perfection: 3, Imperfection: 3},
		{Scale: scale.Morimic, Name: "Morimic", Number: 3161, Perfection: 3, Imperfection: 3},

		// 6.19
		{Scale: scale.Aerycrimic, Name: "Aerycrimic", Number: 2233, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ganimic, Name: "Ganimic", Number: 2968, Perfection: 3, Imperfection: 3},
		{Scale: scale.Eparimic, Name: "Eparimic", Number: 3682, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lyrimic, Name: "Lyrimic", Number: 3269, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phraptimic, Name: "Phraptimic", Number: 2443, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bacrimic, Name: "Bacrimic", Number: 3164, Perfection: 3, Imperfection: 3},

		// 6.20
		{Scale: scale.Phralimic, Name: "Phralimic", Number: 2251, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrogimic, Name: "Phrogimic", Number: 3256, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rathimic, Name: "Rathimic", Number: 2417, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katocrimic, Name: "Katocrimic", Number: 2956, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phryptimic, Name: "Phryptimic", Number: 3634, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katynimic, Name: "Katynimic", Number: 3173, Perfection: 3, Imperfection: 3},

		// 6.21
		{Scale: scale.Solimic, Name: "Solimic", Number: 2253, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionolimic, Name: "Ionolimic", Number: 3288, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionophimic, Name: "Ionophimic", Number: 2481, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeologimic, Name: "Aeologimic", Number: 3468, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zadimic, Name: "Zadimic", Number: 2841, Perfection: 3, Imperfection: 3},
		{Scale: scale.Sygimic, Name: "Sygimic", Number: 3174, Perfection: 3, Imperfection: 3},

		// 6.22
		{Scale: scale.Thogimic, Name: "Thogimic", Number: 2254, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rythimic, Name: "Rythimic", Number: 3304, Perfection: 3, Imperfection: 3},
		{Scale: scale.Donimic, Name: "Donimic", Number: 2513, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeoloptimic, Name: "Aeoloptimic", Number: 3724, Perfection: 3, Imperfection: 3},
		{Scale: scale.Panimic, Name: "Panimic", Number: 3353, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lodimic, Name: "Lodimic", Number: 2611, Perfection: 3, Imperfection: 3},

		// 6.23
		{Scale: scale.Laptimic, Name: "Laptimic", Number: 2265, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lygimic, Name: "Lygimic", Number: 3480, Perfection: 3, Imperfection: 3},
		{Scale: scale.Logimic, Name: "Logimic", Number: 2865, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lalimic, Name: "Lalimic", Number: 3270, Perfection: 3, Imperfection: 3},
		{Scale: scale.Sothimic, Name: "Sothimic", Number: 2445, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrocrimic, Name: "Phrocrimic", Number: 3180, Perfection: 3, Imperfection: 3},

		// 6.24
		{Scale: scale.Modimic, Name: "Modimic", Number: 2266, Perfection: 3, Imperfection: 3},
		{Scale: scale.Barimic, Name: "Barimic", Number: 3496, Perfection: 3, Imperfection: 3},
		{Scale: scale.Poptimic, Name: "Poptimic", Number: 2897, Perfection: 3, Imperfection: 3},
		{Scale: scale.Sagimic, Name: "Sagimic", Number: 3398, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aelothimic, Name: "Aelothimic", Number: 2701, Perfection: 3, Imperfection: 3},
		{Scale: scale.Socrimic, Name: "Socrimic", Number: 2614, Perfection: 3, Imperfection: 3},

		// 6.25
		{Scale: scale.Syrimic, Name: "Syrimic", Number: 2268, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stodimic, Name: "Stodimic", Number: 3528, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionocrimic, Name: "Ionocrimic", Number: 2961, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zycrimic, Name: "Zycrimic", Number: 3654, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionygimic, Name: "Ionygimic", Number: 3213, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katathimic, Name: "Katathimic", Number: 2331, Perfection: 3, Imperfection: 3},

		// 6.26
		{Scale: scale.Bolimic, Name: "Bolimic", Number: 2278, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bothimic, Name: "Bothimic", Number: 3688, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katadimic, Name: "Katadimic", Number: 3281, Perfection: 3, Imperfection: 3},
		{Scale: scale.Kodimic, Name: "Kodimic", Number: 2467, Perfection: 3, Imperfection: 3},
		{Scale: scale.Tholimic, Name: "Tholimic", Number: 3356, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ralimic, Name: "Ralimic", Number: 2617, Perfection: 3, Imperfection: 3},

		// 6.27
		{Scale: scale.Kanimic, Name: "Kanimic", Number: 2281, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zylimic, Name: "Zylimic", Number: 3736, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zodimic, Name: "Zodimic", Number: 3377, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zarimic, Name: "Zarimic", Number: 2659, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrythimic, Name: "Phrythimic", Number: 2446, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rorimic, Name: "Rorimic", Number: 3188, Perfection: 3, Imperfection: 3},

		// 6.28
		{Scale: scale.Pynimic, Name: "Pynimic", Number: 2290, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zanimic, Name: "Zanimic", Number: 3880, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ranimic, Name: "Ranimic", Number: 3665, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ladimic, Name: "Ladimic", Number: 3235, Perfection: 3, Imperfection: 3},
		{Scale: scale.Podimic, Name: "Podimic", Number: 2375, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionothimic, Name: "Ionothimic", Number: 2620, Perfection: 3, Imperfection: 3},

		// 6.29
		{Scale: scale.Kytrimic, Name: "Kytrimic", Number: 2292, Perfection: 3, Imperfection: 3},
		{Scale: scale.Golimic, Name: "Golimic", Number: 3912, Perfection: 3, Imperfection: 3},
		{Scale: scale.Dyptimic, Name: "Dyptimic", Number: 3729, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ryrimic, Name: "Ryrimic", Number: 3363, Perfection: 3, Imperfection: 3},
		{Scale: scale.Gylimic, Name: "Gylimic", Number: 2631, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeolycrimic, Name: "Aeolycrimic", Number: 2334, Perfection: 3, Imperfection: 3},

		// 6.30
		{Scale: scale.Palimic, Name: "Palimic", Number: 2347, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stothimic, Name: "Stothimic", Number: 2396, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeronimic, Name: "Aeronimic", Number: 2788, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katagimic, Name: "Katagimic", Number: 2962, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phronimic, Name: "Phronimic", Number: 3658, Perfection: 3, Imperfection: 3},
		{Scale: scale.Banimic, Name: "Banimic", Number: 3221, Perfection: 3, Imperfection: 3},

		// 6.31
		{Scale: scale.Ionodimic, Name: "Ionodimic", Number: 2355, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bogimic, Name: "Bogimic", Number: 2460, Perfection: 3, Imperfection: 3},
		{Scale: scale.Mogimic, Name: "Mogimic", Number: 3300, Perfection: 3, Imperfection: 3},
		{Scale: scale.Docrimic, Name: "Docrimic", Number: 2505, Perfection: 3, Imperfection: 3},
		{Scale: scale.Epadimic, Name: "Epadimic", Number: 3660, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerynimic, Name: "Aerynimic", Number: 3225, Perfection: 3, Imperfection: 3},

		// 6.32
		{Scale: scale.Mydimic, Name: "Mydimic", Number: 2361, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thyptimic, Name: "Thyptimic", Number: 2508, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrothimic, Name: "Phrothimic", Number: 3684, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katycrimic, Name: "Katycrimic", Number: 3273, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionalimic, Name: "Ionalimic", Number: 2451, Perfection: 3, Imperfection: 3},
		{Scale: scale.Loptimic, Name: "Loptimic", Number: 3228, Perfection: 3, Imperfection: 3},

		// 6.33
		{Scale: scale.Zagimic, Name: "Zagimic", Number: 2362, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lagimic, Name: "Lagimic", Number: 2516, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thyrimic, Name: "Thyrimic", Number: 3748, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thothimic, Name: "Thothimic", Number: 3401, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bycrimic, Name: "Bycrimic", Number: 2707, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pathimic, Name: "Pathimic", Number: 2638, Perfection: 3, Imperfection: 3},

		// 6.34
		{Scale: scale.Mothimic, Name: "Mothimic", Number: 2393, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeranimic, Name: "Aeranimic", Number: 2764, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ragimic, Name: "Ragimic", Number: 2866, Perfection: 3, Imperfection: 3},
		{Scale: scale.Dolimic, Name: "Dolimic", Number: 3274, Perfection: 3, Imperfection: 3},
		{Scale: scale.Porimic, Name: "Porimic", Number: 2453, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerylimic, Name: "Aerylimic", Number: 3244, Perfection: 3, Imperfection: 3},

		// 6.35
		{Scale: scale.Bocrimic, Name: "Bocrimic", Number: 2406, Perfection: 3, Imperfection: 3},
		{Scale: scale.Gythimic, Name: "Gythimic", Number: 2868, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pagimic, Name: "Pagimic", Number: 3282, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeolythimic, Name: "Aeolythimic", Number: 2469, Perfection: 3, Imperfection: 3},
		{Scale: scale.Molimic, Name: "Molimic", Number: 3372, Perfection: 3, Imperfection: 3},
		{Scale: scale.Staptimic, Name: "Staptimic", Number: 2649, Perfection: 3, Imperfection: 3},

		// 6.36
		{Scale: scale.Zacrimic, Name: "Zacrimic", Number: 2409, Perfection: 3, Imperfection: 3},
		{Scale: scale.Larimic, Name: "Larimic", Number: 2892, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thacrimic, Name: "Thacrimic", Number: 3378, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stydimic, Name: "Stydimic", Number: 2661, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lorimic, Name: "Lorimic", Number: 2454, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionadimic, Name: "Ionadimic", Number: 3252, Perfection: 3, Imperfection: 3},

		// 6.37
		{Scale: scale.Ionythimic, Name: "Ionythimic", Number: 2457, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerythimic, Name: "Aerythimic", Number: 3276, Perfection: 3, Imperfection: 3},

		// 6.38
		{Scale: scale.Dynimic, Name: "Dynimic", Number: 2458, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zydimic, Name: "Zydimic", Number: 3284, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zathimic, Name: "Zathimic", Number: 2473, Perfection: 3, Imperfection: 3},
		{Scale: scale.Radimic, Name: "Radimic", Number: 3404, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stonimic, Name: "Stonimic", Number: 2713, Perfection: 3, Imperfection: 3},
		{Scale: scale.Syptimic, Name: "Syptimic", Number: 2662, Perfection: 3, Imperfection: 3},

		// 6.39
		{Scale: scale.Ponimic, Name: "Ponimic", Number: 2191, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kadimic, Name: "Kadimic", Number: 2296, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gynimic, Name: "Gynimic", Number: 3976, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thydimic, Name: "Thydimic", Number: 3857, Perfection: 2, Imperfection: 4},
		{Scale: scale.Polimic, Name: "Polimic", Number: 3619, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thanimic, Name: "Thanimic", Number: 3143, Perfection: 2, Imperfection: 4},

		// 6.40
		{Scale: scale.Lathimic, Name: "Lathimic", Number: 2203, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeralimic, Name: "Aeralimic", Number: 2488, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kynimic, Name: "Kynimic", Number: 3524, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stynimic, Name: "Stynimic", Number: 2953, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epytimic, Name: "Epytimic", Number: 3622, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katoptimic, Name: "Katoptimic", Number: 3149, Perfection: 2, Imperfection: 4},

		// 6.41
		{Scale: scale.Galimic, Name: "Galimic", Number: 2206, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kathimic, Name: "Kathimic", Number: 2536, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lylimic, Name: "Lylimic", Number: 3908, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epalimic, Name: "Epalimic", Number: 3721, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epacrimic, Name: "Epacrimic", Number: 3347, Perfection: 2, Imperfection: 4},
		{Scale: scale.Sathimic, Name: "Sathimic", Number: 2599, Perfection: 2, Imperfection: 4},

		// 6.42
		{Scale: scale.Katanimic, Name: "Katanimic", Number: 2219, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katyrimic, Name: "Katyrimic", Number: 2744, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rynimic, Name: "Rynimic", Number: 2786, Perfection: 2, Imperfection: 4},
		{Scale: scale.Pogimic, Name: "Pogimic", Number: 2954, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeraptimic, Name: "Aeraptimic", Number: 3626, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epylimic, Name: "Epylimic", Number: 3157, Perfection: 2, Imperfection: 4},

		// 6.43
		{Scale: scale.Manimic, Name: "Manimic", Number: 2230, Perfection: 2, Imperfection: 4},
		{Scale: scale.Marimic, Name: "Marimic", Number: 2920, Perfection: 2, Imperfection: 4},
		{Scale: scale.Locrimic, Name: "Locrimic", Number: 3490, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rylimic, Name: "Rylimic", Number: 2885, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epatimic, Name: "Epatimic", Number: 3350, Perfection: 2, Imperfection: 4},
		{Scale: scale.Byrimic, Name: "Byrimic", Number: 2605, Perfection: 2, Imperfection: 4},

		// 6.44
		{Scale: scale.Kocrimic, Name: "Kocrimic", Number: 2236, Perfection: 2, Imperfection: 4},
		{Scale: scale.Korimic, Name: "Korimic", Number: 3016, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lynimic, Name: "Lynimic", Number: 3874, Perfection: 2, Imperfection: 4},
		{Scale: scale.Malimic, Name: "Malimic", Number: 3653, Perfection: 2, Imperfection: 4},
		{Scale: scale.Synimic, Name: "Synimic", Number: 3211, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phragimic, Name: "Phragimic", Number: 2327, Perfection: 2, Imperfection: 4},

		// 6.45
		{Scale: scale.Mycrimic, Name: "Mycrimic", Number: 2282, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionorimic, Name: "Ionorimic", Number: 3752, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phrydimic, Name: "Phrydimic", Number: 3409, Perfection: 2, Imperfection: 4},
		{Scale: scale.Zyptimic, Name: "Zyptimic", Number: 2723, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katothimic, Name: "Katothimic", Number: 2702, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phrylimic, Name: "Phrylimic", Number: 2618, Perfection: 2, Imperfection: 4},

		// 6.46
		{Scale: scale.Aerothimic, Name: "Aerothimic", Number: 2284, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stagimic, Name: "Stagimic", Number: 3784, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dorimic, Name: "Dorimic", Number: 3473, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phrycrimic, Name: "Phrycrimic", Number: 2851, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kyptimic, Name: "Kyptimic", Number: 3214, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionylimic, Name: "Ionylimic", Number: 2333, Perfection: 2, Imperfection: 4},

		// 6.47
		{Scale: scale.Epynimic, Name: "Epynimic", Number: 2343, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionogimic, Name: "Ionogimic", Number: 2364, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kydimic, Name: "Kydimic", Number: 2532, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gaptimic, Name: "Gaptimic", Number: 3876, Perfection: 2, Imperfection: 4},
		{Scale: scale.Tharimic, Name: "Tharimic", Number: 3657, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionaphimic, Name: "Ionaphimic", Number: 3219, Perfection: 2, Imperfection: 4},

		// 6.48
		{Scale: scale.Thoptimic, Name: "Thoptimic", Number: 2349, Perfection: 2, Imperfection: 4},
		{Scale: scale.Bagimic, Name: "Bagimic", Number: 2412, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kyrimic, Name: "Kyrimic", Number: 2916, Perfection: 2, Imperfection: 4},
		{Scale: scale.Sonimic, Name: "Sonimic", Number: 3474, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolonimic, Name: "Aeolonimic", Number: 2853, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rygimic, Name: "Rygimic", Number: 3222, Perfection: 2, Imperfection: 4},

		// 6.49
		{Scale: scale.Thagimic, Name: "Thagimic", Number: 2350, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kolimic, Name: "Kolimic", Number: 2420, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dycrimic, Name: "Dycrimic", Number: 2980, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epycrimic, Name: "Epycrimic", Number: 3730, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gocrimic, Name: "Gocrimic", Number: 3365, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katolimic, Name: "Katolimic", Number: 2635, Perfection: 2, Imperfection: 4},

		// 6.50
		{Scale: scale.Dagimic, Name: "Dagimic", Number: 2357, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolydimic, Name: "Aeolydimic", Number: 2476, Perfection: 2, Imperfection: 4},
		{Scale: scale.Parimic, Name: "Parimic", Number: 3428, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionaptimic, Name: "Ionaptimic", Number: 2761, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thylimic, Name: "Thylimic", Number: 2854, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lolimic, Name: "Lolimic", Number: 3226, Perfection: 2, Imperfection: 4},

		// 6.51
		{Scale: scale.Thalimic, Name: "Thalimic", Number: 2358, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stygimic, Name: "Stygimic", Number: 2484, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolygimic, Name: "Aeolygimic", Number: 3492, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aerogimic, Name: "Aerogimic", Number: 2889, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dacrimic, Name: "Dacrimic", Number: 3366, Perfection: 2, Imperfection: 4},
		{Scale: scale.Baptimic, Name: "Baptimic", Number: 2637, Perfection: 2, Imperfection: 4},

		// 6.52
		{Scale: scale.Stythimic, Name: "Stythimic", Number: 2381, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kothimic, Name: "Kothimic", Number: 2668, Perfection: 2, Imperfection: 4},
		{Scale: scale.Pygimic, Name: "Pygimic", Number: 2482, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rodimic, Name: "Rodimic", Number: 3476, Perfection: 2, Imperfection: 4},
		{Scale: scale.Sorimic, Name: "Sorimic", Number: 2857, Perfection: 2, Imperfection: 4},
		{Scale: scale.Monimic, Name: "Monimic", Number: 3238, Perfection: 2, Imperfection: 4},

		// 6.53
		{Scale: scale.Aeragimic, Name: "Aeragimic", Number: 2389, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epothimic, Name: "Epothimic", Number: 2732, Perfection: 2, Imperfection: 4},
		{Scale: scale.Salimic, Name: "Salimic", Number: 2738, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lyptimic, Name: "Lyptimic", Number: 2762, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katonimic, Name: "Katonimic", Number: 2858, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gygimic, Name: "Gygimic", Number: 3242, Perfection: 2, Imperfection: 4},

		// 6.54
		{Scale: scale.Aeradimic, Name: "Aeradimic", Number: 2405, Perfection: 2, Imperfection: 4},
		{Scale: scale.Zyrimic, Name: "Zyrimic", Number: 2860, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stylimic, Name: "Stylimic", Number: 3250, Perfection: 2, Imperfection: 4},

		// 6.55
		{Scale: scale.Lythimic, Name: "Lythimic", Number: 2470, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dodimic, Name: "Dodimic", Number: 3380, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katalimic, Name: "Katalimic", Number: 2665, Perfection: 2, Imperfection: 4},

		// 6.56
		{Scale: scale.Boptimic, Name: "Boptimic", Number: 2474, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stogimic, Name: "Stogimic", Number: 3412, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thynimic, Name: "Thynimic", Number: 2729, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolathimic, Name: "Aeolathimic", Number: 2726, Perfection: 2, Imperfection: 4},
		{Scale: scale.Bythimic, Name: "Bythimic", Number: 2714, Perfection: 2, Imperfection: 4},
		{Scale: scale.Padimic, Name: "Padimic", Number: 2666, Perfection: 2, Imperfection: 4},

		// 6.57
		{Scale: scale.Dathimic, Name: "Dathimic", Number: 2222, Perfection: 1, Imperfection: 5},
		{Scale: scale.Epagimic, Name: "Epagimic", Number: 2792, Perfection: 1, Imperfection: 5},
		{Scale: scale.Raptimic, Name: "Raptimic", Number: 2978, Perfection: 1, Imperfection: 5},
		{Scale: scale.Epolimic, Name: "Epolimic", Number: 3722, Perfection: 1, Imperfection: 5},
		{Scale: scale.Sythimic, Name: "Sythimic", Number: 3349, Perfection: 1, Imperfection: 5},
		{Scale: scale.Sydimic, Name: "Sydimic", Number: 2603, Perfection: 1, Imperfection: 5},

		// 6.58
		{Scale: scale.Gacrimic, Name: "Gacrimic", Number: 2234, Perfection: 1, Imperfection: 5},
		{Scale: scale.Borimic, Name: "Borimic", Number: 2984, Perfection: 1, Imperfection: 5},
		{Scale: scale.Sycrimic, Name: "Sycrimic", Number: 3746, Perfection: 1, Imperfection: 5},
		{Scale: scale.Gadimic, Name: "Gadimic", Number: 3397, Perfection: 1, Imperfection: 5},
		{Scale: scale.Aeolocrimic, Name: "Aeolocrimic", Number: 2699, Perfection: 1, Imperfection: 5},
		{Scale: scale.Phrygimic, Name: "Phrygimic", Number: 2606, Perfection: 1, Imperfection: 5},

		// 6.59
		{Scale: scale.WholeTone, Name: "WholeTone", Number: 2730, Perfection: 0, Imperfection: 6},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Number, tc.Scale.Number())

			result := tc.Scale.Perfection()
			assert.Equal(t, tc.Perfection, result.Perfection)
			assert.Equal(t, tc.Imperfection, result.Imperfection)
		})
	}
}

func TestScale_Pitches(t *testing.T) {
	type testCase struct {
		Scale   scale.Type
		Tonic   pitch.Type
		Pitches []pitch.Type
	}

	testCases := []testCase{
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.CNatural,
			Pitches: []pitch.Type{pitch.CNatural, pitch.ENatural, pitch.GSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.CSharp,
			Pitches: []pitch.Type{pitch.CSharp, pitch.FNatural, pitch.ANatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.DNatural,
			Pitches: []pitch.Type{pitch.DNatural, pitch.FSharp, pitch.ASharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.DSharp,
			Pitches: []pitch.Type{pitch.DSharp, pitch.GNatural, pitch.BNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.ENatural,
			Pitches: []pitch.Type{pitch.ENatural, pitch.GSharp, pitch.CNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.FNatural,
			Pitches: []pitch.Type{pitch.FNatural, pitch.ANatural, pitch.CSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.FSharp,
			Pitches: []pitch.Type{pitch.FSharp, pitch.ASharp, pitch.DNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.GNatural,
			Pitches: []pitch.Type{pitch.GNatural, pitch.BNatural, pitch.DSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.GSharp,
			Pitches: []pitch.Type{pitch.GSharp, pitch.CNatural, pitch.ENatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.ANatural,
			Pitches: []pitch.Type{pitch.ANatural, pitch.CSharp, pitch.FNatural},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.ASharp,
			Pitches: []pitch.Type{pitch.ASharp, pitch.DNatural, pitch.FSharp},
		},
		{
			Scale:   scale.Minoric,
			Tonic:   pitch.BNatural,
			Pitches: []pitch.Type{pitch.BNatural, pitch.DSharp, pitch.GNatural},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Scale%sWithTonic%s", tc.Scale, tc.Tonic), func(t *testing.T) {
			assert.Equal(t, tc.Pitches, tc.Scale.Pitches(tc.Tonic))
		})
	}
}

func TestAllScales(t *testing.T) {
	assert.NotEmpty(t, scale.AllScales())
}
