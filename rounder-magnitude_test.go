package quantity

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMagnitudeRounder(t *testing.T) {

	assertFloat := func(expected string, actual float64) (err error) {
		actuals := fmt.Sprintf("%0.8g", actual)
		if actuals != expected {
			err = fmt.Errorf("expected '%s', got '%s'", expected, actuals)
		}
		return
	}

	var err error
	r := NewTwentiethRounder()
	var f float64

	if err = assertFloat("12500", r.RoundFloat(12345.67)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("1250", r.RoundFloat(1234.567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("125", r.RoundFloat(123.4567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("12.5", r.RoundFloat(12.34567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("1.25", r.RoundFloat(1.234567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("0.125", r.RoundFloat(0.1234567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("0.0125", r.RoundFloat(0.01234567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("0.00125", r.RoundFloat(0.001234567)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("0.000125", r.RoundFloat(0.0001234567)); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(1234567, 1)).Float64()
	if err = assertFloat("1250000", f); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(1234567, 10)).Float64()
	if err = assertFloat("125000", f); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(1234567, 100)).Float64()
	if err = assertFloat("12500", f); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(1234567, 1000)).Float64()
	if err = assertFloat("1250", f); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(40654321, 1000)).Float64()
	if err = assertFloat("40500", f); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(40876543, 1000)).Float64()
	if err = assertFloat("41000", f); err != nil {
		t.Fatal(err.Error())
	}

	f, _ = r.RoundBigRat(big.NewRat(1678901, 1000)).Float64()
	if err = assertFloat("1700", f); err != nil {
		t.Fatal(err.Error())
	}
}
