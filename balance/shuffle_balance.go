package balance

import (
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())

	//shuffle
	for i := 0; i < size/2; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		instances[a], instances[b] = instances[b], instances[a]
	}
	return instances[rand.Perm(size)[0]], nil
}
