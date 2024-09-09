package d8convert

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func checkLimits(n int64, cLen, cDec int) error {
	limit := int64(math.Pow10(cLen - cDec))
	if n > (limit-1) || n < (limit*-1)+1 {
		return fmt.Errorf("d8convert: exceed max. lenght in picture:value (%v)", n)
	}
	return nil
}
func justifyDisplay(n int64, l int) (ndisplay string) {
	if n < 0 {
		n *= -1
	}
	ndisplay = strconv.FormatInt(n, 10)
	for i := l - len(ndisplay); i > 0; i-- {
		ndisplay = "0" + ndisplay
	}
	return
}
func formatDisplay(s string, cLen, cDec int, cSign bool) (int64, error) {
	// check empty string
	if s == "" {
		s = "0"
	}
	// Check format number
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("d8convert: usage display: %s", err)
	}
	// Check sign mask
	if !cSign && n < 0 {
		return 0, fmt.Errorf("d8convert: negative value in unsigned variable: value(%v)", n)
	}
	// Check COBOL mask limits
	if err := checkLimits(int64(n), cLen, cDec); err != nil {
		return 0, err
	}
	// Align number to decimal mask
	return int64(n * math.Pow10(cDec)), nil
}

func usage0(s string, cLen, cDec int, cSign bool) (string, error) {
	var negDecimals = [10]byte{0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79}

	// Align number to picture
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
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(f))
	return buf[:]
}

func usage2(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
func usage3(s string, cLen, cDec int, cSign bool) ([]byte, error) {
	// align number to picture
	nAlign, err := formatDisplay(s, cLen, cDec, cSign)
	if err != nil {
		return []byte{}, err
	}
	sl := strings.Split(strconv.FormatInt(nAlign, 10), "")
	// add sign +/- to last element in list
	if nAlign >= 0 {
		sl = append(sl, "c")
	} else {
		sl = append(sl[1:], "d")
	}
	// odd lenght add zero to list
	if len(sl)%2 != 0 {
		sl = append(append(sl, "0"), sl...)[len(sl):]
	}
	// justify number to picture length ... add zeroes
	x := ((cLen / 2) + 1) - (len(sl) / 2)
	for i := x * 2; i != 0; i-- {
		sl = append(append(sl, "0"), sl...)[len(sl):]
	}
	pdisplay, err := hex.DecodeString(strings.Join(sl, ""))
	if err != nil {
		return []byte{}, err
	}
	return pdisplay, nil
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
	return []byte{}, fmt.Errorf("d8convert: invalid length: %v", cLen)
}
func usage5Udouble(f uint64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], f)
	return buf[:]
}
func usage5Double(f int64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(f))
	return buf[:]
}
func usage5Ulong(f uint32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], f)
	return buf[:]
}
func usage5Long(f int32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(f))
	return buf[:]
}
func usage5Short(f int16) []byte {
	var buf [2]byte
	binary.LittleEndian.PutUint16(buf[:], uint16(f))
	return buf[:]
}
func usage5Ushort(f uint16) []byte {
	var buf [2]byte
	binary.LittleEndian.PutUint16(buf[:], f)
	return buf[:]
}
func usage9(s string, l int) []byte {
	var buf []byte
	for i := l - len(s); i > 0; i-- {
		s += " "
	}
	buf = []byte(s)
	return buf[0:l]
}
