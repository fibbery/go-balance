package balance

import (
	"fmt"
	"math/rand"
	"time"
)

type Shuffle2Balance struct {
}

func init() {
	RegisterBalance("shuffle2", &Shuffle2Balance{})
}

func (s *Shuffle2Balance) DoBalance(instances []*Instance, keys ...string) (inst *Instance, err error) {
	size := len(instances)
	if size == 0 {
		return nil, fmt.Errorf("no instance found")
	}
	//fisher-yates算法[修正后的洗牌算法]
	//主要思路为每次随机挑选一个值，放在数组末尾。然后在n-1个元素的数组中再随机挑选一个值，放在数组末尾，以此类推。
	// 可以用 instances[rand.Perm(size)[0]]替代
	rand.Seed(time.Now().UnixNano())
	for i:= size; i > 0; i--{
		lastIdx := i - 1
		idx := rand.Intn(i)
		instances[lastIdx],instances[idx] = instances[idx],instances[lastIdx]
	}
	return instances[0], nil
}
