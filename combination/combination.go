package combination

import (
	"fmt"
	"sort"
)

type Load struct {
	Name string
}

func (l Load) String() string {
	return l.Name
}

type Multiplication struct {
	Factor   float64
	LoadPart *Load
}

func (m Multiplication) String() string {
	return fmt.Sprintf("%.2f %s", m.Factor, m.LoadPart)
}

type Summ []Multiplication

func (s Summ) Equal(so Summ) bool {
	s1 := []Multiplication(s)
	s2 := []Multiplication(so)
	if len(s1) != len(s2) {
		return false
	}
	size := len(s1)
	for i := 0; i < size; i++ {
		if s1[i].Factor != s2[i].Factor {
			return false
		}
		if s1[i].LoadPart != s2[i].LoadPart {
			return false
		}
	}
	return true
}

func (s Summ) String() string {
	var str string
	for _, m := range []Multiplication(s) {
		str += fmt.Sprintf(" %s", m.String())
	}
	return str
}

func GenerateMain(Pd Load, Pl, Pt []Load) (combs []Summ) {
	generate := func(P []Load, ψ []float64) (comb []Summ) {
		for _, indexes := range generateIndex(len(P), len(ψ)) {
			var s Summ
			for pos, i := range indexes {
				f := ψ[len(ψ)-1] // default last factor
				if pos < len(ψ) {
					f = ψ[pos]
				}
				s = append(s, Multiplication{
					Factor:   f,
					LoadPart: &P[i],
				})
			}
			// append
			comb = append(comb, s)
		}

		// sorting by name
		for ic := range comb {
			sort.Slice([]Multiplication(comb[ic]), func(i, j int) bool {
				if comb[ic][i].Factor != comb[ic][j].Factor {
					return false
				}
				return comb[ic][i].LoadPart.Name < comb[ic][j].LoadPart.Name
			})
		}
		// check on unique summ
		var combFilter []Summ
		for i := 0; i < len(comb); i++ {
			isUnique := true
			for j := 0; j < len(combFilter); j++ {
				if comb[i].Equal(combFilter[j]) {
					isUnique = false
					break
				}
			}
			if !isUnique {
				continue
			}
			combFilter = append(combFilter, comb[i])
		}
		comb = combFilter
		return
	}

	// generate long loads
	ψl := []float64{1.00, 0.95}
	combl := generate(Pl, ψl)

	// generate short loads
	ψt := []float64{1.00, 0.90, 0.70}
	combt := generate(Pt, ψt)

	// create combinations
	combs = append(combs, []Multiplication{
		Multiplication{
			Factor:   1.0,
			LoadPart: &Pd,
		},
	})

	{ // long load
		var newcombs []Summ
		for j := range combs {
			for i := 0; i < len(combl); i++ {
				newcombs = append(newcombs, append(combs[j], combl[i]...))
			}
		}
		combs = newcombs
	}
	{ // short load
		var newcombs []Summ
		for j := range combs {
			for i := 0; i < len(combt); i++ {
				newcombs = append(newcombs, append(combs[j], combt[i]...))
			}
		}
		combs = newcombs
	}

	return
}

func generateIndex(size, limit int) (combs [][]int) {
	positions := make([]int, size)
	for i := range positions {
		positions[i] = i
	}
	permutation(positions, func(a []int) {
		p := make([]int, size)
		copy(p, a)
		combs = append(combs, p)
	}, 0)
	return
}

// Permute the values at index i to len(a)-1.
func permutation(a []int, f func([]int), i int) {
	if len(a) < i {
		f(a)
		return
	}
	permutation(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permutation(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
