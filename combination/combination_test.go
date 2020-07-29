package combination_test

import (
	"fmt"
	"os"

	"github.com/Konstantin8105/load/combination"
)

type load struct {
	name string
}

func (l load) LoadName() string {
	return l.name
}

func Example() {
	var (
		DL    = combination.Load(load{name: "DL"})
		Ls    = []combination.Load{load{name: "L1"}, load{name: "L2"}, load{name: "L3"}}
		Ts    = []combination.Load{load{name: "T1"}, load{name: "T2"}, load{name: "T3"}, load{name: "T4"}}
		combs = combination.GenerateMain(DL, Ls, Ts)
	)
	for i := range combs {
		fmt.Fprintf(os.Stdout, "COMB %d\n%s\n", i+1, combs[i])
	}
	// Output:
	// COMB 1
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T1 0.90 T2 0.70 T3 0.70 T4
	// COMB 2
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T1 0.90 T3 0.70 T2 0.70 T4
	// COMB 3
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T1 0.90 T4 0.70 T2 0.70 T3
	// COMB 4
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T2 0.90 T1 0.70 T3 0.70 T4
	// COMB 5
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T2 0.90 T3 0.70 T1 0.70 T4
	// COMB 6
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T2 0.90 T4 0.70 T1 0.70 T3
	// COMB 7
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T3 0.90 T2 0.70 T1 0.70 T4
	// COMB 8
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T3 0.90 T1 0.70 T2 0.70 T4
	// COMB 9
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T3 0.90 T4 0.70 T1 0.70 T2
	// COMB 10
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T4 0.90 T2 0.70 T1 0.70 T3
	// COMB 11
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T4 0.90 T3 0.70 T1 0.70 T2
	// COMB 12
	//  1.00 DL 1.00 L1 0.95 L2 0.95 L3 1.00 T4 0.90 T1 0.70 T2 0.70 T3
	// COMB 13
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T1 0.90 T2 0.70 T3 0.70 T4
	// COMB 14
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T1 0.90 T3 0.70 T2 0.70 T4
	// COMB 15
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T1 0.90 T4 0.70 T2 0.70 T3
	// COMB 16
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T2 0.90 T1 0.70 T3 0.70 T4
	// COMB 17
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T2 0.90 T3 0.70 T1 0.70 T4
	// COMB 18
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T2 0.90 T4 0.70 T1 0.70 T3
	// COMB 19
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T3 0.90 T2 0.70 T1 0.70 T4
	// COMB 20
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T3 0.90 T1 0.70 T2 0.70 T4
	// COMB 21
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T3 0.90 T4 0.70 T1 0.70 T2
	// COMB 22
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T4 0.90 T2 0.70 T1 0.70 T3
	// COMB 23
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T4 0.90 T3 0.70 T1 0.70 T2
	// COMB 24
	//  1.00 DL 1.00 L2 0.95 L1 0.95 L3 1.00 T4 0.90 T1 0.70 T2 0.70 T3
	// COMB 25
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T1 0.90 T2 0.70 T3 0.70 T4
	// COMB 26
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T1 0.90 T3 0.70 T2 0.70 T4
	// COMB 27
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T1 0.90 T4 0.70 T2 0.70 T3
	// COMB 28
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T2 0.90 T1 0.70 T3 0.70 T4
	// COMB 29
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T2 0.90 T3 0.70 T1 0.70 T4
	// COMB 30
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T2 0.90 T4 0.70 T1 0.70 T3
	// COMB 31
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T3 0.90 T2 0.70 T1 0.70 T4
	// COMB 32
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T3 0.90 T1 0.70 T2 0.70 T4
	// COMB 33
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T3 0.90 T4 0.70 T1 0.70 T2
	// COMB 34
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T4 0.90 T2 0.70 T1 0.70 T3
	// COMB 35
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T4 0.90 T3 0.70 T1 0.70 T2
	// COMB 36
	//  1.00 DL 1.00 L3 0.95 L1 0.95 L2 1.00 T4 0.90 T1 0.70 T2 0.70 T3
}
