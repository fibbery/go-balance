package balance

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
)

type HashConsistentBalance struct {
	table *crc32.Table
}

func init() {
	RegisterBalance("hashconsistent", &HashConsistentBalance{
		table: crc32.MakeTable(crc32.IEEE),
	})
}

func (h *HashConsistentBalance) DoBalance(instances []*Instance, keys ...string) (inst *Instance, err error) {
	size := len(instances)
	if size == 0 {
		err = fmt.Errorf("no instances")
		return
	}
	hashKey := strconv.Itoa(rand.Int())
	if len(keys) > 0 {
		hashKey = keys[0]
	}
	hashcode := crc32.Checksum([]byte(hashKey), h.table)
	index := int(hashcode) % size
	return instances[index], nil
}
