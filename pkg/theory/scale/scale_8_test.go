package scale_test

import (
	"testing"

	"github.com/edipermadi/music-db/pkg/theory/scale"
	"github.com/stretchr/testify/assert"
)

func TestScaleWith8Notes_Values(t *testing.T) {
	type testCase struct {
		Scale        scale.Type
		Name         string
		Cardinality  int
		Number       int
		Perfection   int
		Imperfection int
	}

	testCases := []testCase{
		// 8.1
		{Scale: scale.Aerycryllic, Name: "Aerycryllic", Cardinality: 8, Number: 2775, Perfection: 7, Imperfection: 1},
		{Scale: scale.Gadyllic, Name: "Gadyllic", Cardinality: 8, Number: 2910, Perfection: 7, Imperfection: 1},
		{Scale: scale.Solyllic, Name: "Solyllic", Cardinality: 8, Number: 3450, Perfection: 7, Imperfection: 1},
		{Scale: scale.Zylyllic, Name: "Zylyllic", Cardinality: 8, Number: 2805, Perfection: 7, Imperfection: 1},
		{Scale: scale.Mixodyllic, Name: "Mixodyllic", Cardinality: 8, Number: 3030, Perfection: 7, Imperfection: 1},
		{Scale: scale.Soryllic, Name: "Soryllic", Cardinality: 8, Number: 3930, Perfection: 7, Imperfection: 1},
		{Scale: scale.Godyllic, Name: "Godyllic", Cardinality: 8, Number: 3765, Perfection: 7, Imperfection: 1},
		{Scale: scale.Epiphyllic, Name: "Epiphyllic", Cardinality: 8, Number: 3435, Perfection: 7, Imperfection: 1},

		// 8.2
		{Scale: scale.Pynyllic, Name: "Pynyllic", Cardinality: 8, Number: 2295, Perfection: 6, Imperfection: 2},
		{Scale: scale.Bocryllic, Name: "Bocryllic", Cardinality: 8, Number: 3960, Perfection: 6, Imperfection: 2},
		{Scale: scale.Kogyllic, Name: "Kogyllic", Cardinality: 8, Number: 3825, Perfection: 6, Imperfection: 2},
		{Scale: scale.Raryllic, Name: "Raryllic", Cardinality: 8, Number: 3555, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zycryllic, Name: "Zycryllic", Cardinality: 8, Number: 3015, Perfection: 6, Imperfection: 2},
		{Scale: scale.Mycryllic, Name: "Mycryllic", Cardinality: 8, Number: 3870, Perfection: 6, Imperfection: 2},
		{Scale: scale.Laptyllic, Name: "Laptyllic", Cardinality: 8, Number: 3645, Perfection: 6, Imperfection: 2},
		{Scale: scale.Pylyllic, Name: "Pylyllic", Cardinality: 8, Number: 3195, Perfection: 6, Imperfection: 2},

		// 8.3
		{Scale: scale.Pothyllic, Name: "Pothyllic", Cardinality: 8, Number: 2427, Perfection: 6, Imperfection: 2},
		{Scale: scale.Phronyllic, Name: "Phronyllic", Cardinality: 8, Number: 3036, Perfection: 6, Imperfection: 2},
		{Scale: scale.Stynyllic, Name: "Stynyllic", Cardinality: 8, Number: 3954, Perfection: 6, Imperfection: 2},
		{Scale: scale.Rathyllic, Name: "Rathyllic", Cardinality: 8, Number: 3813, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aeryptyllic, Name: "Aeryptyllic", Cardinality: 8, Number: 3531, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zydyllic, Name: "Zydyllic", Cardinality: 8, Number: 2967, Perfection: 6, Imperfection: 2},
		{Scale: scale.Katolyllic, Name: "Katolyllic", Cardinality: 8, Number: 3678, Perfection: 6, Imperfection: 2},
		{Scale: scale.Rythyllic, Name: "Rythyllic", Cardinality: 8, Number: 3261, Perfection: 6, Imperfection: 2},

		// 8.4
		{Scale: scale.Locryllic, Name: "Locryllic", Cardinality: 8, Number: 2511, Perfection: 6, Imperfection: 2},
		{Scale: scale.Bylyllic, Name: "Bylyllic", Cardinality: 8, Number: 3708, Perfection: 6, Imperfection: 2},
		{Scale: scale.Sogyllic, Name: "Sogyllic", Cardinality: 8, Number: 3321, Perfection: 6, Imperfection: 2},
		{Scale: scale.Ionycryllic, Name: "Ionycryllic", Cardinality: 8, Number: 2547, Perfection: 6, Imperfection: 2},
		{Scale: scale.Koptyllic, Name: "Koptyllic", Cardinality: 8, Number: 3996, Perfection: 6, Imperfection: 2},
		{Scale: scale.Epyryllic, Name: "Epyryllic", Cardinality: 8, Number: 3897, Perfection: 6, Imperfection: 2},
		{Scale: scale.Soptyllic, Name: "Soptyllic", Cardinality: 8, Number: 3699, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aeolylyllic, Name: "Aeolylyllic", Cardinality: 8, Number: 3303, Perfection: 6, Imperfection: 2},

		// 8.5
		{Scale: scale.Aeracryllic, Name: "Aeracryllic", Cardinality: 8, Number: 2519, Perfection: 6, Imperfection: 2},
		{Scale: scale.Epygyllic, Name: "Epygyllic", Cardinality: 8, Number: 3772, Perfection: 6, Imperfection: 2},
		{Scale: scale.Thonyllic, Name: "Thonyllic", Cardinality: 8, Number: 3449, Perfection: 6, Imperfection: 2},
		{Scale: scale.Lanyllic, Name: "Lanyllic", Cardinality: 8, Number: 2803, Perfection: 6, Imperfection: 2},
		{Scale: scale.Phrynyllic, Name: "Phrynyllic", Cardinality: 8, Number: 3022, Perfection: 6, Imperfection: 2},
		{Scale: scale.Lycryllic, Name: "Lycryllic", Cardinality: 8, Number: 3898, Perfection: 6, Imperfection: 2},
		{Scale: scale.Ionyptyllic, Name: "Ionyptyllic", Cardinality: 8, Number: 3701, Perfection: 6, Imperfection: 2},
		{Scale: scale.Epathyllic, Name: "Epathyllic", Cardinality: 8, Number: 3307, Perfection: 6, Imperfection: 2},

		// 8.6
		{Scale: scale.Dydyllic, Name: "Dydyllic", Cardinality: 8, Number: 2523, Perfection: 6, Imperfection: 2},
		{Scale: scale.Thogyllic, Name: "Thogyllic", Cardinality: 8, Number: 3804, Perfection: 6, Imperfection: 2},
		{Scale: scale.Rygyllic, Name: "Rygyllic", Cardinality: 8, Number: 3513, Perfection: 6, Imperfection: 2},
		{Scale: scale.Bycryllic, Name: "Bycryllic", Cardinality: 8, Number: 2931, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zacryllic, Name: "Zacryllic", Cardinality: 8, Number: 3534, Perfection: 6, Imperfection: 2},
		{Scale: scale.Panyllic, Name: "Panyllic", Cardinality: 8, Number: 2973, Perfection: 6, Imperfection: 2},
		{Scale: scale.Dyryllic, Name: "Dyryllic", Cardinality: 8, Number: 3702, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zathyllic, Name: "Zathyllic", Cardinality: 8, Number: 3309, Perfection: 6, Imperfection: 2},

		// 8.7
		{Scale: scale.Dagyllic, Name: "Dagyllic", Cardinality: 8, Number: 2526, Perfection: 6, Imperfection: 2},
		{Scale: scale.Katalyllic, Name: "Katalyllic", Cardinality: 8, Number: 3828, Perfection: 6, Imperfection: 2},
		{Scale: scale.Katoryllic, Name: "Katoryllic", Cardinality: 8, Number: 3561, Perfection: 6, Imperfection: 2},
		{Scale: scale.Dodyllic, Name: "Dodyllic", Cardinality: 8, Number: 3027, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zogyllic, Name: "Zogyllic", Cardinality: 8, Number: 3918, Perfection: 6, Imperfection: 2},
		{Scale: scale.Madyllic, Name: "Madyllic", Cardinality: 8, Number: 3741, Perfection: 6, Imperfection: 2},
		{Scale: scale.Dycryllic, Name: "Dycryllic", Cardinality: 8, Number: 3387, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aeologyllic, Name: "Aeologyllic", Cardinality: 8, Number: 2679, Perfection: 6, Imperfection: 2},

		// 8.8
		{Scale: scale.Sydyllic, Name: "Sydyllic", Cardinality: 8, Number: 2535, Perfection: 6, Imperfection: 2},
		{Scale: scale.Katogyllic, Name: "Katogyllic", Cardinality: 8, Number: 3900, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zygyllic, Name: "Zygyllic", Cardinality: 8, Number: 3705, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aeralyllic, Name: "Aeralyllic", Cardinality: 8, Number: 3315, Perfection: 6, Imperfection: 2},

		// 8.9
		{Scale: scale.Bacryllic, Name: "Bacryllic", Cardinality: 8, Number: 2539, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aerygyllic, Name: "Aerygyllic", Cardinality: 8, Number: 3932, Perfection: 6, Imperfection: 2},
		{Scale: scale.Dathyllic, Name: "Dathyllic", Cardinality: 8, Number: 3769, Perfection: 6, Imperfection: 2},
		{Scale: scale.Boptyllic, Name: "Boptyllic", Cardinality: 8, Number: 3443, Perfection: 6, Imperfection: 2},
		{Scale: scale.Bagyllic, Name: "Bagyllic", Cardinality: 8, Number: 2791, Perfection: 6, Imperfection: 2},
		{Scale: scale.Mathyllic, Name: "Mathyllic", Cardinality: 8, Number: 2974, Perfection: 6, Imperfection: 2},
		{Scale: scale.Styptyllic, Name: "Styptyllic", Cardinality: 8, Number: 3706, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zolyllic, Name: "Zolyllic", Cardinality: 8, Number: 3317, Perfection: 6, Imperfection: 2},

		// 8.10
		{Scale: scale.Rocryllic, Name: "Rocryllic", Cardinality: 8, Number: 2743, Perfection: 6, Imperfection: 2},
		{Scale: scale.Zyryllic, Name: "Zyryllic", Cardinality: 8, Number: 2782, Perfection: 6, Imperfection: 2},
		{Scale: scale.Sagyllic, Name: "Sagyllic", Cardinality: 8, Number: 2938, Perfection: 6, Imperfection: 2},
		{Scale: scale.Epinyllic, Name: "Epinyllic", Cardinality: 8, Number: 3562, Perfection: 6, Imperfection: 2},
		{Scale: scale.Katagyllic, Name: "Katagyllic", Cardinality: 8, Number: 3029, Perfection: 6, Imperfection: 2},
		{Scale: scale.Ragyllic, Name: "Ragyllic", Cardinality: 8, Number: 3926, Perfection: 6, Imperfection: 2},
		{Scale: scale.Gothyllic, Name: "Gothyllic", Cardinality: 8, Number: 3757, Perfection: 6, Imperfection: 2},
		{Scale: scale.Lythyllic, Name: "Lythyllic", Cardinality: 8, Number: 3419, Perfection: 6, Imperfection: 2},

		// 8.11
		{Scale: scale.Ionocryllic, Name: "Ionocryllic", Cardinality: 8, Number: 2749, Perfection: 6, Imperfection: 2},
		{Scale: scale.Gocryllic, Name: "Gocryllic", Cardinality: 8, Number: 2806, Perfection: 6, Imperfection: 2},
		{Scale: scale.Epiryllic, Name: "Epiryllic", Cardinality: 8, Number: 3034, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aeradyllic, Name: "Aeradyllic", Cardinality: 8, Number: 3946, Perfection: 6, Imperfection: 2},
		{Scale: scale.Staptyllic, Name: "Staptyllic", Cardinality: 8, Number: 3797, Perfection: 6, Imperfection: 2},
		{Scale: scale.Danyllic, Name: "Danyllic", Cardinality: 8, Number: 3499, Perfection: 6, Imperfection: 2},
		{Scale: scale.Goptyllic, Name: "Goptyllic", Cardinality: 8, Number: 2903, Perfection: 6, Imperfection: 2},
		{Scale: scale.Epocryllic, Name: "Epocryllic", Cardinality: 8, Number: 3422, Perfection: 6, Imperfection: 2},

		// 8.12
		{Scale: scale.Ionoptyllic, Name: "Ionoptyllic", Cardinality: 8, Number: 2781, Perfection: 6, Imperfection: 2},
		{Scale: scale.Aeoloryllic, Name: "Aeoloryllic", Cardinality: 8, Number: 2934, Perfection: 6, Imperfection: 2},
		{Scale: scale.Thydyllic, Name: "Thydyllic", Cardinality: 8, Number: 3546, Perfection: 6, Imperfection: 2},
		{Scale: scale.Gycryllic, Name: "Gycryllic", Cardinality: 8, Number: 2997, Perfection: 6, Imperfection: 2},
		{Scale: scale.Lyryllic, Name: "Lyryllic", Cardinality: 8, Number: 3798, Perfection: 6, Imperfection: 2},
		{Scale: scale.Mogyllic, Name: "Mogyllic", Cardinality: 8, Number: 3501, Perfection: 6, Imperfection: 2},
		{Scale: scale.Katodyllic, Name: "Katodyllic", Cardinality: 8, Number: 2907, Perfection: 6, Imperfection: 2},
		{Scale: scale.Moptyllic, Name: "Moptyllic", Cardinality: 8, Number: 3438, Perfection: 6, Imperfection: 2},

		// 8.13
		{Scale: scale.Dolyllic, Name: "Dolyllic", Cardinality: 8, Number: 2271, Perfection: 5, Imperfection: 3},
		{Scale: scale.Moryllic, Name: "Moryllic", Cardinality: 8, Number: 3576, Perfection: 5, Imperfection: 3},
		{Scale: scale.Bydyllic, Name: "Bydyllic", Cardinality: 8, Number: 3057, Perfection: 5, Imperfection: 3},
		{Scale: scale.Pocryllic, Name: "Pocryllic", Cardinality: 8, Number: 4038, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phracryllic, Name: "Phracryllic", Cardinality: 8, Number: 3981, Perfection: 5, Imperfection: 3},
		{Scale: scale.Gyryllic, Name: "Gyryllic", Cardinality: 8, Number: 3867, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrygyllic, Name: "Phrygyllic", Cardinality: 8, Number: 3639, Perfection: 5, Imperfection: 3},
		{Scale: scale.Dogyllic, Name: "Dogyllic", Cardinality: 8, Number: 3183, Perfection: 5, Imperfection: 3},

		// 8.14
		{Scale: scale.Thagyllic, Name: "Thagyllic", Cardinality: 8, Number: 2287, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thoptyllic, Name: "Thoptyllic", Cardinality: 8, Number: 3832, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phraptyllic, Name: "Phraptyllic", Cardinality: 8, Number: 3569, Perfection: 5, Imperfection: 3},
		{Scale: scale.Gylyllic, Name: "Gylyllic", Cardinality: 8, Number: 3043, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phralyllic, Name: "Phralyllic", Cardinality: 8, Number: 3982, Perfection: 5, Imperfection: 3},
		{Scale: scale.Dygyllic, Name: "Dygyllic", Cardinality: 8, Number: 3869, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ronyllic, Name: "Ronyllic", Cardinality: 8, Number: 3643, Perfection: 5, Imperfection: 3},
		{Scale: scale.Epogyllic, Name: "Epogyllic", Cardinality: 8, Number: 3191, Perfection: 5, Imperfection: 3},

		// 8.15
		{Scale: scale.Aeoladyllic, Name: "Aeoladyllic", Cardinality: 8, Number: 2299, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kocryllic, Name: "Kocryllic", Cardinality: 8, Number: 4024, Perfection: 5, Imperfection: 3},
		{Scale: scale.Lodyllic, Name: "Lodyllic", Cardinality: 8, Number: 3953, Perfection: 5, Imperfection: 3},
		{Scale: scale.Bynyllic, Name: "Bynyllic", Cardinality: 8, Number: 3811, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kydyllic, Name: "Kydyllic", Cardinality: 8, Number: 3527, Perfection: 5, Imperfection: 3},
		{Scale: scale.Bygyllic, Name: "Bygyllic", Cardinality: 8, Number: 2959, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phryptyllic, Name: "Phryptyllic", Cardinality: 8, Number: 3646, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionayllic, Name: "Ionayllic", Cardinality: 8, Number: 3197, Perfection: 5, Imperfection: 3},

		// 8.16
		{Scale: scale.Phroryllic, Name: "Phroryllic", Cardinality: 8, Number: 2301, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thyphyllic, Name: "Thyphyllic", Cardinality: 8, Number: 4056, Perfection: 5, Imperfection: 3},
		{Scale: scale.Poptyllic, Name: "Poptyllic", Cardinality: 8, Number: 4017, Perfection: 5, Imperfection: 3},
		{Scale: scale.Mixonyllic, Name: "Mixonyllic", Cardinality: 8, Number: 3939, Perfection: 5, Imperfection: 3},
		{Scale: scale.Paptyllic, Name: "Paptyllic", Cardinality: 8, Number: 3783, Perfection: 5, Imperfection: 3},
		{Scale: scale.Storyllic, Name: "Storyllic", Cardinality: 8, Number: 3471, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrycryllic, Name: "Phrycryllic", Cardinality: 8, Number: 2847, Perfection: 5, Imperfection: 3},
		{Scale: scale.Palyllic, Name: "Palyllic", Cardinality: 8, Number: 3198, Perfection: 5, Imperfection: 3},

		// 8.17
		{Scale: scale.Phranyllic, Name: "Phranyllic", Cardinality: 8, Number: 2399, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stydyllic, Name: "Stydyllic", Cardinality: 8, Number: 2812, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zadyllic, Name: "Zadyllic", Cardinality: 8, Number: 3058, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zalyllic, Name: "Zalyllic", Cardinality: 8, Number: 4042, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zocryllic, Name: "Zocryllic", Cardinality: 8, Number: 3989, Perfection: 5, Imperfection: 3},
		{Scale: scale.Katocryllic, Name: "Katocryllic", Cardinality: 8, Number: 3883, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aerathyllic, Name: "Aerathyllic", Cardinality: 8, Number: 3671, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stoptyllic, Name: "Stoptyllic", Cardinality: 8, Number: 3247, Perfection: 5, Imperfection: 3},

		// 8.18
		{Scale: scale.Lydyllic, Name: "Lydyllic", Cardinality: 8, Number: 2415, Perfection: 5, Imperfection: 3},
		{Scale: scale.Radyllic, Name: "Radyllic", Cardinality: 8, Number: 2940, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stagyllic, Name: "Stagyllic", Cardinality: 8, Number: 3570, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionoryllic, Name: "Ionoryllic", Cardinality: 8, Number: 3045, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrodyllic, Name: "Phrodyllic", Cardinality: 8, Number: 3990, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeragyllic, Name: "Aeragyllic", Cardinality: 8, Number: 3885, Perfection: 5, Imperfection: 3},
		{Scale: scale.Banyllic, Name: "Banyllic", Cardinality: 8, Number: 3675, Perfection: 5, Imperfection: 3},
		{Scale: scale.Epothyllic, Name: "Epothyllic", Cardinality: 8, Number: 3255, Perfection: 5, Imperfection: 3},

		// 8.19
		{Scale: scale.Zoryllic, Name: "Zoryllic", Cardinality: 8, Number: 2423, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrolyllic, Name: "Phrolyllic", Cardinality: 8, Number: 3004, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kolyllic, Name: "Kolyllic", Cardinality: 8, Number: 3826, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thodyllic, Name: "Thodyllic", Cardinality: 8, Number: 3557, Perfection: 5, Imperfection: 3},
		{Scale: scale.Socryllic, Name: "Socryllic", Cardinality: 8, Number: 3019, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolyllic, Name: "Aeolyllic", Cardinality: 8, Number: 3886, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zythyllic, Name: "Zythyllic", Cardinality: 8, Number: 3677, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeoryllic, Name: "Aeoryllic", Cardinality: 8, Number: 3259, Perfection: 5, Imperfection: 3},

		// 8.20
		{Scale: scale.Mixolydyllic, Name: "Mixolydyllic", Cardinality: 8, Number: 2430, Perfection: 5, Imperfection: 3},
		{Scale: scale.Mixonyphyllic, Name: "Mixonyphyllic", Cardinality: 8, Number: 3060, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolanyllic, Name: "Aeolanyllic", Cardinality: 8, Number: 4050, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thocryllic, Name: "Thocryllic", Cardinality: 8, Number: 4005, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kygyllic, Name: "Kygyllic", Cardinality: 8, Number: 3915, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionagyllic, Name: "Ionagyllic", Cardinality: 8, Number: 3735, Perfection: 5, Imperfection: 3},
		{Scale: scale.Gogyllic, Name: "Gogyllic", Cardinality: 8, Number: 3375, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phradyllic, Name: "Phradyllic", Cardinality: 8, Number: 2655, Perfection: 5, Imperfection: 3},

		// 8.21
		{Scale: scale.Ioniptyllic, Name: "Ioniptyllic", Cardinality: 8, Number: 2463, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kycryllic, Name: "Kycryllic", Cardinality: 8, Number: 3324, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolaptyllic, Name: "Aeolaptyllic", Cardinality: 8, Number: 2553, Perfection: 5, Imperfection: 3},
		{Scale: scale.Rodyllic, Name: "Rodyllic", Cardinality: 8, Number: 4044, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionathyllic, Name: "Ionathyllic", Cardinality: 8, Number: 3993, Perfection: 5, Imperfection: 3},
		{Scale: scale.Pythyllic, Name: "Pythyllic", Cardinality: 8, Number: 3891, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zonyllic, Name: "Zonyllic", Cardinality: 8, Number: 3687, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ryryllic, Name: "Ryryllic", Cardinality: 8, Number: 3279, Perfection: 5, Imperfection: 3},

		// 8.22
		{Scale: scale.Aeolothyllic, Name: "Aeolothyllic", Cardinality: 8, Number: 2479, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionyryllic, Name: "Ionyryllic", Cardinality: 8, Number: 3452, Perfection: 5, Imperfection: 3},
		{Scale: scale.Rydyllic, Name: "Rydyllic", Cardinality: 8, Number: 2809, Perfection: 5, Imperfection: 3},
		{Scale: scale.Gonyllic, Name: "Gonyllic", Cardinality: 8, Number: 3046, Perfection: 5, Imperfection: 3},
		{Scale: scale.Rolyllic, Name: "Rolyllic", Cardinality: 8, Number: 3994, Perfection: 5, Imperfection: 3},
		{Scale: scale.Katydyllic, Name: "Katydyllic", Cardinality: 8, Number: 3893, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zyptyllic, Name: "Zyptyllic", Cardinality: 8, Number: 3691, Perfection: 5, Imperfection: 3},
		{Scale: scale.Modyllic, Name: "Modyllic", Cardinality: 8, Number: 3287, Perfection: 5, Imperfection: 3},

		// 8.23
		{Scale: scale.Maptyllic, Name: "Maptyllic", Cardinality: 8, Number: 2487, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeraptyllic, Name: "Aeraptyllic", Cardinality: 8, Number: 3516, Perfection: 5, Imperfection: 3},
		{Scale: scale.Katadyllic, Name: "Katadyllic", Cardinality: 8, Number: 2937, Perfection: 5, Imperfection: 3},
		{Scale: scale.Magyllic, Name: "Magyllic", Cardinality: 8, Number: 3558, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrylyllic, Name: "Phrylyllic", Cardinality: 8, Number: 3021, Perfection: 5, Imperfection: 3},
		{Scale: scale.Epigyllic, Name: "Epigyllic", Cardinality: 8, Number: 3894, Perfection: 5, Imperfection: 3},
		{Scale: scale.Molyllic, Name: "Molyllic", Cardinality: 8, Number: 3693, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ponyllic, Name: "Ponyllic", Cardinality: 8, Number: 3291, Perfection: 5, Imperfection: 3},

		// 8.24
		{Scale: scale.Thyptyllic, Name: "Thyptyllic", Cardinality: 8, Number: 2491, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionogyllic, Name: "Ionogyllic", Cardinality: 8, Number: 3548, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolaryllic, Name: "Aeolaryllic", Cardinality: 8, Number: 3001, Perfection: 5, Imperfection: 3},
		{Scale: scale.Katygyllic, Name: "Katygyllic", Cardinality: 8, Number: 3814, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ganyllic, Name: "Ganyllic", Cardinality: 8, Number: 3533, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kyptyllic, Name: "Kyptyllic", Cardinality: 8, Number: 2971, Perfection: 5, Imperfection: 3},
		{Scale: scale.Salyllic, Name: "Salyllic", Cardinality: 8, Number: 3694, Perfection: 5, Imperfection: 3},
		{Scale: scale.Sanyllic, Name: "Sanyllic", Cardinality: 8, Number: 3293, Perfection: 5, Imperfection: 3},

		// 8.25
		{Scale: scale.Doptyllic, Name: "Doptyllic", Cardinality: 8, Number: 2493, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionilyllic, Name: "Ionilyllic", Cardinality: 8, Number: 3564, Perfection: 5, Imperfection: 3},
		{Scale: scale.Manyllic, Name: "Manyllic", Cardinality: 8, Number: 3033, Perfection: 5, Imperfection: 3},
		{Scale: scale.Polyllic, Name: "Polyllic", Cardinality: 8, Number: 3942, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stanyllic, Name: "Stanyllic", Cardinality: 8, Number: 3789, Perfection: 5, Imperfection: 3},
		{Scale: scale.Mixotharyllic, Name: "Mixotharyllic", Cardinality: 8, Number: 3483, Perfection: 5, Imperfection: 3},
		{Scale: scale.Eporyllic, Name: "Eporyllic", Cardinality: 8, Number: 2871, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aerynyllic, Name: "Aerynyllic", Cardinality: 8, Number: 3294, Perfection: 5, Imperfection: 3},

		// 8.26
		{Scale: scale.Lonyllic, Name: "Lonyllic", Cardinality: 8, Number: 2525, Perfection: 5, Imperfection: 3},
		{Scale: scale.Sathyllic, Name: "Sathyllic", Cardinality: 8, Number: 3820, Perfection: 5, Imperfection: 3},
		{Scale: scale.Layllic, Name: "Layllic", Cardinality: 8, Number: 3545, Perfection: 5, Imperfection: 3},
		{Scale: scale.Saryllic, Name: "Saryllic", Cardinality: 8, Number: 2995, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thacryllic, Name: "Thacryllic", Cardinality: 8, Number: 3790, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolynyllic, Name: "Aeolynyllic", Cardinality: 8, Number: 3485, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thadyllic, Name: "Thadyllic", Cardinality: 8, Number: 2875, Perfection: 5, Imperfection: 3},
		{Scale: scale.Lynyllic, Name: "Lynyllic", Cardinality: 8, Number: 3310, Perfection: 5, Imperfection: 3},

		// 8.27
		{Scale: scale.Aeolathyllic, Name: "Aeolathyllic", Cardinality: 8, Number: 2541, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolocryllic, Name: "Aeolocryllic", Cardinality: 8, Number: 3948, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phroptyllic, Name: "Phroptyllic", Cardinality: 8, Number: 3801, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kodyllic, Name: "Kodyllic", Cardinality: 8, Number: 3507, Perfection: 5, Imperfection: 3},
		{Scale: scale.Epaptyllic, Name: "Epaptyllic", Cardinality: 8, Number: 2919, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionoyllic, Name: "Ionoyllic", Cardinality: 8, Number: 3486, Perfection: 5, Imperfection: 3},
		{Scale: scale.Gyptyllic, Name: "Gyptyllic", Cardinality: 8, Number: 2877, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aerythyllic, Name: "Aerythyllic", Cardinality: 8, Number: 3318, Perfection: 5, Imperfection: 3},

		// 8.28
		{Scale: scale.Zagyllic, Name: "Zagyllic", Cardinality: 8, Number: 2542, Perfection: 5, Imperfection: 3},
		{Scale: scale.Epacryllic, Name: "Epacryllic", Cardinality: 8, Number: 3956, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thorcryllic, Name: "Thorcryllic", Cardinality: 8, Number: 3817, Perfection: 5, Imperfection: 3},
		{Scale: scale.Loptyllic, Name: "Loptyllic", Cardinality: 8, Number: 3539, Perfection: 5, Imperfection: 3},
		{Scale: scale.Katylyllic, Name: "Katylyllic", Cardinality: 8, Number: 2983, Perfection: 5, Imperfection: 3},
		{Scale: scale.Malyllic, Name: "Malyllic", Cardinality: 8, Number: 3742, Perfection: 5, Imperfection: 3},
		{Scale: scale.Mydyllic, Name: "Mydyllic", Cardinality: 8, Number: 3389, Perfection: 5, Imperfection: 3},
		{Scale: scale.Thycryllic, Name: "Thycryllic", Cardinality: 8, Number: 2683, Perfection: 5, Imperfection: 3},

		// 8.29
		{Scale: scale.Gythyllic, Name: "Gythyllic", Cardinality: 8, Number: 2549, Perfection: 5, Imperfection: 3},
		{Scale: scale.Pyryllic, Name: "Pyryllic", Cardinality: 8, Number: 4012, Perfection: 5, Imperfection: 3},
		{Scale: scale.Rycryllic, Name: "Rycryllic", Cardinality: 8, Number: 3929, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrathyllic, Name: "Phrathyllic", Cardinality: 8, Number: 3763, Perfection: 5, Imperfection: 3},
		{Scale: scale.Badyllic, Name: "Badyllic", Cardinality: 8, Number: 3431, Perfection: 5, Imperfection: 3},
		{Scale: scale.Phrocryllic, Name: "Phrocryllic", Cardinality: 8, Number: 2767, Perfection: 5, Imperfection: 3},
		{Scale: scale.Staryllic, Name: "Staryllic", Cardinality: 8, Number: 2878, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zothyllic, Name: "Zothyllic", Cardinality: 8, Number: 3322, Perfection: 5, Imperfection: 3},

		// 8.30
		{Scale: scale.Tharyllic, Name: "Tharyllic", Cardinality: 8, Number: 2550, Perfection: 5, Imperfection: 3},
		{Scale: scale.Sylyllic, Name: "Sylyllic", Cardinality: 8, Number: 4020, Perfection: 5, Imperfection: 3},
		{Scale: scale.Lothyllic, Name: "Lothyllic", Cardinality: 8, Number: 3945, Perfection: 5, Imperfection: 3},
		{Scale: scale.Daryllic, Name: "Daryllic", Cardinality: 8, Number: 3795, Perfection: 5, Imperfection: 3},
		{Scale: scale.Monyllic, Name: "Monyllic", Cardinality: 8, Number: 3495, Perfection: 5, Imperfection: 3},
		{Scale: scale.Styryllic, Name: "Styryllic", Cardinality: 8, Number: 2895, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolacryllic, Name: "Aeolacryllic", Cardinality: 8, Number: 3390, Perfection: 5, Imperfection: 3},
		{Scale: scale.Raptyllic, Name: "Raptyllic", Cardinality: 8, Number: 2685, Perfection: 5, Imperfection: 3},

		// 8.31
		{Scale: scale.Kataryllic, Name: "Kataryllic", Cardinality: 8, Number: 2554, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aerocryllic, Name: "Aerocryllic", Cardinality: 8, Number: 4052, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zanyllic, Name: "Zanyllic", Cardinality: 8, Number: 4009, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolonyllic, Name: "Aeolonyllic", Cardinality: 8, Number: 3923, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeonyllic, Name: "Aeonyllic", Cardinality: 8, Number: 3751, Perfection: 5, Imperfection: 3},
		{Scale: scale.Kyryllic, Name: "Kyryllic", Cardinality: 8, Number: 3407, Perfection: 5, Imperfection: 3},
		{Scale: scale.Sythyllic, Name: "Sythyllic", Cardinality: 8, Number: 2719, Perfection: 5, Imperfection: 3},
		{Scale: scale.Katycryllic, Name: "Katycryllic", Cardinality: 8, Number: 2686, Perfection: 5, Imperfection: 3},

		// 8.32
		{Scale: scale.Stogyllic, Name: "Stogyllic", Cardinality: 8, Number: 2779, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionidyllic, Name: "Ionidyllic", Cardinality: 8, Number: 2926, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stonyllic, Name: "Stonyllic", Cardinality: 8, Number: 3514, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stalyllic, Name: "Stalyllic", Cardinality: 8, Number: 2933, Perfection: 5, Imperfection: 3},
		{Scale: scale.Poryllic, Name: "Poryllic", Cardinality: 8, Number: 3542, Perfection: 5, Imperfection: 3},
		{Scale: scale.Mocryllic, Name: "Mocryllic", Cardinality: 8, Number: 2989, Perfection: 5, Imperfection: 3},
		{Scale: scale.Aeolyryllic, Name: "Aeolyryllic", Cardinality: 8, Number: 3766, Perfection: 5, Imperfection: 3},
		{Scale: scale.Baryllic, Name: "Baryllic", Cardinality: 8, Number: 3437, Perfection: 5, Imperfection: 3},

		// 8.33
		{Scale: scale.Dalyllic, Name: "Dalyllic", Cardinality: 8, Number: 2797, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionyphyllic, Name: "Ionyphyllic", Cardinality: 8, Number: 2998, Perfection: 5, Imperfection: 3},
		{Scale: scale.Zaptyllic, Name: "Zaptyllic", Cardinality: 8, Number: 3802, Perfection: 5, Imperfection: 3},
		{Scale: scale.Garyllic, Name: "Garyllic", Cardinality: 8, Number: 3509, Perfection: 5, Imperfection: 3},
		{Scale: scale.Gathyllic, Name: "Gathyllic", Cardinality: 8, Number: 2923, Perfection: 5, Imperfection: 3},
		{Scale: scale.Mixopyryllic, Name: "Mixopyryllic", Cardinality: 8, Number: 3502, Perfection: 5, Imperfection: 3},
		{Scale: scale.Ionacryllic, Name: "Ionacryllic", Cardinality: 8, Number: 2909, Perfection: 5, Imperfection: 3},
		{Scale: scale.Stylyllic, Name: "Stylyllic", Cardinality: 8, Number: 3446, Perfection: 5, Imperfection: 3},

		// 8.34
		{Scale: scale.Stycryllic, Name: "Stycryllic", Cardinality: 8, Number: 2239, Perfection: 4, Imperfection: 4},
		{Scale: scale.Ionothyllic, Name: "Ionothyllic", Cardinality: 8, Number: 3064, Perfection: 4, Imperfection: 4},
		{Scale: scale.Mythyllic, Name: "Mythyllic", Cardinality: 8, Number: 4066, Perfection: 4, Imperfection: 4},
		{Scale: scale.Aerylyllic, Name: "Aerylyllic", Cardinality: 8, Number: 4037, Perfection: 4, Imperfection: 4},
		{Scale: scale.Bonyllic, Name: "Bonyllic", Cardinality: 8, Number: 3979, Perfection: 4, Imperfection: 4},
		{Scale: scale.Tholyllic, Name: "Tholyllic", Cardinality: 8, Number: 3863, Perfection: 4, Imperfection: 4},
		{Scale: scale.Katyryllic, Name: "Katyryllic", Cardinality: 8, Number: 3631, Perfection: 4, Imperfection: 4},
		{Scale: scale.Sadyllic, Name: "Sadyllic", Cardinality: 8, Number: 3167, Perfection: 4, Imperfection: 4},

		// 8.35
		{Scale: scale.Stolyllic, Name: "Stolyllic", Cardinality: 8, Number: 2302, Perfection: 4, Imperfection: 4},
		{Scale: scale.Logyllic, Name: "Logyllic", Cardinality: 8, Number: 4072, Perfection: 4, Imperfection: 4},
		{Scale: scale.Dacryllic, Name: "Dacryllic", Cardinality: 8, Number: 4049, Perfection: 4, Imperfection: 4},
		{Scale: scale.Thynyllic, Name: "Thynyllic", Cardinality: 8, Number: 4003, Perfection: 4, Imperfection: 4},
		{Scale: scale.Gydyllic, Name: "Gydyllic", Cardinality: 8, Number: 3911, Perfection: 4, Imperfection: 4},
		{Scale: scale.Eparyllic, Name: "Eparyllic", Cardinality: 8, Number: 3727, Perfection: 4, Imperfection: 4},
		{Scale: scale.Dynyllic, Name: "Dynyllic", Cardinality: 8, Number: 3359, Perfection: 4, Imperfection: 4},
		{Scale: scale.Ionyllic, Name: "Ionyllic", Cardinality: 8, Number: 2623, Perfection: 4, Imperfection: 4},

		// 8.36
		{Scale: scale.Zaryllic, Name: "Zaryllic", Cardinality: 8, Number: 2367, Perfection: 4, Imperfection: 4},
		{Scale: scale.Dythyllic, Name: "Dythyllic", Cardinality: 8, Number: 2556, Perfection: 4, Imperfection: 4},
		{Scale: scale.Ionaryllic, Name: "Ionaryllic", Cardinality: 8, Number: 4068, Perfection: 4, Imperfection: 4},
		{Scale: scale.Laryllic, Name: "Laryllic", Cardinality: 8, Number: 4041, Perfection: 4, Imperfection: 4},
		{Scale: scale.Kataptyllic, Name: "Kataptyllic", Cardinality: 8, Number: 3987, Perfection: 4, Imperfection: 4},
		{Scale: scale.Sonyllic, Name: "Sonyllic", Cardinality: 8, Number: 3879, Perfection: 4, Imperfection: 4},
		{Scale: scale.Pathyllic, Name: "Pathyllic", Cardinality: 8, Number: 3663, Perfection: 4, Imperfection: 4},
		{Scale: scale.Loryllic, Name: "Loryllic", Cardinality: 8, Number: 3231, Perfection: 4, Imperfection: 4},

		// 8.37
		{Scale: scale.Aeronyllic, Name: "Aeronyllic", Cardinality: 8, Number: 2429, Perfection: 4, Imperfection: 4},
		{Scale: scale.Pycryllic, Name: "Pycryllic", Cardinality: 8, Number: 3052, Perfection: 4, Imperfection: 4},
		{Scale: scale.Mygyllic, Name: "Mygyllic", Cardinality: 8, Number: 4018, Perfection: 4, Imperfection: 4},
		{Scale: scale.Lylyllic, Name: "Lylyllic", Cardinality: 8, Number: 3941, Perfection: 4, Imperfection: 4},
		{Scale: scale.Daptyllic, Name: "Daptyllic", Cardinality: 8, Number: 3787, Perfection: 4, Imperfection: 4},
		{Scale: scale.Ioninyllic, Name: "Ioninyllic", Cardinality: 8, Number: 3479, Perfection: 4, Imperfection: 4},
		{Scale: scale.Epaphyllic, Name: "Epaphyllic", Cardinality: 8, Number: 2863, Perfection: 4, Imperfection: 4},
		{Scale: scale.Lolyllic, Name: "Lolyllic", Cardinality: 8, Number: 3262, Perfection: 4, Imperfection: 4},

		// 8.38
		{Scale: scale.Stacryllic, Name: "Stacryllic", Cardinality: 8, Number: 2494, Perfection: 4, Imperfection: 4},
		{Scale: scale.Doryllic, Name: "Doryllic", Cardinality: 8, Number: 3572, Perfection: 4, Imperfection: 4},
		{Scale: scale.Kadyllic, Name: "Kadyllic", Cardinality: 8, Number: 3049, Perfection: 4, Imperfection: 4},
		{Scale: scale.Rynyllic, Name: "Rynyllic", Cardinality: 8, Number: 4006, Perfection: 4, Imperfection: 4},
		{Scale: scale.Aerogyllic, Name: "Aerogyllic", Cardinality: 8, Number: 3917, Perfection: 4, Imperfection: 4},
		{Scale: scale.Rothyllic, Name: "Rothyllic", Cardinality: 8, Number: 3739, Perfection: 4, Imperfection: 4},
		{Scale: scale.Kagyllic, Name: "Kagyllic", Cardinality: 8, Number: 3383, Perfection: 4, Imperfection: 4},
		{Scale: scale.Stathyllic, Name: "Stathyllic", Cardinality: 8, Number: 2671, Perfection: 4, Imperfection: 4},

		// 8.39
		{Scale: scale.Thyryllic, Name: "Thyryllic", Cardinality: 8, Number: 2735, Perfection: 4, Imperfection: 4},
		{Scale: scale.Gygyllic, Name: "Gygyllic", Cardinality: 8, Number: 2750, Perfection: 4, Imperfection: 4},
		{Scale: scale.Sodyllic, Name: "Sodyllic", Cardinality: 8, Number: 2810, Perfection: 4, Imperfection: 4},
		{Scale: scale.Goryllic, Name: "Goryllic", Cardinality: 8, Number: 3050, Perfection: 4, Imperfection: 4},
		{Scale: scale.Bothyllic, Name: "Bothyllic", Cardinality: 8, Number: 4010, Perfection: 4, Imperfection: 4},
		{Scale: scale.Gynyllic, Name: "Gynyllic", Cardinality: 8, Number: 3925, Perfection: 4, Imperfection: 4},
		{Scale: scale.Ionaptyllic, Name: "Ionaptyllic", Cardinality: 8, Number: 3755, Perfection: 4, Imperfection: 4},
		{Scale: scale.Phryryllic, Name: "Phryryllic", Cardinality: 8, Number: 3415, Perfection: 4, Imperfection: 4},

		// 8.40
		{Scale: scale.Racryllic, Name: "Racryllic", Cardinality: 8, Number: 2747, Perfection: 4, Imperfection: 4},
		{Scale: scale.Epicryllic, Name: "Epicryllic", Cardinality: 8, Number: 2798, Perfection: 4, Imperfection: 4},
		{Scale: scale.Stygyllic, Name: "Stygyllic", Cardinality: 8, Number: 3002, Perfection: 4, Imperfection: 4},
		{Scale: scale.Syryllic, Name: "Syryllic", Cardinality: 8, Number: 3818, Perfection: 4, Imperfection: 4},
		{Scale: scale.Stythyllic, Name: "Stythyllic", Cardinality: 8, Number: 3541, Perfection: 4, Imperfection: 4},
		{Scale: scale.Aerothyllic, Name: "Aerothyllic", Cardinality: 8, Number: 2987, Perfection: 4, Imperfection: 4},
		{Scale: scale.Mixoryllic, Name: "Mixoryllic", Cardinality: 8, Number: 3758, Perfection: 4, Imperfection: 4},
		{Scale: scale.Thanyllic, Name: "Thanyllic", Cardinality: 8, Number: 3421, Perfection: 4, Imperfection: 4},

		// 8.41
		{Scale: scale.Roryllic, Name: "Roryllic", Cardinality: 8, Number: 2795, Perfection: 4, Imperfection: 4},
		{Scale: scale.Epotyllic, Name: "Epotyllic", Cardinality: 8, Number: 2990, Perfection: 4, Imperfection: 4},
		{Scale: scale.Epidyllic, Name: "Epidyllic", Cardinality: 8, Number: 3770, Perfection: 4, Imperfection: 4},
		{Scale: scale.Kaptyllic, Name: "Kaptyllic", Cardinality: 8, Number: 3445, Perfection: 4, Imperfection: 4},

		// 8.42
		{Scale: scale.MajorDiminished, Name: "MajorDiminished", Cardinality: 8, Number: 2925, Perfection: 4, Imperfection: 4},
		{Scale: scale.MinorDiminished, Name: "MinorDiminished", Cardinality: 8, Number: 3510, Perfection: 4, Imperfection: 4},
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
