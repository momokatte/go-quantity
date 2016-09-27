package quantity

import (
	"math/big"
)

type FixedRounder struct {
	RoundingBase *big.Rat
}

func NewFixedRounder(base *big.Rat) (r Rounder) {
	r = &FixedRounder{
		RoundingBase: (&big.Rat{}).Set(base),
	}
	return
}

func (r *FixedRounder) RoundBigFloat(x *big.Float) (z *big.Float) {
	xf, _ := x.Float64()
	z = big.NewFloat(r.RoundFloat(xf))
	return
}

func (r *FixedRounder) RoundFloat(x float64) (z float64) {
	bf, _ := r.RoundingBase.Float64()
	z = RoundFloat(x, bf)
	return
}

func (r *FixedRounder) RoundBigInt(x *big.Int) (z *big.Int) {
	if r.RoundingBase.IsInt() {
		z = RoundBigInt(x, r.RoundingBase.Num())
	} else {
		z = &big.Int{}
		z.Set(x)
	}
	return
}

func (r *FixedRounder) RoundInt(x int64) (z int64) {
	return r.RoundBigInt(big.NewInt(x)).Int64()
}

func (r *FixedRounder) RoundBigRat(x *big.Rat) (z *big.Rat) {
	z = RoundBigRat(x, r.RoundingBase)
	return
}
