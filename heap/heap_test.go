package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Heap(t *testing.T) {
	heap := &Heap{}
	heap = heap.InitHeap(10)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(5)
	res := heap.Top()
	assert.Equal(t, res, 2)
	heap.DelTop()
	res = heap.Top()
	assert.Equal(t, res, 3)

	heap.DeleteK(1)
	res = heap.Top()
	assert.Equal(t, res, 4)
}
