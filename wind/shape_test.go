package wind

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/Konstantin8105/compare"
)

func Test(t *testing.T) {
	t.Run("frame", func(t *testing.T) {
		var buf bytes.Buffer
		Wsum := Frame(&buf, ZoneA, RegionII, LogDecriment15, 18.965, []float64{1.393})
		for _, z := range []float64{2.8, 5, 10} {
			fmt.Fprintf(&buf, "Wsum[z = %6.3f m] = %6.1f Pa\n", z, Wsum(z))
		}
		compare.Test(t, "test.frame", buf.Bytes())
	})
	t.Run("rectangle", func(t *testing.T) {
		var buf bytes.Buffer
		Wsum := Rectangle(&buf, ZoneA, RegionII, LogDecriment15, 5.38, 7.32, 18.965, 0.000, []float64{1.393})
		for _, z := range []float64{10, 15} {
			fmt.Fprintf(&buf, "z = %6.3f m\n", z)
			for _, side := range ListRectangleSides() {
				fmt.Fprintf(&buf, "Wsum[%s] = %6.1f Pa\n", side,
					Wsum[side](z))
			}
			fmt.Fprintf(&buf, "Summary D and E:\n")
			fmt.Fprintf(&buf, "Wsum[z = %6.3f m] = %6.1f Pa\n", z,
				math.Abs(Wsum[SideD](z))+math.Abs(Wsum[SideE](z)))
			fmt.Fprintf(&buf, "\n")
		}
		compare.Test(t, "test.rectangle", buf.Bytes())
	})
	t.Run("cylinder", func(t *testing.T) {
		var buf bytes.Buffer
		Wsum := Cylinder(&buf, ZoneA, RegionII, LogDecriment15, 0.200, 4.710, 10.100, 2.800,
			[]float64{3.091, 3.414, 3.719})
		for _, z := range []float64{2.8, 5, 10} {
			fmt.Fprintf(&buf, "Wsum[z = %6.3f m] = %6.1f Pa\n", z, Wsum(z))
		}
		compare.Test(t, "test.cylinder", buf.Bytes())
	})
}
