package deformation

import (
	"fmt"
	"math"
)

// СП 20.13330.2016
// Д.2 Предельные прогибы
// Д.2.1 Вертикальные предельные прогибы элементов конструкций
// Элементы конструкций: 2
//
// length unit: meter
func Vertical(l float64) (deformMax float64, err error) {
	l = math.Abs(l)
	if eps := 1e-5; l < eps {
		l = eps
	}
	ds := [][2]float64{
		{0.000, 1.0 / 120.0},
		{1.000, 1.0 / 120.0},
		{3.000, 1.0 / 150.0},
		{6.000, 1.0 / 200.0},
		{12.000, 1.0 / 250.0},
		{24.000, 1.0 / 300.0},
	}
	for i := range ds {
		ds[i][1] *= l
	}
	for i := range ds {
		if i == 0 {
			continue
		}
		if ds[i-1][0] <= l && l <= ds[i][0] {
			deform := ds[i-1][1] + (ds[i][1]-ds[i-1][1])*(l-ds[i-1][0])/(ds[i][0]-ds[i-1][0])
			return deform, nil
		}
	}
	err = fmt.Errorf("approximation is not valid for %.1e", l)
	return
}
