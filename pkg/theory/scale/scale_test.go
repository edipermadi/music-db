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
