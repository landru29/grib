package griblib

import (
	"fmt"
	"math"
)

type floatParameters struct {
	hasSign      bool
	exponentSize byte
	mantissaSize byte
}

var floatingtypes = map[int]floatParameters{
	10: floatParameters{
		hasSign:      false,
		exponentSize: 5,
		mantissaSize: 5,
	},
	11: floatParameters{
		hasSign:      false,
		exponentSize: 5,
		mantissaSize: 6,
	},
	14: floatParameters{
		hasSign:      false,
		exponentSize: 5,
		mantissaSize: 9,
	},
	16: floatParameters{
		hasSign:      true,
		exponentSize: 5,
		mantissaSize: 10,
	},
	32: floatParameters{
		hasSign:      true,
		exponentSize: 8,
		mantissaSize: 32,
	},
	64: floatParameters{
		hasSign:      true,
		exponentSize: 11,
		mantissaSize: 52,
	},
}

func getMask(nBits byte) uint64 {
	mask := uint64(0)
	for i := byte(0); i < nBits; i++ {
		mask = (mask << 1) + 1
	}
	return mask
}

// GetFloat performs a binary cast
func GetFloat(data uint64, bits int) (float64, error) {
	parameters, ok := floatingtypes[bits]
	if !ok {
		return 0, fmt.Errorf("Unrecongnize floating bit size")
	}

	negative := parameters.hasSign && ((data>>(parameters.exponentSize+parameters.mantissaSize))&1 == 1)

	exponentVal := (data >> parameters.mantissaSize) & getMask(parameters.exponentSize)
	mantissaVal := data & getMask(parameters.mantissaSize)

	fmt.Printf("Mask %d %X\n", data>>parameters.mantissaSize, getMask(parameters.exponentSize))

	fmt.Printf("%d: %d - %d\n", data, exponentVal, mantissaVal)

	mantissa := float64(0)
	for i := byte(1); i <= parameters.mantissaSize; i++ {
		if (mantissaVal>>(parameters.mantissaSize-i))&1 != 0 {
			mantissa = mantissa + math.Pow(2, -float64(i))
			fmt.Println(mantissa)
		}
	}

	fmt.Printf("mantissa : %f\n", mantissa)

	if exponentVal == 0 {
		exponentVal = 1
	} else {
		mantissa++
	}

	fmt.Printf("toto %d\n", int64(1<<(parameters.exponentSize-1))-1)

	exponent := int64(exponentVal) - (int64(1<<(parameters.exponentSize-1)) - 1)

	fmt.Printf("exponent : %d\n", exponent)

	val := mantissa * math.Pow(2, float64(exponent))
	if negative {
		return -val, nil
	}

	return val, nil
}
