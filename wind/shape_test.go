package wind

func ExampleRectangle() {
	Rectangle(ZoneA, RegionII, LogDecriment15, 5.38, 7.32, 18.965, []float64{1.393})

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
	// b  5.380 m
	// d  7.320 m
	// h 18.965 m
	//
	// | side z      ze     Kz     ζ      ξ      | cx   ρ      χ      ν      Wm     Wp     Wsum   |
	// |                                         |                                                |
	// | A     0.000  7.320  0.911  0.796  0.945 | -1.0  2.928 18.965  0.860 -273.2 -176.9 -450.1 |
	// | A     5.000  7.320  0.911  0.796  1.532 | -1.0  2.928 18.965  0.860 -273.2 -286.8 -560.0 |
	// | A    10.000 10.000  1.000  0.760  1.564 | -1.0  2.928 18.965  0.860 -300.0 -306.8 -606.8 |
	// | A    15.000 18.965  1.212  0.690  1.582 | -1.0  2.928 18.965  0.860 -363.5 -341.6 -705.1 |
	// | A    18.965 18.965  1.212  0.690  1.592 | -1.0  2.928 18.965  0.860 -363.5 -343.8 -707.3 |
	// |                                         |                                                |
	// | B     0.000  7.320  0.911  0.796  0.945 | -0.8  2.928 18.965  0.860 -218.6 -141.5 -360.1 |
	// | B     5.000  7.320  0.911  0.796  1.532 | -0.8  2.928 18.965  0.860 -218.6 -229.5 -448.0 |
	// | B    10.000 10.000  1.000  0.760  1.564 | -0.8  2.928 18.965  0.860 -240.0 -245.5 -485.5 |
	// | B    15.000 18.965  1.212  0.690  1.582 | -0.8  2.928 18.965  0.860 -290.8 -273.3 -564.1 |
	// | B    18.965 18.965  1.212  0.690  1.592 | -0.8  2.928 18.965  0.860 -290.8 -275.0 -565.8 |
	// |                                         |                                                |
	// | C     0.000  7.320  0.911  0.796  0.945 | -0.5  2.928 18.965  0.860 -136.6  -88.5 -225.1 |
	// | C     5.000  7.320  0.911  0.796  1.532 | -0.5  2.928 18.965  0.860 -136.6 -143.4 -280.0 |
	// | C    10.000 10.000  1.000  0.760  1.564 | -0.5  2.928 18.965  0.860 -150.0 -153.4 -303.4 |
	// | C    15.000 18.965  1.212  0.690  1.582 | -0.5  2.928 18.965  0.860 -181.8 -170.8 -352.6 |
	// | C    18.965 18.965  1.212  0.690  1.592 | -0.5  2.928 18.965  0.860 -181.8 -171.9 -353.6 |
	// |                                         |                                                |
	// | D     0.000  5.380  0.830  0.834  0.945 |  0.8  5.380 18.965  0.841  199.3  132.1  331.3 |
	// | D     5.000  5.380  0.830  0.834  1.532 |  0.8  5.380 18.965  0.841  199.3  214.1  413.4 |
	// | D    10.000 10.000  1.000  0.760  1.564 |  0.8  5.380 18.965  0.841  240.0  239.9  479.9 |
	// | D    15.000 18.965  1.212  0.690  1.582 |  0.8  5.380 18.965  0.841  290.8  267.1  557.9 |
	// | D    18.965 18.965  1.212  0.690  1.592 |  0.8  5.380 18.965  0.841  290.8  268.8  559.6 |
	// |                                         |                                                |
	// | E     0.000  5.380  0.830  0.834  0.945 | -0.5  5.380 18.965  0.841 -124.5  -82.5 -207.1 |
	// | E     5.000  5.380  0.830  0.834  1.532 | -0.5  5.380 18.965  0.841 -124.5 -133.8 -258.4 |
	// | E    10.000 10.000  1.000  0.760  1.564 | -0.5  5.380 18.965  0.841 -150.0 -149.9 -299.9 |
	// | E    15.000 18.965  1.212  0.690  1.582 | -0.5  5.380 18.965  0.841 -181.8 -166.9 -348.7 |
	// | E    18.965 18.965  1.212  0.690  1.592 | -0.5  5.380 18.965  0.841 -181.8 -168.0 -349.7 |
	//
	//    Ws on top    |----------->             |--------->
	//                 |          /              |         |
	//                 |--------->    Ws average |--------->
	//                 |        /                |         |
	//    Ws on zero   |------->                 |--------->
	//               --------------- ground ------------------
	//
	// | side   | width  | Ws on zero | Ws on top | Center of Ws | Wind average |
	// |        |        | elevation  | elevation |              | h/2 elev.    |
	// |        | meter  | Pa         | Pa        | meter        | Pa           |
	// |        |        |            |           |              |              |
	// |      A |  1.076 |   -473.8   |   -741.3  | 10.178       |   -652.1     |
	// |      B |  4.304 |   -379.1   |   -593.0  | 10.178       |   -521.7     |
	// |      C |  1.940 |   -236.9   |   -370.6  | 10.178       |   -326.1     |
	// |      D |  5.380 |   +344.6   |   +594.5  | 10.324       |   +511.2     |
	// |      E |  5.380 |   -215.4   |   -371.6  | 10.324       |   -319.5     |
}
