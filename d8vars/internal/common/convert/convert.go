package d8convert

import (
	"fmt"
	"reflect"

	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	"github.com/driver8soft/examples/d8vars/model/request"
	"github.com/driver8soft/examples/d8vars/model/response"
)

func ParseCob(req *request.Request) (string, error) {
	// reflect struct request
	rValue := reflect.ValueOf(req).Elem()
	// create commarea from cobol copybook initialized
	commareaIn := make([]byte, len(d8conf.Cobol.Commarea.String()))
	copy(commareaIn[:], d8conf.Cobol.Commarea.String())

	// iterate over struct request fields
	for i := 0; i < rValue.NumField(); i++ {
		field := reflect.ValueOf(req).Elem().Type().Field(i).Name

		// get cobol copybook definition (config yaml file)
		copyItem, exists := d8conf.Cobol.Cvars[field]
		if !exists {
			return "", fmt.Errorf("field %s not found in Cobol config yaml file", field)
		}
		cType, cLen, cDec, cInit, cEnd := copyItem[0], copyItem[1], copyItem[2], copyItem[4], copyItem[5]
		cSign := copyItem[3] != 0

		// switch cobol type variable to convert
		switch cType {
		// display
		case 0:
			ndisplay, err := usage0(rValue.Field(i).String(), cLen, cDec, cSign)
			if err != nil {
				return "", fmt.Errorf("%s field: (%s)", err, field)
			}
			copy(commareaIn[cInit:cEnd], ndisplay)
		// comp1
		case 1:
			nfloat := usage1(rValue.Field(i).Interface().(float32))
			copy(commareaIn[cInit:cEnd], nfloat)
		// comp2
		case 2:
			nfloat := usage2(rValue.Field(i).Interface().(float64))
			copy(commareaIn[cInit:cEnd], nfloat)
		// comp3 packed decimal
		case 3:
			npacked, err := usage3(rValue.Field(i).String(), cLen, cDec, cSign)
			if err != nil {
				return "", fmt.Errorf("%s field: (%s)", err, field)
			}
			copy(commareaIn[cInit:cEnd], npacked)
		// comp binary comp4
		case 4:
			ncomp, err := usage4(rValue.Field(i).String(), cLen, cDec, cSign)
			if err != nil {
				return "", fmt.Errorf("%s field: (%s)", err, field)
			}
			copy(commareaIn[cInit:cEnd], ncomp)
		// comp5
		case 5:
			// if length is between 1 and 4 = short
			if cLen > 0 && cLen < 5 {
				if cSign {
					nshort := usage5Short(rValue.Field(i).Interface().(int16))
					copy(commareaIn[cInit:cEnd], nshort)
				} else {
					nushort := usage5Ushort(rValue.Field(i).Interface().(uint16))
					copy(commareaIn[cInit:cEnd], nushort)
				}
			}
			// if length is between 5 and 9 = long
			if cLen > 4 && cLen < 10 {
				if cSign {
					nlong := usage5Long(rValue.Field(i).Interface().(int32))
					copy(commareaIn[cInit:cEnd], nlong)
				} else {
					nulong := usage5Ulong(rValue.Field(i).Interface().(uint32))
					copy(commareaIn[cInit:cEnd], nulong)
				}
			}
			// if length is between 10 and 18 = double
			if cLen > 9 && cLen < 19 {
				if cSign {
					ndouble := usage5Double(rValue.Field(i).Interface().(int64))
					copy(commareaIn[cInit:cEnd], ndouble)
				} else {
					nudouble := usage5Udouble(rValue.Field(i).Interface().(uint64))
					copy(commareaIn[cInit:cEnd], nudouble)
				}
			}
			// if length is greater than 18 = error
			if cLen > 18 {
				return "", fmt.Errorf("variable %s max length exceed", field)

			}
		// char alphanumeric
		case 9:
			nchar := usage9(rValue.Field(i).String(), cLen)
			copy(commareaIn[cInit:cEnd], nchar)
		default:
			return "", fmt.Errorf("variable %s type %v not implemented", field, cType)
		}
	}
	return string(commareaIn), nil
}

func ParseGo(commareaOut string) (*response.Response, error) {
	res := response.Response{}

	// reflect struct response
	rValue := reflect.ValueOf(&res).Elem()
	// iterate over struct response fields
	for i := 0; i < rValue.NumField(); i++ {
		field := reflect.ValueOf(&res).Elem().Type().Field(i).Name
		// get cobol copybook definition (config yaml file)
		copyItem := d8conf.Cobol.Cvars[field]
		cType, cLen, cDec, cInit, cEnd := copyItem[0], copyItem[1], copyItem[2], copyItem[4], copyItem[5]
		cSign := copyItem[3] != 0
		// get cobol variable value from commarea
		fieldValue := commareaOut[cInit:cEnd]
		// switch cobol type variable to convert
		switch cType {
		// string
		case 0, 3, 9:
			goValue := usageStr(cType, cDec, fieldValue)
			reflect.ValueOf(&res).Elem().FieldByName(field).SetString(goValue)
		// float32
		case 1:
			goValue := usageFloat32(fieldValue)
			reflect.ValueOf(&res).Elem().FieldByName(field).SetFloat(float64(goValue))
		// float64
		case 2:
			goValue := usageFloat64(fieldValue)
			reflect.ValueOf(&res).Elem().FieldByName(field).SetFloat(goValue)
		// comp to string
		case 4:
			goValue := usageComp(cLen, cDec, cSign, fieldValue)
			reflect.ValueOf(&res).Elem().FieldByName(field).SetString(goValue)
		// int
		case 5:
			// if length is between 1 and 4 = int16
			if cLen > 0 && cLen < 5 {
				if cSign {
					goValue := usageInt16(fieldValue)
					reflect.ValueOf(&res).Elem().FieldByName(field).SetInt(int64(goValue))
				} else {
					goValue := usageUint16(fieldValue)
					reflect.ValueOf(&res).Elem().FieldByName(field).SetUint(uint64(goValue))
				}
			}
			// if length is between 5 and 9 = int32
			if cLen > 4 && cLen < 10 {
				if cSign {
					goValue := usageInt32(fieldValue)
					reflect.ValueOf(&res).Elem().FieldByName(field).SetInt(int64(goValue))
				} else {
					goValue := usageUint32(fieldValue)
					reflect.ValueOf(&res).Elem().FieldByName(field).SetUint(uint64(goValue))
				}
			}
			// if length is between 10 and 18 = int64
			if cLen > 9 && cLen < 19 {
				if cSign {
					goValue := usageInt64(fieldValue)
					reflect.ValueOf(&res).Elem().FieldByName(field).SetInt(goValue)
				} else {
					goValue := usageUint64(fieldValue)
					reflect.ValueOf(&res).Elem().FieldByName(field).SetUint(goValue)
				}
			}
			// if length is greater than 18 = error
			if cLen > 18 {
				return &res, fmt.Errorf("variable %s max length exceed", field)
			}
		default:
			return &res, fmt.Errorf("variable %s type %v not implemented", field, cType)
		}

	}
	return &res, nil
}
