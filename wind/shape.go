package wind

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"text/tabwriter"
)

// Sphere сфера by B.1.11
func Sphere(zone Zone, wr Region, zg, d, Δ float64) (cx, cz, Re, ν float64) {
	ze := zg + d/2.0
	Wo := float64(wr)
	Kz := FactorKz(zone, ze)
	Re = 0.88 * d * math.Sqrt(Wo*Kz*γf) * 1e5
	cx = GraphB14(d, Δ, Re)
	ν = FactorNu(0.7*d, 0.7*d)

	// коэффициент подъемной силы сферы
	if d/2 < zg {
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
func Frame(out io.Writer, zone Zone, wr Region, ld LogDecriment, h float64, hzs []float64) (
	WsZ func(z float64) float64,
) {

	var buf bytes.Buffer
	defer func() {
		if out != nil {
			fmt.Fprintf(out, "%s", buf.String())
		}
	}()
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
          *     |
          *     |
          *     |
          *     |
   ---- ground ----

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
	zs := SplitHeigth(zo, h)

	separator := func() {
		fmt.Fprintf(w, "\t|")
	}
	fmt.Fprintf(w, "\t|\tz\tze\tKz\tζ\tξ\t|\tWm\tWp\tWsum\t|\n")
	fmt.Fprintf(w, "\t|\tm\tm\t\t\t\t|\tPa\tPa\tPa\t|\n")
	fmt.Fprintf(w, "\t|\t\t\t\t\t\t|\t\t\t\t|\n")
	WsZ = func(z float64) float64 {
		if z < zo || h < z {
			panic(fmt.Errorf("not acceptable: z(%.2e) < zo(%.2e) || h(%.2e) < z(%.2e) ", z, zo, h, z))
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
	return
}

// Cylinder return Wsum dependency of height.
// Acceptable for vertical and horizontal duct.
func Cylinder(out io.Writer, zone Zone, wr Region, ld LogDecriment, Δ, d, h float64, zo float64, hzs []float64) (
	WsZ func(z float64) float64,
) {
	var buf bytes.Buffer
	defer func() {
		if out != nil {
			fmt.Fprintf(out, "%s", buf.String())
		}
	}()
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
		_, _, Re, _ := Sphere(zone, wr, ze, d, Δ)
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
	zs := SplitHeigth(zo, h)

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

	fmt.Fprintf(w, "\t|\tz\tze\tKz\tζ\tξ\t|\tWm\tWp\tWsum\t|\n")
	fmt.Fprintf(w, "\t|\tm\tm\t\t\t\t|\tPa\tPa\tPa\t|\n")
	fmt.Fprintf(w, "\t|\t\t\t\t\t\t|\t\t\t\t|\n")
	section := func(w io.Writer, z float64) float64 {
		if w == nil {
			var buf bytes.Buffer
			w = &buf
		}
		if z < zo || h < z {
			panic(fmt.Errorf("not acceptable: z(%.2e) < zo(%.2e) || h(%.2e) < z(%.2e) ", z, zo, h, z))
		}
		fmt.Fprintf(w, "\t|") // separator
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
		fmt.Fprintf(w, "\t|") // separator
		// Wm
		Wm := Wo * Kz * cx
		fmt.Fprintf(w, "\t%6.1f", Wm)
		// Wp
		Wp := Wm * ξ * ζ * ν
		fmt.Fprintf(w, "\t%6.1f", Wp)
		// Wsum
		Wsum := Wm + Wp
		fmt.Fprintf(w, "\t%6.1f", Wsum)
		fmt.Fprintf(w, "\t|") // separator

		fmt.Fprintf(w, "\n")
		return Wsum
	}
	WsZ = func(z float64) float64 {
		return section(nil, z)
	}
	for _, z := range zs {
		_ = section(w, z)
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
	fmt.Fprintf(w, "|\tside\t|\twidth\t|\tCenter of Ws\t|\tWs average\t|\n")
	Wa, Z := averageW(zo, h, WsZ)
	fmt.Fprintf(w, "|\t%6s\t|\t%6.3f\t|\t%6.3f\t|\t%+8.1f\t|\n",
		"front", d, Z, Wa)
	w.Flush()
	return
}

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

func SplitHeigth(zo, h float64) (zs []float64) {
	if zo < 0.000 {
		panic("zo is negative")
	}
	if h < zo {
		panic("not valid h < zo")
	}
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
func Rectangle(out io.Writer, zone Zone, wr Region, ld LogDecriment, b, d, h float64, zo float64, hzs []float64) (
	WsZ [SideSize]func(z float64) float64,
) {
	// generate height sections
	zs := SplitHeigth(zo, h)

	var buf bytes.Buffer
	defer func() {
		if out != nil {
			fmt.Fprintf(out, "%s", buf.String())
		}
	}()
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
				panic(fmt.Errorf("not acceptable: z(%.2e) < zo(%.2e) || h(%.2e) < z(%.2e) ", z, zo, h, z))
			}
			return section(nil, z, side)
		}
	}

	// average value
	type average struct{ center, value float64 }
	av := make([]average, len(ListRectangleSides()))

	fmt.Fprintf(w, "\t|\tside\tz\tze\tKz\tζ\tξ\t|\tcx\tρ\tχ\tν\tWm\tWp\tWsum\t|\n")
	fmt.Fprintf(w, "\t|\t\t\t\t\t\t\t|\t\t\t\t\tPa\tPa\tPa\t|\n")
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
	fmt.Fprintf(w, "|\tside\t|\twidth\t|\tCenter of Ws\t|\tWs average\t|\n")
	for _, side := range ListRectangleSides() {
		Wa, Z := averageW(zo, h, WsZ[side])
		wd := width(side)
		fmt.Fprintf(w, "|\t%6s\t|\t%6.3f\t|\t%6.3f\t|\t%+8.1f\t|\n",
			side.Name(), wd, Z, Wa)
	}

	w.Flush()
	return
}

func averageW(z0, z1 float64, w func(z float64) float64) (Wa, Z float64) {
	// center mass of shape:
	// Z := sum(mi*zi)/sum(mi)
	zh := SplitHeigth(z0, z1)
	var up, down float64
	for i := range zh {
		if i == 0 {
			continue
		}
		z0, z1 := zh[i-1], zh[i]
		w0, w1 := w(z0), w(z1)
		dz := z1 - z0
		mi := (w0 + w1) / 2.0 * dz                                  // area of trapezoid
		zi := (w0*(z0+1.0/3.0*dz) + w1*(z0+2.0/3.0*dz)) / (w0 + w1) // center of trapezoid
		up += mi * zi
		down += mi
	}
	return down / (z1 - z0), up / down
}
