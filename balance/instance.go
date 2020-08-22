package balance

import (
	"fmt"
)

type Instance struct {
	host   string
	port   int
	weight int
}

func NewInstance(host string, port int, weight int) *Instance {
	return &Instance{host: host, port: port, weight: weight}
}

func (inst *Instance) Port() int {
	return inst.port
}

func (inst *Instance) Host() string {
	return inst.host
}

func (inst *Instance) String() string {
	return fmt.Sprintf("%s:%d", inst.host, inst.port)
}
