package main

import (
	"fmt"
	"github.com/fibbery/go-balance/balance"
	"math/rand"
	"strconv"
)

func main() {
	instances := make([]*balance.Instance, 0)
	for i := 0 ; i < 10; i++ {
		host := fmt.Sprintf("192.168.%s.%s",strconv.Itoa(rand.Intn(255)),strconv.Itoa(rand.Intn(rand.Intn(255))))
		port := rand.Intn(65536)
		instances = append(instances,balance.NewInstance(host, port))
	}
	fmt.Println(instances)

}