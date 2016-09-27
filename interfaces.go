package quantity

import (
	"math/big"
)

type Rounder interface {
	RoundBigFloat(x *big.Float) *big.Float
	RoundFloat(x float64) float64
	RoundBigInt(x *big.Int) *big.Int
	RoundInt(x int64) int64
	RoundBigRat(x *big.Rat) *big.Rat
}

type Scaler interface {
	ScaleBigFloat(x *big.Float) ScaledFloat
	ScaleFloat(x float64) ScaledFloat
	ScaleBigInt(x *big.Int) ScaledRat
	ScaleInt(x int64) ScaledRat
	ScaleBigRat(x *big.Rat) ScaledRat
}
