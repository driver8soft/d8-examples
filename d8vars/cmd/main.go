package main

import (
	"log"

	d8service "github.com/driver8soft/examples/d8vars/internal/service"
	d8test "github.com/driver8soft/examples/d8vars/test"
)

func main() {

	if err := d8service.NewService(d8test.New()); err != nil {
		log.Println(err)
	}

}
