package snow

import (
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"
)

func Example() {
	sgs := ListSg()
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "factor Sg\n")

	fmt.Fprintf(w, "region")
	for _, sg := range sgs {
		fmt.Fprintf(w, "\t%8s", sg.Name())
	}
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "Sg, kPa")
	for _, sg := range sgs {
		fmt.Fprintf(w, "\t%8.2f", sg.Value()/1000.0)
	}
	fmt.Fprintf(w, "\n")

	for _, sg := range sgs {
		fmt.Fprintf(w, "%s\n", sg)
	}

	w.Flush()
	fmt.Fprintf(os.Stdout, "%s", buf.String())
	// Output:
	// factor Sg
	// region         I       II      III       IV        V       VI      VII     VIII
	// Sg, kPa     0.50     1.00     1.50     2.00     2.50     3.00     3.50     4.00
	// Snow region:    I with value = 500.0 Pa
	// Snow region:   II with value = 1000.0 Pa
	// Snow region:  III with value = 1500.0 Pa
	// Snow region:   IV with value = 2000.0 Pa
	// Snow region:    V with value = 2500.0 Pa
	// Snow region:   VI with value = 3000.0 Pa
	// Snow region:  VII with value = 3500.0 Pa
	// Snow region: VIII with value = 4000.0 Pa
}
