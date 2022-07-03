package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith7Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
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
