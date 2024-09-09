package d8test

import (
	"fmt"
	"log"
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
	// Set of characters used to generate random strings
	Chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func New() *request.Request {

	// load env variables & copybook definition
	if err := d8conf.InitConfig(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("loading variables Env:", d8conf.EnvC.AppEnv)
	}

	r1 = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	p := new(request.Request)

	s := reflect.ValueOf(p).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		if typeOfT.Field(i).Name != d8conf.CobC.Copy[i].Field {
			panic(1)
		}
		switch typeOfT.Field(i).Type.String() {
		case "string":
			switch d8conf.CobC.Copy[i].Type {
			case 0, 3, 4:
				s.Field(i).SetString(randDisplay(d8conf.CobC.Copy[i].Length,
					d8conf.CobC.Copy[i].Decimal))
			case 9:
				s.Field(i).SetString(randString(d8conf.CobC.Copy[i].Length))
			default:
			}
		case "float64":
			s.Field(i).SetFloat(randFloat64())
		case "float32":
			s.Field(i).SetFloat(float64(randFloat32()))
		case "int64":
			s.Field(i).SetInt(randInt64())
		case "int32":
			s.Field(i).SetInt(int64(randInt32()))
		case "int16":
			s.Field(i).SetInt(int64(randInt16()))
		case "uint64":
			s.Field(i).SetUint(randUint64())
		case "uint32":
			s.Field(i).SetUint(uint64(randUint32()))
		case "uint16":
			s.Field(i).SetUint(uint64(randUint16()))
		default:
			fmt.Println("ERROR")
		}
	}

	return p

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
func randString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = Chars[r1.Intn(len(Chars))]
	}
	return string(b)
}
func randDisplay(l, d int) string {
	max := int(math.Pow10(l))
	result := float64(r1.Intn(max*2)-max) / math.Pow10(d)
	return strconv.FormatFloat(result, 'f', -1, 64)

}
