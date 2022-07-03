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
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		{Scale: scale.Invalid, Name: "Invalid"},

		// 3 notes
		{Scale: scale.Minoric, Name: "Minoric", Cardinality: 3, Number: 2184, Imperfection: 3},

		// 4 notes

		// 4.1
		{Scale: scale.Thaptic, Name: "Thaptic", Cardinality: 4, Number: 2193, Perfection: 2, Imperfection: 2},
		{Scale: scale.Lothic, Name: "Lothic", Cardinality: 4, Number: 2328, Perfection: 2, Imperfection: 2},
		{Scale: scale.Phratic, Name: "Phratic", Cardinality: 4, Number: 2244, Perfection: 2, Imperfection: 2},
		{Scale: scale.Aerathic, Name: "Aerathic", Cardinality: 4, Number: 3144, Perfection: 2, Imperfection: 2},

		// 4.2
		{Scale: scale.Epathic, Name: "Epathic", Cardinality: 4, Number: 2196, Perfection: 2, Imperfection: 2},
		{Scale: scale.Mynic, Name: "Mynic", Cardinality: 4, Number: 2376, Perfection: 2, Imperfection: 2},
		{Scale: scale.Rothic, Name: "Rothic", Cardinality: 4, Number: 2628, Perfection: 2, Imperfection: 2},
		{Scale: scale.Eporic, Name: "Eporic", Cardinality: 4, Number: 2322, Perfection: 2, Imperfection: 2},

		// 4.3
		{Scale: scale.Zyphic, Name: "Zyphic", Cardinality: 4, Number: 2185, Perfection: 1, Imperfection: 3},
		{Scale: scale.Epogic, Name: "Epogic", Cardinality: 4, Number: 2200, Perfection: 1, Imperfection: 3},
		{Scale: scale.Lanic, Name: "Lanic", Cardinality: 4, Number: 2440, Perfection: 1, Imperfection: 3},
		{Scale: scale.Pyrric, Name: "Pyrric", Cardinality: 4, Number: 3140, Perfection: 1, Imperfection: 3},

		// 4.4
		{Scale: scale.Aeoloric, Name: "Aeoloric", Cardinality: 4, Number: 2188, Perfection: 1, Imperfection: 3},
		{Scale: scale.Gonic, Name: "Gonic", Cardinality: 4, Number: 2248, Perfection: 1, Imperfection: 3},
		{Scale: scale.Dalic, Name: "Dalic", Cardinality: 4, Number: 3208, Perfection: 1, Imperfection: 3},
		{Scale: scale.Dygic, Name: "Dygic", Cardinality: 4, Number: 2321, Perfection: 1, Imperfection: 3},

		// 4.5
		{Scale: scale.Daric, Name: "Daric", Cardinality: 4, Number: 2194, Perfection: 1, Imperfection: 3},
		{Scale: scale.Lonic, Name: "Lonic", Cardinality: 4, Number: 2344, Perfection: 1, Imperfection: 3},
		{Scale: scale.Phradic, Name: "Phradic", Cardinality: 4, Number: 2372, Perfection: 1, Imperfection: 3},
		{Scale: scale.Bolic, Name: "Bolic", Cardinality: 4, Number: 2596, Perfection: 1, Imperfection: 3},

		// 4.6
		{Scale: scale.Saric, Name: "Saric", Cardinality: 4, Number: 2212, Perfection: 1, Imperfection: 3},
		{Scale: scale.Zoptic, Name: "Zoptic", Cardinality: 4, Number: 2632, Perfection: 1, Imperfection: 3},
		{Scale: scale.Aeraphic, Name: "Aeraphic", Cardinality: 4, Number: 2338, Perfection: 1, Imperfection: 3},
		{Scale: scale.Byptic, Name: "Byptic", Cardinality: 4, Number: 2324, Perfection: 1, Imperfection: 3},

		// 4.7
		{Scale: scale.Aeolic, Name: "Aeolic", Cardinality: 4, Number: 2186, Perfection: 0, Imperfection: 4},
		{Scale: scale.Koptic, Name: "Koptic", Cardinality: 4, Number: 2216, Perfection: 0, Imperfection: 4},
		{Scale: scale.Mixolyric, Name: "Mixolyric", Cardinality: 4, Number: 2696, Perfection: 0, Imperfection: 4},
		{Scale: scale.Lydic, Name: "Lydic", Cardinality: 4, Number: 2594, Perfection: 0, Imperfection: 4},

		// 4.8
		{Scale: scale.Stathic, Name: "Stathic", Cardinality: 4, Number: 2210, Perfection: 0, Imperfection: 4},
		{Scale: scale.Dadic, Name: "Dadic", Cardinality: 4, Number: 2600, Perfection: 0, Imperfection: 4},

		// 4.9
		{Scale: scale.Phrynic, Name: "Phrynic", Cardinality: 4, Number: 2340, Perfection: 0, Imperfection: 4},

		// 5 notes

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

		// 6.1
		{Scale: scale.Epathimic, Name: "Epathimic", Cardinality: 6, Number: 2394, Perfection: 5, Imperfection: 1},
		{Scale: scale.Mynimic, Name: "Mynimic", Cardinality: 6, Number: 2772, Perfection: 5, Imperfection: 1},
		{Scale: scale.Rocrimic, Name: "Rocrimic", Cardinality: 6, Number: 2898, Perfection: 5, Imperfection: 1},
		{Scale: scale.Eporimic, Name: "Eporimic", Cardinality: 6, Number: 3402, Perfection: 5, Imperfection: 1},
		{Scale: scale.Thaptimic, Name: "Thaptimic", Cardinality: 6, Number: 2709, Perfection: 5, Imperfection: 1},
		{Scale: scale.Lothimic, Name: "Lothimic", Cardinality: 6, Number: 2646, Perfection: 5, Imperfection: 1},

		// 6.2
		{Scale: scale.Dyrimic, Name: "Dyrimic", Cardinality: 6, Number: 2229, Perfection: 4, Imperfection: 2},
		{Scale: scale.Koptimic, Name: "Koptimic", Cardinality: 6, Number: 2904, Perfection: 4, Imperfection: 2},
		{Scale: scale.Thocrimic, Name: "Thocrimic", Cardinality: 6, Number: 3426, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolanimic, Name: "Aeolanimic", Cardinality: 6, Number: 2757, Perfection: 4, Imperfection: 2},
		{Scale: scale.Danimic, Name: "Danimic", Cardinality: 6, Number: 2838, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionarimic, Name: "Ionarimic", Cardinality: 6, Number: 3162, Perfection: 4, Imperfection: 2},

		// 6.3
		{Scale: scale.Daptimic, Name: "Daptimic", Cardinality: 6, Number: 2247, Perfection: 4, Imperfection: 2},
		{Scale: scale.Kygimic, Name: "Kygimic", Cardinality: 6, Number: 3192, Perfection: 4, Imperfection: 2},
		{Scale: scale.Mocrimic, Name: "Mocrimic", Cardinality: 6, Number: 2289, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zynimic, Name: "Zynimic", Cardinality: 6, Number: 3864, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolimic, Name: "Aeolimic", Cardinality: 6, Number: 3633, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zythimic, Name: "Zythimic", Cardinality: 6, Number: 3171, Perfection: 4, Imperfection: 2},

		// 6.4
		{Scale: scale.Epygimic, Name: "Epygimic", Cardinality: 6, Number: 2259, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zaptimic, Name: "Zaptimic", Cardinality: 6, Number: 3384, Perfection: 4, Imperfection: 2},
		{Scale: scale.Kagimic, Name: "Kagimic", Cardinality: 6, Number: 2673, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zogimic, Name: "Zogimic", Cardinality: 6, Number: 2502, Perfection: 4, Imperfection: 2},
		{Scale: scale.Epyrimic, Name: "Epyrimic", Cardinality: 6, Number: 3636, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lycrimic, Name: "Lycrimic", Cardinality: 6, Number: 3177, Perfection: 4, Imperfection: 2},

		// 6.5
		{Scale: scale.Bylimic, Name: "Bylimic", Cardinality: 6, Number: 2261, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zothimic, Name: "Zothimic", Cardinality: 6, Number: 3416, Perfection: 4, Imperfection: 2},
		{Scale: scale.Phrolimic, Name: "Phrolimic", Cardinality: 6, Number: 2737, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionagimic, Name: "Ionagimic", Cardinality: 6, Number: 2758, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolaphimic, Name: "Aeolaphimic", Cardinality: 6, Number: 2842, Perfection: 4, Imperfection: 2},
		{Scale: scale.Kycrimic, Name: "Kycrimic", Cardinality: 6, Number: 3178, Perfection: 4, Imperfection: 2},

		// 6.6
		{Scale: scale.Garimic, Name: "Garimic", Cardinality: 6, Number: 2262, Perfection: 4, Imperfection: 2},
		{Scale: scale.Soptimic, Name: "Soptimic", Cardinality: 6, Number: 3432, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionyptimic, Name: "Ionyptimic", Cardinality: 6, Number: 2769, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gyrimic, Name: "Gyrimic", Cardinality: 6, Number: 2886, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zalimic, Name: "Zalimic", Cardinality: 6, Number: 3354, Perfection: 4, Imperfection: 2},
		{Scale: scale.Stolimic, Name: "Stolimic", Cardinality: 6, Number: 2613, Perfection: 4, Imperfection: 2},

		// 6.7
		{Scale: scale.Thonimic, Name: "Thonimic", Cardinality: 6, Number: 2275, Perfection: 4, Imperfection: 2},
		{Scale: scale.Stadimic, Name: "Stadimic", Cardinality: 6, Number: 3640, Perfection: 4, Imperfection: 2},
		{Scale: scale.Thodimic, Name: "Thodimic", Cardinality: 6, Number: 3185, Perfection: 4, Imperfection: 2},

		// 6.8
		{Scale: scale.Mythimic, Name: "Mythimic", Cardinality: 6, Number: 2277, Perfection: 4, Imperfection: 2},
		{Scale: scale.Sogimic, Name: "Sogimic", Cardinality: 6, Number: 3672, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gogimic, Name: "Gogimic", Cardinality: 6, Number: 3249, Perfection: 4, Imperfection: 2},
		{Scale: scale.Rothimic, Name: "Rothimic", Cardinality: 6, Number: 2403, Perfection: 4, Imperfection: 2},
		{Scale: scale.Katarimic, Name: "Katarimic", Cardinality: 6, Number: 2844, Perfection: 4, Imperfection: 2},
		{Scale: scale.Sylimic, Name: "Sylimic", Cardinality: 6, Number: 3186, Perfection: 4, Imperfection: 2},

		// 6.9
		{Scale: scale.Mixolimic, Name: "Mixolimic", Cardinality: 6, Number: 2379, Perfection: 4, Imperfection: 2},
		{Scale: scale.Dadimic, Name: "Dadimic", Cardinality: 6, Number: 2652, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolyphimic, Name: "Aeolyphimic", Cardinality: 6, Number: 2418, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gycrimic, Name: "Gycrimic", Cardinality: 6, Number: 2964, Perfection: 4, Imperfection: 2},
		{Scale: scale.Pyrimic, Name: "Pyrimic", Cardinality: 6, Number: 3666, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lydimic, Name: "Lydimic", Cardinality: 6, Number: 3237, Perfection: 4, Imperfection: 2},

		// 6.10
		{Scale: scale.Ionacrimic, Name: "Ionacrimic", Cardinality: 6, Number: 2382, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gathimic, Name: "Gathimic", Cardinality: 6, Number: 2676, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionynimic, Name: "Ionynimic", Cardinality: 6, Number: 2514, Perfection: 4, Imperfection: 2},
		{Scale: scale.Phrynimic, Name: "Phrynimic", Cardinality: 6, Number: 3732, Perfection: 4, Imperfection: 2},
		{Scale: scale.Stathimic, Name: "Stathimic", Cardinality: 6, Number: 3369, Perfection: 4, Imperfection: 2},
		{Scale: scale.Thatimic, Name: "Thatimic", Cardinality: 6, Number: 2643, Perfection: 4, Imperfection: 2},

		// 6.11
		{Scale: scale.Dalimic, Name: "Dalimic", Cardinality: 6, Number: 2387, Perfection: 4, Imperfection: 2},
		{Scale: scale.Dygimic, Name: "Dygimic", Cardinality: 6, Number: 2716, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zolimic, Name: "Zolimic", Cardinality: 6, Number: 2674, Perfection: 4, Imperfection: 2},
		{Scale: scale.Epogimic, Name: "Epogimic", Cardinality: 6, Number: 2506, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lanimic, Name: "Lanimic", Cardinality: 6, Number: 3668, Perfection: 4, Imperfection: 2},
		{Scale: scale.Paptimic, Name: "Paptimic", Cardinality: 6, Number: 3241, Perfection: 4, Imperfection: 2},

		// 6.12
		{Scale: scale.Darmic, Name: "Darmic", Cardinality: 6, Number: 2390, Perfection: 4, Imperfection: 2},
		{Scale: scale.Lonimic, Name: "Lonimic", Cardinality: 6, Number: 2740, Perfection: 4, Imperfection: 2},
		{Scale: scale.Ionycrimic, Name: "Ionycrimic", Cardinality: 6, Number: 2770, Perfection: 4, Imperfection: 2},
		{Scale: scale.Phradimic, Name: "Phradimic", Cardinality: 6, Number: 2890, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aeolorimic, Name: "Aeolorimic", Cardinality: 6, Number: 3370, Perfection: 4, Imperfection: 2},
		{Scale: scale.Gonimic, Name: "Gonimic", Cardinality: 6, Number: 2645, Perfection: 4, Imperfection: 2},

		// 6.13
		{Scale: scale.Phracrimic, Name: "Phracrimic", Cardinality: 6, Number: 2410, Perfection: 4, Imperfection: 2},
		{Scale: scale.Aerathimic, Name: "Aerathimic", Cardinality: 6, Number: 2900, Perfection: 4, Imperfection: 2},
		{Scale: scale.Sarimic, Name: "Sarimic", Cardinality: 6, Number: 3410, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zoptimic, Name: "Zoptimic", Cardinality: 6, Number: 2725, Perfection: 4, Imperfection: 2},
		{Scale: scale.Zeracrimic, Name: "Zeracrimic", Cardinality: 6, Number: 2710, Perfection: 4, Imperfection: 2},
		{Scale: scale.Byptimic, Name: "Byptimic", Cardinality: 6, Number: 2650, Perfection: 4, Imperfection: 2},

		// 6.14
		{Scale: scale.Starimic, Name: "Starimic", Cardinality: 6, Number: 2199, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrathimic, Name: "Phrathimic", Cardinality: 6, Number: 2424, Perfection: 3, Imperfection: 3},
		{Scale: scale.Saptimic, Name: "Saptimic", Cardinality: 6, Number: 3012, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerodimic, Name: "Aerodimic", Cardinality: 6, Number: 3858, Perfection: 3, Imperfection: 3},
		{Scale: scale.Macrimic, Name: "Macrimic", Cardinality: 6, Number: 3621, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rogimic, Name: "Rogimic", Cardinality: 6, Number: 3147, Perfection: 3, Imperfection: 3},

		// 6.15
		{Scale: scale.Bygimic, Name: "Bygimic", Cardinality: 6, Number: 2205, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thycrimic, Name: "Thycrimic", Cardinality: 6, Number: 2520, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeoladimic, Name: "Aeoladimic", Cardinality: 6, Number: 3780, Perfection: 3, Imperfection: 3},
		{Scale: scale.Dylimic, Name: "Dylimic", Cardinality: 6, Number: 3465, Perfection: 3, Imperfection: 3},
		{Scale: scale.Eponimic, Name: "Eponimic", Cardinality: 6, Number: 2835, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katygimic, Name: "Katygimic", Cardinality: 6, Number: 3150, Perfection: 3, Imperfection: 3},

		// 6.16
		{Scale: scale.Stalimic, Name: "Stalimic", Cardinality: 6, Number: 2215, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stoptimic, Name: "Stoptimic", Cardinality: 6, Number: 2680, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zygimic, Name: "Zygimic", Cardinality: 6, Number: 2530, Perfection: 3, Imperfection: 3},
		{Scale: scale.Kataptimic, Name: "Kataptimic", Cardinality: 6, Number: 3860, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeolaptimic, Name: "Aeolaptimic", Cardinality: 6, Number: 3625, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pothimic, Name: "Pothimic", Cardinality: 6, Number: 3155, Perfection: 3, Imperfection: 3},

		// 6.17
		{Scale: scale.Rycrimic, Name: "Rycrimic", Cardinality: 6, Number: 2221, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ronimic, Name: "Ronimic", Cardinality: 6, Number: 2776, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stycrimic, Name: "Stycrimic", Cardinality: 6, Number: 2914, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katorimic, Name: "Katorimic", Cardinality: 6, Number: 3466, Perfection: 3, Imperfection: 3},
		{Scale: scale.Epythimic, Name: "Epythimic", Cardinality: 6, Number: 2837, Perfection: 3, Imperfection: 3},
		{Scale: scale.Kaptimic, Name: "Kaptimic", Cardinality: 6, Number: 3158, Perfection: 3, Imperfection: 3},

		// 6.18
		{Scale: scale.Katythimic, Name: "Katythimic", Cardinality: 6, Number: 2227, Perfection: 3, Imperfection: 3},
		{Scale: scale.Madimic, Name: "Madimic", Cardinality: 6, Number: 2872, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerygimic, Name: "Aerygimic", Cardinality: 6, Number: 3298, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pylimic, Name: "Pylimic", Cardinality: 6, Number: 2501, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionathimic, Name: "Ionathimic", Cardinality: 6, Number: 3628, Perfection: 3, Imperfection: 3},
		{Scale: scale.Morimic, Name: "Morimic", Cardinality: 6, Number: 3161, Perfection: 3, Imperfection: 3},

		// 6.19
		{Scale: scale.Aerycrimic, Name: "Aerycrimic", Cardinality: 6, Number: 2233, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ganimic, Name: "Ganimic", Cardinality: 6, Number: 2968, Perfection: 3, Imperfection: 3},
		{Scale: scale.Eparimic, Name: "Eparimic", Cardinality: 6, Number: 3682, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lyrimic, Name: "Lyrimic", Cardinality: 6, Number: 3269, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phraptimic, Name: "Phraptimic", Cardinality: 6, Number: 2443, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bacrimic, Name: "Bacrimic", Cardinality: 6, Number: 3164, Perfection: 3, Imperfection: 3},

		// 6.20
		{Scale: scale.Phralimic, Name: "Phralimic", Cardinality: 6, Number: 2251, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrogimic, Name: "Phrogimic", Cardinality: 6, Number: 3256, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rathimic, Name: "Rathimic", Cardinality: 6, Number: 2417, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katocrimic, Name: "Katocrimic", Cardinality: 6, Number: 2956, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phryptimic, Name: "Phryptimic", Cardinality: 6, Number: 3634, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katynimic, Name: "Katynimic", Cardinality: 6, Number: 3173, Perfection: 3, Imperfection: 3},

		// 6.21
		{Scale: scale.Solimic, Name: "Solimic", Cardinality: 6, Number: 2253, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionolimic, Name: "Ionolimic", Cardinality: 6, Number: 3288, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionophimic, Name: "Ionophimic", Cardinality: 6, Number: 2481, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeologimic, Name: "Aeologimic", Cardinality: 6, Number: 3468, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zadimic, Name: "Zadimic", Cardinality: 6, Number: 2841, Perfection: 3, Imperfection: 3},
		{Scale: scale.Sygimic, Name: "Sygimic", Cardinality: 6, Number: 3174, Perfection: 3, Imperfection: 3},

		// 6.22
		{Scale: scale.Thogimic, Name: "Thogimic", Cardinality: 6, Number: 2254, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rythimic, Name: "Rythimic", Cardinality: 6, Number: 3304, Perfection: 3, Imperfection: 3},
		{Scale: scale.Donimic, Name: "Donimic", Cardinality: 6, Number: 2513, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeoloptimic, Name: "Aeoloptimic", Cardinality: 6, Number: 3724, Perfection: 3, Imperfection: 3},
		{Scale: scale.Panimic, Name: "Panimic", Cardinality: 6, Number: 3353, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lodimic, Name: "Lodimic", Cardinality: 6, Number: 2611, Perfection: 3, Imperfection: 3},

		// 6.23
		{Scale: scale.Laptimic, Name: "Laptimic", Cardinality: 6, Number: 2265, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lygimic, Name: "Lygimic", Cardinality: 6, Number: 3480, Perfection: 3, Imperfection: 3},
		{Scale: scale.Logimic, Name: "Logimic", Cardinality: 6, Number: 2865, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lalimic, Name: "Lalimic", Cardinality: 6, Number: 3270, Perfection: 3, Imperfection: 3},
		{Scale: scale.Sothimic, Name: "Sothimic", Cardinality: 6, Number: 2445, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrocrimic, Name: "Phrocrimic", Cardinality: 6, Number: 3180, Perfection: 3, Imperfection: 3},

		// 6.24
		{Scale: scale.Modimic, Name: "Modimic", Cardinality: 6, Number: 2266, Perfection: 3, Imperfection: 3},
		{Scale: scale.Barimic, Name: "Barimic", Cardinality: 6, Number: 3496, Perfection: 3, Imperfection: 3},
		{Scale: scale.Poptimic, Name: "Poptimic", Cardinality: 6, Number: 2897, Perfection: 3, Imperfection: 3},
		{Scale: scale.Sagimic, Name: "Sagimic", Cardinality: 6, Number: 3398, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aelothimic, Name: "Aelothimic", Cardinality: 6, Number: 2701, Perfection: 3, Imperfection: 3},
		{Scale: scale.Socrimic, Name: "Socrimic", Cardinality: 6, Number: 2614, Perfection: 3, Imperfection: 3},

		// 6.25
		{Scale: scale.Syrimic, Name: "Syrimic", Cardinality: 6, Number: 2268, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stodimic, Name: "Stodimic", Cardinality: 6, Number: 3528, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionocrimic, Name: "Ionocrimic", Cardinality: 6, Number: 2961, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zycrimic, Name: "Zycrimic", Cardinality: 6, Number: 3654, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionygimic, Name: "Ionygimic", Cardinality: 6, Number: 3213, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katathimic, Name: "Katathimic", Cardinality: 6, Number: 2331, Perfection: 3, Imperfection: 3},

		// 6.26
		{Scale: scale.Bolimic, Name: "Bolimic", Cardinality: 6, Number: 2278, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bothimic, Name: "Bothimic", Cardinality: 6, Number: 3688, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katadimic, Name: "Katadimic", Cardinality: 6, Number: 3281, Perfection: 3, Imperfection: 3},
		{Scale: scale.Kodimic, Name: "Kodimic", Cardinality: 6, Number: 2467, Perfection: 3, Imperfection: 3},
		{Scale: scale.Tholimic, Name: "Tholimic", Cardinality: 6, Number: 3356, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ralimic, Name: "Ralimic", Cardinality: 6, Number: 2617, Perfection: 3, Imperfection: 3},

		// 6.27
		{Scale: scale.Kanimic, Name: "Kanimic", Cardinality: 6, Number: 2281, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zylimic, Name: "Zylimic", Cardinality: 6, Number: 3736, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zodimic, Name: "Zodimic", Cardinality: 6, Number: 3377, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zarimic, Name: "Zarimic", Cardinality: 6, Number: 2659, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrythimic, Name: "Phrythimic", Cardinality: 6, Number: 2446, Perfection: 3, Imperfection: 3},
		{Scale: scale.Rorimic, Name: "Rorimic", Cardinality: 6, Number: 3188, Perfection: 3, Imperfection: 3},

		// 6.28
		{Scale: scale.Pynimic, Name: "Pynimic", Cardinality: 6, Number: 2290, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zanimic, Name: "Zanimic", Cardinality: 6, Number: 3880, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ranimic, Name: "Ranimic", Cardinality: 6, Number: 3665, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ladimic, Name: "Ladimic", Cardinality: 6, Number: 3235, Perfection: 3, Imperfection: 3},
		{Scale: scale.Podimic, Name: "Podimic", Cardinality: 6, Number: 2375, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionothimic, Name: "Ionothimic", Cardinality: 6, Number: 2620, Perfection: 3, Imperfection: 3},

		// 6.29
		{Scale: scale.Kytrimic, Name: "Kytrimic", Cardinality: 6, Number: 2292, Perfection: 3, Imperfection: 3},
		{Scale: scale.Golimic, Name: "Golimic", Cardinality: 6, Number: 3912, Perfection: 3, Imperfection: 3},
		{Scale: scale.Dyptimic, Name: "Dyptimic", Cardinality: 6, Number: 3729, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ryrimic, Name: "Ryrimic", Cardinality: 6, Number: 3363, Perfection: 3, Imperfection: 3},
		{Scale: scale.Gylimic, Name: "Gylimic", Cardinality: 6, Number: 2631, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeolycrimic, Name: "Aeolycrimic", Cardinality: 6, Number: 2334, Perfection: 3, Imperfection: 3},

		// 6.30
		{Scale: scale.Palimic, Name: "Palimic", Cardinality: 6, Number: 2347, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stothimic, Name: "Stothimic", Cardinality: 6, Number: 2396, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeronimic, Name: "Aeronimic", Cardinality: 6, Number: 2788, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katagimic, Name: "Katagimic", Cardinality: 6, Number: 2962, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phronimic, Name: "Phronimic", Cardinality: 6, Number: 3658, Perfection: 3, Imperfection: 3},
		{Scale: scale.Banimic, Name: "Banimic", Cardinality: 6, Number: 3221, Perfection: 3, Imperfection: 3},

		// 6.31
		{Scale: scale.Ionodimic, Name: "Ionodimic", Cardinality: 6, Number: 2355, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bogimic, Name: "Bogimic", Cardinality: 6, Number: 2460, Perfection: 3, Imperfection: 3},
		{Scale: scale.Mogimic, Name: "Mogimic", Cardinality: 6, Number: 3300, Perfection: 3, Imperfection: 3},
		{Scale: scale.Docrimic, Name: "Docrimic", Cardinality: 6, Number: 2505, Perfection: 3, Imperfection: 3},
		{Scale: scale.Epadimic, Name: "Epadimic", Cardinality: 6, Number: 3660, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerynimic, Name: "Aerynimic", Cardinality: 6, Number: 3225, Perfection: 3, Imperfection: 3},

		// 6.32
		{Scale: scale.Mydimic, Name: "Mydimic", Cardinality: 6, Number: 2361, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thyptimic, Name: "Thyptimic", Cardinality: 6, Number: 2508, Perfection: 3, Imperfection: 3},
		{Scale: scale.Phrothimic, Name: "Phrothimic", Cardinality: 6, Number: 3684, Perfection: 3, Imperfection: 3},
		{Scale: scale.Katycrimic, Name: "Katycrimic", Cardinality: 6, Number: 3273, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionalimic, Name: "Ionalimic", Cardinality: 6, Number: 2451, Perfection: 3, Imperfection: 3},
		{Scale: scale.Loptimic, Name: "Loptimic", Cardinality: 6, Number: 3228, Perfection: 3, Imperfection: 3},

		// 6.33
		{Scale: scale.Zagimic, Name: "Zagimic", Cardinality: 6, Number: 2362, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lagimic, Name: "Lagimic", Cardinality: 6, Number: 2516, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thyrimic, Name: "Thyrimic", Cardinality: 6, Number: 3748, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thothimic, Name: "Thothimic", Cardinality: 6, Number: 3401, Perfection: 3, Imperfection: 3},
		{Scale: scale.Bycrimic, Name: "Bycrimic", Cardinality: 6, Number: 2707, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pathimic, Name: "Pathimic", Cardinality: 6, Number: 2638, Perfection: 3, Imperfection: 3},

		// 6.34
		{Scale: scale.Mothimic, Name: "Mothimic", Cardinality: 6, Number: 2393, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeranimic, Name: "Aeranimic", Cardinality: 6, Number: 2764, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ragimic, Name: "Ragimic", Cardinality: 6, Number: 2866, Perfection: 3, Imperfection: 3},
		{Scale: scale.Dolimic, Name: "Dolimic", Cardinality: 6, Number: 3274, Perfection: 3, Imperfection: 3},
		{Scale: scale.Porimic, Name: "Porimic", Cardinality: 6, Number: 2453, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerylimic, Name: "Aerylimic", Cardinality: 6, Number: 3244, Perfection: 3, Imperfection: 3},

		// 6.35
		{Scale: scale.Bocrimic, Name: "Bocrimic", Cardinality: 6, Number: 2406, Perfection: 3, Imperfection: 3},
		{Scale: scale.Gythimic, Name: "Gythimic", Cardinality: 6, Number: 2868, Perfection: 3, Imperfection: 3},
		{Scale: scale.Pagimic, Name: "Pagimic", Cardinality: 6, Number: 3282, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aeolythimic, Name: "Aeolythimic", Cardinality: 6, Number: 2469, Perfection: 3, Imperfection: 3},
		{Scale: scale.Molimic, Name: "Molimic", Cardinality: 6, Number: 3372, Perfection: 3, Imperfection: 3},
		{Scale: scale.Staptimic, Name: "Staptimic", Cardinality: 6, Number: 2649, Perfection: 3, Imperfection: 3},

		// 6.36
		{Scale: scale.Zacrimic, Name: "Zacrimic", Cardinality: 6, Number: 2409, Perfection: 3, Imperfection: 3},
		{Scale: scale.Larimic, Name: "Larimic", Cardinality: 6, Number: 2892, Perfection: 3, Imperfection: 3},
		{Scale: scale.Thacrimic, Name: "Thacrimic", Cardinality: 6, Number: 3378, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stydimic, Name: "Stydimic", Cardinality: 6, Number: 2661, Perfection: 3, Imperfection: 3},
		{Scale: scale.Lorimic, Name: "Lorimic", Cardinality: 6, Number: 2454, Perfection: 3, Imperfection: 3},
		{Scale: scale.Ionadimic, Name: "Ionadimic", Cardinality: 6, Number: 3252, Perfection: 3, Imperfection: 3},

		// 6.37
		{Scale: scale.Ionythimic, Name: "Ionythimic", Cardinality: 6, Number: 2457, Perfection: 3, Imperfection: 3},
		{Scale: scale.Aerythimic, Name: "Aerythimic", Cardinality: 6, Number: 3276, Perfection: 3, Imperfection: 3},

		// 6.38
		{Scale: scale.Dynimic, Name: "Dynimic", Cardinality: 6, Number: 2458, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zydimic, Name: "Zydimic", Cardinality: 6, Number: 3284, Perfection: 3, Imperfection: 3},
		{Scale: scale.Zathimic, Name: "Zathimic", Cardinality: 6, Number: 2473, Perfection: 3, Imperfection: 3},
		{Scale: scale.Radimic, Name: "Radimic", Cardinality: 6, Number: 3404, Perfection: 3, Imperfection: 3},
		{Scale: scale.Stonimic, Name: "Stonimic", Cardinality: 6, Number: 2713, Perfection: 3, Imperfection: 3},
		{Scale: scale.Syptimic, Name: "Syptimic", Cardinality: 6, Number: 2662, Perfection: 3, Imperfection: 3},

		// 6.39
		{Scale: scale.Ponimic, Name: "Ponimic", Cardinality: 6, Number: 2191, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kadimic, Name: "Kadimic", Cardinality: 6, Number: 2296, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gynimic, Name: "Gynimic", Cardinality: 6, Number: 3976, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thydimic, Name: "Thydimic", Cardinality: 6, Number: 3857, Perfection: 2, Imperfection: 4},
		{Scale: scale.Polimic, Name: "Polimic", Cardinality: 6, Number: 3619, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thanimic, Name: "Thanimic", Cardinality: 6, Number: 3143, Perfection: 2, Imperfection: 4},

		// 6.40
		{Scale: scale.Lathimic, Name: "Lathimic", Cardinality: 6, Number: 2203, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeralimic, Name: "Aeralimic", Cardinality: 6, Number: 2488, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kynimic, Name: "Kynimic", Cardinality: 6, Number: 3524, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stynimic, Name: "Stynimic", Cardinality: 6, Number: 2953, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epytimic, Name: "Epytimic", Cardinality: 6, Number: 3622, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katoptimic, Name: "Katoptimic", Cardinality: 6, Number: 3149, Perfection: 2, Imperfection: 4},

		// 6.41
		{Scale: scale.Galimic, Name: "Galimic", Cardinality: 6, Number: 2206, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kathimic, Name: "Kathimic", Cardinality: 6, Number: 2536, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lylimic, Name: "Lylimic", Cardinality: 6, Number: 3908, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epalimic, Name: "Epalimic", Cardinality: 6, Number: 3721, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epacrimic, Name: "Epacrimic", Cardinality: 6, Number: 3347, Perfection: 2, Imperfection: 4},
		{Scale: scale.Sathimic, Name: "Sathimic", Cardinality: 6, Number: 2599, Perfection: 2, Imperfection: 4},

		// 6.42
		{Scale: scale.Katanimic, Name: "Katanimic", Cardinality: 6, Number: 2219, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katyrimic, Name: "Katyrimic", Cardinality: 6, Number: 2744, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rynimic, Name: "Rynimic", Cardinality: 6, Number: 2786, Perfection: 2, Imperfection: 4},
		{Scale: scale.Pogimic, Name: "Pogimic", Cardinality: 6, Number: 2954, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeraptimic, Name: "Aeraptimic", Cardinality: 6, Number: 3626, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epylimic, Name: "Epylimic", Cardinality: 6, Number: 3157, Perfection: 2, Imperfection: 4},

		// 6.43
		{Scale: scale.Manimic, Name: "Manimic", Cardinality: 6, Number: 2230, Perfection: 2, Imperfection: 4},
		{Scale: scale.Marimic, Name: "Marimic", Cardinality: 6, Number: 2920, Perfection: 2, Imperfection: 4},
		{Scale: scale.Locrimic, Name: "Locrimic", Cardinality: 6, Number: 3490, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rylimic, Name: "Rylimic", Cardinality: 6, Number: 2885, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epatimic, Name: "Epatimic", Cardinality: 6, Number: 3350, Perfection: 2, Imperfection: 4},
		{Scale: scale.Byrimic, Name: "Byrimic", Cardinality: 6, Number: 2605, Perfection: 2, Imperfection: 4},

		// 6.44
		{Scale: scale.Kocrimic, Name: "Kocrimic", Cardinality: 6, Number: 2236, Perfection: 2, Imperfection: 4},
		{Scale: scale.Korimic, Name: "Korimic", Cardinality: 6, Number: 3016, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lynimic, Name: "Lynimic", Cardinality: 6, Number: 3874, Perfection: 2, Imperfection: 4},
		{Scale: scale.Malimic, Name: "Malimic", Cardinality: 6, Number: 3653, Perfection: 2, Imperfection: 4},
		{Scale: scale.Synimic, Name: "Synimic", Cardinality: 6, Number: 3211, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phragimic, Name: "Phragimic", Cardinality: 6, Number: 2327, Perfection: 2, Imperfection: 4},

		// 6.45
		{Scale: scale.Mycrimic, Name: "Mycrimic", Cardinality: 6, Number: 2282, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionorimic, Name: "Ionorimic", Cardinality: 6, Number: 3752, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phrydimic, Name: "Phrydimic", Cardinality: 6, Number: 3409, Perfection: 2, Imperfection: 4},
		{Scale: scale.Zyptimic, Name: "Zyptimic", Cardinality: 6, Number: 2723, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katothimic, Name: "Katothimic", Cardinality: 6, Number: 2702, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phrylimic, Name: "Phrylimic", Cardinality: 6, Number: 2618, Perfection: 2, Imperfection: 4},

		// 6.46
		{Scale: scale.Aerothimic, Name: "Aerothimic", Cardinality: 6, Number: 2284, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stagimic, Name: "Stagimic", Cardinality: 6, Number: 3784, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dorimic, Name: "Dorimic", Cardinality: 6, Number: 3473, Perfection: 2, Imperfection: 4},
		{Scale: scale.Phrycrimic, Name: "Phrycrimic", Cardinality: 6, Number: 2851, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kyptimic, Name: "Kyptimic", Cardinality: 6, Number: 3214, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionylimic, Name: "Ionylimic", Cardinality: 6, Number: 2333, Perfection: 2, Imperfection: 4},

		// 6.47
		{Scale: scale.Epynimic, Name: "Epynimic", Cardinality: 6, Number: 2343, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionogimic, Name: "Ionogimic", Cardinality: 6, Number: 2364, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kydimic, Name: "Kydimic", Cardinality: 6, Number: 2532, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gaptimic, Name: "Gaptimic", Cardinality: 6, Number: 3876, Perfection: 2, Imperfection: 4},
		{Scale: scale.Tharimic, Name: "Tharimic", Cardinality: 6, Number: 3657, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionaphimic, Name: "Ionaphimic", Cardinality: 6, Number: 3219, Perfection: 2, Imperfection: 4},

		// 6.48
		{Scale: scale.Thoptimic, Name: "Thoptimic", Cardinality: 6, Number: 2349, Perfection: 2, Imperfection: 4},
		{Scale: scale.Bagimic, Name: "Bagimic", Cardinality: 6, Number: 2412, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kyrimic, Name: "Kyrimic", Cardinality: 6, Number: 2916, Perfection: 2, Imperfection: 4},
		{Scale: scale.Sonimic, Name: "Sonimic", Cardinality: 6, Number: 3474, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolonimic, Name: "Aeolonimic", Cardinality: 6, Number: 2853, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rygimic, Name: "Rygimic", Cardinality: 6, Number: 3222, Perfection: 2, Imperfection: 4},

		// 6.49
		{Scale: scale.Thagimic, Name: "Thagimic", Cardinality: 6, Number: 2350, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kolimic, Name: "Kolimic", Cardinality: 6, Number: 2420, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dycrimic, Name: "Dycrimic", Cardinality: 6, Number: 2980, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epycrimic, Name: "Epycrimic", Cardinality: 6, Number: 3730, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gocrimic, Name: "Gocrimic", Cardinality: 6, Number: 3365, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katolimic, Name: "Katolimic", Cardinality: 6, Number: 2635, Perfection: 2, Imperfection: 4},

		// 6.50
		{Scale: scale.Dagimic, Name: "Dagimic", Cardinality: 6, Number: 2357, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolydimic, Name: "Aeolydimic", Cardinality: 6, Number: 2476, Perfection: 2, Imperfection: 4},
		{Scale: scale.Parimic, Name: "Parimic", Cardinality: 6, Number: 3428, Perfection: 2, Imperfection: 4},
		{Scale: scale.Ionaptimic, Name: "Ionaptimic", Cardinality: 6, Number: 2761, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thylimic, Name: "Thylimic", Cardinality: 6, Number: 2854, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lolimic, Name: "Lolimic", Cardinality: 6, Number: 3226, Perfection: 2, Imperfection: 4},

		// 6.51
		{Scale: scale.Thalimic, Name: "Thalimic", Cardinality: 6, Number: 2358, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stygimic, Name: "Stygimic", Cardinality: 6, Number: 2484, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolygimic, Name: "Aeolygimic", Cardinality: 6, Number: 3492, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aerogimic, Name: "Aerogimic", Cardinality: 6, Number: 2889, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dacrimic, Name: "Dacrimic", Cardinality: 6, Number: 3366, Perfection: 2, Imperfection: 4},
		{Scale: scale.Baptimic, Name: "Baptimic", Cardinality: 6, Number: 2637, Perfection: 2, Imperfection: 4},

		// 6.52
		{Scale: scale.Stythimic, Name: "Stythimic", Cardinality: 6, Number: 2381, Perfection: 2, Imperfection: 4},
		{Scale: scale.Kothimic, Name: "Kothimic", Cardinality: 6, Number: 2668, Perfection: 2, Imperfection: 4},
		{Scale: scale.Pygimic, Name: "Pygimic", Cardinality: 6, Number: 2482, Perfection: 2, Imperfection: 4},
		{Scale: scale.Rodimic, Name: "Rodimic", Cardinality: 6, Number: 3476, Perfection: 2, Imperfection: 4},
		{Scale: scale.Sorimic, Name: "Sorimic", Cardinality: 6, Number: 2857, Perfection: 2, Imperfection: 4},
		{Scale: scale.Monimic, Name: "Monimic", Cardinality: 6, Number: 3238, Perfection: 2, Imperfection: 4},

		// 6.53
		{Scale: scale.Aeragimic, Name: "Aeragimic", Cardinality: 6, Number: 2389, Perfection: 2, Imperfection: 4},
		{Scale: scale.Epothimic, Name: "Epothimic", Cardinality: 6, Number: 2732, Perfection: 2, Imperfection: 4},
		{Scale: scale.Salimic, Name: "Salimic", Cardinality: 6, Number: 2738, Perfection: 2, Imperfection: 4},
		{Scale: scale.Lyptimic, Name: "Lyptimic", Cardinality: 6, Number: 2762, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katonimic, Name: "Katonimic", Cardinality: 6, Number: 2858, Perfection: 2, Imperfection: 4},
		{Scale: scale.Gygimic, Name: "Gygimic", Cardinality: 6, Number: 3242, Perfection: 2, Imperfection: 4},

		// 6.54
		{Scale: scale.Aeradimic, Name: "Aeradimic", Cardinality: 6, Number: 2405, Perfection: 2, Imperfection: 4},
		{Scale: scale.Zyrimic, Name: "Zyrimic", Cardinality: 6, Number: 2860, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stylimic, Name: "Stylimic", Cardinality: 6, Number: 3250, Perfection: 2, Imperfection: 4},

		// 6.55
		{Scale: scale.Lythimic, Name: "Lythimic", Cardinality: 6, Number: 2470, Perfection: 2, Imperfection: 4},
		{Scale: scale.Dodimic, Name: "Dodimic", Cardinality: 6, Number: 3380, Perfection: 2, Imperfection: 4},
		{Scale: scale.Katalimic, Name: "Katalimic", Cardinality: 6, Number: 2665, Perfection: 2, Imperfection: 4},

		// 6.56
		{Scale: scale.Boptimic, Name: "Boptimic", Cardinality: 6, Number: 2474, Perfection: 2, Imperfection: 4},
		{Scale: scale.Stogimic, Name: "Stogimic", Cardinality: 6, Number: 3412, Perfection: 2, Imperfection: 4},
		{Scale: scale.Thynimic, Name: "Thynimic", Cardinality: 6, Number: 2729, Perfection: 2, Imperfection: 4},
		{Scale: scale.Aeolathimic, Name: "Aeolathimic", Cardinality: 6, Number: 2726, Perfection: 2, Imperfection: 4},
		{Scale: scale.Bythimic, Name: "Bythimic", Cardinality: 6, Number: 2714, Perfection: 2, Imperfection: 4},
		{Scale: scale.Padimic, Name: "Padimic", Cardinality: 6, Number: 2666, Perfection: 2, Imperfection: 4},

		// 6.57
		{Scale: scale.Dathimic, Name: "Dathimic", Cardinality: 6, Number: 2222, Perfection: 1, Imperfection: 5},
		{Scale: scale.Epagimic, Name: "Epagimic", Cardinality: 6, Number: 2792, Perfection: 1, Imperfection: 5},
		{Scale: scale.Raptimic, Name: "Raptimic", Cardinality: 6, Number: 2978, Perfection: 1, Imperfection: 5},
		{Scale: scale.Epolimic, Name: "Epolimic", Cardinality: 6, Number: 3722, Perfection: 1, Imperfection: 5},
		{Scale: scale.Sythimic, Name: "Sythimic", Cardinality: 6, Number: 3349, Perfection: 1, Imperfection: 5},
		{Scale: scale.Sydimic, Name: "Sydimic", Cardinality: 6, Number: 2603, Perfection: 1, Imperfection: 5},

		// 6.58
		{Scale: scale.Gacrimic, Name: "Gacrimic", Cardinality: 6, Number: 2234, Perfection: 1, Imperfection: 5},
		{Scale: scale.Borimic, Name: "Borimic", Cardinality: 6, Number: 2984, Perfection: 1, Imperfection: 5},
		{Scale: scale.Sycrimic, Name: "Sycrimic", Cardinality: 6, Number: 3746, Perfection: 1, Imperfection: 5},
		{Scale: scale.Gadimic, Name: "Gadimic", Cardinality: 6, Number: 3397, Perfection: 1, Imperfection: 5},
		{Scale: scale.Aeolocrimic, Name: "Aeolocrimic", Cardinality: 6, Number: 2699, Perfection: 1, Imperfection: 5},
		{Scale: scale.Phrygimic, Name: "Phrygimic", Cardinality: 6, Number: 2606, Perfection: 1, Imperfection: 5},

		// 6.59
		{Scale: scale.WholeTone, Name: "WholeTone", Cardinality: 6, Number: 2730, Perfection: 0, Imperfection: 6},

		// 7.1
		{Scale: scale.Lydian, Name: "Lydian", Cardinality: 7, Number: 2741, Perfection: 6, Imperfection: 1},
		{Scale: scale.Mixolydian, Name: "Mixolydian", Cardinality: 7, Number: 2774, Perfection: 6, Imperfection: 1},
		{Scale: scale.Aeolian, Name: "Aeolian", Cardinality: 7, Number: 2906, Perfection: 6, Imperfection: 1},
		{Scale: scale.Locrian, Name: "Locrian", Cardinality: 7, Number: 3434, Perfection: 6, Imperfection: 1},
		{Scale: scale.Ionian, Name: "Ionian", Cardinality: 7, Number: 2773, Perfection: 6, Imperfection: 1},
		{Scale: scale.Dorian, Name: "Dorian", Cardinality: 7, Number: 2902, Perfection: 6, Imperfection: 1},
		{Scale: scale.Phrygian, Name: "Phrygian", Cardinality: 7, Number: 3418, Perfection: 6, Imperfection: 1},

		// 7.2
		{Scale: scale.Ionythian, Name: "Ionythian", Cardinality: 7, Number: 2263, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeolyrian, Name: "Aeolyrian", Cardinality: 7, Number: 3448, Perfection: 5, Imperfection: 2},
		{Scale: scale.Gorian, Name: "Gorian", Cardinality: 7, Number: 2801, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeolodian, Name: "Aeolodian", Cardinality: 7, Number: 3014, Perfection: 5, Imperfection: 2},
		{Scale: scale.Doptian, Name: "Doptian", Cardinality: 7, Number: 3866, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeraphian, Name: "Aeraphian", Cardinality: 7, Number: 3637, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zacrian, Name: "Zacrian", Cardinality: 7, Number: 3179, Perfection: 5, Imperfection: 2},

		// 7.3
		{Scale: scale.Ionarian, Name: "Ionarian", Cardinality: 7, Number: 2279, Perfection: 5, Imperfection: 2},
		{Scale: scale.Dynian, Name: "Dynian", Cardinality: 7, Number: 3704, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zydian, Name: "Zydian", Cardinality: 7, Number: 3313, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zathian, Name: "Zathian", Cardinality: 7, Number: 2531, Perfection: 5, Imperfection: 2},
		{Scale: scale.Radian, Name: "Radian", Cardinality: 7, Number: 3868, Perfection: 5, Imperfection: 2},
		{Scale: scale.Stonian, Name: "Stonian", Cardinality: 7, Number: 3641, Perfection: 5, Imperfection: 2},
		{Scale: scale.Syptian, Name: "Syptian", Cardinality: 7, Number: 3187, Perfection: 5, Imperfection: 2},

		// 7.4
		{Scale: scale.Aeolacrian, Name: "Aeolacrian", Cardinality: 7, Number: 2291, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zythian, Name: "Zythian", Cardinality: 7, Number: 3896, Perfection: 5, Imperfection: 2},
		{Scale: scale.Dyrian, Name: "Dyrian", Cardinality: 7, Number: 3697, Perfection: 5, Imperfection: 2},
		{Scale: scale.Koptian, Name: "Koptian", Cardinality: 7, Number: 3299, Perfection: 5, Imperfection: 2},
		{Scale: scale.Thocrian, Name: "Thocrian", Cardinality: 7, Number: 2503, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeolanian, Name: "Aeolanian", Cardinality: 7, Number: 3644, Perfection: 5, Imperfection: 2},
		{Scale: scale.Danian, Name: "Danian", Cardinality: 7, Number: 3193, Perfection: 5, Imperfection: 2},

		// 7.5
		{Scale: scale.Zogian, Name: "Zogian", Cardinality: 7, Number: 2293, Perfection: 5, Imperfection: 2},
		{Scale: scale.Epyrian, Name: "Epyrian", Cardinality: 7, Number: 3928, Perfection: 5, Imperfection: 2},
		{Scale: scale.Lycrian, Name: "Lycrian", Cardinality: 7, Number: 3761, Perfection: 5, Imperfection: 2},
		{Scale: scale.Daptian, Name: "Daptian", Cardinality: 7, Number: 3427, Perfection: 5, Imperfection: 2},
		{Scale: scale.Kygian, Name: "Kygian", Cardinality: 7, Number: 2759, Perfection: 5, Imperfection: 2},
		{Scale: scale.Mocrian, Name: "Mocrian", Cardinality: 7, Number: 2846, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zynian, Name: "Zynian", Cardinality: 7, Number: 3194, Perfection: 5, Imperfection: 2},

		// 7.6
		{Scale: scale.Phrolian, Name: "Phrolian", Cardinality: 7, Number: 2395, Perfection: 5, Imperfection: 2},
		{Scale: scale.Ionagian, Name: "Ionagian", Cardinality: 7, Number: 2780, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeodian, Name: "Aeodian", Cardinality: 7, Number: 2930, Perfection: 5, Imperfection: 2},
		{Scale: scale.Kycrian, Name: "Kycrian", Cardinality: 7, Number: 3530, Perfection: 5, Imperfection: 2},
		{Scale: scale.Epygian, Name: "Epygian", Cardinality: 7, Number: 2965, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zaptian, Name: "Zaptian", Cardinality: 7, Number: 3670, Perfection: 5, Imperfection: 2},
		{Scale: scale.Kagian, Name: "Kagian", Cardinality: 7, Number: 3245, Perfection: 5, Imperfection: 2},

		// 7.7
		{Scale: scale.Soptian, Name: "Soptian", Cardinality: 7, Number: 2398, Perfection: 5, Imperfection: 2},
		{Scale: scale.Ionyptian, Name: "Ionyptian", Cardinality: 7, Number: 2804, Perfection: 5, Imperfection: 2},
		{Scale: scale.Gyrian, Name: "Gyrian", Cardinality: 7, Number: 3026, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zalian, Name: "Zalian", Cardinality: 7, Number: 3914, Perfection: 5, Imperfection: 2},
		{Scale: scale.Stolian, Name: "Stolian", Cardinality: 7, Number: 3733, Perfection: 5, Imperfection: 2},
		{Scale: scale.Bylian, Name: "Bylian", Cardinality: 7, Number: 3371, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zothian, Name: "Zothian", Cardinality: 7, Number: 2647, Perfection: 5, Imperfection: 2},

		// 7.8
		{Scale: scale.Thonian, Name: "Thonian", Cardinality: 7, Number: 2411, Perfection: 5, Imperfection: 2},
		{Scale: scale.Phrorian, Name: "Phrorian", Cardinality: 7, Number: 2908, Perfection: 5, Imperfection: 2},
		{Scale: scale.Stadian, Name: "Stadian", Cardinality: 7, Number: 3442, Perfection: 5, Imperfection: 2},
		{Scale: scale.Thodian, Name: "Thodian", Cardinality: 7, Number: 2789, Perfection: 5, Imperfection: 2},
		{Scale: scale.Dogian, Name: "Dogian", Cardinality: 7, Number: 2966, Perfection: 5, Imperfection: 2},
		{Scale: scale.Mixopyrian, Name: "Mixopyrian", Cardinality: 7, Number: 3674, Perfection: 5, Imperfection: 2},
		{Scale: scale.Garian, Name: "Garian", Cardinality: 7, Number: 3253, Perfection: 5, Imperfection: 2},

		// 7.9
		{Scale: scale.Epathian, Name: "Epathian", Cardinality: 7, Number: 2419, Perfection: 5, Imperfection: 2},
		{Scale: scale.Mythian, Name: "Mythian", Cardinality: 7, Number: 2972, Perfection: 5, Imperfection: 2},
		{Scale: scale.Sogian, Name: "Sogian", Cardinality: 7, Number: 3698, Perfection: 5, Imperfection: 2},
		{Scale: scale.Gogian, Name: "Gogian", Cardinality: 7, Number: 3301, Perfection: 5, Imperfection: 2},
		{Scale: scale.Rothian, Name: "Rothian", Cardinality: 7, Number: 2507, Perfection: 5, Imperfection: 2},
		{Scale: scale.Katarian, Name: "Katarian", Cardinality: 7, Number: 3676, Perfection: 5, Imperfection: 2},
		{Scale: scale.Stylian, Name: "Stylian", Cardinality: 7, Number: 3257, Perfection: 5, Imperfection: 2},

		// 7.10
		{Scale: scale.Stathian, Name: "Stathian", Cardinality: 7, Number: 2426, Perfection: 5, Imperfection: 2},
		{Scale: scale.Mixonyphian, Name: "Mixonyphian", Cardinality: 7, Number: 3028, Perfection: 5, Imperfection: 2},
		{Scale: scale.Magian, Name: "Magian", Cardinality: 7, Number: 3922, Perfection: 5, Imperfection: 2},
		{Scale: scale.Dadian, Name: "Dadian", Cardinality: 7, Number: 3749, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeolylian, Name: "Aeolylian", Cardinality: 7, Number: 3403, Perfection: 5, Imperfection: 2},
		{Scale: scale.Gycrian, Name: "Gycrian", Cardinality: 7, Number: 2711, Perfection: 5, Imperfection: 2},
		{Scale: scale.Pyrian, Name: "Pyrian", Cardinality: 7, Number: 2654, Perfection: 5, Imperfection: 2},

		// 7.11
		{Scale: scale.Epogian, Name: "Epogian", Cardinality: 7, Number: 2510, Perfection: 5, Imperfection: 2},
		{Scale: scale.Lanian, Name: "Lanian", Cardinality: 7, Number: 3700, Perfection: 5, Imperfection: 2},
		{Scale: scale.Paptian, Name: "Paptian", Cardinality: 7, Number: 3305, Perfection: 5, Imperfection: 2},
		{Scale: scale.Ionacrian, Name: "Ionacrian", Cardinality: 7, Number: 2515, Perfection: 5, Imperfection: 2},
		{Scale: scale.Gathian, Name: "Gathian", Cardinality: 7, Number: 3740, Perfection: 5, Imperfection: 2},
		{Scale: scale.Ionyphian, Name: "Ionyphian", Cardinality: 7, Number: 3385, Perfection: 5, Imperfection: 2},
		{Scale: scale.Phrynian, Name: "Phrynian", Cardinality: 7, Number: 2675, Perfection: 5, Imperfection: 2},

		// 7.12
		{Scale: scale.Ionycrian, Name: "Ionycrian", Cardinality: 7, Number: 2518, Perfection: 5, Imperfection: 2},
		{Scale: scale.Phradian, Name: "Phradian", Cardinality: 7, Number: 3764, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeolorian, Name: "Aeolorian", Cardinality: 7, Number: 3433, Perfection: 5, Imperfection: 2},
		{Scale: scale.Gonian, Name: "Gonian", Cardinality: 7, Number: 2771, Perfection: 5, Imperfection: 2},
		{Scale: scale.Dalian, Name: "Dalian", Cardinality: 7, Number: 2894, Perfection: 5, Imperfection: 2},
		{Scale: scale.Dygian, Name: "Dygian", Cardinality: 7, Number: 3386, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zolian, Name: "Zolian", Cardinality: 7, Number: 2677, Perfection: 5, Imperfection: 2},

		// 7.13
		{Scale: scale.Aerathian, Name: "Aerathian", Cardinality: 7, Number: 2522, Perfection: 5, Imperfection: 2},
		{Scale: scale.Sarian, Name: "Sarian", Cardinality: 7, Number: 3796, Perfection: 5, Imperfection: 2},
		{Scale: scale.Zoptian, Name: "Zoptian", Cardinality: 7, Number: 3497, Perfection: 5, Imperfection: 2},
		{Scale: scale.Aeracrian, Name: "Aeracrian", Cardinality: 7, Number: 2899, Perfection: 5, Imperfection: 2},
		{Scale: scale.Byptian, Name: "Byptian", Cardinality: 7, Number: 3406, Perfection: 5, Imperfection: 2},
		{Scale: scale.Darian, Name: "Darian", Cardinality: 7, Number: 2717, Perfection: 5, Imperfection: 2},
		{Scale: scale.Lonian, Name: "Lonian", Cardinality: 7, Number: 2678, Perfection: 5, Imperfection: 2},

		// 7.14
		{Scale: scale.Aeopian, Name: "Aeopian", Cardinality: 7, Number: 2231, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rygian, Name: "Rygian", Cardinality: 7, Number: 2936, Perfection: 4, Imperfection: 3},
		{Scale: scale.Epynian, Name: "Epynian", Cardinality: 7, Number: 3554, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionogian, Name: "Ionogian", Cardinality: 7, Number: 3013, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kydian, Name: "Kydian", Cardinality: 7, Number: 3862, Perfection: 4, Imperfection: 3},
		{Scale: scale.Gaptian, Name: "Gaptian", Cardinality: 7, Number: 3629, Perfection: 4, Imperfection: 3},
		{Scale: scale.Tharian, Name: "Tharian", Cardinality: 7, Number: 3163, Perfection: 4, Imperfection: 3},

		// 7.15
		{Scale: scale.Epycrian, Name: "Epycrian", Cardinality: 7, Number: 2237, Perfection: 4, Imperfection: 3},
		{Scale: scale.Gocrian, Name: "Gocrian", Cardinality: 7, Number: 3032, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katolian, Name: "Katolian", Cardinality: 7, Number: 3938, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thoptian, Name: "Thoptian", Cardinality: 7, Number: 3781, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bagian, Name: "Bagian", Cardinality: 7, Number: 3467, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kyrian, Name: "Kyrian", Cardinality: 7, Number: 2839, Perfection: 4, Imperfection: 3},
		{Scale: scale.Sonian, Name: "Sonian", Cardinality: 7, Number: 3166, Perfection: 4, Imperfection: 3},

		// 7.16
		{Scale: scale.Parian, Name: "Parian", Cardinality: 7, Number: 2255, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionaptian, Name: "Ionaptian", Cardinality: 7, Number: 3320, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thylian, Name: "Thylian", Cardinality: 7, Number: 2545, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lolian, Name: "Lolian", Cardinality: 7, Number: 3980, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thagian, Name: "Thagian", Cardinality: 7, Number: 3865, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kolian, Name: "Kolian", Cardinality: 7, Number: 3635, Perfection: 4, Imperfection: 3},
		{Scale: scale.Dycrian, Name: "Dycrian", Cardinality: 7, Number: 3175, Perfection: 4, Imperfection: 3},

		// 7.17
		{Scale: scale.Stygian, Name: "Stygian", Cardinality: 7, Number: 2267, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolygian, Name: "Aeolygian", Cardinality: 7, Number: 3512, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerogian, Name: "Aerogian", Cardinality: 7, Number: 2929, Perfection: 4, Imperfection: 3},
		{Scale: scale.Dacrian, Name: "Dacrian", Cardinality: 7, Number: 3526, Perfection: 4, Imperfection: 3},
		{Scale: scale.Baptian, Name: "Baptian", Cardinality: 7, Number: 2957, Perfection: 4, Imperfection: 3},
		{Scale: scale.Dagian, Name: "Dagian", Cardinality: 7, Number: 3638, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolydian, Name: "Aeolydian", Cardinality: 7, Number: 3181, Perfection: 4, Imperfection: 3},

		// 7.18
		{Scale: scale.Stythian, Name: "Stythian", Cardinality: 7, Number: 2269, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kothian, Name: "Kothian", Cardinality: 7, Number: 3544, Perfection: 4, Imperfection: 3},
		{Scale: scale.Pygian, Name: "Pygian", Cardinality: 7, Number: 2993, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rodian, Name: "Rodian", Cardinality: 7, Number: 3782, Perfection: 4, Imperfection: 3},
		{Scale: scale.Sorian, Name: "Sorian", Cardinality: 7, Number: 3469, Perfection: 4, Imperfection: 3},
		{Scale: scale.Monian, Name: "Monian", Cardinality: 7, Number: 2843, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thalian, Name: "Thalian", Cardinality: 7, Number: 3182, Perfection: 4, Imperfection: 3},

		// 7.19
		{Scale: scale.Zorian, Name: "Zorian", Cardinality: 7, Number: 2270, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeragian, Name: "Aeragian", Cardinality: 7, Number: 3560, Perfection: 4, Imperfection: 3},
		{Scale: scale.Epothian, Name: "Epothian", Cardinality: 7, Number: 3025, Perfection: 4, Imperfection: 3},
		{Scale: scale.Salian, Name: "Salian", Cardinality: 7, Number: 3910, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lyptian, Name: "Lyptian", Cardinality: 7, Number: 3725, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katonian, Name: "Katonian", Cardinality: 7, Number: 3355, Perfection: 4, Imperfection: 3},
		{Scale: scale.Gyphian, Name: "Gyphian", Cardinality: 7, Number: 2615, Perfection: 4, Imperfection: 3},

		// 7.20
		{Scale: scale.Thacrian, Name: "Thacrian", Cardinality: 7, Number: 2283, Perfection: 4, Imperfection: 3},
		{Scale: scale.Dodian, Name: "Dodian", Cardinality: 7, Number: 3768, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolyptian, Name: "Aeolyptian", Cardinality: 7, Number: 3441, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolonian, Name: "Aeolonian", Cardinality: 7, Number: 2787, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeradian, Name: "Aeradian", Cardinality: 7, Number: 2958, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolagian, Name: "Aeolagian", Cardinality: 7, Number: 3642, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zyrian, Name: "Zyrian", Cardinality: 7, Number: 3189, Perfection: 4, Imperfection: 3},

		// 7.21
		{Scale: scale.Aeolathian, Name: "Aeolathian", Cardinality: 7, Number: 2285, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bythian, Name: "Bythian", Cardinality: 7, Number: 3800, Perfection: 4, Imperfection: 3},
		{Scale: scale.Padian, Name: "Padian", Cardinality: 7, Number: 3505, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rolian, Name: "Rolian", Cardinality: 7, Number: 2915, Perfection: 4, Imperfection: 3},
		{Scale: scale.Pydian, Name: "Pydian", Cardinality: 7, Number: 3470, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thygian, Name: "Thygian", Cardinality: 7, Number: 2845, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katalian, Name: "Katalian", Cardinality: 7, Number: 3190, Perfection: 4, Imperfection: 3},

		// 7.22
		{Scale: scale.Saptian, Name: "Saptian", Cardinality: 7, Number: 2294, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerodian, Name: "Aerodian", Cardinality: 7, Number: 3944, Perfection: 4, Imperfection: 3},
		{Scale: scale.Macrian, Name: "Macrian", Cardinality: 7, Number: 3793, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rogian, Name: "Rogian", Cardinality: 7, Number: 3491, Perfection: 4, Imperfection: 3},
		{Scale: scale.Boptian, Name: "Boptian", Cardinality: 7, Number: 2887, Perfection: 4, Imperfection: 3},
		{Scale: scale.Stogian, Name: "Stogian", Cardinality: 7, Number: 3358, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thynian, Name: "Thynian", Cardinality: 7, Number: 2621, Perfection: 4, Imperfection: 3},

		// 7.23
		{Scale: scale.Thycrian, Name: "Thycrian", Cardinality: 7, Number: 2297, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeoladian, Name: "Aeoladian", Cardinality: 7, Number: 3992, Perfection: 4, Imperfection: 3},
		{Scale: scale.Dylian, Name: "Dylian", Cardinality: 7, Number: 3889, Perfection: 4, Imperfection: 3},
		{Scale: scale.Eponian, Name: "Eponian", Cardinality: 7, Number: 3683, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katygian, Name: "Katygian", Cardinality: 7, Number: 3271, Perfection: 4, Imperfection: 3},
		{Scale: scale.Starian, Name: "Starian", Cardinality: 7, Number: 2447, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phrathian, Name: "Phrathian", Cardinality: 7, Number: 3196, Perfection: 4, Imperfection: 3},

		// 7.24
		{Scale: scale.Stalian, Name: "Stalian", Cardinality: 7, Number: 2363, Perfection: 4, Imperfection: 3},
		{Scale: scale.Stoptian, Name: "Stoptian", Cardinality: 7, Number: 2524, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zygian, Name: "Zygian", Cardinality: 7, Number: 3812, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kataptian, Name: "Kataptian", Cardinality: 7, Number: 3529, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolaptian, Name: "Aeolaptian", Cardinality: 7, Number: 2963, Perfection: 4, Imperfection: 3},
		{Scale: scale.Pothian, Name: "Pothian", Cardinality: 7, Number: 3662, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bygian, Name: "Bygian", Cardinality: 7, Number: 3229, Perfection: 4, Imperfection: 3},

		// 7.25
		{Scale: scale.Morian, Name: "Morian", Cardinality: 7, Number: 2383, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rycrian, Name: "Rycrian", Cardinality: 7, Number: 2684, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ronian, Name: "Ronian", Cardinality: 7, Number: 2546, Perfection: 4, Imperfection: 3},
		{Scale: scale.Stycrian, Name: "Stycrian", Cardinality: 7, Number: 3988, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katorian, Name: "Katorian", Cardinality: 7, Number: 3881, Perfection: 4, Imperfection: 3},
		{Scale: scale.Epythian, Name: "Epythian", Cardinality: 7, Number: 3667, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kaptian, Name: "Kaptian", Cardinality: 7, Number: 3239, Perfection: 4, Imperfection: 3},

		// 7.26
		{Scale: scale.Phraptian, Name: "Phraptian", Cardinality: 7, Number: 2391, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bacrian, Name: "Bacrian", Cardinality: 7, Number: 2748, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katythian, Name: "Katythian", Cardinality: 7, Number: 2802, Perfection: 4, Imperfection: 3},
		{Scale: scale.Madian, Name: "Madian", Cardinality: 7, Number: 3018, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerygian, Name: "Aerygian", Cardinality: 7, Number: 3882, Perfection: 4, Imperfection: 3},
		{Scale: scale.Pylian, Name: "Pylian", Cardinality: 7, Number: 3669, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionathian, Name: "Ionathian", Cardinality: 7, Number: 3243, Perfection: 4, Imperfection: 3},

		// 7.27
		{Scale: scale.Katocrian, Name: "Katocrian", Cardinality: 7, Number: 2407, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phryptian, Name: "Phryptian", Cardinality: 7, Number: 2876, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katynian, Name: "Katynian", Cardinality: 7, Number: 3314, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerycrian, Name: "Aerycrian", Cardinality: 7, Number: 2533, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ganian, Name: "Ganian", Cardinality: 7, Number: 3884, Perfection: 4, Imperfection: 3},
		{Scale: scale.Eparian, Name: "Eparian", Cardinality: 7, Number: 3673, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lyrian, Name: "Lyrian", Cardinality: 7, Number: 3251, Perfection: 4, Imperfection: 3},

		// 7.28
		{Scale: scale.Ionopian, Name: "Ionopian", Cardinality: 7, Number: 2414, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeologian, Name: "Aeologian", Cardinality: 7, Number: 2932, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zadian, Name: "Zadian", Cardinality: 7, Number: 3538, Perfection: 4, Imperfection: 3},
		{Scale: scale.Sygian, Name: "Sygian", Cardinality: 7, Number: 2981, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phralian, Name: "Phralian", Cardinality: 7, Number: 3734, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phrogian, Name: "Phrogian", Cardinality: 7, Number: 3373, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rathian, Name: "Rathian", Cardinality: 7, Number: 2651, Perfection: 4, Imperfection: 3},

		// 7.29
		{Scale: scale.Rythian, Name: "Rythian", Cardinality: 7, Number: 2422, Perfection: 4, Imperfection: 3},
		{Scale: scale.Donian, Name: "Donian", Cardinality: 7, Number: 2996, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeoloptian, Name: "Aeoloptian", Cardinality: 7, Number: 3794, Perfection: 4, Imperfection: 3},
		{Scale: scale.Panian, Name: "Panian", Cardinality: 7, Number: 3493, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lodian, Name: "Lodian", Cardinality: 7, Number: 2891, Perfection: 4, Imperfection: 3},
		{Scale: scale.Solian, Name: "Solian", Cardinality: 7, Number: 3374, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionolian, Name: "Ionolian", Cardinality: 7, Number: 2653, Perfection: 4, Imperfection: 3},

		// 7.30
		{Scale: scale.Laptian, Name: "Laptian", Cardinality: 7, Number: 2425, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lygian, Name: "Lygian", Cardinality: 7, Number: 3020, Perfection: 4, Imperfection: 3},
		{Scale: scale.Logian, Name: "Logian", Cardinality: 7, Number: 3890, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lalian, Name: "Lalian", Cardinality: 7, Number: 3685, Perfection: 4, Imperfection: 3},
		{Scale: scale.Sothian, Name: "Sothian", Cardinality: 7, Number: 3275, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phrocrian, Name: "Phrocrian", Cardinality: 7, Number: 2455, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thogian, Name: "Thogian", Cardinality: 7, Number: 3260, Perfection: 4, Imperfection: 3},

		// 7.31
		{Scale: scale.Katathian, Name: "Katathian", Cardinality: 7, Number: 2459, Perfection: 4, Imperfection: 3},
		{Scale: scale.Modian, Name: "Modian", Cardinality: 7, Number: 3292, Perfection: 4, Imperfection: 3},
		{Scale: scale.Barian, Name: "Barian", Cardinality: 7, Number: 2489, Perfection: 4, Imperfection: 3},
		{Scale: scale.Mixolocrian, Name: "Mixolocrian", Cardinality: 7, Number: 3532, Perfection: 4, Imperfection: 3},
		{Scale: scale.Sagian, Name: "Sagian", Cardinality: 7, Number: 2969, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolothian, Name: "Aeolothian", Cardinality: 7, Number: 3686, Perfection: 4, Imperfection: 3},
		{Scale: scale.Socrian, Name: "Socrian", Cardinality: 7, Number: 3277, Perfection: 4, Imperfection: 3},

		// 7.32
		{Scale: scale.Tholian, Name: "Tholian", Cardinality: 7, Number: 2461, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ralian, Name: "Ralian", Cardinality: 7, Number: 3308, Perfection: 4, Imperfection: 3},
		{Scale: scale.Syrian, Name: "Syrian", Cardinality: 7, Number: 2521, Perfection: 4, Imperfection: 3},
		{Scale: scale.Stodian, Name: "Stodian", Cardinality: 7, Number: 3788, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionocrian, Name: "Ionocrian", Cardinality: 7, Number: 3481, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zycrian, Name: "Zycrian", Cardinality: 7, Number: 2867, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionygian, Name: "Ionygian", Cardinality: 7, Number: 3278, Perfection: 4, Imperfection: 3},

		// 7.33
		{Scale: scale.Zarian, Name: "Zarian", Cardinality: 7, Number: 2462, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phrythian, Name: "Phrythian", Cardinality: 7, Number: 3316, Perfection: 4, Imperfection: 3},
		{Scale: scale.Rorian, Name: "Rorian", Cardinality: 7, Number: 2537, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bolian, Name: "Bolian", Cardinality: 7, Number: 3916, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bothian, Name: "Bothian", Cardinality: 7, Number: 3737, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katadian, Name: "Katadian", Cardinality: 7, Number: 3379, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kodian, Name: "Kodian", Cardinality: 7, Number: 2663, Perfection: 4, Imperfection: 3},

		// 7.34
		{Scale: scale.Ranian, Name: "Ranian", Cardinality: 7, Number: 2471, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ladian, Name: "Ladian", Cardinality: 7, Number: 3388, Perfection: 4, Imperfection: 3},
		{Scale: scale.Podian, Name: "Podian", Cardinality: 7, Number: 2681, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionothian, Name: "Ionothian", Cardinality: 7, Number: 2534, Perfection: 4, Imperfection: 3},
		{Scale: scale.Kanian, Name: "Kanian", Cardinality: 7, Number: 3892, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zylian, Name: "Zylian", Cardinality: 7, Number: 3689, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zodian, Name: "Zodian", Cardinality: 7, Number: 3283, Perfection: 4, Imperfection: 3},

		// 7.35
		{Scale: scale.Golian, Name: "Golian", Cardinality: 7, Number: 2475, Perfection: 4, Imperfection: 3},
		{Scale: scale.Dyptian, Name: "Dyptian", Cardinality: 7, Number: 3420, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ryphian, Name: "Ryphian", Cardinality: 7, Number: 2745, Perfection: 4, Imperfection: 3},
		{Scale: scale.Gylian, Name: "Gylian", Cardinality: 7, Number: 2790, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolycrian, Name: "Aeolycrian", Cardinality: 7, Number: 2970, Perfection: 4, Imperfection: 3},
		{Scale: scale.Pynian, Name: "Pynian", Cardinality: 7, Number: 3690, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zanian, Name: "Zanian", Cardinality: 7, Number: 3285, Perfection: 4, Imperfection: 3},

		// 7.36
		{Scale: scale.Palian, Name: "Palian", Cardinality: 7, Number: 2477, Perfection: 4, Imperfection: 3},
		{Scale: scale.Stothian, Name: "Stothian", Cardinality: 7, Number: 3436, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerorian, Name: "Aerorian", Cardinality: 7, Number: 2777, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katagian, Name: "Katagian", Cardinality: 7, Number: 2918, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phronian, Name: "Phronian", Cardinality: 7, Number: 3482, Perfection: 4, Imperfection: 3},
		{Scale: scale.Banian, Name: "Banian", Cardinality: 7, Number: 2869, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeronian, Name: "Aeronian", Cardinality: 7, Number: 3286, Perfection: 4, Imperfection: 3},

		// 7.37
		{Scale: scale.Loptian, Name: "Loptian", Cardinality: 7, Number: 2483, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionodian, Name: "Ionodian", Cardinality: 7, Number: 3484, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bogian, Name: "Bogian", Cardinality: 7, Number: 2873, Perfection: 4, Imperfection: 3},
		{Scale: scale.Mogian, Name: "Mogian", Cardinality: 7, Number: 3302, Perfection: 4, Imperfection: 3},
		{Scale: scale.Docrian, Name: "Docrian", Cardinality: 7, Number: 2509, Perfection: 4, Imperfection: 3},
		{Scale: scale.Epadian, Name: "Epadian", Cardinality: 7, Number: 3692, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerynian, Name: "Aerynian", Cardinality: 7, Number: 3289, Perfection: 4, Imperfection: 3},

		// 7.38
		{Scale: scale.Bycrian, Name: "Bycrian", Cardinality: 7, Number: 2485, Perfection: 4, Imperfection: 3},
		{Scale: scale.Pathian, Name: "Pathian", Cardinality: 7, Number: 3500, Perfection: 4, Imperfection: 3},
		{Scale: scale.Mydian, Name: "Mydian", Cardinality: 7, Number: 2905, Perfection: 4, Imperfection: 3},
		{Scale: scale.Thyptian, Name: "Thyptian", Cardinality: 7, Number: 3430, Perfection: 4, Imperfection: 3},
		{Scale: scale.Phrothian, Name: "Phrothian", Cardinality: 7, Number: 2765, Perfection: 4, Imperfection: 3},
		{Scale: scale.Katycrian, Name: "Katycrian", Cardinality: 7, Number: 2870, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionalian, Name: "Ionalian", Cardinality: 7, Number: 3290, Perfection: 4, Imperfection: 3},

		// 7.39
		{Scale: scale.Dolian, Name: "Dolian", Cardinality: 7, Number: 2517, Perfection: 4, Imperfection: 3},
		{Scale: scale.Porian, Name: "Porian", Cardinality: 7, Number: 3756, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aerylian, Name: "Aerylian", Cardinality: 7, Number: 3417, Perfection: 4, Imperfection: 3},
		{Scale: scale.Zagian, Name: "Zagian", Cardinality: 7, Number: 2739, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lagian, Name: "Lagian", Cardinality: 7, Number: 2766, Perfection: 4, Imperfection: 3},
		{Scale: scale.Tyrian, Name: "Tyrian", Cardinality: 7, Number: 2874, Perfection: 4, Imperfection: 3},
		{Scale: scale.Mixonorian, Name: "Mixonorian", Cardinality: 7, Number: 3306, Perfection: 4, Imperfection: 3},

		// 7.40
		{Scale: scale.Pagian, Name: "Pagian", Cardinality: 7, Number: 2538, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeolythian, Name: "Aeolythian", Cardinality: 7, Number: 3924, Perfection: 4, Imperfection: 3},
		{Scale: scale.Molian, Name: "Molian", Cardinality: 7, Number: 3753, Perfection: 4, Imperfection: 3},
		{Scale: scale.Staptian, Name: "Staptian", Cardinality: 7, Number: 3411, Perfection: 4, Imperfection: 3},
		{Scale: scale.Mothian, Name: "Mothian", Cardinality: 7, Number: 2727, Perfection: 4, Imperfection: 3},
		{Scale: scale.Aeranian, Name: "Aeranian", Cardinality: 7, Number: 2718, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ragian, Name: "Ragian", Cardinality: 7, Number: 2682, Perfection: 4, Imperfection: 3},

		// 7.41
		{Scale: scale.Larian, Name: "Larian", Cardinality: 7, Number: 2733, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lythian, Name: "Lythian", Cardinality: 7, Number: 2742, Perfection: 4, Imperfection: 3},
		{Scale: scale.Stydian, Name: "Stydian", Cardinality: 7, Number: 2778, Perfection: 4, Imperfection: 3},
		{Scale: scale.Lorian, Name: "Lorian", Cardinality: 7, Number: 2922, Perfection: 4, Imperfection: 3},
		{Scale: scale.Ionadian, Name: "Ionadian", Cardinality: 7, Number: 3498, Perfection: 4, Imperfection: 3},
		{Scale: scale.Bocrian, Name: "Bocrian", Cardinality: 7, Number: 2901, Perfection: 4, Imperfection: 3},
		{Scale: scale.Mixolythian, Name: "Mixolythian", Cardinality: 7, Number: 3414, Perfection: 4, Imperfection: 3},

		// 7.42
		{Scale: scale.Thadian, Name: "Thadian", Cardinality: 7, Number: 2207, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sanian, Name: "Sanian", Cardinality: 7, Number: 2552, Perfection: 3, Imperfection: 4},
		{Scale: scale.Ionydian, Name: "Ionydian", Cardinality: 7, Number: 4036, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epydian, Name: "Epydian", Cardinality: 7, Number: 3977, Perfection: 3, Imperfection: 4},
		{Scale: scale.Katydian, Name: "Katydian", Cardinality: 7, Number: 3859, Perfection: 3, Imperfection: 4},
		{Scale: scale.Mathian, Name: "Mathian", Cardinality: 7, Number: 3623, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aeryptian, Name: "Aeryptian", Cardinality: 7, Number: 3151, Perfection: 3, Imperfection: 4},

		// 7.43
		{Scale: scale.Pythian, Name: "Pythian", Cardinality: 7, Number: 2223, Perfection: 3, Imperfection: 4},
		{Scale: scale.Katylian, Name: "Katylian", Cardinality: 7, Number: 2808, Perfection: 3, Imperfection: 4},
		{Scale: scale.Bydian, Name: "Bydian", Cardinality: 7, Number: 3042, Perfection: 3, Imperfection: 4},
		{Scale: scale.Bynian, Name: "Bynian", Cardinality: 7, Number: 3978, Perfection: 3, Imperfection: 4},
		{Scale: scale.Galian, Name: "Galian", Cardinality: 7, Number: 3861, Perfection: 3, Imperfection: 4},
		{Scale: scale.Zonian, Name: "Zonian", Cardinality: 7, Number: 3627, Perfection: 3, Imperfection: 4},
		{Scale: scale.Myrian, Name: "Myrian", Cardinality: 7, Number: 3159, Perfection: 3, Imperfection: 4},

		// 7.44
		{Scale: scale.Katogian, Name: "Katogian", Cardinality: 7, Number: 2235, Perfection: 3, Imperfection: 4},
		{Scale: scale.Stacrian, Name: "Stacrian", Cardinality: 7, Number: 3000, Perfection: 3, Imperfection: 4},
		{Scale: scale.Styrian, Name: "Styrian", Cardinality: 7, Number: 3810, Perfection: 3, Imperfection: 4},
		{Scale: scale.Ionyrian, Name: "Ionyrian", Cardinality: 7, Number: 3525, Perfection: 3, Imperfection: 4},
		{Scale: scale.Phrodian, Name: "Phrodian", Cardinality: 7, Number: 2955, Perfection: 3, Imperfection: 4},
		{Scale: scale.Pycrian, Name: "Pycrian", Cardinality: 7, Number: 3630, Perfection: 3, Imperfection: 4},
		{Scale: scale.Gyptian, Name: "Gyptian", Cardinality: 7, Number: 3165, Perfection: 3, Imperfection: 4},

		// 7.45
		{Scale: scale.Katacrian, Name: "Katacrian", Cardinality: 7, Number: 2286, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sodian, Name: "Sodian", Cardinality: 7, Number: 3816, Perfection: 3, Imperfection: 4},
		{Scale: scale.Bathian, Name: "Bathian", Cardinality: 7, Number: 3537, Perfection: 3, Imperfection: 4},
		{Scale: scale.Mylian, Name: "Mylian", Cardinality: 7, Number: 2979, Perfection: 3, Imperfection: 4},
		{Scale: scale.Godian, Name: "Godian", Cardinality: 7, Number: 3726, Perfection: 3, Imperfection: 4},
		{Scale: scale.Thorian, Name: "Thorian", Cardinality: 7, Number: 3357, Perfection: 3, Imperfection: 4},
		{Scale: scale.Zocrian, Name: "Zocrian", Cardinality: 7, Number: 2619, Perfection: 3, Imperfection: 4},

		// 7.46
		{Scale: scale.Stanian, Name: "Stanian", Cardinality: 7, Number: 2298, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epanian, Name: "Epanian", Cardinality: 7, Number: 4008, Perfection: 3, Imperfection: 4},
		{Scale: scale.Konian, Name: "Konian", Cardinality: 7, Number: 3921, Perfection: 3, Imperfection: 4},
		{Scale: scale.Stocrian, Name: "Stocrian", Cardinality: 7, Number: 3747, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kalian, Name: "Kalian", Cardinality: 7, Number: 3399, Perfection: 3, Imperfection: 4},
		{Scale: scale.Phroptian, Name: "Phroptian", Cardinality: 7, Number: 2703, Perfection: 3, Imperfection: 4},
		{Scale: scale.Dydian, Name: "Dydian", Cardinality: 7, Number: 2622, Perfection: 3, Imperfection: 4},

		// 7.47
		{Scale: scale.Katyptian, Name: "Katyptian", Cardinality: 7, Number: 2300, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epodian, Name: "Epodian", Cardinality: 7, Number: 4040, Perfection: 3, Imperfection: 4},
		{Scale: scale.Mygian, Name: "Mygian", Cardinality: 7, Number: 3985, Perfection: 3, Imperfection: 4},
		{Scale: scale.Pacrian, Name: "Pacrian", Cardinality: 7, Number: 3875, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aerocrian, Name: "Aerocrian", Cardinality: 7, Number: 3655, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aeolarian, Name: "Aeolarian", Cardinality: 7, Number: 3215, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kythian, Name: "Kythian", Cardinality: 7, Number: 2335, Perfection: 3, Imperfection: 4},

		// 7.48
		{Scale: scale.Bonian, Name: "Bonian", Cardinality: 7, Number: 2351, Perfection: 3, Imperfection: 4},
		{Scale: scale.Badian, Name: "Badian", Cardinality: 7, Number: 2428, Perfection: 3, Imperfection: 4},
		{Scale: scale.Katodian, Name: "Katodian", Cardinality: 7, Number: 3044, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sadian, Name: "Sadian", Cardinality: 7, Number: 3986, Perfection: 3, Imperfection: 4},
		{Scale: scale.Dothian, Name: "Dothian", Cardinality: 7, Number: 3877, Perfection: 3, Imperfection: 4},
		{Scale: scale.Moptian, Name: "Moptian", Cardinality: 7, Number: 3659, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aeryrian, Name: "Aeryrian", Cardinality: 7, Number: 3223, Perfection: 3, Imperfection: 4},

		// 7.49
		{Scale: scale.Epagian, Name: "Epagian", Cardinality: 7, Number: 2359, Perfection: 3, Imperfection: 4},
		{Scale: scale.Raptian, Name: "Raptian", Cardinality: 7, Number: 2492, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epolian, Name: "Epolian", Cardinality: 7, Number: 3556, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sythian, Name: "Sythian", Cardinality: 7, Number: 3017, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sydian, Name: "Sydian", Cardinality: 7, Number: 3878, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epocrian, Name: "Epocrian", Cardinality: 7, Number: 3661, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kylian, Name: "Kylian", Cardinality: 7, Number: 3227, Perfection: 3, Imperfection: 4},

		// 7.50
		{Scale: scale.Gacrian, Name: "Gacrian", Cardinality: 7, Number: 2365, Perfection: 3, Imperfection: 4},
		{Scale: scale.Borian, Name: "Borian", Cardinality: 7, Number: 2540, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sycrian, Name: "Sycrian", Cardinality: 7, Number: 3940, Perfection: 3, Imperfection: 4},
		{Scale: scale.Gadian, Name: "Gadian", Cardinality: 7, Number: 3785, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aeolocrian, Name: "Aeolocrian", Cardinality: 7, Number: 3475, Perfection: 3, Imperfection: 4},
		{Scale: scale.Mixodorian, Name: "Mixodorian", Cardinality: 7, Number: 2855, Perfection: 3, Imperfection: 4},
		{Scale: scale.Dathian, Name: "Dathian", Cardinality: 7, Number: 3230, Perfection: 3, Imperfection: 4},

		// 7.51
		{Scale: scale.Katoptian, Name: "Katoptian", Cardinality: 7, Number: 2366, Perfection: 3, Imperfection: 4},
		{Scale: scale.Ponian, Name: "Ponian", Cardinality: 7, Number: 2548, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kadian, Name: "Kadian", Cardinality: 7, Number: 4004, Perfection: 3, Imperfection: 4},
		{Scale: scale.Gynian, Name: "Gynian", Cardinality: 7, Number: 3913, Perfection: 3, Imperfection: 4},
		{Scale: scale.Thyphian, Name: "Thyphian", Cardinality: 7, Number: 3731, Perfection: 3, Imperfection: 4},
		{Scale: scale.Polian, Name: "Polian", Cardinality: 7, Number: 3367, Perfection: 3, Imperfection: 4},
		{Scale: scale.Thanian, Name: "Thanian", Cardinality: 7, Number: 2639, Perfection: 3, Imperfection: 4},

		// 7.52
		{Scale: scale.Epacrian, Name: "Epacrian", Cardinality: 7, Number: 2397, Perfection: 3, Imperfection: 4},
		{Scale: scale.Sathian, Name: "Sathian", Cardinality: 7, Number: 2796, Perfection: 3, Imperfection: 4},
		{Scale: scale.Lathian, Name: "Lathian", Cardinality: 7, Number: 2994, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aeralian, Name: "Aeralian", Cardinality: 7, Number: 3786, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kynian, Name: "Kynian", Cardinality: 7, Number: 3477, Perfection: 3, Imperfection: 4},
		{Scale: scale.Stynian, Name: "Stynian", Cardinality: 7, Number: 2859, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epyphian, Name: "Epyphian", Cardinality: 7, Number: 3246, Perfection: 3, Imperfection: 4},

		// 7.53
		{Scale: scale.Pogian, Name: "Pogian", Cardinality: 7, Number: 2413, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aeraptian, Name: "Aeraptian", Cardinality: 7, Number: 2924, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epylian, Name: "Epylian", Cardinality: 7, Number: 3506, Perfection: 3, Imperfection: 4},
		{Scale: scale.Gamian, Name: "Gamian", Cardinality: 7, Number: 2917, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kathian, Name: "Kathian", Cardinality: 7, Number: 3478, Perfection: 3, Imperfection: 4},
		{Scale: scale.Lylian, Name: "Lylian", Cardinality: 7, Number: 2861, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epalian, Name: "Epalian", Cardinality: 7, Number: 3254, Perfection: 3, Imperfection: 4},

		// 7.54
		{Scale: scale.Eporian, Name: "Eporian", Cardinality: 7, Number: 2421, Perfection: 3, Imperfection: 4},
		{Scale: scale.Rylian, Name: "Rylian", Cardinality: 7, Number: 2988, Perfection: 3, Imperfection: 4},
		{Scale: scale.Epaptian, Name: "Epaptian", Cardinality: 7, Number: 3762, Perfection: 3, Imperfection: 4},
		{Scale: scale.Byrian, Name: "Byrian", Cardinality: 7, Number: 3429, Perfection: 3, Imperfection: 4},
		{Scale: scale.Katanian, Name: "Katanian", Cardinality: 7, Number: 2763, Perfection: 3, Imperfection: 4},
		{Scale: scale.Katyrian, Name: "Katyrian", Cardinality: 7, Number: 2862, Perfection: 3, Imperfection: 4},
		{Scale: scale.Rynian, Name: "Rynian", Cardinality: 7, Number: 3258, Perfection: 3, Imperfection: 4},

		// 7.55
		{Scale: scale.Korian, Name: "Korian", Cardinality: 7, Number: 2478, Perfection: 3, Imperfection: 4},
		{Scale: scale.Lynian, Name: "Lynian", Cardinality: 7, Number: 3444, Perfection: 3, Imperfection: 4},
		{Scale: scale.Malian, Name: "Malian", Cardinality: 7, Number: 2793, Perfection: 3, Imperfection: 4},
		{Scale: scale.Synian, Name: "Synian", Cardinality: 7, Number: 2982, Perfection: 3, Imperfection: 4},
		{Scale: scale.Phragian, Name: "Phragian", Cardinality: 7, Number: 3738, Perfection: 3, Imperfection: 4},
		{Scale: scale.Manian, Name: "Manian", Cardinality: 7, Number: 3381, Perfection: 3, Imperfection: 4},
		{Scale: scale.Marian, Name: "Marian", Cardinality: 7, Number: 2667, Perfection: 3, Imperfection: 4},

		// 7.56
		{Scale: scale.Mycrian, Name: "Mycrian", Cardinality: 7, Number: 2486, Perfection: 3, Imperfection: 4},
		{Scale: scale.Ionorian, Name: "Ionorian", Cardinality: 7, Number: 3508, Perfection: 3, Imperfection: 4},
		{Scale: scale.Phrydian, Name: "Phrydian", Cardinality: 7, Number: 2921, Perfection: 3, Imperfection: 4},
		{Scale: scale.Zyptian, Name: "Zyptian", Cardinality: 7, Number: 3494, Perfection: 3, Imperfection: 4},
		{Scale: scale.Katothian, Name: "Katothian", Cardinality: 7, Number: 2893, Perfection: 3, Imperfection: 4},
		{Scale: scale.Phrylian, Name: "Phrylian", Cardinality: 7, Number: 3382, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kocrian, Name: "Kocrian", Cardinality: 7, Number: 2669, Perfection: 3, Imperfection: 4},

		// 7.57
		{Scale: scale.Ionanian, Name: "Ionanian", Cardinality: 7, Number: 2490, Perfection: 3, Imperfection: 4},
		{Scale: scale.Aerothian, Name: "Aerothian", Cardinality: 7, Number: 3540, Perfection: 3, Imperfection: 4},
		{Scale: scale.Stagian, Name: "Stagian", Cardinality: 7, Number: 2985, Perfection: 3, Imperfection: 4},
		{Scale: scale.Lothian, Name: "Lothian", Cardinality: 7, Number: 3750, Perfection: 3, Imperfection: 4},
		{Scale: scale.Phrycrian, Name: "Phrycrian", Cardinality: 7, Number: 3405, Perfection: 3, Imperfection: 4},
		{Scale: scale.Kyptian, Name: "Kyptian", Cardinality: 7, Number: 2715, Perfection: 3, Imperfection: 4},
		{Scale: scale.Ionylian, Name: "Ionylian", Cardinality: 7, Number: 2670, Perfection: 3, Imperfection: 4},

		// 7.58
		{Scale: scale.Gydian, Name: "Gydian", Cardinality: 7, Number: 2238, Perfection: 2, Imperfection: 5},
		{Scale: scale.Kogian, Name: "Kogian", Cardinality: 7, Number: 3048, Perfection: 2, Imperfection: 5},
		{Scale: scale.Rarian, Name: "Rarian", Cardinality: 7, Number: 4002, Perfection: 2, Imperfection: 5},
		{Scale: scale.Aerolian, Name: "Aerolian", Cardinality: 7, Number: 3909, Perfection: 2, Imperfection: 5},
		{Scale: scale.Karian, Name: "Karian", Cardinality: 7, Number: 3723, Perfection: 2, Imperfection: 5},
		{Scale: scale.Myptian, Name: "Myptian", Cardinality: 7, Number: 3351, Perfection: 2, Imperfection: 5},
		{Scale: scale.Rydian, Name: "Rydian", Cardinality: 7, Number: 2607, Perfection: 2, Imperfection: 5},

		// 7.59
		{Scale: scale.Aeolynian, Name: "Aeolynian", Cardinality: 7, Number: 2731, Perfection: 2, Imperfection: 5},
		{Scale: scale.Aeroptian, Name: "Aeroptian", Cardinality: 7, Number: 2734, Perfection: 2, Imperfection: 5},
		{Scale: scale.Phryrian, Name: "Phryrian", Cardinality: 7, Number: 2746, Perfection: 2, Imperfection: 5},
		{Scale: scale.Gothian, Name: "Gothian", Cardinality: 7, Number: 2794, Perfection: 2, Imperfection: 5},
		{Scale: scale.Storian, Name: "Storian", Cardinality: 7, Number: 2986, Perfection: 2, Imperfection: 5},
		{Scale: scale.Pyptian, Name: "Pyptian", Cardinality: 7, Number: 3754, Perfection: 2, Imperfection: 5},
		{Scale: scale.Thydian, Name: "Thydian", Cardinality: 7, Number: 3413, Perfection: 2, Imperfection: 5},
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
