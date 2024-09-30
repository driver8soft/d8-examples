package d8convert

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var negDecimals = [10]byte{0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79}

func formatStr(s string, cDec int) string {
	v, _ := strconv.ParseFloat(s, 64)
	v = v / math.Pow10(cDec)
	return strconv.FormatFloat(v, 'f', -1, 64)
}
func checkLimits(n int64, cLen, cDec int) error {
	limit := int64(math.Pow10(cLen - cDec))
	if n > (limit-1) || n < (limit*-1)+1 {
		return fmt.Errorf("exceed max. lenght in picture:value (%v)", n)
	}
	return nil
}
func justifyDisplay(n int64, cLen int) string {
	var count int
	if n < 0 {
		n = -n
	}
	ndisplay := strconv.FormatInt(n, 10)
	if count = cLen - len(ndisplay); count <= 0 {
		return ndisplay
	}
	var builder strings.Builder
	builder.Grow(cLen)
	for i := 0; i < count; i++ {
		builder.WriteByte('0')
	}
	builder.WriteString(ndisplay)
	return builder.String()
}
func formatDisplay(s string, cLen, cDec int, cSign bool) (nAlign int64, err error) {
	var n float64
	// check empty string
	if s == "" {
		s = "0"
	}
	// Check format number
	if n, err = strconv.ParseFloat(s, 64); err != nil {
		return 0, fmt.Errorf("usage display: %s", err)
	}
	// Check sign mask
	if !cSign && n < 0 {
		return 0, fmt.Errorf("negative value in unsigned variable: value(%v)", n)
	}
	// Check COBOL mask limits
	if err := checkLimits(int64(n), cLen, cDec); err != nil {
		return 0, err
	}
	n = n * math.Pow10(cDec)
	n = math.Round(n*100) / 100
	return int64(n), nil
}

func usage0(s string, cLen, cDec int, cSign bool) (string, error) {
	// Align number to picture clause
	nAlign, err := formatDisplay(s, cLen, cDec, cSign)
	if err != nil {
		return "", err
	}
	// justify number to picture length ... add zeroes
	ndisplay := justifyDisplay(nAlign, cLen)

	if nAlign >= 0 {
		return ndisplay, nil
	} else {
		i, _ := strconv.Atoi(ndisplay[len(ndisplay)-1:])
		return ndisplay[:len(ndisplay)-1] + string(negDecimals[i]), nil
	}
}
func usage1(f float32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, math.Float32bits(f))
	return buf
}
func usage2(f float64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(f))
	return buf
}
func usage3(s string, cLen, cDec int, cSign bool) ([]byte, error) {
	// align number to picture
	nAlign, err := formatDisplay(s, cLen, cDec, cSign)
	if err != nil {
		return nil, err
	}
	// justify number to picture length ... add zeroes
	nPacked := justifyDisplay(nAlign, cLen)

	var builder strings.Builder
	builder.Grow(cLen)

	// odd lenght add zero
	if len(nPacked)%2 == 0 {
		builder.WriteByte('0')
	}
	builder.WriteString(nPacked)

	// add sign +/- to last byte
	if nAlign >= 0 {
		builder.WriteByte('C')
	} else {
		builder.WriteByte('D')
	}
	result, _ := hex.DecodeString(builder.String())

	return result, nil
}
func usage4(s string, cLen, cDec int, cSign bool) ([]byte, error) {
	// align number to picture
	nAlign, err := formatDisplay(s, cLen, cDec, cSign)
	if err != nil {
		return []byte{}, err
	}
	if cSign {
		if cLen > 0 && cLen < 5 {
			return usage5Short(int16(nAlign)), nil
		}
		if cLen > 4 && cLen < 10 {
			return usage5Long(int32(nAlign)), nil
		}
		if cLen > 9 && cLen < 18 {
			return usage5Double(int64(nAlign)), nil
		}
	} else {
		if cLen > 0 && cLen < 5 {
			return usage5Ushort(uint16(nAlign)), nil
		}
		if cLen > 4 && cLen < 10 {
			return usage5Ulong(uint32(nAlign)), nil
		}
		if cLen > 9 && cLen < 18 {
			return usage5Udouble(uint64(nAlign)), nil
		}
	}
	return []byte{}, fmt.Errorf("invalid length: %v", cLen)
}
func usage5Udouble(f uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, f)
	return buf
}
func usage5Double(f int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(f))
	return buf
}
func usage5Ulong(f uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, f)
	return buf
}
func usage5Long(f int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(f))
	return buf
}
func usage5Short(f int16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(f))
	return buf
}
func usage5Ushort(f uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, f)
	return buf
}
func usage9(s string, cLen int) string {
	var count int
	if count = cLen - len(s); count <= 0 {
		return s
	}
	var builder strings.Builder
	builder.Grow(cLen)
	builder.WriteString(s)
	for i := 0; i < count; i++ {
		builder.WriteByte(' ')
	}
	return builder.String()
}
func usageStr(cType, cDec int, fieldValue string) string {
	switch cType {
	case 0:
		lastByte := fieldValue[len(fieldValue)-1:]
		for i, j := range negDecimals {
			if lastByte == string(j) {
				fieldValue = "-" + fieldValue[:len(fieldValue)-1] + strconv.Itoa(i)
			}
		}
		return formatStr(fieldValue, cDec)
	case 3:
		s := hex.EncodeToString([]byte(fieldValue))
		if s[len(s)-1:] == "c" {
			// Positive number
			s = s[0 : len(s)-1]
		} else {
			// Negative number
			s = "-" + s[0:len(s)-1]
		}
		return formatStr(s, cDec)
	case 9:
		return strings.TrimRight(fieldValue, " ")
	default:
		panic(1)
	}
}
func usageComp(cLen, cDec int, cSign bool, cobValue string) (s string) {
	if cLen > 0 && cLen < 5 {
		if cSign {
			s = strconv.Itoa(int(usageInt16(cobValue)))
		} else {
			s = strconv.Itoa(int(usageInt16(cobValue)))
		}
	}
	if cLen > 4 && cLen < 10 {
		if cSign {
			s = strconv.Itoa(int(usageInt32(cobValue)))
		} else {
			s = strconv.Itoa(int(usageUint32(cobValue)))
		}
	}
	if cLen > 9 && cLen < 19 {
		if cSign {
			s = strconv.Itoa(int(usageInt64(cobValue)))
		} else {
			s = strconv.Itoa(int(usageUint64(cobValue)))
		}
	}
	if cLen > 18 {
		fmt.Println("variable max length exceed")
		return ""
	}
	return formatStr(s, cDec)
}

func usageFloat32(cobValue string) float32 {
	bits := binary.LittleEndian.Uint32([]byte(cobValue))
	return math.Float32frombits(bits)
}
func usageFloat64(cobValue string) float64 {
	bits := binary.LittleEndian.Uint64([]byte(cobValue))
	return math.Float64frombits(bits)
}
func usageUint16(cobValue string) uint16 {
	return binary.LittleEndian.Uint16([]byte(cobValue))
}
func usageUint32(cobValue string) uint32 {
	return binary.LittleEndian.Uint32([]byte(cobValue))
}
func usageUint64(cobValue string) uint64 {
	return binary.LittleEndian.Uint64([]byte(cobValue))
}
func usageInt16(cobValue string) int16 {
	bits := binary.LittleEndian.Uint16([]byte(cobValue))
	return int16(bits)
}
func usageInt32(cobValue string) int32 {
	bits := binary.LittleEndian.Uint32([]byte(cobValue))
	return int32(bits)
}
func usageInt64(cobValue string) int64 {
	bits := binary.LittleEndian.Uint64([]byte(cobValue))
	return int64(bits)
}
