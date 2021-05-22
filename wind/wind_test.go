package wind

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"testing"
	"text/tabwriter"
)

func ExampleRegion() {
	wos := ListWo()
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "factor Wo\n")

	fmt.Fprintf(w, "region")
	for _, wo := range wos {
		fmt.Fprintf(w, "\t%8s", wo.Name())
	}
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Wo, kPa")
	for _, wo := range wos {
		fmt.Fprintf(w, "\t%8.2f", float64(wo)/1000.0)
	}
	fmt.Fprintf(w, "\n")

	for _, wo := range wos {
		fmt.Fprintf(w, "%s\n", wo)
	}

	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor Wo
	// region        Ia        I       II      III       IV        V       VI      VII
	// Wo, kPa     0.17     0.23     0.30     0.38     0.48     0.60     0.73     0.85
	// Wind region:  Ia with value = 170.0 Pa
	// Wind region:   I with value = 230.0 Pa
	// Wind region:  II with value = 300.0 Pa
	// Wind region: III with value = 380.0 Pa
	// Wind region:  IV with value = 480.0 Pa
	// Wind region:   V with value = 600.0 Pa
	// Wind region:  VI with value = 730.0 Pa
	// Wind region: VII with value = 850.0 Pa
}

func TestFactorXi(t *testing.T) {
	tcs := []struct {
		e          float64
		xi30, xi15 float64
	}{
		{0.000000, 1.0000, 1.0000},
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
		if eps < act {
			return fmt.Errorf("Not valid precision: %.2f%% (%.4f,%.4f)", act*100.0, x1, x2)
		}
		return nil
	}
	for index, tc := range tcs {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			xi30 := factorXi(LogDecriment30, tc.e)
			xi15 := factorXi(LogDecriment15, tc.e)
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
			kz := FactorKz(zone, ze)
			fmt.Fprintf(w, "\t%.2f", kz)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor kz
	// ze  A    B    C
	// 5   0.75 0.50 0.40
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
			zeta := FactorZeta(zone, ze)
			fmt.Fprintf(w, "\t%.2f", zeta)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor zeta
	// ze  A    B    C
	// 5   0.85 1.22 1.78
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
			v := FactorNu(ro, xi)
			fmt.Fprintf(w, "\t%.2f", v)
		}
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "factor Nu with middle points\n")
	for _, ro := range []float64{0.1, 3, 5, 7, 10, 15, 20, 25, 40, 60, 80, 100, 160} {
		fmt.Fprintf(w, "| %6.2f |", ro)
		for _, xi := range []float64{5, 7, 10, 15, 20, 30, 40, 60, 80, 120, 160, 300, 350} {
			v := FactorNu(ro, xi)
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

func ExampleGraphB14() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "Graph B14\n")
	Res := []float64{1e5, 3e5, 4e5, 5e5, 6e5, 1e6, 1e7, 1e8}
	fmt.Fprintf(w, "| %7s |", "-")
	for _, re := range Res {
		fmt.Fprintf(w, "\t%6.0e", re)
	}
	fmt.Fprintf(w, "\n")
	for _, delta := range []float64{1e-6, 1e-5, 5e-5, 10e-5, 12e-5, 50e-5, 100e-5, 1e-2} {
		fmt.Fprintf(w, "| %6.1e |", delta)
		for _, re := range Res {
			d := 1.4
			cx := GraphB14(d, delta*d, re)
			fmt.Fprintf(w, "\t%6.4f", cx)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// Graph B14
	// |       - |  1e+05  3e+05  4e+05  5e+05  6e+05  1e+06  1e+07  1e+08
	// | 1.0e-06 | 0.6000 0.6000 0.6000 0.6000 0.2000 0.2000 0.2000 0.2000
	// | 1.0e-05 | 0.6000 0.6000 0.6000 0.6000 0.2000 0.2000 0.2000 0.2000
	// | 5.0e-05 | 0.6000 0.6000 0.6000 0.6000 0.2690 0.2690 0.2690 0.2690
	// | 1.0e-04 | 0.6000 0.6000 0.6000 0.6000 0.3000 0.3000 0.3000 0.3000
	// | 1.2e-04 | 0.6000 0.6000 0.6000 0.6000 0.3070 0.3070 0.3070 0.3070
	// | 5.0e-04 | 0.6000 0.6000 0.6000 0.6000 0.3690 0.3690 0.3690 0.3690
	// | 1.0e-03 | 0.6000 0.6000 0.6000 0.6000 0.4000 0.4000 0.4000 0.4000
	// | 1.0e-02 | 0.6000 0.6000 0.6000 0.6000 0.4000 0.4000 0.4000 0.4000
}

func ExampleGraphB17() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "Graph B15\n")
	Res := []float64{1e4, 1e5, 2e5, 4e5, 6e5, 8e5, 1e6, 5e6, 1e7, 5e7, 1e8}
	fmt.Fprintf(w, "| %7s |", "-")
	for _, re := range Res {
		fmt.Fprintf(w, "\t%6.0e", re)
	}
	fmt.Fprintf(w, "\n")
	for _, delta := range []float64{1e-1, 1e-2, 1e-3, 1e-4, 1e-5, 1e-6} {
		fmt.Fprintf(w, "| %6.1e |", delta)
		for _, re := range Res {
			d := 1.4
			cx := GraphB17(d, delta*d, re)
			fmt.Fprintf(w, "\t%6.4f", cx)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// Graph B15
	// |       - |  1e+04  1e+05  2e+05  4e+05  6e+05  8e+05  1e+06  5e+06  1e+07  5e+07  1e+08
	// | 1.0e-01 | 1.2000 1.2000 1.2000 1.2000 1.2000 1.2000 1.2000 1.2000 1.2000 1.2000 1.2000
	// | 1.0e-02 | 1.2000 1.2000 0.9484 0.9891 1.0099 1.0232 1.0328 1.0819 1.0922 1.0906 1.0790
	// | 1.0e-03 | 1.2000 1.2000 0.7067 0.7721 0.8063 0.8287 0.8450 0.9360 0.9605 0.9837 0.9790
	// | 1.0e-04 | 1.2000 1.2000 0.5734 0.5551 0.6027 0.6342 0.6573 0.7900 0.8289 0.8767 0.8790
	// | 1.0e-05 | 1.2000 1.2000 0.5734 0.4000 0.4000 0.4396 0.4695 0.6441 0.6973 0.7698 0.7790
	// | 1.0e-06 | 1.2000 1.2000 0.5734 0.4000 0.4000 0.4396 0.4695 0.6441 0.6973 0.7698 0.7790
}

func ExampleGraphB23() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "Graph B23\n")
	λes := []float64{1, 5, 10, 50, 200, 500}
	ϕs := []float64{0.1, 0.5, 0.9, 0.95, 1.0}
	fmt.Fprintf(w, "| %7s |", "-")
	for _, λe := range λes {
		fmt.Fprintf(w, "\t%6.0e", λe)
	}
	fmt.Fprintf(w, "\n")
	for _, ϕ := range ϕs {
		fmt.Fprintf(w, "| %6.1e |", ϕ)
		for _, λe := range λes {
			Kλ := GraphB23(λe, ϕ)
			fmt.Fprintf(w, "\t%6.4f", Kλ)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// Graph B23
	// |       - |  1e+00  5e+00  1e+01  5e+01  2e+02  5e+02
	// | 1.0e-01 | 0.9800 0.9870 0.9900 0.9900 1.0000 1.0000
	// | 5.0e-01 | 0.8800 0.9010 0.9100 0.9589 1.0000 1.0000
	// | 9.0e-01 | 0.8200 0.8549 0.8700 0.9399 1.0000 1.0000
	// | 9.5e-01 | 0.7300 0.7789 0.8000 0.9118 1.0000 1.0000
	// | 1.0e+00 | 0.6000 0.6699 0.7000 0.8747 1.0000 1.0000
}

func ExampleNuPlates() {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "Table 11.7\n")
	fmt.Fprintf(w, "Plate\tρ\tχ\n")
	b, h, a := 1.0, 2.0, 3.0
	for _, pl := range []Plate{ZOY, ZOX, XOY} {
		ρ, χ := NuPlates(b, h, a, pl)
		fmt.Fprintf(w, "%s\t%.2f\t%.2f\n", pl, ρ, χ)
	}
	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// Table 11.7
	// Plate ρ    χ
	// ZOY   1.00 2.00
	// ZOX   1.20 2.00
	// XOY   1.00 3.00
}
