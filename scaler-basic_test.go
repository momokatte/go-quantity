package quantity

import (
	"fmt"
	"testing"
)

func TestBigNumberScaler(t *testing.T) {

	assertString := func(expected string, actual string) (err error) {
		if actual != expected {
			err = fmt.Errorf("expected '%s', got '%s'", expected, actual)
		}
		return
	}

	var err error
	scaler := NewBigNumberScaler()
	var sf ScaledFloat

	sf = scaler.ScaleFloat(1234567890000.0)
	if err = assertString("1.23T", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(12345678900.0)
	if err = assertString("12.3B", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1234567890.0)
	if err = assertString("1.23B", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(123456789.0)
	if err = assertString("123M", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(12345678.90)
	if err = assertString("12.3M", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1234567.89)
	if err = assertString("1.23M", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(123456.789)
	if err = assertString("123K", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(12345.6789)
	if err = assertString("12.3K", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1234.56789)
	if err = assertString("1.23K", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(123.456789)
	if err = assertString("123", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(12.3456789)
	if err = assertString("12.3", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1.23456789)
	if err = assertString("1.23", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(0.123456789)
	if err = assertString("0.123", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(0.0123456789)
	if err = assertString("0.0123", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(0.00123456789)
	if err = assertString("0.00123", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

}

func TestSIUnitScaler(t *testing.T) {

	assertString := func(expected string, actual string) (err error) {
		if actual != expected {
			err = fmt.Errorf("expected '%s', got '%s'", expected, actual)
		}
		return
	}

	var err error
	scaler := NewSIUnitScaler()
	var sf ScaledFloat

	sf = scaler.ScaleFloat(1234567890000.0)
	if err = assertString("1.235T", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1234567890.0)
	if err = assertString("1.235G", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1234567.89)
	if err = assertString("1.235M", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1234.56789)
	if err = assertString("1.235k", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1.23456789)
	if err = assertString("1.235", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(0.00123456789)
	if err = assertString("1.235m", sf.String()); err != nil {
		t.Fatal(err.Error())
	}
}
