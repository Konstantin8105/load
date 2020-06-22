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
