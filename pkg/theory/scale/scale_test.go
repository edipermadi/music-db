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
		Scale  scale.Type
		Name   string
		Number int
	}

	testCases := []testCase{
		{scale.Invalid, "InvalidScale", 0},

		// 3 notes
		{scale.Minoric, "Minoric", 2184},

		// 4 notes

		// 4.1
		{scale.Thaptic, "Thaptic", 2193},
		{scale.Lothic, "Lothic", 2328},
		{scale.Phratic, "Phratic", 2244},
		{scale.Aerathic, "Aerathic", 3144},

		// 4.2
		{scale.Epathic, "Epathic", 2196},
		{scale.Mynic, "Mynic", 2376},
		{scale.Rothic, "Rothic", 2628},
		{scale.Eporic, "Eporic", 2322},

		// 4.3
		{Scale: scale.Zyphic, Name: "Zyphic", Number: 2185},
		{Scale: scale.Epogic, Name: "Epogic", Number: 2200},
		{Scale: scale.Lanic, Name: "Lanic", Number: 2440},
		{Scale: scale.Pyrric, Name: "Pyrric", Number: 3140},

		// 4.4
		{Scale: scale.Aeoloric, Name: "Aeoloric", Number: 2188},
		{Scale: scale.Gonic, Name: "Gonic", Number: 2248},
		{Scale: scale.Dalic, Name: "Dalic", Number: 3208},
		{Scale: scale.Dygic, Name: "Dygic", Number: 2321},

		// 4.5
		{Scale: scale.Daric, Name: "Daric", Number: 2194},
		{Scale: scale.Lonic, Name: "Lonic", Number: 2344},
		{Scale: scale.Phradic, Name: "Phradic", Number: 2372},
		{Scale: scale.Bolic, Name: "Bolic", Number: 2596},

		// 4.6
		{Scale: scale.Saric, Name: "Saric", Number: 2212},
		{Scale: scale.Zoptic, Name: "Zoptic", Number: 2632},
		{Scale: scale.Aeraphic, Name: "Aeraphic", Number: 2338},
		{Scale: scale.Byptic, Name: "Byptic", Number: 2324},

		// 4.7
		{Scale: scale.Aeolic, Name: "Aeolic", Number: 2186},
		{Scale: scale.Koptic, Name: "Koptic", Number: 2216},
		{Scale: scale.Mixolyric, Name: "Mixolyric", Number: 2696},
		{Scale: scale.Lydic, Name: "Lydic", Number: 2594},

		// 4.8
		{Scale: scale.Stathic, Name: "Stathic", Number: 2210},
		{Scale: scale.Dadic, Name: "Dadic", Number: 2600},

		// 4.9
		{Scale: scale.Phrynic, Name: "Phrynic", Number: 2340},

		// 5 notes

		// 5.1
		{Scale: scale.Epathitonic, Name: "Epathitonic", Number: 2378},
		{Scale: scale.Mynitonic, Name: "Mynitonic", Number: 2644},
		{Scale: scale.Rocritonic, Name: "Rocritonic", Number: 2386},
		{Scale: scale.Pentatonic, Name: "Pentatonic", Number: 2708},
		{Scale: scale.Thaptitonic, Name: "Thaptitonic", Number: 2642},

		// 5.2
		{Scale: scale.Magitonic, Name: "Magitonic", Number: 2197},
		{Scale: scale.Daditonic, Name: "Daditonic", Number: 2392},
		{Scale: scale.Aeolyphritonic, Name: "Aeolyphritonic", Number: 2756},
		{Scale: scale.Gycritonic, Name: "Gycritonic", Number: 2834},
		{Scale: scale.Pyritonic, Name: "Pyritonic", Number: 3146},

		// 5.3
		{Scale: scale.Gathitonic, Name: "Gathitonic", Number: 2213},
		{Scale: scale.Ionitonic, Name: "Ionitonic", Number: 2648},
		{Scale: scale.Phrynitonic, Name: "Phrynitonic", Number: 2402},
		{Scale: scale.Stathitonic, Name: "Stathitonic", Number: 2836},
		{Scale: scale.Thalitonic, Name: "Thalitonic", Number: 3154},

		// 5.4
		{Scale: scale.Zolitonic, Name: "Zolitonic", Number: 2225},
		{Scale: scale.Epogitonic, Name: "Epogitonic", Number: 2840},
		{Scale: scale.Lanitonic, Name: "Lanitonic", Number: 3170},
		{Scale: scale.Paptitonic, Name: "Paptitonic", Number: 2245},
		{Scale: scale.Ionacritonic, Name: "Ionacritonic", Number: 3160},

		// 5.5
		{Scale: scale.Phraditonic, Name: "Phraditonic", Number: 2246},
		{Scale: scale.Aeoloritonic, Name: "Aeoloritonic", Number: 3176},
		{Scale: scale.Gonitonic, Name: "Gonitonic", Number: 2257},
		{Scale: scale.Dalitonic, Name: "Dalitonic", Number: 3352},
		{Scale: scale.Dygitonic, Name: "Dygitonic", Number: 2609},

		// 5.6
		{Scale: scale.Aeracritonic, Name: "Aeracritonic", Number: 2258},
		{Scale: scale.Byptitonic, Name: "Byptitonic", Number: 3368},
		{Scale: scale.Daritonic, Name: "Daritonic", Number: 2641},
		{Scale: scale.Lonitonic, Name: "Lonitonic", Number: 2374},
		{Scale: scale.Ionycritonic, Name: "Ionycritonic", Number: 2612},

		// 5.7
		{Scale: scale.Lothitonic, Name: "Lothitonic", Number: 2260},
		{Scale: scale.Phratonic, Name: "Phratonic", Number: 3400},
		{Scale: scale.Aerathitonic, Name: "Aerathitonic", Number: 2705},
		{Scale: scale.Saritonic, Name: "Saritonic", Number: 2630},
		{Scale: scale.Zoptitonic, Name: "Zoptitonic", Number: 2330},

		// 5.8
		{Scale: scale.Dolitonic, Name: "Dolitonic", Number: 2189},
		{Scale: scale.Poritonic, Name: "Poritonic", Number: 2264},
		{Scale: scale.Aerylitonic, Name: "Aerylitonic", Number: 3464},
		{Scale: scale.Zagitonic, Name: "Zagitonic", Number: 2833},
		{Scale: scale.Lagitonic, Name: "Lagitonic", Number: 3142},

		// 5.9
		{Scale: scale.Molitonic, Name: "Molitonic", Number: 2195},
		{Scale: scale.Staptitonic, Name: "Staptitonic", Number: 2360},
		{Scale: scale.Mothitonic, Name: "Mothitonic", Number: 2500},
		{Scale: scale.Aeritonic, Name: "Aeritonic", Number: 3620},
		{Scale: scale.Ragitonic, Name: "Ragitonic", Number: 3145},

		// 5.10
		{Scale: scale.Ionaditonic, Name: "Ionaditonic", Number: 2198},
		{Scale: scale.Bocritonic, Name: "Bocritonic", Number: 2408},
		{Scale: scale.Gythitonic, Name: "Gythitonic", Number: 2884},
		{Scale: scale.Pagitonic, Name: "Pagitonic", Number: 3346},
		{Scale: scale.Aeolythitonic, Name: "Aeolythitonic", Number: 2597},

		// 5.11
		{Scale: scale.Zacritonic, Name: "Zacritonic", Number: 2201},
		{Scale: scale.Laritonic, Name: "Laritonic", Number: 2456},
		{Scale: scale.Thacritonic, Name: "Thacritonic", Number: 3268},
		{Scale: scale.Styditonic, Name: "Styditonic", Number: 2441},
		{Scale: scale.Loritonic, Name: "Loritonic", Number: 3148},

		// 5.12
		{Scale: scale.Aeolyritonic, Name: "Aeolyritonic", Number: 2204},
		{Scale: scale.Goritonic, Name: "Goritonic", Number: 2504},
		{Scale: scale.Aeoloditonic, Name: "Aeoloditonic", Number: 3652},
		{Scale: scale.Doptitonic, Name: "Doptitonic", Number: 3209},
		{Scale: scale.Aeraphitonic, Name: "Aeraphitonic", Number: 2323},

		// 5.13
		{Scale: scale.Zathitonic, Name: "Zathitonic", Number: 2211},
		{Scale: scale.Raditonic, Name: "Raditonic", Number: 2616},
		{Scale: scale.Stonitonic, Name: "Stonitonic", Number: 2274},
		{Scale: scale.Syptitonic, Name: "Syptitonic", Number: 3624},
		{Scale: scale.Ionythitonic, Name: "Ionythitonic", Number: 3153},

		// 5.14
		{Scale: scale.Aeolanitonic, Name: "Aeolanitonic", Number: 2217},
		{Scale: scale.Danitonic, Name: "Danitonic", Number: 2712},
		{Scale: scale.Ionaritonic, Name: "Ionaritonic", Number: 2658},
		{Scale: scale.Dynitonic, Name: "Dynitonic", Number: 2442},
		{Scale: scale.Zyditonic, Name: "Zyditonic", Number: 3156},

		// 5.15
		{Scale: scale.Aeolacritonic, Name: "Aeolacritonic", Number: 2228},
		{Scale: scale.Zythitonic, Name: "Zythitonic", Number: 2888},
		{Scale: scale.Dyritonic, Name: "Dyritonic", Number: 3362},
		{Scale: scale.Koptitonic, Name: "Koptitonic", Number: 2629},
		{Scale: scale.Thocritonic, Name: "Thocritonic", Number: 2326},

		// 5.16
		{Scale: scale.Lycritonic, Name: "Lycritonic", Number: 2249},
		{Scale: scale.Daptitonic, Name: "Daptitonic", Number: 3224},
		{Scale: scale.Kygitonic, Name: "Kygitonic", Number: 2353},
		{Scale: scale.Mocritonic, Name: "Mocritonic", Number: 2444},
		{Scale: scale.Zynitonic, Name: "Zynitonic", Number: 3172},

		// 5.17
		{Scale: scale.Epygitonic, Name: "Epygitonic", Number: 2250},
		{Scale: scale.Zaptitonic, Name: "Zaptitonic", Number: 3240},
		{Scale: scale.Kagitonic, Name: "Kagitonic", Number: 2385},
		{Scale: scale.Zogitonic, Name: "Zogitonic", Number: 2700},
		{Scale: scale.Epyritonic, Name: "Epyritonic", Number: 2610},

		// 5.18
		{Scale: scale.Zothitonic, Name: "Zothitonic", Number: 2252},
		{Scale: scale.Phrolitonic, Name: "Phrolitonic", Number: 3272},
		{Scale: scale.Ionagitonic, Name: "Ionagitonic", Number: 2449},
		{Scale: scale.Aeolapritonic, Name: "Aeolapritonic", Number: 3212},
		{Scale: scale.Kyritonic, Name: "Kyritonic", Number: 2329},

		// 5.19
		{Scale: scale.Ionyptitonic, Name: "Ionyptitonic", Number: 2276},
		{Scale: scale.Gyritonic, Name: "Gyritonic", Number: 3656},
		{Scale: scale.Zalitonic, Name: "Zalitonic", Number: 3217},
		{Scale: scale.Stolitonic, Name: "Stolitonic", Number: 2339},
		{Scale: scale.Bylitonic, Name: "Bylitonic", Number: 2332},

		// 5.20
		{Scale: scale.Thoditonic, Name: "Thoditonic", Number: 2345},
		{Scale: scale.Dogitonic, Name: "Dogitonic", Number: 2380},
		{Scale: scale.Phralitonic, Name: "Phralitonic", Number: 2660},
		{Scale: scale.Garitonic, Name: "Garitonic", Number: 2450},
		{Scale: scale.Soptitonic, Name: "Soptitonic", Number: 3220},

		// 5.21
		{Scale: scale.Kataritonic, Name: "Kataritonic", Number: 2346},
		{Scale: scale.Sylitonic, Name: "Sylitonic", Number: 2388},
		{Scale: scale.Thonitonic, Name: "Thonitonic", Number: 2724},
		{Scale: scale.Phropitonic, Name: "Phropitonic", Number: 2706},
		{Scale: scale.Staditonic, Name: "Staditonic", Number: 2634},

		// 5.22
		{Scale: scale.Lyditonic, Name: "Lyditonic", Number: 2354},
		{Scale: scale.Mythitonic, Name: "Mythitonic", Number: 2452},
		{Scale: scale.Sogitonic, Name: "Sogitonic", Number: 3236},
		{Scale: scale.Gothitonic, Name: "Gothitonic", Number: 2377},
		{Scale: scale.Rothitonic, Name: "Rothitonic", Number: 2636},

		// 5.23
		{Scale: scale.Zylitonic, Name: "Zylitonic", Number: 2187},
		{Scale: scale.Zoditonic, Name: "Zoditonic", Number: 2232},
		{Scale: scale.Zaritonic, Name: "Zaritonic", Number: 2952},
		{Scale: scale.Phrythitonic, Name: "Phrythitonic", Number: 3618},
		{Scale: scale.Rolitonic, Name: "Rolitonic", Number: 3141},

		// 5.24
		{Scale: scale.Ranitonic, Name: "Ranitonic", Number: 2190},
		{Scale: scale.Laditonic, Name: "Laditonic", Number: 2280},
		{Scale: scale.Poditonic, Name: "Poditonic", Number: 3720},
		{Scale: scale.Ionothitonic, Name: "Ionothitonic", Number: 3345},
		{Scale: scale.Kanitonic, Name: "Kanitonic", Number: 2595},

		// 5.25
		{Scale: scale.Ryphitonic, Name: "Ryphitonic", Number: 2202},
		{Scale: scale.Gylitonic, Name: "Gylitonic", Number: 2472},
		{Scale: scale.Aeolycritonic, Name: "Aeolycritonic", Number: 3396},
		{Scale: scale.Pynitonic, Name: "Pynitonic", Number: 2697},
		{Scale: scale.Zanitonic, Name: "Zanitonic", Number: 2598},

		// 5.26
		{Scale: scale.Phronitonic, Name: "Phronitonic", Number: 2214},
		{Scale: scale.Banitonic, Name: "Banitonic", Number: 2664},
		{Scale: scale.Aeronitonic, Name: "Aeronitonic", Number: 2466},
		{Scale: scale.Golitonic, Name: "Golitonic", Number: 3348},
		{Scale: scale.Dyptitonic, Name: "Dyptitonic", Number: 2601},

		// 5.27
		{Scale: scale.Aerynitonic, Name: "Aerynitonic", Number: 2220},
		{Scale: scale.Palitonic, Name: "Palitonic", Number: 2760},
		{Scale: scale.Stothitonic, Name: "Stothitonic", Number: 2850},
		{Scale: scale.Aerophitonic, Name: "Aerophitonic", Number: 3210},
		{Scale: scale.Katagitonic, Name: "Katagitonic", Number: 2325},

		// 5.28
		{Scale: scale.Ionoditonic, Name: "Ionoditonic", Number: 2226},
		{Scale: scale.Bogitonic, Name: "Bogitonic", Number: 2856},
		{Scale: scale.Mogitonic, Name: "Mogitonic", Number: 3234},
		{Scale: scale.Docritonic, Name: "Docritonic", Number: 2373},
		{Scale: scale.Epaditonic, Name: "Epaditonic", Number: 2604},

		// 5.29
		{Scale: scale.Mixitonic, Name: "Mixitonic", Number: 2341},
		{Scale: scale.Phrothitonic, Name: "Phrothitonic", Number: 2348},
		{Scale: scale.Katycritonic, Name: "Katycritonic", Number: 2404},
		{Scale: scale.Ionalitonic, Name: "Ionalitonic", Number: 2852},
		{Scale: scale.Loptitonic, Name: "Loptitonic", Number: 3218},

		// 5.30
		{Scale: scale.Thyritonic, Name: "Thyritonic", Number: 2342},
		{Scale: scale.Thoptitonic, Name: "Thoptitonic", Number: 2356},
		{Scale: scale.Bycritonic, Name: "Bycritonic", Number: 2468},
		{Scale: scale.Pathitonic, Name: "Pathitonic", Number: 3364},
		{Scale: scale.Myditonic, Name: "Myditonic", Number: 2633},

		// 5.31
		{Scale: scale.Bolitonic, Name: "Bolitonic", Number: 2218},
		{Scale: scale.Bothitonic, Name: "Bothitonic", Number: 2728},
		{Scale: scale.Kataditonic, Name: "Kataditonic", Number: 2722},
		{Scale: scale.Koditonic, Name: "Koditonic", Number: 2698},
		{Scale: scale.Tholitonic, Name: "Tholitonic", Number: 2602},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Name, tc.Scale.String())
			assert.Equal(t, tc.Number, tc.Scale.Number())
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
