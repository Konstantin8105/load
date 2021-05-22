// wind package for calculate the wind load on buildings.
// Main code : SP20.13330.2016
// Primary language for output: english
// Language for comments: any

package wind

import (
	"fmt"
	"math"
	"sort"

	"github.com/Konstantin8105/pow"
)

// Region is wind region. Ветровые районы (принимаются по карте 2 приложения Е)
type Region float64

// Unit: Pa
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

// ListWo is list of all wind regions
func ListWo() []Region {
	return []Region{
		RegionIa,
		RegionI,
		RegionII,
		RegionIII,
		RegionIV,
		RegionV,
		RegionVI,
		RegionVII,
	}
}

// Name of wind region
func (wr Region) Name() string {
	for _, v := range []struct {
		r    Region
		name string
	}{
		{RegionIa, "Ia"},
		{RegionI, "I"},
		{RegionII, "II"},
		{RegionIII, "III"},
		{RegionIV, "IV"},
		{RegionV, "V"},
		{RegionVI, "VI"},
		{RegionVII, "VII"},
	} {
		if v.r == wr {
			return v.name
		}
	}
	return "Undefined"
}

// String implementation of Stringer interface
func (wr Region) String() string {
	return fmt.Sprintf("Wind region: %3s with value = %.1f Pa", wr.Name(), float64(wr))
}

// Zone - тип местности
type Zone byte

const (
	// ZoneA - открытые побережья морей, озер и водохранилищ, сельские
	// местности, в том числе с постройками высотой менее 10 м, пустыни,
	// степи, лесостепи, тундра;
	ZoneA Zone = 'A'

	// ZoneB - городские территории, лесные массивы и другие местности,
	// равномерно покрытые препятствиями высотой более 10 м;
	ZoneB Zone = 'B'

	// ZoneC - городские районы с плотной застройкой зданиями высотой более 25м
	ZoneC Zone = 'C'
)

// String implementation of Stringer interface
func (z Zone) String() string {
	return "Wind zone: " + string(z)
}

// constants by table 11.3 SP20.13330.2016
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

// LogDecriment - Значение логарифмического декремента колебаний
type LogDecriment float64

const (
	// LogDecriment15 для стальных сооружений, футерованных дымовых труб,
	// аппаратов колонного типа, в том числе на железобетонных постаментах
	LogDecriment15 LogDecriment = 0.15

	// LogDecriment22 для стекла, а также смешанных сооружений, имеющих
	// одновременно стальные и железобетонные несущие конструкции
	LogDecriment22 LogDecriment = 0.22

	// LogDecriment30 для железобетонных и каменных сооружений, а также
	// для зданий со стальным каркасом при наличии ограждающих конструкций
	LogDecriment30 LogDecriment = 0.30
)

// Name of log decriment
func (ld LogDecriment) Name() string {
	return fmt.Sprintf("δ = %.2f", float64(ld))
}

// String implementation of Stringer interface
func (ld LogDecriment) String() string {
	return fmt.Sprintf("Wind log decrement: %s", ld.Name())
}

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
	if h-d <= z {
		return h
	}
	if d <= z && z <= h-d {
		return z
	}
	if 0 <= z && z <= d {
		return d
	}
	panic(fmt.Errorf("not implemented %v %v %v", z, d, h))
}

// see 11.1.12
const γf = 1.40

// FactorKz by table 11.2 and formula 11.4
func FactorKz(zone Zone, ze float64) (kz float64) {
	if !(zone == ZoneA || zone == ZoneB || zone == ZoneC) {
		panic("undefined zone")
	}
	if 300.0 < ze {
		panic(fmt.Errorf("ze = %f is too big", ze))
	}
	if ze <= 5.0 {
		ze = 5.0
	}
	if ze < 10 {
		// см. примечение к формуле 11.4
		var k5, k10 float64
		switch zone {
		case ZoneA:
			k5, k10 = 0.75, 1.00
		case ZoneB:
			k5, k10 = 0.50, 0.65
		case ZoneC:
			k5, k10 = 0.40, 0.40
		}
		return k5 + (k10-k5)*(ze-5.0)/(10.0-5.0)
	}
	α, k10, _ := zone.constants()
	return k10 * math.Pow(ze/10.0, 2.0*α)
}

// FactorZeta by table 11.4 and formula 11.6
func FactorZeta(zone Zone, ze float64) (ζ float64) {
	if !(zone == ZoneA || zone == ZoneB || zone == ZoneC) {
		panic("undefined zone")
	}
	if 300.0 < ze {
		panic(fmt.Errorf("ze = %f is too big", ze))
	}
	if ze <= 5.0 {
		ze = 5.0
	}
	if ze < 10 {
		// см. примечение к формуле 11.4
		var c5, c10 float64
		switch zone {
		case ZoneA:
			c5, c10 = 0.85, 0.76
		case ZoneB:
			c5, c10 = 1.22, 1.06
		case ZoneC:
			c5, c10 = 1.78, 1.78
		}
		return c5 + (c10-c5)*(ze-5.0)/(10.0-5.0)
	}
	α, _, ζ10 := zone.constants()
	return ζ10 * math.Pow(ze/10.0, -α)
}

// Limit dimensionless period - by table 11.5
func DimlessPeriodLimit(ld LogDecriment) float64 {
	switch ld {
	case LogDecriment15:
		return 0.0077
	case LogDecriment22:
		return 0.0140
	case LogDecriment30:
		return 0.0230
	}
	panic("not implemented")
}

// NaturalFrequencyLimit by formula 11.9a
func NaturalFrequencyLimit(zone Zone, wr Region, ld LogDecriment, z float64) (Flim float64) {

	wo := float64(wr)
	Tglim := DimlessPeriodLimit(ld)

	// Для конструктивных элементов zэк — высота z, на которой они расположены;
	// для зданий и сооружений zэк = 0,8h, где h — высота сооружений;
	zek := z
	k := FactorKz(zone, zek)

	return math.Sqrt(wo*k*γf) / (940.0 * Tglim)
}

// FactorXiHz - коэффициент динамичности by pic 11.1 c учетом
// динамической реации по s собственным формам
//	isBuilding = true  - для зданий и сооружений zэк = 0.8*h
//	isBuilding = false - для конструктивных элементов zэк — высота z, на
//		которой они расположены
//	hzs - список частот собственных колебаний
func FactorXiHz(wr Region, zone Zone, ld LogDecriment, isBuilding bool, z float64, hzs []float64) (ξ float64) {
	defer func() {
		// round
		ξ *= 1000.0
		ξi := int64(ξ)
		ξ = float64(ξi) / 1000.0
	}()
	flim := NaturalFrequencyLimit(zone, wr, ld, z)
	// filter of natural frequency
	{
		s := make([]float64, len(hzs))
		copy(s, hzs)
		hzs = s
	again:
		for i := range hzs {
			if flim < hzs[i] {
				hzs = append(hzs[:i], hzs[i+1:]...)
				goto again
			}
		}
	}
	// sort
	sort.Float64s(hzs)
	// Kz
	if isBuilding {
		z = 0.8 * z
	}
	Kz := FactorKz(zone, z)
	Wo := float64(wr)
	ξ = 1.0 // by default if flim < f
	if 0 < len(hzs) {
		ξ = 0.0 // reset value
		for _, hz := range hzs {
			Tgi := math.Sqrt(Wo*Kz*γf) / (940.0 * hz) // see formula 11.8a
			ξi := factorXi(ld, Tgi)
			ξ += pow.E2(ξi)
		}
		ξ = math.Sqrt(ξ)
	}
	if ξ < 0.8 {
		panic(fmt.Errorf("%v %v %v", hzs, flim, ξ))
	}
	return
}

// factorXi - коэффициент динамичности by pic 11.1
func factorXi(ld LogDecriment, Tg float64) (ξ float64) {
	defer func() {
		// предполагается, что коэффициент динамичности не может быть менее 1.0
		if ξ < 1.0 {
			ξ = 1.0
		}
	}()
	var graph []float64
	switch ld {
	case LogDecriment15:
		graph = []float64{
			0.000000, 1.00489,
			0.002802, 1.19363,
			0.006955, 1.33462,
			0.010190, 1.40781,
			0.020562, 1.59802,
			0.050304, 1.96331,
			0.100067, 2.32763,
			0.150742, 2.58187,
			0.199989, 2.77751,
			0.250704, 2.92906,
			0.300563, 3.04398,
		}

	case LogDecriment22:
		graph = []float64{
			0.000000, 1.00489,
			0.002959, 1.14337,
			0.006493, 1.22779,
			0.010324, 1.29807,
			0.020567, 1.43646,
			0.050480, 1.69741,
			0.100372, 1.98558,
			0.150347, 2.18566,
			0.200044, 2.33796,
			0.250480, 2.44505,
			0.299672, 2.53425,
		}

	case LogDecriment30:
		graph = []float64{
			0.000000, 1.00489,
			0.003089, 1.08927,
			0.006737, 1.16176,
			0.010499, 1.22196,
			0.020573, 1.32552,
			0.050631, 1.54261,
			0.100455, 1.77071,
			0.150220, 1.91787,
			0.200159, 2.04105,
			0.250603, 2.12852,
			0.299873, 2.20233,
		}

	default:
		panic("not implemented")
	}
	for i := 2; i < len(graph); i += 2 {
		x0, y0 := graph[i-2], graph[i-1]
		x2, y2 := graph[i], graph[i+1]
		if x0 <= Tg && Tg <= x2 {
			return interpol(x0, Tg, x2, y0, y2)
		}
	}
	panic("not implemented")
}

// Plate плоскость
type Plate string

// Плоскости
const (
	ZOY Plate = "ZOY"
	ZOX Plate = "ZOX"
	XOY Plate = "XOY"
)

// NuPlates calculate dimention values.
// see table 11.7
func NuPlates(b, h, a float64, pl Plate) (ρ, χ float64) {
	switch pl {
	case ZOY:
		ρ, χ = b, h
	case ZOX:
		ρ, χ = 0.4*a, h
	case XOY:
		ρ, χ = b, a
	default:
		panic("not implemented")
	}
	return
}

// FactorNu Коэффициент пространственной корреляции пульсаций давления
// by table 11.6
func FactorNu(ρ, χ float64) (ν float64) {
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
	if header[col-1] < χ {
		χ = header[col-1]
	}
	if ρ < ro[0] {
		ρ = ro[0]
	}
	if ro[row-1] < ρ {
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

// GraphB14 Реализованый алгоритм упрощенный, но в худшую сторону.
// Аэродинамические коэффициенты лобового сопротивления сх сферы
func GraphB14(d, Δ, Re float64) (cx float64) {
	defer func() {
		// round
		cx *= 1000.0
		cxi := int64(cx)
		cx = float64(cxi) / 1000.0
	}()
	if Re < 6.0e5 {
		return 0.6
	}
	dd := Δ / d
	if 1e-3 < dd {
		return 0.4
	}
	if dd < 1e-5 {
		return 0.2
	}
	ddp := math.Log10(dd)
	return 0.2 + (0.4-0.2)*(ddp-(-5))/(-3-(-5))
}

// GraphB17 Значение коэффициента Cx
func GraphB17(d, Δ, Re float64) (Cx float64) {
	defer func() {
		if 1.2 < Cx {
			Cx = 1.2
		}
		if Cx < 0.4 {
			Cx = 0.4
		}
	}()

	if Re < 1e5 {
		return 1.2
	}

	x := math.Log10(Re)
	Δd := Δ / d
	if Δd < 1e-5 {
		Δd = 1e-5
	}
	if 1e-2 < Δd {
		return 1.2
	}

	Cx = 0.4
	if Re < 3.0e5 {
		Cx = math.Max(Cx, 3.01192*pow.E2(x)-33.6354*x+94.2379)
	}

	Cx5 := -0.0730890*pow.E2(x) + 1.178010*x - 3.967380
	Cx2 := -0.0362124*pow.E2(x) + 0.530084*x - 0.844029
	power := math.Log10(Δd) // for example: -5...-2

	return math.Max(Cx, Cx5+(Cx2-Cx5)*(power-(-5))/(-2-(-5)))
}

// GraphB23 - В.1.15 Учет относительного удлинения
func GraphB23(λe, ϕ float64) (Kλ float64) {
	if λe < 1 {
		λe = 1.0
	}
	if 200 < λe {
		λe = 200
	}
	type line struct {
		ϕ                       float64
		Kλ1, Kλ10, Kλ100, Kλ200 float64
	}
	K := []line{
		{ϕ: 0.00, Kλ1: 1.00, Kλ10: 1.00, Kλ100: 1.00, Kλ200: 1.00},
		{ϕ: 0.10, Kλ1: 0.98, Kλ10: 0.99, Kλ100: 0.99, Kλ200: 1.00},
		{ϕ: 0.50, Kλ1: 0.88, Kλ10: 0.91, Kλ100: 0.98, Kλ200: 1.00},
		{ϕ: 0.90, Kλ1: 0.82, Kλ10: 0.87, Kλ100: 0.97, Kλ200: 1.00},
		{ϕ: 0.95, Kλ1: 0.73, Kλ10: 0.80, Kλ100: 0.96, Kλ200: 1.00},
		{ϕ: 1.00, Kλ1: 0.60, Kλ10: 0.70, Kλ100: 0.95, Kλ200: 1.00},
	}
	var l line // actual line
	for index := 1; index < len(K); index++ {
		if K[index-1].ϕ <= ϕ && ϕ <= K[index].ϕ {
			l.Kλ1 = K[index-1].Kλ1 + (K[index].Kλ1-K[index-1].Kλ1)*
				(ϕ-K[index-1].ϕ)/(K[index].ϕ-K[index-1].ϕ)
			l.Kλ10 = K[index-1].Kλ10 + (K[index].Kλ10-K[index-1].Kλ10)*
				(ϕ-K[index-1].ϕ)/(K[index].ϕ-K[index-1].ϕ)
			l.Kλ100 = K[index-1].Kλ100 + (K[index].Kλ100-K[index-1].Kλ100)*
				(ϕ-K[index-1].ϕ)/(K[index].ϕ-K[index-1].ϕ)
			l.Kλ200 = K[index-1].Kλ200 + (K[index].Kλ200-K[index-1].Kλ200)*
				(ϕ-K[index-1].ϕ)/(K[index].ϕ-K[index-1].ϕ)
		}
	}
	lλe := math.Log10(λe)
	if 1 <= λe && λe <= 10 {
		return l.Kλ1 + (l.Kλ10-l.Kλ1)*(lλe-math.Log10(1.0))/
			(math.Log10(10.0)-math.Log10(1.0))
	} else if 10 <= λe && λe <= 100 {
		return l.Kλ10 + (l.Kλ100-l.Kλ10)*(lλe-math.Log10(10.0))/
			(math.Log10(100.0)-math.Log10(10.0))
	} else if 100 <= λe && λe <= 200 {
		return l.Kλ100 + (l.Kλ200-l.Kλ100)*(lλe-math.Log10(100.0))/
			(math.Log10(200.0)-math.Log10(100.0))
	}

	return 1.0
}

const (
	TableB10Col1 float64 = 0.5
	TableB10Col2 float64 = 1.0
	TableB10Col3 float64 = 2.0
	TableB10Col4 float64 = 1e100 // Infinity : math.Inf(1)
)

type Struhale int

const (
	StCylinder Struhale = iota
	StRectangle
)

func (st Struhale) Value() float64 {
	switch st {
	case StCylinder:
		return 0.2
	case StRectangle:
		return 0.11
	}
	panic("not implemented")
}

func (st Struhale) Name() string {
	return fmt.Sprintf("St = %4.2f", st.Value())
}

func (st Struhale) String() string {
	return "Struhale number: " + st.Name()
}

func interpol(x0, x1, x2, y0, y2 float64) (y1 float64) {
	return y0 + (y2-y0)*(x1-x0)/(x2-x0)
}
