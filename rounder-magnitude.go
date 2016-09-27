package quantity

import (
	"math"
	"math/big"
)

type MagnitudeRounder struct {
	RoundingFactor *big.Rat
}

func NewMagnitudeRounder(factor *big.Rat) (r Rounder) {
	r = &MagnitudeRounder{
		RoundingFactor: factor,
	}
	return
}

func NewTwentiethRounder() (r Rounder) {
	r = NewMagnitudeRounder(big.NewRat(1, 20))
	return
}

func NewHundredthRounder() (r Rounder) {
	r = NewMagnitudeRounder(big.NewRat(1, 100))
	return
}

func (r *MagnitudeRounder) RoundBigFloat(x *big.Float) (z *big.Float) {
	z = &big.Float{}
	if x.IsInt() {
		xr, _ := x.Rat(nil)
		z = z.SetRat(r.RoundBigRat(xr))
	} else {
		xf, _ := x.Float64()
		z = z.SetFloat64(r.RoundFloat(xf))
	}
	return
}

func (r *MagnitudeRounder) RoundFloat(x float64) (z float64) {
	base, _ := r.CalcFloatBase(x).Float64()
	z = RoundFloat(x, base)
	return
}

func (r *MagnitudeRounder) RoundBigInt(x *big.Int) (z *big.Int) {
	base := r.CalcIntBase(x.Int64())
	if base.IsInt() {
		z = RoundBigInt(x, base.Num())
	} else {
		xr := big.NewRat(x.Int64(), 1)
		z = RoundBigRat(xr, base).Num()
	}
	return
}

func (r *MagnitudeRounder) RoundInt(x int64) (z int64) {
	xi := big.NewInt(x)
	z = r.RoundBigInt(xi).Int64()
	return
}

func (r *MagnitudeRounder) RoundBigRat(x *big.Rat) (z *big.Rat) {
	base := r.CalcRatBase(x)
	z = RoundBigRat(x, base)
	return
}

func (r *MagnitudeRounder) CalcFloatBase(x float64) (z *big.Rat) {
	lt := math.Log10(x)
	z = &big.Rat{}
	if lt >= 0 {
		lti := int(math.Trunc(lt))
		z.SetInt64(int64(math.Pow10(lti)))
	} else {
		lti := int(math.Trunc(math.Abs(lt))) + 1
		z.SetFrac64(1, int64(math.Pow10(lti)))
	}
	z = z.Mul(z, r.RoundingFactor)
	return
}

func (r *MagnitudeRounder) CalcIntBase(x int64) (z *big.Rat) {
	z = big.NewRat(x, 1)
	z = r.CalcRatBase(z)
	return
}

func (r *MagnitudeRounder) CalcRatBase(x *big.Rat) (z *big.Rat) {
	xf, _ := x.Float64()
	z = r.CalcFloatBase(xf)
	return
}
