package main

import (
	"fmt"
	"log"

	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	d8service "github.com/driver8soft/examples/d8vars/internal/service"
	d8test "github.com/driver8soft/examples/d8vars/test"
)

func main() {
	// load env variables & copybook definition
	d8conf.InitConfig()

	// run COBOL program
	if res, err := d8service.NewService(d8test.New()); err != nil {
		log.Println(err)
	} else {
		fmt.Println(res)
	}
}
