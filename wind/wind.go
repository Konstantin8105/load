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
//	z - height from ground
//	d - dimension of building (perpendicular)
//	h - heigth of building
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

//
// double SNiP2_01_07_p6_7b_Eta(double Wo, double Frequency, )
// {
//     double eta = sqrt(1.4 * Wo)/(940. * Frequency);
//     return eta;
// };
//
// double SNiP2_01_07_Schema12a_Re(double Diameter,double Wo, double Kz, )
// {
//     double Re = 0.88*Diameter*sqrt(Wo*Kz*1.4)*1e5;
//     return Re;
// }
//
// double SNiP2_01_07_Schema12b_K1(double h1, double d)
// {
//     double Otn[14]={
//         0.2,    0.8,
//         0.5,    0.9,
//         1.0,    0.95,
//         2.0,    1.0,
//         5.0,    1.1,
//         10.,    1.15,
//         25.,    1.2
//     };
//     if(h1/d < Otn[0*2+0])return Otn[0*2+1];
//     if(h1/d > Otn[6*2+0])return Otn[6*2+1];
//     type_LLU i;
//     for(i=0;i<7;i++)
//     {
//         if(h1/d == Otn[i*2+0]) return Otn[i*2 + 1];
//         if(h1/d <  Otn[i*2+0]) break;
//     }
//     double k1 = LinearInter(Otn[i*2+1],Otn[i*2+0],Otn[(i-1)*2+1],Otn[(i-1)*2+0],h1/d);
//     return k1;
// }
//
// double SNiP2_01_07_Table9_Epsilon(double ro, double hi, )
// {
//     double ArrEpsilon[8*8]={
//         0.0,    5.0,    10.,    20.,    40.,    80.,    160,    350,
//         0.1,    0.95,   0.92,   0.88,   0.83,   0.76,   0.67,   0.56,
//         5.0,    0.89,   0.87,   0.84,   0.80,   0.73,   0.65,   0.54,
//         10.,    0.85,   0.84,   0.81,   0.77,   0.71,   0.64,   0.53,
//         20.,    0.80,   0.78,   0.76,   0.73,   0.68,   0.61,   0.51,
//         40.,    0.72,   0.72,   0.70,   0.67,   0.63,   0.57,   0.48,
//         80.,    0.63,   0.63,   0.61,   0.59,   0.56,   0.51,   0.44,
//         100,    0.53,   0.53,   0.52,   0.50,   0.47,   0.44,   0.38};
//     if(ro < ArrEpsilon[1*8+0] || ro > ArrEpsilon[7*8+0])
//     {
//         print_name("Add function in SNiP2_01_07_Table9_Epsilon for ro. 1");
//         printf("ro = %f\thi = %f\n",ro,hi);
//         FATAL();
//     }
//     if(hi < ArrEpsilon[0*8+1] || hi > ArrEpsilon[0*8+7])
//     {
//         print_name("Add function in SNiP2_01_07_Table9_Epsilon for ro. 2");
//         printf("ro = %f\thi = %f\n",ro,hi);
//         FATAL();
//     }
//     type_LLU i,j;
//     for(i=0;i<7;i++)
//         if(ro < ArrEpsilon [(i+1)*8+0]) break;
//     for(j=0;j<7;j++)
//         if(hi < ArrEpsilon [0*8+(j+1)]) break;
//     i++;j++;
//     //   p1   p5   p2    //
//     //   *    |    *     //
//     //                   //
//     //   *    |    *     //
//     //   p3   p6   p4    //
//     double p1 = ArrEpsilon[(i-1)*8+(j-1)];
//     double p2 = ArrEpsilon[(i-1)*8+(j-0)];
//     double p3 = ArrEpsilon[(i-0)*8+(j-1)];
//     double p4 = ArrEpsilon[(i-0)*8+(j-0)];
//     double p5 = LinearInter(p2,ArrEpsilon [0*8+j],p1,ArrEpsilon [0*8+(j-1)],hi);
//     double p6 = LinearInter(p4,ArrEpsilon [0*8+j],p3,ArrEpsilon [0*8+(j-1)],hi);
//     double eps = LinearInter(p6,ArrEpsilon [i*8+0],p5,ArrEpsilon [(i-1)*8+0],ro);
//     return eps;
// };
//
// double SNiP2_01_07_Schema12b_Ce1(double angle, double h1, double d, )
// {
//     if(angle == 0. ) return 1.0;
//     if(angle < 0.  ) return SNiP2_01_07_Schema12b_Ce1(-angle                     ,h1,d,OUT);
//     if(angle > 360.) return SNiP2_01_07_Schema12b_Ce1(angle - int(angle/360.)*360,h1,d,OUT);
//     if(angle > 180.) return SNiP2_01_07_Schema12b_Ce1(180.-(angle-180.)          ,h1,d,OUT);
//     double Otn[30] = {
//         00.,    1.0,
//         10.,    1.0,
//         20.,    0.8,
//         30.,    0.4,
//         40.,    0.0,
//         50.,    -0.6,
//         60.,    -1.2,
//         70.,    -1.3,
//         80.,    -1.2,
//         90.,    -1.0,
//         100.,   -0.8,
//         110.,   -0.4,
//         120.,   -0.4,
//         130.,   -0.4,
//         180.1,   -0.4};
// //    if(angle < Otn[00*2+0])return Otn[00*2+1];
// //    if(angle > Otn[14*2+0])return Otn[14*2+1];
//     type_LLU i;
//     for(i=0;i<15;i++)
//     {
//         //if(angle == Otn[i*2+0]) return Otn[i*2 + 1];
//         if(angle <  Otn[i*2+0]) break;
//     }
//     double Cbetta = LinearInter(Otn[i*2+1],Otn[i*2+0],Otn[(i-1)*2+1],Otn[(i-1)*2+0],angle);
//     double k1     = SNiP2_01_07_Schema12b_K1(h1, d);
//     if(Cbetta > 0) k1 = 1;
//     double Ce1 = Cbetta * k1;
//     return Ce1;
// }
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
