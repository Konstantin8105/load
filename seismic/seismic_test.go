package seismic

func Example() {
	f := Factors{
		DesignSeismicSite: 8,
		Grount:            2,
		K0row:             1,
		K1:                0.25,
		Kψrow:             3,
	}
	acs, rs := f.Acceleration()
	_ = acs
	_ = rs
	// Output:
	// Расчет ускорений для сейсмической нагрузки:
	// Расчетная сейсмичность площадки строительства: 8    Баллов |
	// Категория грунта по таблице 4.1:               2           |
	// Строка по таблице 4.2:                         1    строка |
	// Коэффициент K0 при расчете на РЗ:              1.20        |
	// Коэффициент K0 при расчете на КЗ:              2.00        |
	// Коэффициент K1 по таблице 5.2:                 0.25        |
	// Строка по таблице 5.3:                         3    строка |
	// Коэффициент Kψ по таблице 5.3:                 1.00        |
	//
	// Расчёт:
	// Коэффициент A по пункту 5.5: 2.00 м/кв.сек
	//
	// Период   | β      | Ускорение | Ускорение |
	//          |        | при РЗ    | при РЗ    |
	//          |        | K0=1.20   | K0=1.20   |
	//          |        | K1=0.25   | K1=1.00   |
	// сек      |        | м/сек^2   | м/сек^2   |
	//     0.00 | 1.0000 |  0.600    |  2.400    |
	//     0.10 | 2.5000 |  1.500    |  6.000    |
	//     0.40 | 2.5000 |  1.500    |  6.000    |
	//     0.45 | 2.3570 |  1.414    |  5.657    |
	//     0.50 | 2.2361 |  1.342    |  5.367    |
	//     0.55 | 2.1320 |  1.279    |  5.117    |
	//     0.60 | 2.0412 |  1.225    |  4.899    |
	//     0.65 | 1.9612 |  1.177    |  4.707    |
	//     0.70 | 1.8898 |  1.134    |  4.536    |
	//     0.75 | 1.8257 |  1.095    |  4.382    |
	//     0.80 | 1.7678 |  1.061    |  4.243    |
	//     0.85 | 1.7150 |  1.029    |  4.116    |
	//     0.90 | 1.6667 |  1.000    |  4.000    |
	//     0.95 | 1.6222 |  0.973    |  3.893    |
	//     1.00 | 1.5811 |  0.949    |  3.795    |
	//     1.05 | 1.5430 |  0.926    |  3.703    |
	//     1.10 | 1.5076 |  0.905    |  3.618    |
	//     1.15 | 1.4744 |  0.885    |  3.539    |
	//     1.20 | 1.4434 |  0.866    |  3.464    |
	//     1.25 | 1.4142 |  0.849    |  3.394    |
	//     1.30 | 1.3868 |  0.832    |  3.328    |
	//     1.35 | 1.3608 |  0.816    |  3.266    |
	//     1.40 | 1.3363 |  0.802    |  3.207    |
	//     1.45 | 1.3131 |  0.788    |  3.151    |
	//     1.50 | 1.2910 |  0.775    |  3.098    |
	//     1.55 | 1.2700 |  0.762    |  3.048    |
	//     1.60 | 1.2500 |  0.750    |  3.000    |
	//     1.65 | 1.2309 |  0.739    |  2.954    |
	//     1.70 | 1.2127 |  0.728    |  2.910    |
	//     1.75 | 1.1952 |  0.717    |  2.869    |
	//     1.80 | 1.1785 |  0.707    |  2.828    |
	//     1.85 | 1.1625 |  0.697    |  2.790    |
	//     1.90 | 1.1471 |  0.688    |  2.753    |
	//     1.95 | 1.1323 |  0.679    |  2.717    |
	//     2.00 | 1.1180 |  0.671    |  2.683    |
	//     2.05 | 1.1043 |  0.663    |  2.650    |
	//     2.10 | 1.0911 |  0.655    |  2.619    |
	//     2.15 | 1.0783 |  0.647    |  2.588    |
	//     2.20 | 1.0660 |  0.640    |  2.558    |
	//     2.25 | 1.0541 |  0.632    |  2.530    |
	//     2.30 | 1.0426 |  0.626    |  2.502    |
	//     2.35 | 1.0314 |  0.619    |  2.475    |
	//     2.40 | 1.0206 |  0.612    |  2.449    |
	//     2.45 | 1.0102 |  0.606    |  2.424    |
	//     2.50 | 1.0000 |  0.600    |  2.400    |
	//     2.60 | 0.9806 |  0.588    |  2.353    |
	//     2.70 | 0.9623 |  0.577    |  2.309    |
	//     2.80 | 0.9449 |  0.567    |  2.268    |
	//     2.90 | 0.9285 |  0.557    |  2.228    |
	//     3.00 | 0.9129 |  0.548    |  2.191    |
	//     3.10 | 0.8980 |  0.539    |  2.155    |
	//     3.20 | 0.8839 |  0.530    |  2.121    |
	//     3.30 | 0.8704 |  0.522    |  2.089    |
	//     3.40 | 0.8575 |  0.514    |  2.058    |
	//     3.50 | 0.8452 |  0.507    |  2.028    |
	//     3.60 | 0.8333 |  0.500    |  2.000    |
	//     3.70 | 0.8220 |  0.493    |  1.973    |
	//     3.80 | 0.8111 |  0.487    |  1.947    |
	//     3.90 | 0.8006 |  0.480    |  1.922    |
	//  2561.00 | 0.8000 |  0.480    |  1.920    |
	//
	// Коэффициенты соотношения к расчету РЗ:
	// (ускорение для РЗ при K1 = K1)  / (ускорение для РЗ при K1 = 1.0) 0.250
	// (ускорение для РЗ при K1 = 1.0) / (ускорение для РЗ при K1 = 1.0) 1.000
}
