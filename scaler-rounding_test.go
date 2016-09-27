package quantity

import (
	"fmt"
	"testing"
)

func TestRoundingScaler(t *testing.T) {

	assertString := func(expected string, actual string) (err error) {
		if actual != expected {
			err = fmt.Errorf("expected '%s', got '%s'", expected, actual)
		}
		return
	}

	var err error
	scaler := NewTwentiethRoundingScaler(NewBigNumberScaler())
	var sf ScaledFloat

	sf = scaler.ScaleFloat(1234567890000.0)
	if err = assertString("1.25T", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1214567890000.0)
	if err = assertString("1.2T", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1276567890000.0)
	if err = assertString("1.3T", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(127656789000.0)
	if err = assertString("130B", sf.String()); err != nil {
		t.Fatal(err.Error())
	}

	sf = scaler.ScaleFloat(1276567890.0)
	if err = assertString("1.3B", sf.String()); err != nil {
		t.Fatal(err.Error())
	}
}
