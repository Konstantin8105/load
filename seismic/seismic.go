package seismic

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

// Factors in according to code SP14.13330.2018 without Rev.01
type Factors struct {
	// Расчетная сейсмичность площадки строительства
	// Баллы : 7,8,9
	DesignSeismicSite uint8

	// Категория грунта по сейсмическим свойствам
	// Таблица : 4.1
	// Значения : 1,2,3,4
	Grount uint8

	// Классификация объектов в сейсмических районах
	// по их назначению.
	// Таблица : 4.2
	// Строки : 1,2,3,4
	K0row uint8

	// Коэффициент К1, учитывающий допускаемые повреждения
	// зданий и сооружений.
	// Таблица : 5.2
	// Значения меняются от 0.12 до 1.0
	K1 float64

	// Коэффициент Kψ, учитывающий способность здания и
	// сооружений к рассеиванию энергии
	// Таблица : 5.3
	// Строки : 1,2,3
	Kψrow uint8
}

// maximal period for generation betta graph
var PeriodMax = 2560.0 // sec

func Betta(ground uint8) (data [][2]float64) {

	var T []float64
	var lenT int
	switch ground {
	case 1, 2:
		T = []float64{0.0, 0.1, 0.4}
		lenT = len(T)
	case 3, 4:
		T = []float64{0.0, 0.1, 0.8}
		lenT = len(T)
	}

	for {
		Tlast := T[len(T)-1]
		if PeriodMax < Tlast {
			break
		}
		if Tlast < 16 {
			T = append(T, Tlast+0.05)
			continue
		}
		Tlast = float64(int(Tlast))
		T = append(T, Tlast+1.0)
	}

	β := func(period float64) (β float64) {
		if period < 0 {
			panic(fmt.Errorf("not valid period: %.2e", period))
		}
		switch ground {
		case 1, 2:
			if period <= 0.1 {
				β = 1.0 + 15.0*period
				break
			}
			if period < 0.4 {
				β = 2.5
				break
			}
			β = 2.5 * math.Pow(0.4/period, 0.5)
		case 3, 4:
			if period <= 0.1 {
				β = 1.0 + 15.0*period
				break
			}
			if period < 0.8 {
				β = 2.5
				break
			}
			β = 2.5 * math.Pow(0.8/period, 0.5)
		default:
			panic("not valid ground")
		}
		if β < 0.8 {
			β = 0.8
		}
		return
	}

	for i, period := range T {
		if lenT < i && i != len(T)-1 {
			betta := β(period)
			lastBetta := data[len(data)-1][1]
			delta := math.Abs(lastBetta-betta) / lastBetta
			if delta < 1.0/100 { // procent
				continue
			}
		}
		data = append(data, [2]float64{period, β(period)})
	}

	return
}

type Accelerate struct {
	// Period - период колебания
	Period float64

	// Betta - значение бетта по рисунку 5.2
	Betta float64

	// Value - значение ускорения
	// Value[0] - K0 = K0 при РЗ, K1 = K1
	// Value[1] - K0 = K0 при РЗ, K1 = 1.0 для деформаций
	// Value[2] - K0 = K0 при КЗ, K1 = 1.0 для деформаций
	Value [3]float64
}

// Acceleration - расчитывает ускорения для задания сейсмической нагрузки
func (f Factors) Acceleration() (acs []Accelerate, ratios [3]float64) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "Расчет ускорений для сейсмической нагрузки:\n")
	fmt.Fprintf(w, "Расчетная сейсмичность площадки строительства:\t%d\tБаллов |\n",
		f.DesignSeismicSite)

	fmt.Fprintf(w, "Категория грунта по таблице 4.1:\t%d\t\t|\n", f.Grount)

	fmt.Fprintf(w, "Строка по таблице 4.2:\t%d\tстрока\t|\n", f.K0row)
	var K0 [2]float64 // [0] - РЗ, [1] - КЗ
	switch f.K0row {
	case 1:
		K0 = [2]float64{1.2, 2.0}
	case 2:
		K0 = [2]float64{1.1, 1.5}
	case 3:
		K0 = [2]float64{1.0, 1.0}
	case 4:
		K0 = [2]float64{0.8, 0.0}
	default:
		panic("not valid value of K0row")
	}
	fmt.Fprintf(w, "Коэффициент K0 при расчете на РЗ:\t%.2f\t\t|\n", K0[0])
	fmt.Fprintf(w, "Коэффициент K0 при расчете на КЗ:\t%.2f\t\t|\n", K0[1])

	K1 := f.K1
	fmt.Fprintf(w, "Коэффициент K1 по таблице 5.2:\t%.2f\t\t|\n", K1)

	fmt.Fprintf(w, "Строка по таблице 5.3:\t%d\tстрока\t|\n", f.Kψrow)
	var Kψ float64
	switch f.Kψrow {
	case 1:
		Kψ = 1.5
	case 2:
		Kψ = 1.3
	case 3:
		Kψ = 1.0
	default:
		panic("not valid value of Kψrow")
	}
	fmt.Fprintf(w, "Коэффициент Kψ по таблице 5.3:\t%.2f\t\t|\n", Kψ)

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Расчёт:\n")
	var A float64
	switch f.DesignSeismicSite {
	case 7:
		A = 1.0 // м/кв.сек
	case 8:
		A = 2.0 // м/кв.сек
	case 9:
		A = 4.0 // м/кв.сек
	default:
		panic("not valid value of DesignSeismicSite")
	}
	fmt.Fprintf(w, "Коэффициент A по пункту 5.5:\t%.2f\tм/кв.сек\n", A)

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Период\t|\tβ\t|\tУскорение\t|\tУскорение\t|\n")
	fmt.Fprintf(w, " \t|\t \t|\tпри РЗ\t|\tпри РЗ\t|\n")
	fmt.Fprintf(w, " \t|\t \t|\tK0=%.2f\t|\tK0=%.2f\t|\n",
		K0[0], K0[0])
	fmt.Fprintf(w, " \t|\t \t|\tK1=%.2f\t|\tK1=%.2f\t|\n",
		K1, 1.0)
	fmt.Fprintf(w, "сек\t|\t \t|\tм/сек^2\t|\tм/сек^2\t|\n")
	βs := Betta(f.Grount)
	for _, β := range βs {
		period := β[0]
		β := β[1]
		acceleration := [3]float64{
			K0[0] * K1 * Kψ * A * β,
			K0[0] * 1. * Kψ * A * β, // for deformation, K1 = 1
		}
		fmt.Fprintf(w, "%8.2f\t|\t%6.4f\t|\t%6.3f\t|\t%6.3f\t|\n",
			period, β,
			acceleration[0],
			acceleration[1],
		)
		acs = append(acs, Accelerate{
			Period: period,
			Betta:  β,
			Value:  acceleration,
		})
	}

	ratios = [3]float64{
		K0[0] * K1 * Kψ * A / (K0[0] * 1. * Kψ * A),
		K0[0] * 1. * Kψ * A / (K0[0] * 1. * Kψ * A),
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Коэффициенты соотношения к расчету РЗ:\n")

	fmt.Fprintf(w, "(ускорение для РЗ при K1 = K1)  / (ускорение для РЗ при K1 = 1.0)\t%.3f\n", ratios[0])
	fmt.Fprintf(w, "(ускорение для РЗ при K1 = 1.0) / (ускорение для РЗ при K1 = 1.0)\t%.3f\n", ratios[1])

	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	return
}
