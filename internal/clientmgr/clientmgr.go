package clientmgr

import (
	"fmt"
	// import modbus client lib
)

type ClientDesc struct {
	Name   string
	Host   string
	Port   int
	UnitID int
}

func NewClientFromDesc(d ClientDesc) ( /*client*/ interface{}, error) {
	// ex : create TCP client to d.Host:d.Port and set unit id
	// return an object implementing Read/Write helper methods
	return nil, fmt.Errorf("not implemented")
}
