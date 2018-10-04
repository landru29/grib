package griblib

import (
	"testing"
)

func Test_GetFloat11(t *testing.T) {
	x := uint64(123456789)

	f, err := GetFloat(x, 11)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(42.5) {
		t.Errorf("Expected 42.5, got %f\n", f)
	}

}

func Test_GetFloat11WithoutExponent(t *testing.T) {
	x := uint64(51221)

	f, err := GetFloat(x, 11)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(0.00002002716064453125) {
		t.Errorf("Expected 0.000020, got %f\n", f)
	}
}

func Test_GetFloat11Zero(t *testing.T) {
	x := uint64(0)

	f, err := GetFloat(x, 11)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(0.00000) {
		t.Errorf("Expected 0.00000, got %f\n", f)
	}
}

func Test_GetFloat11WithoutMantissa(t *testing.T) {
	x := uint64(52480)

	f, err := GetFloat(x, 11)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(32.0) {
		t.Errorf("Expected 32.0, got %f\n", f)
	}
}

func Test_GetFloat32(t *testing.T) {

	x := uint64(0x42289000)

	f, err := GetFloat(x, 32)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(42.140625) {
		t.Errorf("Expected 42.140625, got %f\n", f)
	}

}

func Test_GetFloat32Negative(t *testing.T) {

	x := uint64(0xc2289000)

	f, err := GetFloat(x, 32)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(-42.140625) {
		t.Errorf("Expected 42.140625, got %f\n", f)
	}

}
