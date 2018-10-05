package griblib

import (
	"fmt"
	"math"
)

type floatDefinition struct {
	hasSign      bool
	exponentSize byte
	mantissaSize byte
}

type floatParameters struct {
	negative bool
	exponent int64
	mantissa float64
}

var floatingtypes = map[int]floatDefinition{
	10: floatDefinition{
		hasSign:      false,
		exponentSize: 5,
		mantissaSize: 5,
	},
	11: floatDefinition{
		hasSign:      false,
		exponentSize: 5,
		mantissaSize: 6,
	},
	14: floatDefinition{
		hasSign:      false,
		exponentSize: 5,
		mantissaSize: 9,
	},
	16: floatDefinition{
		hasSign:      true,
		exponentSize: 5,
		mantissaSize: 10,
	},
	32: floatDefinition{
		hasSign:      true,
		exponentSize: 8,
		mantissaSize: 23,
	},
	64: floatDefinition{
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

func newFloatParameters(data uint64, bits int) (floatParameters, error) {
	definition, ok := floatingtypes[bits]
	if !ok {
		return floatParameters{}, fmt.Errorf("Unrecongnize floating bit size")
	}

	parameters := floatParameters{
		negative: definition.hasSign && ((data>>(definition.exponentSize+definition.mantissaSize))&1 == 1),
		exponent: 0,
		mantissa: float64(0),
	}

	exponentVal := (data >> definition.mantissaSize) & getMask(definition.exponentSize)
	mantissaVal := data & getMask(definition.mantissaSize)

	for i := byte(1); i <= definition.mantissaSize; i++ {
		if (mantissaVal>>(definition.mantissaSize-i))&1 != 0 {
			parameters.mantissa = parameters.mantissa + math.Pow(2, -float64(i))
		}
	}

	if exponentVal == 0 {
		exponentVal = 1
	} else {
		parameters.mantissa++
	}

	parameters.exponent = int64(exponentVal) - (int64(1<<(definition.exponentSize-1)) - 1)

	return parameters, nil
}

// GetFloat performs a binary cast
// https://www.h-schmidt.net/FloatConverter/IEEE754.html
func GetFloat(data uint64, nBits int) (float64, error) {

	parameters, err := newFloatParameters(data, nBits)
	if err != nil {
		return 0.0, err
	}

	val := parameters.mantissa * math.Pow(2, float64(parameters.exponent))
	if parameters.negative {
		return -val, nil
	}

	return val, nil
}
