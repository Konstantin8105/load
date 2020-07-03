package wind

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"text/tabwriter"
)

// Sphere сфера by B.1.11
func Sphere(zone Zone, wr Region, zg, d, Δ float64) (cx, cz, Re, ν float64, err error) {
	// TODO : add error handling
	ze := zg + d/2.0
	Wo := float64(wr)
	Kz := FactorKz(zone, ze)
	Re = 0.88 * d * math.Sqrt(Wo*Kz*γf) * 1e5
	cx = GraphB14(d, Δ, Re)
	ν = FactorNu(0.7*d, 0.7*d)

	// коэффициент подъемной силы сферы
	if zg > d/2 {
		cz = 0.0
	} else {
		cz = 0.6
	}

	if zg < d/2.0 {
		cx *= 1.6
	}
	return
}

// Frame return Wsum dependency of height
func Frame(zone Zone, wr Region, ld LogDecriment, h float64, hzs []float64) (
	WsZ func(z float64) float64,
) {

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, `Sketch:

          *-------
          *     |
  Wind    *     |
 ----->   *     |
          *     |
          *     |
          *     |
          *     h
          *---  |
            |   |
            zo  |
            |   |
   ---------- ground -------------

`)

	fmt.Fprintf(w, "%s\n", zone.String())
	fmt.Fprintf(w, "%s\n", wr.String())
	fmt.Fprintf(w, "%s\n", ld.String())
	fmt.Fprintf(w, "Natural frequency : %v\n", hzs)
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Dimensions:\n")
	fmt.Fprintf(w, "\th\t%6.3f m\n", h)
	fmt.Fprintf(w, "\n")

	// Аэродинамические коэффициенты лобового сопротивления
	// see par B.1.13
	cx := 1.4
	fmt.Fprintf(w, "Cx  = %6.3f\n", cx)
	fmt.Fprintf(w, "\n")

	// Коэффициент пространственной корреляции пульсаций давлавления
	// ν
	fmt.Fprintf(w, "The spatial correlation coefficient of pressure pulsations:\n")
	ν := func() float64 {
		pl := ZOY
		b := 0.0
		d := 0.0
		ρ, χ := NuPlates(b, h, d, pl)
		fmt.Fprintf(w, "\tρ\t%6.3f\n", ρ)
		fmt.Fprintf(w, "\tχ\t%6.3f\n", χ)
		return FactorNu(ρ, χ)
	}()
	fmt.Fprintf(w, "\tν\t%6.3f\n", ν)
	fmt.Fprintf(w, "\n")

	// generate height sections
	zo := 0.0
	zs := splitHeigth(zo, h)

	separator := func() {
		fmt.Fprintf(w, "\t|")
	}
	fmt.Fprintf(w, "\t|\tz\tze\tKz\tζ\tξ\t|\tWm\tWp\tWsum\t|\n")
	fmt.Fprintf(w, "\t|\tm\tm\t\t\t\t|\tPa\tPa\tPa\t|\n")
	fmt.Fprintf(w, "\t|\t\t\t\t\t\t|\t\t\t\t|\n")
	WsZ = func(z float64) float64 {
		if z < zo || h < z {
			panic("not acceptable")
		}
		separator()
		// Wo
		Wo := float64(wr)
		fmt.Fprintf(w, "\t%6.3f", z)
		// Ze
		ze := func() float64 {
			d := 0.0
			return EffectiveHeigth(z, d, h, true)
		}()
		fmt.Fprintf(w, "\t%6.3f", ze)
		// Kz
		Kz := FactorKz(zone, ze)
		fmt.Fprintf(w, "\t%6.3f", Kz)
		// Zeta
		ζ := FactorZeta(zone, ze)
		fmt.Fprintf(w, "\t%6.3f", ζ)
		// Xi
		ξ := FactorXiHz(wr, zone, ld, true, h, hzs)
		fmt.Fprintf(w, "\t%6.3f", ξ)
		separator()
		// Wm
		Wm := Wo * Kz * cx
		fmt.Fprintf(w, "\t%6.1f", Wm)
		// Wp
		Wp := Wm * ξ * ζ * ν
		fmt.Fprintf(w, "\t%6.1f", Wp)
		// Wsum
		Wsum := Wm + Wp
		fmt.Fprintf(w, "\t%6.1f", Wsum)
		separator()

		fmt.Fprintf(w, "\n")
		return Wsum
	}
	for _, z := range zs {
		_ = WsZ(z)
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	return
}

// TODO: add cylinder
// TODO: add frame
// TODO: add horizontal duct

// Cylinder return Wsum dependency of height
func Cylinder(zone Zone, wr Region, ld LogDecriment, Δ, d, h float64, zo float64, hzs []float64) (
	WsZ func(z float64) float64,
) {

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, `Sketch:

          |<--- d ----->|
          |             |
          ***************-------
          *             *     |
  Wind    *             *     |
 ----->   *             *     |
          *             *     |
          *             *     |
          *             *     |
          *             *     h
          ***************---  |
                          |   |
                          zo  |
                          |   |
   ---------- ground -------------

`)

	fmt.Fprintf(w, "%s\n", zone.String())
	fmt.Fprintf(w, "%s\n", wr.String())
	fmt.Fprintf(w, "%s\n", ld.String())
	fmt.Fprintf(w, "Natural frequency : %v\n", hzs)
	fmt.Fprintf(w, "\n")

	b := d

	fmt.Fprintf(w, "Dimensions:\n")
	fmt.Fprintf(w, "\tb\t%6.3f m\n", b)
	fmt.Fprintf(w, "\td\t%6.3f m\n", d)
	fmt.Fprintf(w, "\tzo\t%6.3f m\n", zo)
	fmt.Fprintf(w, "\th\t%6.3f m\n", h)
	fmt.Fprintf(w, "\n")

	Re := func() float64 {
		// Число Рейнольдса Re определяется по формуле, приведенной в В.1.11
		//	где ze = 0,8h для вертикально расположенных сооружений;
		//	ze равно расстоянию от поверхности земли до оси горизонтально
		//	расположенного сооружения.
		ze := zo + 0.8*(h-zo)
		_, _, Re, _, err := Sphere(zone, wr, ze, d, Δ)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Re = %6.3f*10^5 for ze=0.8*h = %6.3f\n", Re*1e-5, ze)
		return Re
	}()
	fmt.Fprintf(w, "\n")

	CxInfinite := GraphB17(d, Δ, Re)
	fmt.Fprintf(w, "Cx∞ = %6.3f\n", CxInfinite)
	fmt.Fprintf(w, "\n")

	// Учет относительного удлинения
	Kλ := func() float64 {
		λ := (h - zo) / b
		λe := TableB10Col3 * λ
		ϕ := 1.0
		Kλ := GraphB23(λe, ϕ)
		fmt.Fprintf(w, "Elongation:\n")
		fmt.Fprintf(w, "\tλ \t%6.3f\n", λ)
		fmt.Fprintf(w, "\tλe\t%6.3f\n", λe)
		fmt.Fprintf(w, "\tϕ \t%6.3f\n", ϕ)
		fmt.Fprintf(w, "\tKλ\t%6.3f\n", Kλ)
		fmt.Fprintf(w, "\n")
		return Kλ
	}()

	// Аэродинамические коэффициенты лобового сопротивления
	cx := Kλ * CxInfinite
	fmt.Fprintf(w, "Cx  = %6.3f\n", cx)
	fmt.Fprintf(w, "\n")

	// generate height sections
	zs := splitHeigth(zo, h)

	// Коэффициент пространственной корреляции пульсаций давлавления
	// ν
	pl := ZOY
	ρ, χ := NuPlates(b, h, d, pl)
	ν := FactorNu(ρ, χ)
	fmt.Fprintf(w, "The spatial correlation coefficient of pressure pulsations:\n")
	fmt.Fprintf(w, "\tρ\t%6.3f\n", ρ)
	fmt.Fprintf(w, "\tχ\t%6.3f\n", χ)
	fmt.Fprintf(w, "\tν\t%6.3f\n", ν)
	fmt.Fprintf(w, "\n")

	separator := func() {
		fmt.Fprintf(w, "\t|")
	}
	fmt.Fprintf(w, "\t|\tz\tze\tKz\tζ\tξ\t|\tWm\tWp\tWsum\t|\n")
	fmt.Fprintf(w, "\t|\tm\tm\t\t\t\t|\tPa\tPa\tPa\t|\n")
	fmt.Fprintf(w, "\t|\t\t\t\t\t\t|\t\t\t\t|\n")
	WsZ = func(z float64) float64 {
		if z < zo || h < z {
			panic("not acceptable")
		}
		separator()
		// Wo
		Wo := float64(wr)
		fmt.Fprintf(w, "\t%6.3f", z)
		// Ze
		ze := EffectiveHeigth(z, d, h, true)
		fmt.Fprintf(w, "\t%6.3f", ze)
		// Kz
		Kz := FactorKz(zone, ze)
		fmt.Fprintf(w, "\t%6.3f", Kz)
		// Zeta
		ζ := FactorZeta(zone, ze)
		fmt.Fprintf(w, "\t%6.3f", ζ)
		// Xi
		ξ := FactorXiHz(wr, zone, ld, true, h, hzs)
		fmt.Fprintf(w, "\t%6.3f", ξ)
		separator()
		// Wm
		Wm := Wo * Kz * cx
		fmt.Fprintf(w, "\t%6.1f", Wm)
		// Wp
		Wp := Wm * ξ * ζ * ν
		fmt.Fprintf(w, "\t%6.1f", Wp)
		// Wsum
		Wsum := Wm + Wp
		fmt.Fprintf(w, "\t%6.1f", Wsum)
		separator()

		fmt.Fprintf(w, "\n")
		return Wsum
	}
	for _, z := range zs {
		_ = WsZ(z)
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	return
}

// TODO : add stack

type RectangleSide int

const (
	SideA RectangleSide = iota
	SideB
	SideC
	SideD
	SideE
	SideSize // size of sides for rectangle building
)

func ListRectangleSides() [SideSize]RectangleSide {
	return [SideSize]RectangleSide{
		SideA,
		SideB,
		SideC,
		SideD,
		SideE,
	}
}

func (rs RectangleSide) Convert() Plate {
	switch rs {
	case SideA, SideB, SideC:
		return ZOX
	case SideD, SideE:
		return ZOY
	}
	panic("not implemented")
}

func (rs RectangleSide) Value() float64 {
	switch rs {
	case SideA:
		return -1.0
	case SideB:
		return -0.8
	case SideC:
		return -0.5
	case SideD:
		return +0.8
	case SideE:
		return -0.5
	}
	panic("not implemented")
}

func (rs RectangleSide) Name() string {
	switch rs {
	case SideA:
		return "A"
	case SideB:
		return "B"
	case SideC:
		return "C"
	case SideD:
		return "D"
	case SideE:
		return "E"
	}
	panic("not implemented")
}

func (rs RectangleSide) String() string {
	name := rs.Name()
	return fmt.Sprintf("Side of rectangle: %s", name)
}

func splitHeigth(zo, h float64) (zs []float64) {
	zs = append(zs, zo)
	zmin := float64(int(zo/5)+1) * 5.0
	for z := zmin; z < h; z += 5.0 {
		zs = append(zs, z)
	}
	if zs[len(zs)-1] != h {
		zs = append(zs, h)
	}
	return
}

// Rectangle return Wsum dependency of height, see part B.1.2
func Rectangle(zone Zone, wr Region, ld LogDecriment, b, d, h float64, zo float64, hzs []float64) (
	WsZ [SideSize]func(z float64) float64,
) {
	// generate height sections
	zs := splitHeigth(zo, h)

	// TODO: for 2 directions
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, `Sketch:

          |<------- d --------->|
          |                     |
          ***********************------
          *                     *    |
  Wind    *                     *    |
 ----->   *                     *    |
        D *                     * E  b
          *                     *    |
          *                     *    |
          *                     *    |
          ***********************------
          |  A  |    B    |  C  |

`)

	fmt.Fprintf(w, "%s\n", zone.String())
	fmt.Fprintf(w, "%s\n", wr.String())
	fmt.Fprintf(w, "%s\n", ld.String())
	fmt.Fprintf(w, "Natural frequency : %v\n", hzs)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Dimensions:\n")
	fmt.Fprintf(w, "\tb\t%6.3f m\n", b)
	fmt.Fprintf(w, "\td\t%6.3f m\n", d)
	fmt.Fprintf(w, "\tzo\t%6.3f m\n", zo)
	fmt.Fprintf(w, "\th\t%6.3f m\n", h)
	fmt.Fprintf(w, "\n")

	section := func(w io.Writer, z float64, side RectangleSide) (wsum float64) {
		if w == nil {
			var buf bytes.Buffer
			w = &buf
		}
		separator := func() {
			fmt.Fprintf(w, "\t|")
		}

		// separator
		separator()

		// side name
		fmt.Fprintf(w, "\t%s", side.Name())
		// Wo
		Wo := float64(wr)
		fmt.Fprintf(w, "\t%6.3f", z)
		// Ze
		var ze float64
		switch side.Convert() {
		case ZOX:
			ze = EffectiveHeigth(z, d, h, false)
		case ZOY:
			ze = EffectiveHeigth(z, b, h, false)
		}
		fmt.Fprintf(w, "\t%6.3f", ze)
		// Kz
		Kz := FactorKz(zone, ze)
		fmt.Fprintf(w, "\t%6.3f", Kz)
		// Zeta
		ζ := FactorZeta(zone, ze)
		fmt.Fprintf(w, "\t%6.3f", ζ)
		// Xi
		ξ := FactorXiHz(wr, zone, ld, true, h, hzs)
		fmt.Fprintf(w, "\t%6.3f", ξ)
		// separator
		separator()

		// Cx
		cx := side.Value()
		fmt.Fprintf(w, "\t%4.1f", cx)
		// ν
		pl := side.Convert()
		ρ, χ := NuPlates(b, h, d, pl)
		ν := FactorNu(ρ, χ)
		fmt.Fprintf(w, "\t%6.3f", ρ)
		fmt.Fprintf(w, "\t%6.3f", χ)
		fmt.Fprintf(w, "\t%6.3f", ν)
		// Wm
		Wm := Wo * Kz * cx
		fmt.Fprintf(w, "\t%6.1f", Wm)
		// Wp
		Wp := Wm * ξ * ζ * ν
		fmt.Fprintf(w, "\t%6.1f", Wp)
		// Wsum
		Wsum := Wm + Wp
		fmt.Fprintf(w, "\t%6.1f", Wsum)
		// separator
		separator()

		fmt.Fprintf(w, "\n")
		return Wsum
	}

	for _, side := range ListRectangleSides() {
		side := side
		WsZ[side] = func(z float64) float64 {
			if z < zo || h < z {
				panic("not acceptable")
			}
			return section(nil, z, side)
		}
	}

	// average value
	type average struct{ center, value float64 }
	av := make([]average, len(ListRectangleSides()))

	// TODO : add unit
	fmt.Fprintf(w, "\t|\tside\tz\tze\tKz\tζ\tξ\t|\tcx\tρ\tχ\tν\tWm\tWp\tWsum\t|\n")
	for _, side := range ListRectangleSides() {
		fmt.Fprintf(w, "\t|\t \t \t \t \t \t \t|\t \t \t \t \t \t \t \t|\n")
		wmLast := 0.0
		for index, z := range zs {
			wm := section(w, z, side)
			if index == 0 {
				wmLast = wm
				continue
			}
			// local wm
			zLast := zs[index-1]
			area := (wmLast + wm) / 2.0 * (z - zLast)
			center := zLast + (z-zLast)/3.0*(wmLast+2.0*wm)/(wmLast+wm)
			// add to global wm
			av[side].center = (center*area + av[side].center*av[side].value) /
				(area + av[side].value)
			av[side].value += area
			wmLast = wm
		}
	}

	// width
	width := func(side RectangleSide) float64 {
		var width float64
		e := math.Min(b, 2*h)
		switch side {
		case SideA:
			width = math.Min(e/5.0, d)
		case SideB:
			width = math.Max(math.Min(e, d)-e/5.0, 0.0)
		case SideC:
			width = math.Max(0.0, d-e)
		case SideD:
			width = b
		case SideE:
			width = b
		default:
			panic("not implemented")
		}
		return width
	}

	fmt.Fprintf(w, `
   Ws on top    |----------->             |--------->
                |          /              |         |
                |--------->    Ws average |--------->
                |        /                |         |
   Ws on zero   |------->                 |--------->
              --------------- ground ------------------
`)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "|\tside\t|\twidth\t|\tWs on zero\t|\tWs on top\t|\tCenter of Ws\t|\tWs average\t|\n")
	fmt.Fprintf(w, "|\t\t|\t\t|\televation\t|\televation\t|\t\t|\t\t|\n")
	fmt.Fprintf(w, "|\t\t|\tmeter\t|\tPa\t|\tPa\t|\tmeter\t|\tPa\t|\n")
	fmt.Fprintf(w, "|\t\t|\t\t|\t\t|\t\t|\t\t|\t\t|\n")
	for _, side := range ListRectangleSides() {
		// calculate trapezoid wind load
		//	1. area   = (w0+w1)/2 * (h - zo)
		//	2. center = (h - zo)/3*(w0+2*w1)/(w0+w1) + zo
		// solving system:
		//	1. w0+w1   = area/(h - zo)*2
		//	2. w0+2*w1 = (center-zo)*3/(h-zo)*(w0+w1)
		// combile 1 and 2:
		//	1. w0+w1   = area/(h-zo)*2
		//  3. w0+2*w1 = (center-zo)*3/(h-zo)*area/(h-zo)*2
		// rename rigth parts:
		//	1. w0+w1   = r1
		//  3. w0+2*w1 = r2
		// solving:
		//	w0 = r1 - w1
		//	r1 - w1 + 2*w1 = r2
		//	w1 = r2 - r1
		//	w0 = r1 - w1 = r1 - r2 + r1 = 2*r1 - r2
		r1 := av[side].value / (h - zo) * 2.0
		r2 := (av[side].center - zo) * 3.0 / (h - zo) * r1
		w0 := 2*r1 - r2
		w1 := r2 - r1

		// calculate uniform wind load
		// with same moment on ground
		M := av[side].value * av[side].center
		waverage := M / ((h - zo) * ((h-zo)/2.0 + zo))

		{
			// check
			eps := 1e-6
			area := (w0 + w1) / 2 * (h - zo)
			if e := math.Abs((area - av[side].value) / area); e > eps {
				panic(e)
			}
			area = waverage * (h - zo)
			if e := math.Abs((area - av[side].value) / area); math.Abs(area) < math.Abs(av[side].value) {
				panic(fmt.Errorf("%v %v %v", e, area, av[side].value))
			}
		}

		// width
		wd := width(side)
		if wd == 0.0 {
			w0 = 0
			w1 = 0
			av[side].center = 0
			waverage = 0
		}
		fmt.Fprintf(w, "|\t%6s\t|\t%6.3f\t|\t%+8.1f\t|\t%+8.1f\t|\t%6.3f\t|\t%+8.1f\t|\n",
			side.Name(), wd, w0, w1, av[side].center, waverage)
	}

	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	return
}

// TODO : add cylinder

// TODO: add integration test

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
// double SNiP2_01_07_Table_19(double L, bool OUT = true)
// {
//     double fu = 0;
//          if(L <=  1.0) fu = 1./120.;
//     else if(L <=  3.0) fu = 1./150.;
//     else if(L <=  6.0) fu = 1./200.;
//     else if(L <= 12.0) fu = 1./250.;
//     else               fu = 1./300.;
//     if(OUT)
//     {
//         printf("Maximum deflection for L = %.1f m is %.2f mm(fu= %.5f)\n",L,L*fu*1e3,fu);
//     }
//     return fu;
// }
//
