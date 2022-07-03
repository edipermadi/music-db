package scale

// Number return scale number
func (s Type) Number() int {
	return [...]int{
		0,

		// 3 notes, reference: reference https://allthescales.org/scales.php?n=3
		2184,

		// 4 notes, reference: reference https://allthescales.org/scales.php?n=4
		2193, 2328, 2244, 3144,
		2196, 2376, 2628, 2322,
		2185, 2200, 2440, 3140,
		2188, 2248, 3208, 2321,
		2194, 2344, 2372, 2596,
		2212, 2632, 2338, 2324,
		2186, 2216, 2696, 2594,
		2210, 2600,
		2340,

		// 5 notes, reference: reference https://allthescales.org/scales.php?n=5
		2378, 2644, 2386, 2708, 2642,
		2197, 2392, 2756, 2834, 3146,
		2213, 2648, 2402, 2836, 3154,
		2225, 2840, 3170, 2245, 3160,
		2246, 3176, 2257, 3352, 2609,
		2258, 3368, 2641, 2374, 2612,
		2260, 3400, 2705, 2630, 2330,
		2189, 2264, 3464, 2833, 3142,
		2195, 2360, 2500, 3620, 3145,
		2198, 2408, 2884, 3346, 2597,
		2201, 2456, 3268, 2441, 3148,
		2204, 2504, 3652, 3209, 2323,
		2211, 2616, 2274, 3624, 3153,
		2217, 2712, 2658, 2442, 3156,
		2228, 2888, 3362, 2629, 2326,
		2249, 3224, 2353, 2444, 3172,
		2250, 3240, 2385, 2700, 2610,
		2252, 3272, 2449, 3212, 2329,
		2276, 3656, 3217, 2339, 2332,
		2345, 2380, 2660, 2450, 3220,
		2346, 2388, 2724, 2706, 2634,
		2354, 2452, 3236, 2377, 2636,
		2187, 2232, 2952, 3618, 3141,
		2190, 2280, 3720, 3345, 2595,
		2202, 2472, 3396, 2697, 2598,
		2214, 2664, 2466, 3348, 2601,
		2220, 2760, 2850, 3210, 2325,
		2226, 2856, 3234, 2373, 2604,
		2341, 2348, 2404, 2852, 3218,
		2342, 2356, 2468, 3364, 2633,
		2218, 2728, 2722, 2698, 2602,

		2394, 2772, 2898, 3402, 2709, 2646,
		2229, 2904, 3426, 2757, 2838, 3162,
		2247, 3192, 2289, 3864, 3633, 3171,
		2259, 3384, 2673, 2502, 3636, 3177,
		2261, 3416, 2737, 2758, 2842, 3178,
		2262, 3432, 2769, 2886, 3354, 2613,
		2275, 3640, 3185,
		2277, 3672, 3249, 2403, 2844, 3186,
		2379, 2652, 2418, 2964, 3666, 3237,
		2382, 2676, 2514, 3732, 3369, 2643,
		2387, 2716, 2674, 2506, 3668, 3241,
		2390, 2740, 2770, 2890, 3370, 2645,
		2410, 2900, 3410, 2725, 2710, 2650,
		2199, 2424, 3012, 3858, 3621, 3147,
		2205, 2520, 3780, 3465, 2835, 3150,
		2215, 2680, 2530, 3860, 3625, 3155,
		2221, 2776, 2914, 3466, 2837, 3158,
		2227, 2872, 3298, 2501, 3628, 3161,
		2233, 2968, 3682, 3269, 2443, 3164,
		2251, 3256, 2417, 2956, 3634, 3173,
		2253, 3288, 2481, 3468, 2841, 3174,
		2254, 3304, 2513, 3724, 3353, 2611,
		2265, 3480, 2865, 3270, 2445, 3180,
		2266, 3496, 2897, 3398, 2701, 2614,
		2268, 3528, 2961, 3654, 3213, 2331,
		2278, 3688, 3281, 2467, 3356, 2617,
		2281, 3736, 3377, 2659, 2446, 3188,
		2290, 3880, 3665, 3235, 2375, 2620,
		2292, 3912, 3729, 3363, 2631, 2334,
		2347, 2396, 2788, 2962, 3658, 3221,
		2355, 2460, 3300, 2505, 3660, 3225,
		2361, 2508, 3684, 3273, 2451, 3228,
		2362, 2516, 3748, 3401, 2707, 2638,
		2393, 2764, 2866, 3274, 2453, 3244,
		2406, 2868, 3282, 2469, 3372, 2649,
		2409, 2892, 3378, 2661, 2454, 3252,
		2457, 3276,
		2458, 3284, 2473, 3404, 2713, 2662,
		2191, 2296, 3976, 3857, 3619, 3143,
		2203, 2488, 3524, 2953, 3622, 3149,
		2206, 2536, 3908, 3721, 3347, 2599,
		2219, 2744, 2786, 2954, 3626, 3157,
		2230, 2920, 3490, 2885, 3350, 2605,
		2236, 3016, 3874, 3653, 3211, 2327,
		2282, 3752, 3409, 2723, 2702, 2618,
		2284, 3784, 3473, 2851, 3214, 2333,
		2343, 2364, 2532, 3876, 3657, 3219,
		2349, 2412, 2916, 3474, 2853, 3222,
		2350, 2420, 2980, 3730, 3365, 2635,
		2357, 2476, 3428, 2761, 2854, 3226,
		2358, 2484, 3492, 2889, 3366, 2637,
		2381, 2668, 2482, 3476, 2857, 3238,
		2389, 2732, 2738, 2762, 2858, 3242,
		2405, 2860, 3250,
		2470, 3380, 2665,
		2474, 3412, 2729, 2726, 2714, 2666,
		2222, 2792, 2978, 3722, 3349, 2603,
		2234, 2984, 3746, 3397, 2699, 2606,
		2730,
	}[s]
}
