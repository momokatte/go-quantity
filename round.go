package quantity

import (
	"math"
	"math/big"
)

func RoundBigFloat(x *big.Float, base *big.Float) (z *big.Float) {
	xf, _ := x.Float64()
	yf, _ := base.Float64()
	z = big.NewFloat(RoundFloat(xf, yf))
	return
}

func RoundFloat(x float64, base float64) (z float64) {
	z = math.Trunc(x / base)
	if math.Mod(x, base) >= (base / 2) {
		z += 1
	}
	z *= base
	return
}

func RoundBigInt(x *big.Int, base *big.Int) (z *big.Int) {
	z = &big.Int{}
	rem := &big.Int{}
	z, rem = z.QuoRem(x, base, rem)
	if rem.Mul(rem, big.NewInt(2)).Cmp(base) >= 0 {
		z.Add(z, big.NewInt(1))
	}
	z.Mul(z, base)
	return
}

func RoundInt(x int64, base int64) (z int64) {
	z = RoundBigInt(big.NewInt(x), big.NewInt(base)).Int64()
	return
}

func RoundBigRat(x *big.Rat, base *big.Rat) (z *big.Rat) {
	z = &big.Rat{}
	if x.IsInt() && base.IsInt() {
		z.SetInt(RoundBigInt(x.Num(), base.Num()))
	} else {
		xf, _ := x.Float64()
		yf, _ := base.Float64()
		z.SetFloat64(RoundFloat(xf, yf))
	}
	return
}
