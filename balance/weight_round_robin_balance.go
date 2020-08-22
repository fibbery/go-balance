package balance

import "fmt"

type WeightRandRobinBalance struct {
	index  int
	weight int
}

func init() {
	RegisterBalance("weight", &WeightRandRobinBalance{index: -1})
}

func (w *WeightRandRobinBalance) DoBalance(instances []*Instance, key ...string) (inst *Instance, err error) {
	size := len(instances)
	if size == 0 {
		return nil, fmt.Errorf("no instance found")
	}

	// prepare
	var weights []int
	for _, inst := range instances {
		weights = append(weights, inst.weight)
	}
	g := gcdAll(weights)
	maxWeight := getMaxWeight(weights)

	// roundrobin
	for {
		w.index = (w.index + 1) % size
		if w.index == 0 {
			w.weight = w.weight - g
			if w.weight <= 0 {
				w.weight = maxWeight
				if w.weight == 0 {
					return new(Instance), nil
				}
			}
		}

		instance := instances[w.index]
		if instance.weight >= w.weight {
			return instance, nil
		}
	}

}

func getMaxWeight(weights []int) int {
	max := 0
	for _, value := range weights {
		if value > max {
			max = value
		}
	}
	return max
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdAll(all []int) int {
	g := all[0]
	for i := 1; i < len(all); i++ {
		g = gcd(g, all[i])
	}
	return g
}
