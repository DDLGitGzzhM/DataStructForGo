package heap

import "DataStructForGo/pkgs"

/*
1. 实现大根堆，小根堆
2. 构建堆
3. 弹出堆顶数据
*/

var (
	cnt int
	m   int
)

func init() {
	cnt = 0
	m = 0
}

type Heap struct {
	Node []node
	ph   []int
	hp   []int
}

func (h *Heap) NewHeap() *Heap {
	h.Node = make([]node, 10)
	h.hp = make([]int, len(h.Node))
	h.ph = make([]int, len(h.Node))
	return h
}

func (h *Heap) InitHeap(nodeSize int) *Heap {
	h.Node = make([]node, nodeSize)
	h.ph = make([]int, nodeSize)
	h.hp = make([]int, nodeSize)
	return h
}

type node struct {
	Val int
}

func (h *Heap) heap_swap(a, b int) {
	pkgs.Swap(&h.ph[h.hp[a]], &h.ph[h.hp[b]])
	pkgs.Swap(&h.hp[a], &h.hp[b])
	pkgs.Swap(&h.Node[a], &h.Node[b])
}

func (h *Heap) down(u int) {
	temp := u
	if u*2 <= cnt && h.Node[u<<1].Val < h.Node[temp].Val {
		temp = u << 1
	}

	if u*2+1 <= cnt && h.Node[u<<1+1].Val < h.Node[temp].Val {
		temp = u<<1 + 1
	}

	if u != temp {
		h.heap_swap(u, temp)
		h.down(temp)
	}
}

func (h *Heap) up(u int) {
	for u/2 != 0 && h.Node[u].Val < h.Node[u/2].Val {
		h.heap_swap(u, u/2)
		u >>= 1
	}
}

func (h *Heap) Insert(insert_val int) {
	cnt++
	m++
	h.ph[m] = cnt
	h.ph[cnt] = m
	h.Node[cnt].Val = insert_val
	h.up(cnt)
}

func (h *Heap) DeleteK(delPos int) {
	delPos = h.ph[delPos]
	h.heap_swap(delPos, cnt)
	cnt--
	h.up(delPos)
	h.down(delPos)
}

func (h *Heap) DelTop() {
	h.heap_swap(1, cnt)
	cnt--
	h.down(1)
}

func (h *Heap) Top() int {
	return h.Node[1].Val
}
