package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith9Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		// 9.1
		{Scale: scale.Aerycrygic, Name: "Aerycrygic", Cardinality: 9, Number: 2807, Perfection: 8, Imperfection: 1},
		{Scale: scale.Gadygic, Name: "Gadygic", Cardinality: 9, Number: 3038, Perfection: 8, Imperfection: 1},
		{Scale: scale.Solygic, Name: "Solygic", Cardinality: 9, Number: 3962, Perfection: 8, Imperfection: 1},
		{Scale: scale.Zylygic, Name: "Zylygic", Cardinality: 9, Number: 3829, Perfection: 8, Imperfection: 1},
		{Scale: scale.Garygic, Name: "Garygic", Cardinality: 9, Number: 3563, Perfection: 8, Imperfection: 1},
		{Scale: scale.Sorygic, Name: "Sorygic", Cardinality: 9, Number: 3031, Perfection: 8, Imperfection: 1},
		{Scale: scale.Godygic, Name: "Godygic", Cardinality: 9, Number: 3934, Perfection: 8, Imperfection: 1},
		{Scale: scale.Epithygic, Name: "Epithygic", Cardinality: 9, Number: 3773, Perfection: 8, Imperfection: 1},
		{Scale: scale.Ionoptygic, Name: "Ionoptygic", Cardinality: 9, Number: 3451, Perfection: 8, Imperfection: 1},

		// 9.2
		{Scale: scale.Kalygic, Name: "Kalygic", Cardinality: 9, Number: 2527, Perfection: 7, Imperfection: 2},
		{Scale: scale.Ionodygic, Name: "Ionodygic", Cardinality: 9, Number: 3836, Perfection: 7, Imperfection: 2},
		{Scale: scale.Bythygic, Name: "Bythygic", Cardinality: 9, Number: 3577, Perfection: 7, Imperfection: 2},
		{Scale: scale.Epygic, Name: "Epygic", Cardinality: 9, Number: 3059, Perfection: 7, Imperfection: 2},
		{Scale: scale.Marygic, Name: "Marygic", Cardinality: 9, Number: 4046, Perfection: 7, Imperfection: 2},
		{Scale: scale.Gaptygic, Name: "Gaptygic", Cardinality: 9, Number: 3997, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aeroptygic, Name: "Aeroptygic", Cardinality: 9, Number: 3899, Perfection: 7, Imperfection: 2},
		{Scale: scale.Mylygic, Name: "Mylygic", Cardinality: 9, Number: 3703, Perfection: 7, Imperfection: 2},
		{Scale: scale.Galygic, Name: "Galygic", Cardinality: 9, Number: 3311, Perfection: 7, Imperfection: 2},

		// 9.3
		{Scale: scale.Mixolydygic, Name: "Mixolydygic", Cardinality: 9, Number: 2543, Perfection: 7, Imperfection: 2},
		{Scale: scale.Ionycrygic, Name: "Ionycrygic", Cardinality: 9, Number: 3964, Perfection: 7, Imperfection: 2},
		{Scale: scale.Zoptygic, Name: "Zoptygic", Cardinality: 9, Number: 3833, Perfection: 7, Imperfection: 2},
		{Scale: scale.Phrygygic, Name: "Phrygygic", Cardinality: 9, Number: 3571, Perfection: 7, Imperfection: 2},
		{Scale: scale.Locrygic, Name: "Locrygic", Cardinality: 9, Number: 3047, Perfection: 7, Imperfection: 2},
		{Scale: scale.Gonygic, Name: "Gonygic", Cardinality: 9, Number: 3998, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aeracrygic, Name: "Aeracrygic", Cardinality: 9, Number: 3901, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aerathygic, Name: "Aerathygic", Cardinality: 9, Number: 3707, Perfection: 7, Imperfection: 2},
		{Scale: scale.Dorygic, Name: "Dorygic", Cardinality: 9, Number: 3319, Perfection: 7, Imperfection: 2},

		// 9.4
		{Scale: scale.Dycrygic, Name: "Dycrygic", Cardinality: 9, Number: 2551, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aeolygic, Name: "Aeolygic", Cardinality: 9, Number: 4028, Perfection: 7, Imperfection: 2},
		{Scale: scale.Dydygic, Name: "Dydygic", Cardinality: 9, Number: 3961, Perfection: 7, Imperfection: 2},
		{Scale: scale.Tholygic, Name: "Tholygic", Cardinality: 9, Number: 3827, Perfection: 7, Imperfection: 2},
		{Scale: scale.Rynygic, Name: "Rynygic", Cardinality: 9, Number: 3559, Perfection: 7, Imperfection: 2},
		{Scale: scale.Bycrygic, Name: "Bycrygic", Cardinality: 9, Number: 3023, Perfection: 7, Imperfection: 2},
		{Scale: scale.Zacrygic, Name: "Zacrygic", Cardinality: 9, Number: 3902, Perfection: 7, Imperfection: 2},
		{Scale: scale.Panygic, Name: "Panygic", Cardinality: 9, Number: 3709, Perfection: 7, Imperfection: 2},
		{Scale: scale.Dyrygic, Name: "Dyrygic", Cardinality: 9, Number: 3323, Perfection: 7, Imperfection: 2},

		// 9.5
		{Scale: scale.Loptygic, Name: "Loptygic", Cardinality: 9, Number: 2555, Perfection: 7, Imperfection: 2},
		{Scale: scale.Katylygic, Name: "Katylygic", Cardinality: 9, Number: 4060, Perfection: 7, Imperfection: 2},
		{Scale: scale.Phradygic, Name: "Phradygic", Cardinality: 9, Number: 4025, Perfection: 7, Imperfection: 2},
		{Scale: scale.Mixodygic, Name: "Mixodygic", Cardinality: 9, Number: 3955, Perfection: 7, Imperfection: 2},
		{Scale: scale.Katalygic, Name: "Katalygic", Cardinality: 9, Number: 3815, Perfection: 7, Imperfection: 2},
		{Scale: scale.Katorygic, Name: "Katorygic", Cardinality: 9, Number: 3535, Perfection: 7, Imperfection: 2},
		{Scale: scale.Dogygic, Name: "Dogygic", Cardinality: 9, Number: 2975, Perfection: 7, Imperfection: 2},
		{Scale: scale.Zodygic, Name: "Zodygic", Cardinality: 9, Number: 3710, Perfection: 7, Imperfection: 2},
		{Scale: scale.Madygic, Name: "Madygic", Cardinality: 9, Number: 3325, Perfection: 7, Imperfection: 2},

		// 9.6
		{Scale: scale.Bagygic, Name: "Bagygic", Cardinality: 9, Number: 2783, Perfection: 7, Imperfection: 2},
		{Scale: scale.Mathygic, Name: "Mathygic", Cardinality: 9, Number: 2942, Perfection: 7, Imperfection: 2},
		{Scale: scale.Styptygic, Name: "Styptygic", Cardinality: 9, Number: 3578, Perfection: 7, Imperfection: 2},
		{Scale: scale.Zolygic, Name: "Zolygic", Cardinality: 9, Number: 3061, Perfection: 7, Imperfection: 2},
		{Scale: scale.Sydygic, Name: "Sydygic", Cardinality: 9, Number: 4054, Perfection: 7, Imperfection: 2},
		{Scale: scale.Katygic, Name: "Katygic", Cardinality: 9, Number: 4013, Perfection: 7, Imperfection: 2},
		{Scale: scale.Zyphygic, Name: "Zyphygic", Cardinality: 9, Number: 3931, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aeralygic, Name: "Aeralygic", Cardinality: 9, Number: 3767, Perfection: 7, Imperfection: 2},
		{Scale: scale.Ryptygic, Name: "Ryptygic", Cardinality: 9, Number: 3439, Perfection: 7, Imperfection: 2},

		// 9.7
		{Scale: scale.Apinygic, Name: "Apinygic", Cardinality: 9, Number: 2813, Perfection: 7, Imperfection: 2},
		{Scale: scale.Katagygic, Name: "Katagygic", Cardinality: 9, Number: 3062, Perfection: 7, Imperfection: 2},
		{Scale: scale.Radygic, Name: "Radygic", Cardinality: 9, Number: 4058, Perfection: 7, Imperfection: 2},
		{Scale: scale.Gothygic, Name: "Gothygic", Cardinality: 9, Number: 4021, Perfection: 7, Imperfection: 2},
		{Scale: scale.Lythygic, Name: "Lythygic", Cardinality: 9, Number: 3947, Perfection: 7, Imperfection: 2},
		{Scale: scale.Bacrygic, Name: "Bacrygic", Cardinality: 9, Number: 3799, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aerygic, Name: "Aerygic", Cardinality: 9, Number: 3503, Perfection: 7, Imperfection: 2},
		{Scale: scale.Dathygic, Name: "Dathygic", Cardinality: 9, Number: 2911, Perfection: 7, Imperfection: 2},
		{Scale: scale.Boptygic, Name: "Boptygic", Cardinality: 9, Number: 3454, Perfection: 7, Imperfection: 2},

		// 9.8
		{Scale: scale.Epyrygic, Name: "Epyrygic", Cardinality: 9, Number: 2935, Perfection: 7, Imperfection: 2},
		{Scale: scale.Aeradygic, Name: "Aeradygic", Cardinality: 9, Number: 3550, Perfection: 7, Imperfection: 2},
		{Scale: scale.Staptygic, Name: "Staptygic", Cardinality: 9, Number: 3005, Perfection: 7, Imperfection: 2},
		{Scale: scale.Danygic, Name: "Danygic", Cardinality: 9, Number: 3830, Perfection: 7, Imperfection: 2},
		{Scale: scale.Goptygic, Name: "Goptygic", Cardinality: 9, Number: 3565, Perfection: 7, Imperfection: 2},
		{Scale: scale.Epocrygic, Name: "Epocrygic", Cardinality: 9, Number: 3035, Perfection: 7, Imperfection: 2},
		{Scale: scale.Rocrygic, Name: "Rocrygic", Cardinality: 9, Number: 3950, Perfection: 7, Imperfection: 2},
		{Scale: scale.Zyrygic, Name: "Zyrygic", Cardinality: 9, Number: 3805, Perfection: 7, Imperfection: 2},
		{Scale: scale.Sadygic, Name: "Sadygic", Cardinality: 9, Number: 3515, Perfection: 7, Imperfection: 2},

		// 9.9
		{Scale: scale.Aeolorygic, Name: "Aeolorygic", Cardinality: 9, Number: 2939, Perfection: 7, Imperfection: 2},
		{Scale: scale.Thydygic, Name: "Thydygic", Cardinality: 9, Number: 3566, Perfection: 7, Imperfection: 2},
		{Scale: scale.Gycrygic, Name: "Gycrygic", Cardinality: 9, Number: 3037, Perfection: 7, Imperfection: 2},
		{Scale: scale.Lyrygic, Name: "Lyrygic", Cardinality: 9, Number: 3958, Perfection: 7, Imperfection: 2},
		{Scale: scale.Modygic, Name: "Modygic", Cardinality: 9, Number: 3821, Perfection: 7, Imperfection: 2},
		{Scale: scale.Katodygic, Name: "Katodygic", Cardinality: 9, Number: 3547, Perfection: 7, Imperfection: 2},
		{Scale: scale.Moptygic, Name: "Moptygic", Cardinality: 9, Number: 2999, Perfection: 7, Imperfection: 2},
		{Scale: scale.Ionocrygic, Name: "Ionocrygic", Cardinality: 9, Number: 3806, Perfection: 7, Imperfection: 2},
		{Scale: scale.Gocrygic, Name: "Gocrygic", Cardinality: 9, Number: 3517, Perfection: 7, Imperfection: 2},

		// 9.10
		{Scale: scale.Manygic, Name: "Manygic", Cardinality: 9, Number: 2303, Perfection: 6, Imperfection: 3},
		{Scale: scale.Polygic, Name: "Polygic", Cardinality: 9, Number: 4088, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stanygic, Name: "Stanygic", Cardinality: 9, Number: 4081, Perfection: 6, Imperfection: 3},
		{Scale: scale.Thaptygic, Name: "Thaptygic", Cardinality: 9, Number: 4067, Perfection: 6, Imperfection: 3},
		{Scale: scale.Eporygic, Name: "Eporygic", Cardinality: 9, Number: 4039, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aerynygic, Name: "Aerynygic", Cardinality: 9, Number: 3983, Perfection: 6, Imperfection: 3},
		{Scale: scale.Thyptygic, Name: "Thyptygic", Cardinality: 9, Number: 3871, Perfection: 6, Imperfection: 3},
		{Scale: scale.Ionogygic, Name: "Ionogygic", Cardinality: 9, Number: 3647, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolarygic, Name: "Aeolarygic", Cardinality: 9, Number: 3199, Perfection: 6, Imperfection: 3},

		// 9.11
		{Scale: scale.Sathygic, Name: "Sathygic", Cardinality: 9, Number: 2431, Perfection: 6, Imperfection: 3},
		{Scale: scale.Ladygic, Name: "Ladygic", Cardinality: 9, Number: 3068, Perfection: 6, Imperfection: 3},
		{Scale: scale.Sarygic, Name: "Sarygic", Cardinality: 9, Number: 4082, Perfection: 6, Imperfection: 3},
		{Scale: scale.Thacrygic, Name: "Thacrygic", Cardinality: 9, Number: 4069, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolynygic, Name: "Aeolynygic", Cardinality: 9, Number: 4043, Perfection: 6, Imperfection: 3},
		{Scale: scale.Thadygic, Name: "Thadygic", Cardinality: 9, Number: 3991, Perfection: 6, Imperfection: 3},
		{Scale: scale.Lynygic, Name: "Lynygic", Cardinality: 9, Number: 3887, Perfection: 6, Imperfection: 3},
		{Scale: scale.Doptygic, Name: "Doptygic", Cardinality: 9, Number: 3679, Perfection: 6, Imperfection: 3},
		{Scale: scale.Ionilygic, Name: "Ionilygic", Cardinality: 9, Number: 3263, Perfection: 6, Imperfection: 3},

		// 9.12
		{Scale: scale.Phrygic, Name: "Phrygic", Cardinality: 9, Number: 2495, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeranygic, Name: "Aeranygic", Cardinality: 9, Number: 3580, Perfection: 6, Imperfection: 3},
		{Scale: scale.Dothygic, Name: "Dothygic", Cardinality: 9, Number: 3065, Perfection: 6, Imperfection: 3},
		{Scale: scale.Lydygic, Name: "Lydygic", Cardinality: 9, Number: 4070, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stadygic, Name: "Stadygic", Cardinality: 9, Number: 4045, Perfection: 6, Imperfection: 3},
		{Scale: scale.Byptygic, Name: "Byptygic", Cardinality: 9, Number: 3995, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stodygic, Name: "Stodygic", Cardinality: 9, Number: 3895, Perfection: 6, Imperfection: 3},
		{Scale: scale.Zynygic, Name: "Zynygic", Cardinality: 9, Number: 3695, Perfection: 6, Imperfection: 3},
		{Scale: scale.Lonygic, Name: "Lonygic", Cardinality: 9, Number: 3295, Perfection: 6, Imperfection: 3},

		// 9.13
		{Scale: scale.Zothygic, Name: "Zothygic", Cardinality: 9, Number: 2557, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolathygic, Name: "Aeolathygic", Cardinality: 9, Number: 4076, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolocrygic, Name: "Aeolocrygic", Cardinality: 9, Number: 4057, Perfection: 6, Imperfection: 3},
		{Scale: scale.Phroptygic, Name: "Phroptygic", Cardinality: 9, Number: 4019, Perfection: 6, Imperfection: 3},
		{Scale: scale.Kodygic, Name: "Kodygic", Cardinality: 9, Number: 3943, Perfection: 6, Imperfection: 3},
		{Scale: scale.Eparygic, Name: "Eparygic", Cardinality: 9, Number: 3791, Perfection: 6, Imperfection: 3},
		{Scale: scale.Ionygic, Name: "Ionygic", Cardinality: 9, Number: 3487, Perfection: 6, Imperfection: 3},
		{Scale: scale.Gyptygic, Name: "Gyptygic", Cardinality: 9, Number: 2879, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aerythygic, Name: "Aerythygic", Cardinality: 9, Number: 3326, Perfection: 6, Imperfection: 3},

		// 9.14
		{Scale: scale.Aeolacrygic, Name: "Aeolacrygic", Cardinality: 9, Number: 2558, Perfection: 6, Imperfection: 3},
		{Scale: scale.Raptygic, Name: "Raptygic", Cardinality: 9, Number: 4084, Perfection: 6, Imperfection: 3},
		{Scale: scale.Gythygic, Name: "Gythygic", Cardinality: 9, Number: 4073, Perfection: 6, Imperfection: 3},
		{Scale: scale.Pyrygic, Name: "Pyrygic", Cardinality: 9, Number: 4051, Perfection: 6, Imperfection: 3},
		{Scale: scale.Rycrygic, Name: "Rycrygic", Cardinality: 9, Number: 4007, Perfection: 6, Imperfection: 3},
		{Scale: scale.Phrathygic, Name: "Phrathygic", Cardinality: 9, Number: 3919, Perfection: 6, Imperfection: 3},
		{Scale: scale.Badygic, Name: "Badygic", Cardinality: 9, Number: 3743, Perfection: 6, Imperfection: 3},
		{Scale: scale.Phrocrygic, Name: "Phrocrygic", Cardinality: 9, Number: 3391, Perfection: 6, Imperfection: 3},
		{Scale: scale.Starygic, Name: "Starygic", Cardinality: 9, Number: 2687, Perfection: 6, Imperfection: 3},

		// 9.15
		{Scale: scale.Kyrygic, Name: "Kyrygic", Cardinality: 9, Number: 2751, Perfection: 6, Imperfection: 3},
		{Scale: scale.Sythygic, Name: "Sythygic", Cardinality: 9, Number: 2814, Perfection: 6, Imperfection: 3},
		{Scale: scale.Katycrygic, Name: "Katycrygic", Cardinality: 9, Number: 3066, Perfection: 6, Imperfection: 3},
		{Scale: scale.Tharygic, Name: "Tharygic", Cardinality: 9, Number: 4074, Perfection: 6, Imperfection: 3},
		{Scale: scale.Sylygic, Name: "Sylygic", Cardinality: 9, Number: 4053, Perfection: 6, Imperfection: 3},
		{Scale: scale.Lothygic, Name: "Lothygic", Cardinality: 9, Number: 4011, Perfection: 6, Imperfection: 3},
		{Scale: scale.Darygic, Name: "Darygic", Cardinality: 9, Number: 3927, Perfection: 6, Imperfection: 3},
		{Scale: scale.Monygic, Name: "Monygic", Cardinality: 9, Number: 3759, Perfection: 6, Imperfection: 3},
		{Scale: scale.Styrygic, Name: "Styrygic", Cardinality: 9, Number: 3423, Perfection: 6, Imperfection: 3},

		// 9.16
		{Scale: scale.Porygic, Name: "Porygic", Cardinality: 9, Number: 2799, Perfection: 6, Imperfection: 3},
		{Scale: scale.Mocrygic, Name: "Mocrygic", Cardinality: 9, Number: 3006, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolyrygic, Name: "Aeolyrygic", Cardinality: 9, Number: 3834, Perfection: 6, Imperfection: 3},
		{Scale: scale.Barygic, Name: "Barygic", Cardinality: 9, Number: 3573, Perfection: 6, Imperfection: 3},
		{Scale: scale.Katarygic, Name: "Katarygic", Cardinality: 9, Number: 3051, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aerocrygic, Name: "Aerocrygic", Cardinality: 9, Number: 4014, Perfection: 6, Imperfection: 3},
		{Scale: scale.Zanygic, Name: "Zanygic", Cardinality: 9, Number: 3933, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolonygic, Name: "Aeolonygic", Cardinality: 9, Number: 3771, Perfection: 6, Imperfection: 3},
		{Scale: scale.Aeolanygic, Name: "Aeolanygic", Cardinality: 9, Number: 3447, Perfection: 6, Imperfection: 3},

		// 9.17
		{Scale: scale.Kaptygic, Name: "Kaptygic", Cardinality: 9, Number: 2811, Perfection: 6, Imperfection: 3},
		{Scale: scale.Sacrygic, Name: "Sacrygic", Cardinality: 9, Number: 3054, Perfection: 6, Imperfection: 3},
		{Scale: scale.Padygic, Name: "Padygic", Cardinality: 9, Number: 4026, Perfection: 6, Imperfection: 3},
		{Scale: scale.Epilygic, Name: "Epilygic", Cardinality: 9, Number: 3957, Perfection: 6, Imperfection: 3},
		{Scale: scale.Kynygic, Name: "Kynygic", Cardinality: 9, Number: 3819, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stophygic, Name: "Stophygic", Cardinality: 9, Number: 3543, Perfection: 6, Imperfection: 3},
		{Scale: scale.Ionidygic, Name: "Ionidygic", Cardinality: 9, Number: 2991, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stonygic, Name: "Stonygic", Cardinality: 9, Number: 3774, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stalygic, Name: "Stalygic", Cardinality: 9, Number: 3453, Perfection: 6, Imperfection: 3},

		// 9.18
		{Scale: scale.Koptygic, Name: "Koptygic", Cardinality: 9, Number: 2927, Perfection: 6, Imperfection: 3},
		{Scale: scale.Raphygic, Name: "Raphygic", Cardinality: 9, Number: 3518, Perfection: 6, Imperfection: 3},
		{Scale: scale.Zycrygic, Name: "Zycrygic", Cardinality: 9, Number: 2941, Perfection: 6, Imperfection: 3},
		{Scale: scale.Mycrygic, Name: "Mycrygic", Cardinality: 9, Number: 3574, Perfection: 6, Imperfection: 3},
		{Scale: scale.Laptygic, Name: "Laptygic", Cardinality: 9, Number: 3053, Perfection: 6, Imperfection: 3},
		{Scale: scale.Pylygic, Name: "Pylygic", Cardinality: 9, Number: 4022, Perfection: 6, Imperfection: 3},
		{Scale: scale.Rodygic, Name: "Rodygic", Cardinality: 9, Number: 3949, Perfection: 6, Imperfection: 3},
		{Scale: scale.Epolygic, Name: "Epolygic", Cardinality: 9, Number: 3803, Perfection: 6, Imperfection: 3},
		{Scale: scale.Epidygic, Name: "Epidygic", Cardinality: 9, Number: 3511, Perfection: 6, Imperfection: 3},

		// 9.19
		{Scale: scale.Phronygic, Name: "Phronygic", Cardinality: 9, Number: 3003, Perfection: 6, Imperfection: 3},
		{Scale: scale.Stynygic, Name: "Stynygic", Cardinality: 9, Number: 3822, Perfection: 6, Imperfection: 3},
		{Scale: scale.Zydygic, Name: "Zydygic", Cardinality: 9, Number: 3549, Perfection: 6, Imperfection: 3},
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
