package quantity

import (
	"math"
	"math/big"
)

type BasicScaler struct {
	scales    []Scale
	SigDigits int
}

func NewBigNumberScaler() (s Scaler) {
	s = &BasicScaler{
		scales:    MakeBigNumberScales(),
		SigDigits: 3,
	}
	return
}

func NewSIUnitScaler() (s Scaler) {
	s = &BasicScaler{
		scales:    MakeSIUnitScales(),
		SigDigits: 4,
	}
	return
}

func (s *BasicScaler) NearestScale(exponent int) (sc Scale) {
	for _, scale := range s.scales {
		if exponent >= scale.Exponent {
			sc = scale
			break
		}
	}
	return
}

func (s *BasicScaler) ScaleBigFloat(x *big.Float) (sf ScaledFloat) {
	xf, _ := x.Float64()
	sf = s.ScaleFloat(xf)
	return
}

func (s *BasicScaler) ScaleFloat(x float64) (sf ScaledFloat) {
	scale := s.NearestScale(int(math.Trunc(math.Log10(x))))
	sf = ScaleFloat(x, scale)
	sf.SigDigits = s.SigDigits
	return
}

func (s *BasicScaler) ScaleBigInt(x *big.Int) (sr ScaledRat) {
	sr = s.ScaleInt(x.Int64())
	return
}

func (s *BasicScaler) ScaleInt(x int64) (sr ScaledRat) {
	sr = s.ScaleBigRat(big.NewRat(x, 1))
	return
}

func (s *BasicScaler) ScaleBigRat(x *big.Rat) (sr ScaledRat) {
	xf, _ := x.Float64()
	scale := s.NearestScale(int(math.Trunc(math.Log10(xf))))
	sr = ScaleBigRat(x, scale)
	sr.SigDigits = s.SigDigits
	return
}

func MakeBigNumberScales() (scales []Scale) {
	scales = []Scale{
		Scale{12, big.NewRat(1e12, 1), "trillion", "T"},
		Scale{9, big.NewRat(1e9, 1), "billion", "B"},
		Scale{6, big.NewRat(1e6, 1), "million", "M"},
		Scale{3, big.NewRat(1e3, 1), "thousand", "K"},
	}
	return
}

func MakeSIUnitScales() (scales []Scale) {
	scales = []Scale{
		Scale{18, big.NewRat(1e18, 1), "exa", "E"},
		Scale{15, big.NewRat(1e15, 1), "peta", "P"},
		Scale{12, big.NewRat(1e12, 1), "tera", "T"},
		Scale{9, big.NewRat(1e9, 1), "giga", "G"},
		Scale{6, big.NewRat(1e6, 1), "mega", "M"},
		Scale{3, big.NewRat(1e3, 1), "kilo", "k"},
		Scale{0, big.NewRat(0, 1), "", ""},
		Scale{-3, big.NewRat(1, 1e3), "milli", "m"},
		Scale{-6, big.NewRat(1, 1e6), "micro", "µ"},
		Scale{-9, big.NewRat(1, 1e9), "nano", "n"},
		Scale{-12, big.NewRat(1, 1e12), "pico", "p"},
		Scale{-15, big.NewRat(1, 1e15), "femto", "f"},
		Scale{-18, big.NewRat(1, 1e18), "atto", "a"},
	}
	return
}
