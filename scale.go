package quantity

import (
	"fmt"
	"math/big"
)

type Scale struct {
	Exponent int
	Value    *big.Rat
	Name     string
	Symbol   string
}

type ScaledFloat struct {
	Float     big.Float
	Scale     Scale
	SigDigits int
}

func (sf *ScaledFloat) String() string {
	sff, _ := sf.Float.Float64()
	p := "%g%s"
	if sf.SigDigits > 0 {
		p = fmt.Sprintf("%%0.%dg%%s", sf.SigDigits)
	}
	return fmt.Sprintf(p, sff, sf.Scale.Symbol)
}

type ScaledRat struct {
	Rat       big.Rat
	Scale     Scale
	SigDigits int
}

func (sr *ScaledRat) String() string {
	srf, _ := sr.Rat.Float64()
	p := "%g%s"
	if sr.SigDigits > 0 {
		p = fmt.Sprintf("%%0.%dg%%s", sr.SigDigits)
	}
	return fmt.Sprintf(p, srf, sr.Scale.Symbol)
}

func ScaleBigFloat(x *big.Float, scale Scale) (sf ScaledFloat) {
	xf, _ := x.Float64()
	sf = ScaleFloat(xf, scale)
	return
}

func ScaleFloat(x float64, scale Scale) (sf ScaledFloat) {
	z := x
	if scale.Exponent != 0 {
		svf, _ := scale.Value.Float64()
		z /= svf
	}
	sf.Float = *big.NewFloat(z)
	sf.Scale = scale
	return
}

func ScaleBigInt(x *big.Int, scale Scale) (sr ScaledRat) {
	sr = ScaleInt(x.Int64(), scale)
	return
}

func ScaleInt(x int64, scale Scale) (sr ScaledRat) {
	sr = ScaleBigRat(big.NewRat(x, 1), scale)
	return
}

func ScaleBigRat(x *big.Rat, scale Scale) (sr ScaledRat) {
	z := &big.Rat{}
	z.Set(x)
	if scale.Exponent != 0 {
		// shift
		z.Mul(z, scale.Value.Inv(nil))
	}
	sr.Rat = *z
	sr.Scale = scale
	return
}
