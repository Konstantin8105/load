package wind

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"testing"
	"text/tabwriter"
)

func TestFactorXi(t *testing.T) {
	tcs := []struct {
		e          float64
		xi30, xi15 float64
	}{
		{0.000000, 0.9285, 0.9385},
		{0.005237, 1.0274, 1.2491},
		{0.014249, 1.1977, 1.5250},
		{0.035085, 1.5048, 1.8328},
		{0.039403, 1.5563, 1.8790},
		{0.063548, 1.7099, 2.1072},
		{0.065975, 1.7211, 2.1297},
		{0.082751, 1.8027, 2.2833},
		{0.095881, 1.8635, 2.3682},
		{0.106078, 1.9033, 2.4266},
		{0.124738, 1.9659, 2.5280},
		{0.139229, 2.0174, 2.6007},
		{0.158098, 2.0785, 2.6784},
		{0.164401, 2.0944, 2.7031},
		{0.186549, 2.1297, 2.7968},
		{0.200214, 2.1299, 2.8167},
		{0.200469, 2.1297, 2.8178},
	}
	isOk := func(x1, x2 float64, t *testing.T) error {
		eps := 1.0 / 100.0 // 1%
		act := math.Abs((x1 - x2) / x1)
		if act > eps {
			return fmt.Errorf("Not valid precision: %.2f%% (%.4f,%.4f)", act*100.0, x1, x2)
		}
		return nil
	}
	for index, tc := range tcs {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			xi30 := FactorXi(LogDecriment30, tc.e)
			xi15 := FactorXi(LogDecriment15, tc.e)
			if err := isOk(xi30, tc.xi30, t); err != nil {
				t.Errorf("xi30 : %v", err)
			}
			if err := isOk(xi15, tc.xi15, t); err != nil {
				t.Errorf("xi15 : %v", err)
			}
		})
	}
}

func ExampleFactorKz() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "factor kz\n")
	fmt.Fprintf(w, "ze\tA\tB\tC\n")
	for _, ze := range []float64{5, 10, 20, 40, 60, 80, 100, 150, 200, 250, 300} {
		fmt.Fprintf(w, "%.0f", ze)
		for _, zone := range []Zone{ZoneA, ZoneB, ZoneC} {
			kz, err := FactorKz(zone, ze)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "\t%.2f", kz)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor kz
	// ze  A    B    C
	// 5   0.81 0.49 0.28
	// 10  1.00 0.65 0.40
	// 20  1.23 0.86 0.57
	// 40  1.52 1.13 0.80
	// 60  1.71 1.33 0.98
	// 80  1.87 1.49 1.13
	// 100 2.00 1.63 1.26
	// 150 2.25 1.92 1.55
	// 200 2.46 2.15 1.79
	// 250 2.63 2.36 2.00
	// 300 2.77 2.53 2.19
}

func ExampleFactorZeta() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "factor zeta\n")
	fmt.Fprintf(w, "ze\tA\tB\tC\n")
	for _, ze := range []float64{5, 10, 20, 40, 60, 80, 100, 150, 200, 250, 300} {
		fmt.Fprintf(w, "%.0f", ze)
		for _, zone := range []Zone{ZoneA, ZoneB, ZoneC} {
			zeta, err := FactorZeta(zone, ze)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "\t%.2f", zeta)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor zeta
	// ze  A    B    C
	// 5   0.84 1.22 2.12
	// 10  0.76 1.06 1.78
	// 20  0.68 0.92 1.50
	// 40  0.62 0.80 1.26
	// 60  0.58 0.74 1.14
	// 80  0.56 0.70 1.06
	// 100 0.54 0.67 1.00
	// 150 0.51 0.62 0.90
	// 200 0.48 0.58 0.84
	// 250 0.47 0.56 0.80
	// 300 0.46 0.54 0.76
}

func ExampleFactorNu() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "factor Nu\n")
	for _, ro := range []float64{0.1, 5, 10, 20, 40, 80, 160} {
		fmt.Fprintf(w, "| %6.2f |", ro)
		for _, xi := range []float64{5, 10, 20, 40, 80, 160, 350} {
			v := FactoNu(ro, xi)
			fmt.Fprintf(w, "\t%.2f", v)
		}
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "factor Nu with middle points\n")
	for _, ro := range []float64{0.1, 3, 5, 7, 10, 15, 20, 25, 40, 60, 80, 100, 160} {
		fmt.Fprintf(w, "| %6.2f |", ro)
		for _, xi := range []float64{5, 7, 10, 15, 20, 30, 40, 60, 80, 120, 160, 300, 350} {
			v := FactoNu(ro, xi)
			fmt.Fprintf(w, "\t%.2f", v)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor Nu
	// |   0.10 | 0.95 0.92 0.88 0.83 0.76 0.67 0.56
	// |   5.00 | 0.89 0.87 0.84 0.80 0.73 0.65 0.54
	// |  10.00 | 0.85 0.84 0.81 0.77 0.71 0.64 0.53
	// |  20.00 | 0.80 0.78 0.76 0.73 0.68 0.61 0.51
	// |  40.00 | 0.72 0.72 0.70 0.67 0.63 0.57 0.48
	// |  80.00 | 0.63 0.63 0.61 0.59 0.56 0.51 0.44
	// | 160.00 | 0.53 0.53 0.52 0.50 0.47 0.44 0.38
	// factor Nu with middle points
	// |   0.10 | 0.95 0.94 0.92 0.90 0.88 0.85 0.83 0.79 0.76 0.72 0.67 0.59 0.56
	// |   3.00 | 0.91 0.90 0.89 0.87 0.86 0.83 0.81 0.78 0.74 0.70 0.66 0.58 0.55
	// |   5.00 | 0.89 0.88 0.87 0.85 0.84 0.82 0.80 0.77 0.73 0.69 0.65 0.57 0.54
	// |   7.00 | 0.87 0.87 0.86 0.84 0.83 0.81 0.79 0.76 0.72 0.68 0.65 0.56 0.54
	// |  10.00 | 0.85 0.85 0.84 0.82 0.81 0.79 0.77 0.74 0.71 0.68 0.64 0.56 0.53
	// |  15.00 | 0.82 0.82 0.81 0.80 0.79 0.77 0.75 0.72 0.70 0.66 0.62 0.55 0.52
	// |  20.00 | 0.80 0.79 0.78 0.77 0.76 0.74 0.73 0.71 0.68 0.65 0.61 0.54 0.51
	// |  25.00 | 0.78 0.77 0.77 0.76 0.74 0.73 0.71 0.69 0.67 0.63 0.60 0.53 0.50
	// |  40.00 | 0.72 0.72 0.72 0.71 0.70 0.69 0.67 0.65 0.63 0.60 0.57 0.50 0.48
	// |  60.00 | 0.68 0.68 0.68 0.67 0.66 0.64 0.63 0.61 0.59 0.57 0.54 0.48 0.46
	// |  80.00 | 0.63 0.63 0.63 0.62 0.61 0.60 0.59 0.57 0.56 0.54 0.51 0.46 0.44
	// | 100.00 | 0.60 0.60 0.60 0.60 0.59 0.58 0.57 0.55 0.54 0.52 0.49 0.44 0.42
	// | 160.00 | 0.53 0.53 0.53 0.53 0.52 0.51 0.50 0.48 0.47 0.45 0.44 0.40 0.38
}