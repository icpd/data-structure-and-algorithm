package algorithm

import (
	"testing"
)

func TestWeightRoundRobinBalance(t *testing.T) {
	w := new(WeightRoundRobinBalance)
	w.Add("127.0.0.1", 10)
	w.Add("127.0.0.2", 20)
	w.Add("127.0.0.3", 1)

	for i := 0; i < 30; i++ {
		t.Log(w.next())
	}
}
