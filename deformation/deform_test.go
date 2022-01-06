package deformation

import (
	"fmt"
	"os"
	"sort"
)

func Example() {
	// corner lenghts
	L := []float64{0.0001, 1.000, 3.000, 6.000, 12.000, 24.000}
	// middle points
	for i, size := 1, len(L); i < size; i++ {
		L = append(L, (L[i]+L[i-1])*0.5)
	}
	// sort
	sort.Float64s(L)
	// vertical deformation
	fmt.Fprintf(os.Stdout, "%8s %5s\n", "L,mm", "D,mm")
	for i := range L {
		dmax, err := Vertical(L[i])
		if err != nil {
			panic(err)
		}
		ratio := L[i] / dmax
		fmt.Fprintf(os.Stdout, "%8.3f %5.3f 1.0/%5.1f\n", L[i], dmax, ratio)
	}
	// Output:
	// L,mm  D,mm
	//    0.000 0.000 1.0/120.0
	//    0.500 0.004 1.0/120.0
	//    1.000 0.008 1.0/120.0
	//    2.000 0.015 1.0/133.3
	//    3.000 0.020 1.0/150.0
	//    4.500 0.026 1.0/171.4
	//    6.000 0.030 1.0/200.0
	//    9.000 0.041 1.0/222.2
	//   12.000 0.048 1.0/250.0
	//   18.000 0.066 1.0/272.7
	//   24.000 0.080 1.0/300.0
}
