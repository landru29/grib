package griblib

import (
	"encoding/binary"
	"fmt"
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

	if f != float64(0.000020) {
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

	y := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	binary.BigEndian.PutUint64(y, x)

	fmt.Printf("%d %v\n", x, y)

	fmt.Printf("val %X\n", x)

	f, err := GetFloat(x, 32)
	if err != nil {
		t.Error(err)
		return
	}

	if f != float64(42.5) {
		t.Errorf("Expected 42.5, got %f\n", f)
	}

}
