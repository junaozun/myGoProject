package loadbalance

import (
	"math/rand"
	"mygorpc/selector"
	"time"
)
type randomBalancer struct {
}

// 随机
func newRandomBalancer() *randomBalancer {
	return &randomBalancer{}
}

func (r *randomBalancer) Balance(serviceName string, nodes []*selector.Node) *selector.Node {
	if len(nodes) == 0 {
		return nil
	}
	rand.Seed(time.Now().Unix())
	num := rand.Intn(len(nodes))
	return nodes[num]
}
