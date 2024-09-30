package main

import (
	"testing"

	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	d8service "github.com/driver8soft/examples/d8vars/internal/service"
	"github.com/driver8soft/examples/d8vars/model/request"
	d8test "github.com/driver8soft/examples/d8vars/test"

	"github.com/google/go-cmp/cmp"
)

func TestVars(t *testing.T) {
	// load env variables & copybook definition
	d8conf.InitConfig()

	v := request.Request{Char: "hello",
		Comp2: 3.14, Comp1: 3.141618,
		Comp5Double: -123456789, Comp5Long: 2147483647, Comp5Short: -32768,
		Comp5Udouble: 123456789, Comp5Ulong: 3147483647, Comp5Ushort: 65535,
		NumDisplay: "123.34", Comp3: "123.4",
		CompDouble: "123.459", CompLong: "123.45", CompShort: "12.34"}

	_, err := d8service.NewService(&v)
	if err != nil {
		t.Fatalf(`Error:  %v`, err)
	}
}

func TestEmpty(t *testing.T) {
	// load env variables & copybook definition
	d8conf.InitConfig()

	_, err := d8service.NewService(&request.Request{})
	if err != nil {
		t.Fatalf(`Error:  %v`, err)
	}
}

func TestRand(t *testing.T) {
	// load env variables & copybook definition
	d8conf.InitConfig()

	// create random struct
	req := d8test.New()
	// call cobol program
	res, err := d8service.NewService(req)
	if err != nil {
		t.Fatalf(`Error:  %v`, err)
	}
	// compare results
	r1 := request.Request(*res)
	if !cmp.Equal(req, &r1) {
		t.Fatalf(`Not equal:  %v`, cmp.Diff(req, &r1))
	}
}
func BenchmarkRand(b *testing.B) {
	d8conf.InitConfig()

	for i := 0; i < b.N; i++ {
		req := d8test.New()
		res, _ := d8service.NewService(req)
		r1 := request.Request(*res)
		if !cmp.Equal(req, &r1) {
			b.Fatalf(`Not equal:  %v`, cmp.Diff(req, &r1))
		}

	}

}
