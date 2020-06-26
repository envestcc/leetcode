package gol

import "fmt"

type Heap struct {
	elems  []int
	length int
}

func InitHeap(size int) *Heap {
	return &Heap{
		elems:  make([]int, size),
		length: 0,
	}
}

func (h Heap) Print() {
	fmt.Printf("\nHeap:")
	for i := 0; i < h.length; i++ {
		fmt.Printf(" %+v", h.elems[i])
	}
	fmt.Printf("\n")
}

func (h *Heap) Add(a int) {
	if h.length >= len(h.elems) {
		if h.elems[0] >= a {
			return
		}
		h.elems[0] = a
		h.Sink(0)
		return
	}
	h.elems[h.length] = a
	h.length++
	h.Float(h.length - 1)
}

func (h *Heap) Float(id int) {
	// fmt.Printf("Float: index=%+v, value=%+v", id, h.elems[id])
	// h.Print()
	for {
		// fmt.Printf("Float:Loop id=%+v\n", id)
		if id <= 0 {
			return
		}
		parent := (id - 1) / 2
		if h.elems[parent] > h.elems[id] {
			h.elems[parent], h.elems[id] = h.elems[id], h.elems[parent]
		} else {
			return
		}
		id = parent
	}
}

func (h *Heap) Sink(id int) {
	// fmt.Printf("Sink: index=%+v, value=%+v, ", id, h.elems[id])
	// h.Print()
	for {
		// fmt.Printf("Sink:Loop id=%+v\n", id)
		lChild := 2*id + 1
		rChild := 2*id + 2
		if lChild < h.length && rChild < h.length {
			if h.elems[lChild] < h.elems[rChild] {
				if h.elems[id] < h.elems[lChild] {
					return
				}
				if h.elems[id] > h.elems[lChild] {
					h.elems[id], h.elems[lChild] = h.elems[lChild], h.elems[id]
					id = lChild
				} else {
					return
				}
			} else {
				if h.elems[id] < h.elems[rChild] {
					return
				}
				if h.elems[id] > h.elems[rChild] {
					h.elems[id], h.elems[rChild] = h.elems[rChild], h.elems[id]
					id = rChild
				} else {
					return
				}
			}
		} else if lChild < h.length {
			if h.elems[id] < h.elems[lChild] {
				return
			}
			if h.elems[id] > h.elems[lChild] {
				h.elems[id], h.elems[lChild] = h.elems[lChild], h.elems[id]
				id = lChild
			} else {
				return
			}
		} else if rChild < h.length {
			if h.elems[id] < h.elems[rChild] {
				return
			}
			if h.elems[id] > h.elems[rChild] {
				h.elems[id], h.elems[rChild] = h.elems[rChild], h.elems[id]
				id = rChild
			} else {
				return
			}
		} else {
			return
		}
	}
}

func (h *Heap) Min() int {
	return h.elems[0]
}

func FindKthLargest(nums []int, k int) int {
	// Heap k-1
	heap := InitHeap(k)
	for i := range nums {
		heap.Add(nums[i])
	}
	// heap.Print()
	return heap.Min()
}
