package d8convert

import (
	"bytes"
	"fmt"
	"reflect"

	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	"github.com/driver8soft/examples/d8vars/model/request"
)

func ParseCob(g *request.Request) (string, error) {

	rValue := reflect.ValueOf(g).Elem()
	rType := rValue.Type()

	var buf bytes.Buffer

	for i := 0; i < rValue.NumField(); i++ {
		for j := 0; j < len(d8conf.CobC.Copy); j++ {
			if rType.Field(i).Name != d8conf.CobC.Copy[j].Field {
				continue
			} else {
				switch d8conf.CobC.Copy[j].Type {
				case 0:
					ndisplay, err := usage0(rValue.Field(i).String(),
						d8conf.CobC.Copy[j].Length,
						d8conf.CobC.Copy[j].Decimal,
						d8conf.CobC.Copy[j].Sign)
					if err != nil {
						return "", fmt.Errorf("%s field: (%s)", err, d8conf.CobC.Copy[j].Field)
					}
					buf.Write([]byte(ndisplay))
				case 1:
					buf.Write(usage1(rValue.Field(i).Interface().(float32)))
				case 2:
					buf.Write(usage2(rValue.Field(i).Interface().(float64)))
				case 3:
					npacked, err := usage3(rValue.Field(i).String(),
						d8conf.CobC.Copy[j].Length,
						d8conf.CobC.Copy[j].Decimal,
						d8conf.CobC.Copy[j].Sign)
					if err != nil {
						return "", fmt.Errorf("%s field: (%s)", err, d8conf.CobC.Copy[j].Field)
					}
					buf.Write(npacked)
				case 4:
					ncomp, err := usage4(rValue.Field(i).String(),
						d8conf.CobC.Copy[j].Length,
						d8conf.CobC.Copy[j].Decimal,
						d8conf.CobC.Copy[j].Sign)
					if err != nil {
						return "", fmt.Errorf("%s field: (%s)", err, d8conf.CobC.Copy[j].Field)
					}
					buf.Write(ncomp)
				case 5:
					if d8conf.CobC.Copy[j].Length > 0 && d8conf.CobC.Copy[j].Length < 5 {
						if d8conf.CobC.Copy[j].Sign {
							buf.Write(usage5Short(rValue.Field(i).Interface().(int16)))
						} else {
							buf.Write(usage5Ushort(rValue.Field(i).Interface().(uint16)))
						}
					}
					if d8conf.CobC.Copy[j].Length > 4 && d8conf.CobC.Copy[j].Length < 10 {
						if d8conf.CobC.Copy[j].Sign {
							buf.Write(usage5Long(rValue.Field(i).Interface().(int32)))
						} else {
							buf.Write(usage5Ulong(rValue.Field(i).Interface().(uint32)))
						}
					}
					if d8conf.CobC.Copy[j].Length > 9 && d8conf.CobC.Copy[j].Length < 19 {
						if d8conf.CobC.Copy[j].Sign {
							buf.Write(usage5Double(rValue.Field(i).Interface().(int64)))
						} else {
							buf.Write(usage5Udouble(rValue.Field(i).Interface().(uint64)))
						}
					}
					if d8conf.CobC.Copy[j].Length > 18 {
						return "", fmt.Errorf("d8convert: variable %s max length exceed", d8conf.CobC.Copy[j].Field)

					}
				case 9:
					buf.Write(usage9(rValue.Field(i).String(), d8conf.CobC.Copy[j].Length))
				default:
					return "", fmt.Errorf("d8convert: variable %s type %v not implemented", d8conf.CobC.Copy[j].Field, d8conf.CobC.Copy[j].Type)
				}
			}
		}
	}
	return buf.String(), nil
}
