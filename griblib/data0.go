package griblib

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

//Data0 is a Grid point data - simple packing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_temp5-0.shtml
//    | Octet Number | Content
//    -----------------------------------------------------------------------------------------
//    | 12-15	     | Reference value (R) (IEEE 32-bit floating-point value)
//    | 16-17	     | Binary scale factor (E)
//    | 18-19	     | Decimal scale factor (D)
//    | 20	         | Number of bits used for each packed value for simple packing, or for each
//    |              | group reference value for complex packing or spatial differencing
//    | 21           | Type of original field values
//    |              |    - 0 : Floating point
//    |              |    - 1 : Integer
//    |              |    - 2-191 : reserved
//    |              |    - 192-254 : reserved for Local Use
//    |              |    - 255 : missing
type Data0 struct {
	Reference    float32 `json:"reference"`
	BinaryScale  uint16  `json:"binaryScale"`
	DecimalScale uint16  `json:"decimalScale"`
	Bits         uint8   `json:"bits"`
	Type         uint8   `json:"type"`
}

// ParseData0 parses data0 struct from the reader into the an array of floating-point values
func ParseData0(dataReader io.Reader, dataLength int, template *Data0) ([]float64, error) {

	rawData := make([]byte, dataLength)
	dataReader.Read(rawData)

	bscale := math.Pow(2.0, float64(template.BinaryScale))
	dscale := math.Pow(10.0, float64(template.DecimalScale))

	buffer := bytes.NewBuffer(rawData)
	bitReader := newReader(buffer)

	dataSize := int(math.Ceil(
		float64(8*dataLength) / float64(template.Bits),
	))

	fld := make([]float64, dataSize)

	uintDataSlice, errRead := bitReader.readUintsBlock(int(template.Bits), dataSize)
	if errRead != nil {
		return fld, errRead
	}

	switch template.Type {
	case 0: // Float
		for index, uintValue := range uintDataSlice {
			val, err := GetFloat(uintValue, int(template.Bits))
			if err != nil {
				return fld, err
			}
			fld[index] = float64(val)
		}

	case 1: // Integer
		for index, uintValue := range uintDataSlice {
			signed := int64(uintValue)
			fld[index] = float64(signed)
		}
	case 255:
		return []float64{}, fmt.Errorf("Missing data type")
	default:
		return []float64{}, fmt.Errorf("Unsupported data type")
	}

	for i, dataValue := range fld {
		fld[i] = (float64(dataValue)*float64(bscale) + float64(template.Reference)) * float64(dscale)
	}

	return fld, nil
}
