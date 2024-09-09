package main

import (
	"testing"

	d8service "github.com/driver8soft/examples/d8vars/internal/service"
	"github.com/driver8soft/examples/d8vars/model/request"
	d8test "github.com/driver8soft/examples/d8vars/test"
)

// sample struct
func TestVars(t *testing.T) {

	v := request.Request{Char: "hello",
		Comp2: 3.14, Comp1: 3.141618,
		Comp5Double: -123456789, Comp5Long: 2147483647, Comp5Short: -32768,
		Comp5Udouble: 123456789, Comp5Ulong: 3147483647, Comp5Ushort: 65535,
		NumDisplay: "123.34", Comp3: "123.4",
		CompDouble: "123.459", CompLong: "123.45", CompShort: "12.34"}

	err := d8service.NewService(&v)

	if err != nil {
		t.Fatalf(`Error:  %v`, err)
	}
}

// empty struct
func TestEmpty(t *testing.T) {

	v := request.Request{}

	err := d8service.NewService(&v)

	if err != nil {
		t.Fatalf(`Error:  %v`, err)
	}
}

// create random struct
func TestRand(t *testing.T) {

	p := d8test.New()

	err := d8service.NewService(p)

	if err != nil {
		t.Fatalf(`Error:  %v`, err)
	}
}
