package selector

import (
	"sync"
	"time"
)

type weightedRoundRobinBalancer struct {
	pickers  *sync.Map
	duration time.Duration
}

func newWeightedRoundRobinBalancer() *weightedRoundRobinBalancer {
	return &weightedRoundRobinBalancer{
		pickers:  new(sync.Map),
		duration: 3 * time.Minute,
	}
}

type weightedNode struct {
	node            *Node
	weight          int
	effectiveWeight int
	currentWeight   int
}

type wRoundRobinPicker struct {
	nodes          []*weightedNode
	lastUpdateTime time.Time
	duration       time.Duration
}

func (wr *wRoundRobinPicker) pick(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	// update picker after timeout
	if time.Now().Sub(wr.lastUpdateTime) > wr.duration ||
		len(nodes) != len(wr.nodes) {
		wr.nodes = getWeightedNode(nodes)
		wr.lastUpdateTime = time.Now()
	}

	totalWeight := 0
	maxWeight := 0
	index := 0
	for i, node := range wr.nodes {
		node.currentWeight += node.weight
		totalWeight += node.weight
		if node.currentWeight > maxWeight {
			maxWeight = node.currentWeight
			index = i
		}
	}

	wr.nodes[index].currentWeight -= totalWeight

	return wr.nodes[index].node

}

func (w *weightedRoundRobinBalancer) Balance(serviceName string, nodes []*Node) *Node {
	var picker *wRoundRobinPicker
	if len(nodes) == 0 {
		return nil
	}
	if p, ok := w.pickers.Load(serviceName); !ok {
		picker = &wRoundRobinPicker{
			nodes:          getWeightedNode(nodes),
			lastUpdateTime: time.Time{},
			duration:       0,
		}
	} else {
		picker = p.(*wRoundRobinPicker)
	}
	node := picker.pick(nodes)
	w.pickers.Store(serviceName, picker)
	return node
}

func getWeightedNode(nodes []*Node) []*weightedNode {
	var wgs []*weightedNode
	for _, node := range nodes {
		wgs = append(wgs, &weightedNode{
			node:            node,
			weight:          node.weight,
			currentWeight:   node.weight,
			effectiveWeight: node.weight,
		})
	}
	return wgs
}
