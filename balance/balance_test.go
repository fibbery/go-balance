package balance

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
)

var instances []*Instance

func TestMain(m *testing.M) {
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%s.%s", strconv.Itoa(rand.Intn(255)), strconv.Itoa(rand.Intn(rand.Intn(255))))
		port := rand.Intn(65536)
		weight := rand.Intn(10)
		instances = append(instances, NewInstance(host, port, weight))

	}
	os.Exit(m.Run())
}

func TestRoundRobin(t *testing.T) {
	size := len(instances)
	inst, err := DoBalance("roundrobin", instances)
	if err != nil {
		t.Error("do balance error", err)
		return
	}
	for i := 0; i < size-1; i++ {
		DoBalance("roundrobin", instances)
	}
	result, err := DoBalance("roundrobin", instances)
	if err != nil {
		t.Error("do balance error next stage", err)
		return
	}
	if !reflect.DeepEqual(inst, result) {
		t.Errorf("expected %v, result is %v", inst, result)
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < len(instances); i++ {
		balance, err := DoBalance("random", instances)
		if err != nil {
			t.Log("do balance error", err)
		}
		t.Log("select instance " + balance.String())
	}
}

func TestShuffle2(t *testing.T) {
	for i := 0; i < len(instances); i++ {
		balance, err := DoBalance("shuffle2", instances)
		if err != nil {
			t.Log("do balance error", err)
		}
		t.Log("select instance " + balance.String())
	}
}

func TestShuffle(t *testing.T) {
	for i := 0; i < len(instances); i++ {
		balance, err := DoBalance("shuffle", instances)
		if err != nil {
			t.Log("do balance error", err)
		}
		t.Log("select instance " + balance.String())
	}
}

func TestHashconsistent(t *testing.T) {
	for i := 0; i < len(instances); i++ {
		balance, err := DoBalance("hashconsistent", instances)
		if err != nil {
			t.Log("do balance error", err)
		}
		t.Log("select instance " + balance.String())
	}
}

func TestWrr(t *testing.T) {
	counterMap := make(map[string]int)
	instantceMap := make(map[string]*Instance)

	weightSum := 0
	for _, instance := range instances {
		weightSum += instance.weight
		instantceMap[instance.String()] = instance
	}

	for i := 0; i < weightSum; i++ {
		inst, err := DoBalance("weight", instances)
		if err != nil {
			t.Error(err)
			continue
		}
		if _, ok := counterMap[inst.String()]; !ok {
			counterMap[inst.String()] = 0
		}
		t.Logf("select instance %v\n", inst)
		counterMap[inst.String()] ++
	}

	// equal
	for key, count := range counterMap {
		in := instantceMap[key]
		if in.weight != count {
			t.Errorf("Sum : %d, Count: %d, Instance: %s", weightSum, count, in)
		}
	}

}
