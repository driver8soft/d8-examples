package d8service

import (
	"log"

	d8cgo "github.com/driver8soft/examples/d8vars/internal/cgocobol"
	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	d8convert "github.com/driver8soft/examples/d8vars/internal/common/convert"
	"github.com/driver8soft/examples/d8vars/model/request"
)

func NewService(req *request.Request) error {

	// load env variables & copybook definition
	if err := d8conf.InitConfig(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("loading variables Env:", d8conf.EnvC.AppEnv)
	}

	// convert go variables into cobol copybook
	commarea, err := d8convert.ParseCob(req)
	if err != nil {
		log.Fatal(err)
	}

	// call cobol program
	err = d8cgo.CobolCall(d8conf.EnvC.CobolProgram, commarea)
	if err != nil {
		log.Println(err)
	}
	return nil
}
