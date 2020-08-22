package balance

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomBalance struct {
}

func init() {
	RegisterBalance("random", &RandomBalance{})
}

func (r *RandomBalance) DoBalance(instances []*Instance, keys ...string) (inst *Instance, err error) {
	size := len(instances)
	if size == 0 {
		return nil, fmt.Errorf("no instance found")
	}
	rand.Seed(time.Now().UnixNano())
	return instances[rand.Intn(size)], nil
}
