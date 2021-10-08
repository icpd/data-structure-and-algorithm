package algorithm

type WeightNode struct {
	addr            string
	weight          int
	currentWeight   int
	effectiveWeight int
}

type WeightRoundRobinBalance struct {
	curIndex int
	nodes    []*WeightNode
}

func (w *WeightRoundRobinBalance) Add(add string, weight int) {
	w.nodes = append(w.nodes, &WeightNode{
		addr:            add,
		weight:          weight,
		effectiveWeight: weight,
	})
}

func (w *WeightRoundRobinBalance) next() string {
	total := 0
	var node *WeightNode
	for _, n := range w.nodes {
		// 所有节点有效权重和
		total += n.effectiveWeight

		// 更新当前节点临时权重（当前节点临时权重+当前节点有效权重）
		n.currentWeight += n.effectiveWeight

		// 选择最大临时权重节点（权重一样的选第一个）
		if node == nil || n.currentWeight > node.currentWeight {
			node = n
		}
	}

	if node == nil {
		return ""
	}

	// 更新该节点临时权重（该节点临时权重-所有节点有效权重和）
	node.currentWeight -= total
	return node.addr
}
