package balance

import "fmt"

type RoundRobinBalance struct {
	currentIndx int
}

func init() {
	RegisterBalance("roundrobin", &RoundRobinBalance{})
}

func (r *RoundRobinBalance) DoBalance(instances []*Instance, keys ...string) (inst *Instance, err error) {
	size := len(instances)
	if size == 0 {
		return nil, fmt.Errorf("no instance found")
	}
	if r.currentIndx >= size {
		r.currentIndx = 0
	}
	inst = instances[r.currentIndx]
	r.currentIndx = (r.currentIndx + 1) % size
	return
}
