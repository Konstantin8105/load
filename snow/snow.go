package snow

import "fmt"

type Region string

const (
	RegionI    Region = "I"
	RegionII   Region = "II"
	RegionIII  Region = "III"
	RegionIV   Region = "IV"
	RegionV    Region = "V"
	RegionVI   Region = "VI"
	RegionVII  Region = "VII"
	RegionVIII Region = "VIII"
)

func (r Region) Name() string {
	return string(r)
}

func (r Region) Value() float64 {
	switch r {
	case RegionI:
		return 500
	case RegionII:
		return 1000
	case RegionIII:
		return 1500
	case RegionIV:
		return 2000
	case RegionV:
		return 2500
	case RegionVI:
		return 3000
	case RegionVII:
		return 3500
	case RegionVIII:
		return 4000
	}
	panic("not implemented")
}

func ListSg() []Region {
	return []Region{
		RegionI,
		RegionII,
		RegionIII,
		RegionIV,
		RegionV,
		RegionVI,
		RegionVII,
		RegionVIII,
	}
}

func (r Region) String() string {
	name := r.Name()
	return fmt.Sprintf("Snow region: %4s with value = %.1f Pa", name, r.Value())
}
