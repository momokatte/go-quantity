package quantity

import (
	"math/big"
)

type RoundingScaler struct {
	Rounder Rounder
	Scaler  Scaler
}

func NewRoundingScaler(rounder Rounder, scaler Scaler) (rs Scaler) {
	rs = &RoundingScaler{
		Rounder: rounder,
		Scaler:  scaler,
	}
	return
}

func NewTwentiethRoundingScaler(scaler Scaler) Scaler {
	return NewRoundingScaler(NewTwentiethRounder(), scaler)
}

func NewHundredthRoundingScaler(scaler Scaler) Scaler {
	return NewRoundingScaler(NewHundredthRounder(), scaler)
}

func (rs *RoundingScaler) ScaleBigFloat(x *big.Float) ScaledFloat {
	z := rs.Rounder.RoundBigFloat(x)
	return rs.Scaler.ScaleBigFloat(z)
}

func (rs *RoundingScaler) ScaleFloat(x float64) ScaledFloat {
	z := rs.Rounder.RoundFloat(x)
	return rs.Scaler.ScaleFloat(z)
}

func (rs *RoundingScaler) ScaleBigInt(x *big.Int) ScaledRat {
	z := rs.Rounder.RoundBigInt(x)
	return rs.Scaler.ScaleBigInt(z)
}

func (rs *RoundingScaler) ScaleInt(x int64) ScaledRat {
	z := rs.Rounder.RoundInt(x)
	return rs.Scaler.ScaleInt(z)
}

func (rs *RoundingScaler) ScaleBigRat(x *big.Rat) ScaledRat {
	z := rs.Rounder.RoundBigRat(x)
	return rs.Scaler.ScaleBigRat(z)
}
