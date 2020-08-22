package balance

import (
	"fmt"
	"math/rand"
)

type ShuffleBalance struct {
}

func init() {
	RegisterBalance("shuffle", &ShuffleBalance{})
}

func (s *ShuffleBalance) DoBalance(instances []*Instance, keys ...string) (inst *Instance, err error) {
	size := len(instances)
	if size == 0 {
		return nil, fmt.Errorf("no instance found")
	}
	return instances[rand.Perm(size)[0]], nil
}
