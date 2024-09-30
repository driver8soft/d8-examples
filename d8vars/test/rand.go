package d8test

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"

	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	"github.com/driver8soft/examples/d8vars/model/request"

	"golang.org/x/exp/rand"
)

var r1 *rand.Rand

var (
	Chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func New() *request.Request {
	r1 = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	req := new(request.Request)

	rValue := reflect.ValueOf(req).Elem()
	rType := rValue.Type()

	for i := 0; i < rValue.NumField(); i++ {
		copyItem, exists := d8conf.Cobol.Cvars[rType.Field(i).Name]
		if !exists {
			panic(1)
		} else {
			switch rType.Field(i).Type.String() {
			case "string":
				switch copyItem[0] {
				case 0, 3, 4:
					rValue.Field(i).SetString(randDisplay(copyItem[1], copyItem[2]))
				case 9:
					rValue.Field(i).SetString(randString(copyItem[1]))
				default:
				}
			case "float64":
				rValue.Field(i).SetFloat(randFloat64())
			case "float32":
				rValue.Field(i).SetFloat(float64(randFloat32()))
			case "int64":
				rValue.Field(i).SetInt(randInt64())
			case "int32":
				rValue.Field(i).SetInt(int64(randInt32()))
			case "int16":
				rValue.Field(i).SetInt(int64(randInt16()))
			case "uint64":
				rValue.Field(i).SetUint(randUint64())
			case "uint32":
				rValue.Field(i).SetUint(uint64(randUint32()))
			case "uint16":
				rValue.Field(i).SetUint(uint64(randUint16()))
			default:
				fmt.Println("ERROR")
			}
		}
	}
	return req
}
func randFloat32() float32 {
	return r1.Float32()
}
func randFloat64() float64 {
	return r1.Float64()
}
func randInt16() int16 {
	max := int(math.Pow(2, 15))
	return int16(r1.Intn(max*2) - max)
}
func randInt32() int32 {
	max := int(math.Pow(2, 31))
	return int32(r1.Intn(max*2) - max)
}
func randInt64() int64 {
	return r1.Int63()
}
func randUint16() uint16 {
	return uint16(r1.Intn(65535))
}
func randUint32() uint32 {
	return r1.Uint32()
}
func randUint64() uint64 {
	return r1.Uint64()
}
func randString(cLen int) string {
	b := make([]rune, cLen)
	for i := range b {
		b[i] = Chars[r1.Intn(len(Chars))]
	}
	return string(b)
}
func randDisplay(cLen, cDec int) string {
	max := int(math.Pow10(cLen)) - 1
	result := float64(r1.Intn(max*2)-max) / (math.Pow10(cDec))
	return strconv.FormatFloat(result, 'f', -1, 64)
}
