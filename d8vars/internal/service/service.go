package d8service

import (
	d8cgo "github.com/driver8soft/examples/d8vars/internal/cgocobol"
	d8conf "github.com/driver8soft/examples/d8vars/internal/common/config"
	d8conv "github.com/driver8soft/examples/d8vars/internal/common/convert"
	"github.com/driver8soft/examples/d8vars/model/request"
	"github.com/driver8soft/examples/d8vars/model/response"
)

func NewService(req *request.Request) (res *response.Response, err error) {
	// convert go variables into cobol copybook
	commareaIn, err := d8conv.ParseCob(req)
	if err != nil {
		return &response.Response{}, err
	}

	// call cobol program
	commareaOut, err := d8cgo.CobolCall(d8conf.Env.CobolProgram, commareaIn)
	if err != nil {
		return &response.Response{}, err
	}

	// convert cobol variables into go response struct
	res, err = d8conv.ParseGo(commareaOut)
	if err != nil {
		return &response.Response{}, err
	}

	return res, nil
}
