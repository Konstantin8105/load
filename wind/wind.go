package wind

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/pow"
)

type Region float64

// TODO : add unit
const (
	RegionIa  Region = 170.0
	RegionI          = 230.0
	RegionII         = 300.0
	RegionIII        = 380.0
	RegionIV         = 480.0
	RegionV          = 600.0
	RegionVI         = 730.0
	RegionVII        = 850.0
)

func (wr Region) String() string {
	// TODO: add translation
	var name string
	switch wr {
	case RegionIa:
		name = "Ia"
	case RegionI:
		name = "I"
	case RegionII:
		name = "II"
	case RegionIII:
		name = "III"
	case RegionIV:
		name = "IV"
	case RegionV:
		name = "V"
	case RegionVI:
		name = "VI"
	case RegionVII:
		name = "VII"
	default:
		name = "undefine"
		panic(wr)
	}

	return fmt.Sprintf("Wind region: %s with value = %.2f", name, float64(wr)) // TODO: add unit
}

type Zone byte

// add description
const (
	ZoneA = 'A'
	ZoneB = 'B'
	ZoneC = 'C'
)

func (z Zone) String() string {
	// TODO: add translation
	return "Wind zone: " + string(z)
}

func (z Zone) IsValid() bool {
	switch z {
	case ZoneA, ZoneB, ZoneC:
		return true
	}
	return false
}

func (z Zone) constants() (α, k10, ζ10 float64) {
	switch z {
	case ZoneA:
		α, k10, ζ10 = 0.15, 1.00, 0.76
	case ZoneB:
		α, k10, ζ10 = 0.20, 0.65, 1.06
	case ZoneC:
		α, k10, ζ10 = 0.25, 0.40, 1.78
	default:
		panic("not implemented")
	}
	return
}

type LogDecriment float64

const (
	LogDecriment15 LogDecriment = 0.15
	LogDecriment30              = 0.30
)

func (ld LogDecriment) String() string {
	// TODO: add translation
	return fmt.Sprintf("Wind log decriment: %.2f", float64(ld))
}

// TODO : add code name

// EffectiveHeigth by par 11.1.5
//	z - height from ground / высота от поверхности земли
//	d - dimension of building (perpendicular) / размер здания (без учета его
//		стилобатной части) в направлении, перпендикулярном расчетному
//		направлению ветра (поперечный размер).
//	h - heigth of building / высота здания.
func EffectiveHeigth(z, d, h float64, isTower bool) (ze float64) {
	if isTower {
		return z
	}
	if h <= d {
		return h
	}
	if d <= h && h <= 2*d {
		if 0 <= z && z <= h-d {
			return d
		}
		return h
	}
	// 2d < h
	if z >= h-d {
		return h
	}
	if d <= z && z <= h-d {
		return z
	}
	if 0 <= z && z <= d {
		return d
	}
	panic(fmt.Errorf("Not implemented %v %v %v", z, d, h))
}

func FactorKz(zone Zone, ze float64) (float64, error) {
	// TODO: add error handling
	// 	if !zone.IsValid() {
	// 		panic("not valid zone")
	// 	}
	// 	if !validHeigt(ze) {
	// 		panic("not valid heigth")
	// 	}
	// TODO: add link
	// table 11.2
	// formula 11.4
	α, k10, ζ10 := zone.constants()
	_ = ζ10
	return k10 * math.Pow(ze/10.0, 2*α), nil
}

func FactorZeta(zone Zone, ze float64) (float64, error) {
	// TODO: add error handling
	// 	if !zone.IsValid() {
	// 		panic("not valid zone")
	// 	}
	// 	if !validHeigt(ze) {
	// 		panic("not valid heigth")
	// 	}
	// TODO: add link
	// table 11.3
	// formula 11.6
	α, k10, ζ10 := zone.constants()
	_ = k10
	return ζ10 * math.Pow(ze/10.0, -α), nil
}

// Table 11.5
// par 11.1.10
func NaturalFrequencyLimit(wr Region, ld LogDecriment) (float64, error) {
	// TODO: add error handling
	for _, f := range []struct {
		value30, value15 float64
		wr               Region
	}{
		{0.85, 2.60, RegionIa},
		{0.95, 2.90, RegionI},
		{1.10, 3.40, RegionII},
		{1.20, 3.80, RegionIII},
		{1.40, 4.30, RegionIV},
		{1.60, 5.00, RegionV},
		{1.70, 5.60, RegionVI},
		{1.90, 5.90, RegionVII},
	} {
		if wr != f.wr {
			continue
		}
		if ld == LogDecriment30 {
			return f.value30, nil
		}
		if ld == LogDecriment15 {
			return f.value15, nil
		}
	}

	return -1.0, fmt.Errorf("not found")
}

// pic 11.1
func FactorXi(ld LogDecriment, ε float64) (ξ float64) {
	switch ld {
	case LogDecriment30:
		return 189848.0*pow.En(ε, 6) +
			-109948.0*pow.En(ε, 5) +
			+21029.30*pow.En(ε, 4) +
			-1001.640*pow.En(ε, 3) +
			-144.7600*pow.En(ε, 2) +
			+22.09300*pow.En(ε, 1) +
			0.920215
	case LogDecriment15:
		return -2.18484e+8*pow.En(ε, 8) +
			+1.87100e+8*pow.En(ε, 7) +
			-6.63893e+7*pow.En(ε, 6) +
			+1.26134e+7*pow.En(ε, 5) +
			-1.38471e+6*pow.En(ε, 4) +
			+88640.3*pow.En(ε, 3) +
			-3231.09*pow.En(ε, 2) +
			+72.9729*pow.En(ε, 1) +
			0.94572
	}
	panic("not implemented")
}

// TODO: add godoc for all function

func FactoNu(ρ, χ float64) (ν float64) {
	// table 11.6

	const (
		col = 7
		row = 7
	)

	var (
		header = [col]float64{5, 10, 20, 40, 80, 160, 350}
		ro     = [row]float64{0.1, 5, 10, 20, 40, 80, 160}
		data   = [row][col]float64{
			[col]float64{0.95, 0.92, 0.88, 0.83, 0.76, 0.67, 0.56},
			[col]float64{0.89, 0.87, 0.84, 0.80, 0.73, 0.65, 0.54},
			[col]float64{0.85, 0.84, 0.81, 0.77, 0.71, 0.64, 0.53},
			[col]float64{0.80, 0.78, 0.76, 0.73, 0.68, 0.61, 0.51},
			[col]float64{0.72, 0.72, 0.70, 0.67, 0.63, 0.57, 0.48},
			[col]float64{0.63, 0.63, 0.61, 0.59, 0.56, 0.51, 0.44},
			[col]float64{0.53, 0.53, 0.52, 0.50, 0.47, 0.44, 0.38},
		}
	)

	// check outside table
	if χ < header[0] {
		χ = header[0]
	}
	if χ > header[col-1] {
		χ = header[col-1]
	}
	if ρ < ro[0] {
		ρ = ro[0]
	}
	if ρ > ro[row-1] {
		ρ = ro[row-1]
	}
	// parameters now in table

	// generate a column
	var xicol [row]float64
	colIndex := col - 1
	for c := 0; c < col; c++ {
		if χ == header[c] {
			colIndex = c
			break
		}
		if c != 0 {
			if header[c-1] < χ && χ < header[c] {
				colIndex = c
				break
			}
		}
	}
	if colIndex == 0 {
		for r := 0; r < row; r++ {
			xicol[r] = data[r][colIndex]
		}
	} else {
		for r := 0; r < row; r++ {
			xicol[r] = data[r][colIndex-1] +
				(data[r][colIndex]-data[r][colIndex-1])*
					(χ-header[colIndex-1])/
					(header[colIndex]-header[colIndex-1])
		}
	}

	// find by row
	rowIndex := row - 1
	for r := 0; r < row; r++ {
		if ρ == ro[r] {
			rowIndex = r
			break
		}
		if r != 0 {
			if ro[r-1] < ρ && ρ < ro[r] {
				rowIndex = r
				break
			}
		}
	}
	if rowIndex == 0 {
		ν = xicol[rowIndex]
	} else {
		ν = xicol[rowIndex-1] +
			(xicol[rowIndex]-xicol[rowIndex-1])*
				(ρ-ro[rowIndex-1])/
				(ro[rowIndex]-ro[rowIndex-1])
	}
	return
}

//
// double SNiP2_01_07_p6_7b_Eta(double Wo, double Frequency, )
// {
//     double eta = sqrt(1.4 * Wo)/(940. * Frequency);
//     return eta;
// };
//
// double SNiP2_01_07_Formula6_Wn( double Wo, double C, double K, bool OUT=false)
// {
//     double Wm = K * C * Wo ;
//     return Wm;
// };
//
// double SNiP2_01_07_Formula9_Wp( double Wm, double Dzeta, double Ksi, double Eps, bool OUT=false)
// {
//     double Wp = Wm * Dzeta * Ksi * Eps;
//     return Wp;
// };
//
// double SNiP2_01_07_Formula7_Vmax(double Wo, bool OUT=false)
// {
//     double Vmax = sqrt(Wo/0.61);
//     return Vmax;
// }
//
// double SNiP2_01_07_actual_Formula11_13_Vmax(double Wo, double K, bool OUT=false)
// {
//     double Vmax = 1.3*sqrt(Wo*K);
//     return Vmax;
// }