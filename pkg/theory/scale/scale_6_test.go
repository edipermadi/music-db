package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith6Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
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
