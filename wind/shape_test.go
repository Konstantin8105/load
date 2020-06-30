package wind

func ExampleRectangle() {
	Rectangle(ZoneA, RegionII, LogDecriment15, 5.38, 7.32, 18.965, 0.000, []float64{1.393})

	// Output:
	// Sketch:
	//
	//           |<------- d --------->|
	//           |                     |
	//           ***********************------
	//           *                     *    |
	//   Wind    *                     *    |
	//  ----->   *                     *    |
	//         D *                     * E  b
	//           *                     *    |
	//           *                     *    |
	//           *                     *    |
	//           ***********************------
	//           |  A  |    B    |  C  |
	//
	// Wind zone: A
	// Wind region:  II with value = 300.0 Pa
	// Wind log decrement: δ = 0.15
	// Natural frequency : [1.393]
	//
	// Dimensions:
	// b   5.380 m
	// d   7.320 m
	// zo  0.000 m
	// h  18.965 m
	//
	// | side z      ze     Kz     ζ      ξ      | cx   ρ      χ      ν      Wm     Wp     Wsum   |
	// |                                         |                                                |
	// | A     0.000  7.320  0.911  0.796  1.582 | -1.0  2.928 18.965  0.860 -273.2 -296.2 -569.4 |
	// | A     5.000  7.320  0.911  0.796  1.582 | -1.0  2.928 18.965  0.860 -273.2 -296.2 -569.4 |
	// | A    10.000 10.000  1.000  0.760  1.582 | -1.0  2.928 18.965  0.860 -300.0 -310.4 -610.4 |
	// | A    15.000 18.965  1.212  0.690  1.582 | -1.0  2.928 18.965  0.860 -363.5 -341.6 -705.1 |
	// | A    18.965 18.965  1.212  0.690  1.582 | -1.0  2.928 18.965  0.860 -363.5 -341.6 -705.1 |
	// |                                         |                                                |
	// | B     0.000  7.320  0.911  0.796  1.582 | -0.8  2.928 18.965  0.860 -218.6 -236.9 -455.5 |
	// | B     5.000  7.320  0.911  0.796  1.582 | -0.8  2.928 18.965  0.860 -218.6 -236.9 -455.5 |
	// | B    10.000 10.000  1.000  0.760  1.582 | -0.8  2.928 18.965  0.860 -240.0 -248.3 -488.3 |
	// | B    15.000 18.965  1.212  0.690  1.582 | -0.8  2.928 18.965  0.860 -290.8 -273.3 -564.1 |
	// | B    18.965 18.965  1.212  0.690  1.582 | -0.8  2.928 18.965  0.860 -290.8 -273.3 -564.1 |
	// |                                         |                                                |
	// | C     0.000  7.320  0.911  0.796  1.582 | -0.5  2.928 18.965  0.860 -136.6 -148.1 -284.7 |
	// | C     5.000  7.320  0.911  0.796  1.582 | -0.5  2.928 18.965  0.860 -136.6 -148.1 -284.7 |
	// | C    10.000 10.000  1.000  0.760  1.582 | -0.5  2.928 18.965  0.860 -150.0 -155.2 -305.2 |
	// | C    15.000 18.965  1.212  0.690  1.582 | -0.5  2.928 18.965  0.860 -181.8 -170.8 -352.6 |
	// | C    18.965 18.965  1.212  0.690  1.582 | -0.5  2.928 18.965  0.860 -181.8 -170.8 -352.6 |
	// |                                         |                                                |
	// | D     0.000  5.380  0.830  0.834  1.582 |  0.8  5.380 18.965  0.841  199.3  221.1  420.4 |
	// | D     5.000  5.380  0.830  0.834  1.582 |  0.8  5.380 18.965  0.841  199.3  221.1  420.4 |
	// | D    10.000 10.000  1.000  0.760  1.582 |  0.8  5.380 18.965  0.841  240.0  242.6  482.6 |
	// | D    15.000 18.965  1.212  0.690  1.582 |  0.8  5.380 18.965  0.841  290.8  267.1  557.9 |
	// | D    18.965 18.965  1.212  0.690  1.582 |  0.8  5.380 18.965  0.841  290.8  267.1  557.9 |
	// |                                         |                                                |
	// | E     0.000  5.380  0.830  0.834  1.582 | -0.5  5.380 18.965  0.841 -124.5 -138.2 -262.7 |
	// | E     5.000  5.380  0.830  0.834  1.582 | -0.5  5.380 18.965  0.841 -124.5 -138.2 -262.7 |
	// | E    10.000 10.000  1.000  0.760  1.582 | -0.5  5.380 18.965  0.841 -150.0 -151.6 -301.6 |
	// | E    15.000 18.965  1.212  0.690  1.582 | -0.5  5.380 18.965  0.841 -181.8 -166.9 -348.7 |
	// | E    18.965 18.965  1.212  0.690  1.582 | -0.5  5.380 18.965  0.841 -181.8 -166.9 -348.7 |
	//
	//    Ws on top    |----------->             |--------->
	//                 |          /              |         |
	//                 |--------->    Ws average |--------->
	//                 |        /                |         |
	//    Ws on zero   |------->                 |--------->
	//               --------------- ground ------------------
	//
	// | side   | width  | Ws on zero | Ws on top | Center of Ws | Ws average |
	// |        |        | elevation  | elevation |              |            |
	// |        | meter  | Pa         | Pa        | meter        | Pa         |
	// |        |        |            |           |              |            |
	// |      A |  1.076 |   -535.5   |   -717.4  |  9.941       |   -656.8   |
	// |      B |  4.304 |   -428.4   |   -573.9  |  9.941       |   -525.4   |
	// |      C |  1.940 |   -267.8   |   -358.7  |  9.941       |   -328.4   |
	// |      D |  5.380 |   +390.7   |   +576.6  | 10.090       |   +514.6   |
	// |      E |  5.380 |   -244.2   |   -360.4  | 10.090       |   -321.7   |
}

func ExampleCylinder() {

	Cylinder(ZoneA, RegionII, LogDecriment15, 0.200, 4.710, 10.100, 2.800,
		[]float64{3.091, 3.414, 3.719})
	// Output:
	// Sketch:
	//
	//           |<--- d ----->|
	//           |             |
	//           ***************-------
	//           *             *     |
	//   Wind    *             *     |
	//  ----->   *             *     |
	//           *             *     |
	//           *             *     |
	//           *             *     |
	//           *             *     h
	//           ***************---  |
	//                           |   |
	//                           zo  |
	//                           |   |
	//    ---------- ground -------------
	//
	// Wind zone: A
	// Wind region:  II with value = 300.0 Pa
	// Wind log decrement: δ = 0.15
	// Natural frequency : [3.091 3.414 3.719]
	//
	// Dimensions:
	// b   4.710 m
	// d   4.710 m
	// zo  2.800 m
	// h  10.100 m
	//
	// Re = 86.160*10^5 for ze=0.8*h =  8.640
	//
	// Cx∞ =  1.200
	//
	// Elongation:
	// λ   1.550
	// λe  3.100
	// ϕ   1.000
	// Kλ  0.649
	//
	// Cx  =  0.779
	//
	// The spatial correlation coefficient of pressure pulsations:
	// ρ  4.710
	// χ 10.100
	// ν  0.873
	//
	// | z      ze     Kz     ζ      ξ      | Wm     Wp     Wsum   |
	// | m      m                           | Pa     Pa     Pa     |
	// |                                    |                      |
	// |  2.800  2.800  0.683  0.920  1.318 |  159.5  168.8  328.3 |
	// |  5.000  5.000  0.812  0.843  1.318 |  189.8  184.1  373.9 |
	// | 10.000 10.000  1.000  0.760  1.318 |  233.7  204.3  438.0 |
	// | 10.100 10.100  1.003  0.759  1.318 |  234.4  204.6  439.0 |
}
