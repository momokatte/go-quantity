package quantity

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFixedRounder(t *testing.T) {

	assertFloat := func(expected string, actual float64) (err error) {
		actuals := fmt.Sprintf("%0.8g", actual)
		if actuals != expected {
			err = fmt.Errorf("expected '%s', got '%s'", expected, actuals)
		}
		return
	}

	var err error
	r := NewFixedRounder(big.NewRat(1, 100))
	//var f float64

	if err = assertFloat("12345.68", r.RoundFloat(12345.6789)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("1234.57", r.RoundFloat(1234.56789)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("123.46", r.RoundFloat(123.456789)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("12.35", r.RoundFloat(12.3456789)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("1.23", r.RoundFloat(1.23456789)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("0.12", r.RoundFloat(0.123456789)); err != nil {
		t.Fatal(err.Error())
	}

	if err = assertFloat("-0.12", r.RoundFloat(-0.123456789)); err != nil {
		t.Fatal(err.Error())
	}
}
