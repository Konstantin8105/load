// wind package for calculate the wind load on buildings.
// Main code : SP20.13330.2016
// Primary language for output: english
// Language for comments: any

package wind

import (
	"fmt"
	"math"
	"sort"

	"github.com/Konstantin8105/graph"
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

func ListZone() []Zone {
	return []Zone{
		ZoneA,
		ZoneB,
		ZoneC,
	}
}

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

func ListLogDecriment() []LogDecriment {
	return []LogDecriment{
		LogDecriment15,
		LogDecriment22,
		LogDecriment30,
	}
}

// Name of log decriment
func (ld LogDecriment) Name() string {
	return fmt.Sprintf("δ = %.2f", float64(ld))
}

// String implementation of Stringer interface
func (ld LogDecriment) String() string {
	return fmt.Sprintf("Wind log decrement: %s", ld.Name())
}

// EffectiveHeigth by par 11.1.5
//
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
//
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
	var data []graph.Point
	switch ld {
	case LogDecriment15:
		data = []graph.Point{
			{X: 0.000000, Y: 1.00489},
			{X: 0.002802, Y: 1.19363},
			{X: 0.006955, Y: 1.33462},
			{X: 0.010190, Y: 1.40781},
			{X: 0.020562, Y: 1.59802},
			{X: 0.050304, Y: 1.96331},
			{X: 0.100067, Y: 2.32763},
			{X: 0.150742, Y: 2.58187},
			{X: 0.199989, Y: 2.77751},
			{X: 0.250704, Y: 2.92906},
			{X: 0.300000, Y: 3.04398},
		}

	case LogDecriment22:
		data = []graph.Point{
			{X: 0.000000, Y: 1.00489},
			{X: 0.002959, Y: 1.14337},
			{X: 0.006493, Y: 1.22779},
			{X: 0.010324, Y: 1.29807},
			{X: 0.020567, Y: 1.43646},
			{X: 0.050480, Y: 1.69741},
			{X: 0.100372, Y: 1.98558},
			{X: 0.150347, Y: 2.18566},
			{X: 0.200044, Y: 2.33796},
			{X: 0.250480, Y: 2.44505},
			{X: 0.300000, Y: 2.53425},
		}

	case LogDecriment30:
		data = []graph.Point{
			{X: 0.000000, Y: 1.00489},
			{X: 0.003089, Y: 1.08927},
			{X: 0.006737, Y: 1.16176},
			{X: 0.010499, Y: 1.22196},
			{X: 0.020573, Y: 1.32552},
			{X: 0.050631, Y: 1.54261},
			{X: 0.100455, Y: 1.77071},
			{X: 0.150220, Y: 1.91787},
			{X: 0.200159, Y: 2.04105},
			{X: 0.250603, Y: 2.12852},
			{X: 0.300000, Y: 2.20233},
		}

	default:
		panic("not implemented")
	}

	ξ, err := graph.Find(Tg, true, graph.NoCheckSorted, data...)
	if err != nil {
		panic(err)
	}
	return
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
	return interpolTable(χ, ρ, 8, 8, []float64{
		-0.01, 0.00, 5.00, 10.0, 20.0, 40.0, 80.0, 160., 350.,
		0.000, 0.95, 0.95, 0.92, 0.88, 0.83, 0.76, 0.67, 0.56,
		0.100, 0.95, 0.95, 0.92, 0.88, 0.83, 0.76, 0.67, 0.56,
		5.000, 0.89, 0.89, 0.87, 0.84, 0.80, 0.73, 0.65, 0.54,
		10.00, 0.85, 0.85, 0.84, 0.81, 0.77, 0.71, 0.64, 0.53,
		20.00, 0.80, 0.80, 0.78, 0.76, 0.73, 0.68, 0.61, 0.51,
		40.00, 0.72, 0.72, 0.72, 0.70, 0.67, 0.63, 0.57, 0.48,
		80.00, 0.63, 0.63, 0.63, 0.61, 0.59, 0.56, 0.51, 0.44,
		160.0, 0.53, 0.53, 0.53, 0.52, 0.50, 0.47, 0.44, 0.38,
	})
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
		return 1.1
	}

	Cx = 0.4
	if Re < 3.0e5 {
		// TODO: add points
		Cx = math.Max(Cx, 3.01192*pow.E2(x)-33.6354*x+94.2379)
	}

	// TODO: add points
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
	return interpolTable(math.Log10(λe), ϕ, 4, 5, []float64{
		-0.1, 0.00, 1.00, 2.00, 2.30103, //1.00, 10.0, 100., 200.,
		0.10, 0.98, 0.99, 0.99, 1.00,
		0.50, 0.88, 0.91, 0.98, 1.00,
		0.90, 0.82, 0.87, 0.97, 1.00,
		0.95, 0.73, 0.80, 0.96, 1.00,
		1.00, 0.60, 0.70, 0.95, 1.00,
	})
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

// example of data:
//
//	  XXX 1 2 3 4 5
//		4    2 3 3 3 2
//		5    4 5 6 8 8
//
// example of x,y:
//
//	2.5, 4.5
func interpolTable(x, y float64, xSize, ySize int, xyData []float64) (z float64) {
	if len(xyData) != xSize*ySize+xSize+ySize+1 {
		panic("wrong size")
	}
	xList := xyData[1 : xSize+1]
	yList := make([]float64, ySize)
	for i := 0; i < ySize; i++ {
		yList[i] = xyData[(xSize+1)*(i+1)]
	}
	data := make([][]float64, ySize)
	for i := 0; i < ySize; i++ {
		data[i] = make([]float64, xSize)
	}
	for r := 0; r < ySize; r++ {
		for c := 0; c < xSize; c++ {
			data[r][c] = xyData[(xSize+1)*(r+1)+1+c]
		}
	}
	// check outside table
	if x < xList[0] {
		panic(fmt.Errorf("less xList: %e < %e", x, xList[0]))
	}
	if xList[len(xList)-1] < x {
		panic("more xList")
	}
	if y < yList[0] {
		panic("less yList")
	}
	if yList[len(yList)-1] < y {
		panic(fmt.Errorf("less yList: %e < %v", y, yList))
	}
	// parameters now in table
	// generate a column
	for i := 1; i < xSize; i++ {
		if xList[i-1] <= x && x <= xList[i] {
			for j := 1; j < ySize; j++ {
				if yList[j-1] <= y && y <= yList[j] {
					z0 := interpol(xList[i-1], x, xList[i], data[j-1][i-1], data[j-1][i])
					z2 := interpol(xList[i-1], x, xList[i], data[j][i-1], data[j][i])
					z := interpol(yList[j-1], y, yList[j], z0, z2)
					return z
				}
			}
		}
	}
	panic(fmt.Errorf("[%e,%e]", x, y))
}
