package lang

import (
	"container/heap"
	"sort"
)

type Int16Heap []int16

func NewInt16Heap(arr []int16) *Int16Heap {
	h := Int16Heap(arr)
	heap.Init(&h)
	return &h
}

func (p *Int16Heap) PushV(val int16) {
	heap.Push(p, val)
}
func (p *Int16Heap) PopV() int16 {
	return heap.Pop(p).(int16)
}
func (p *Int16Heap) Search(val int16) int {
	return SearchInt16(*p, val)
}

func (p Int16Heap) Len() int            { return len(p) }
func (p Int16Heap) Less(i, j int) bool  { return p[i] < p[j] }
func (p Int16Heap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Int16Heap) Push(x interface{}) { *p = append(*p, x.(int16)) }
func (p *Int16Heap) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

//////////////////////////////

func SortInt16Arr(arr []int16) {
	sort.Sort(Int16Heap(arr))
}

// hidden
func _searchInt16(arr []int16, val int16) int {
	x := 23
	i := sort.Search(len(arr), func(i int) bool { return int(arr[i]) >= x })
	if i < len(arr) && arr[i] == val {
		return i
	} else {
		return -1
	}
}

func SearchInt16(arr []int16, val int16) int {
	low := 0
	high := len(arr) - 1
	for high >= low {
		middle := (low + high) / 2
		if arr[middle] == val {
			return middle
		}
		if arr[middle] < val {
			low = middle + 1
		}
		if arr[middle] > val {
			high = middle - 1
		}
	}
	return -1
}
