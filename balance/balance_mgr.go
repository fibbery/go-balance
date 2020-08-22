package balance

import (
	"fmt"
)

var mgr = &Manager{
	balances: make(map[string]Balance),
}

type Manager struct {
	balances map[string]Balance
}

func RegisterBalance(name string, balance Balance) {
	mgr.balances[name] = balance
}

func DoBalance(name string, instances []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.balances[name]
	if !ok {
		err = fmt.Errorf("%s Balance Not Found", name)
		return
	}
	inst, err = balance.DoBalance(instances)
	return
}
