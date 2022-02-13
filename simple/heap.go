package main

import (
	"encoding/json"
	"fmt"
)

type Heap struct {
	values []int
}

func (h Heap) toString() string {
	s, _ := json.Marshal(h.values)
	return string(s)
}

// ヒープの最後尾に値を追加
func (h *Heap) push(v int) {
	h.values = append(h.values, v)

	// 挿入された箇所の index
	var currentIdx int = len(h.values) - 1
	for currentIdx > 0 {
		var parentIdx int = (currentIdx - 1) / 2
		if h.values[parentIdx] >= v {
			break
		}
		// index の入れ替え
		h.values[currentIdx] = h.values[parentIdx]
		currentIdx = parentIdx
	}
	h.values[currentIdx] = v
}

// ヒープの値から最大値を削除
func (h *Heap) pop() {
	if len(h.values) == 0 {
		return
	}
	// 移動を行う値
	x := h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	var idx int
	for 2*idx+1 < len(h.values) {
		var rightChildIdx int = 2*idx + 1
		var leftChildIdx int = 2*idx + 2
		fmt.Printf("right: %d, left: %d\n", rightChildIdx, leftChildIdx)
		var bigChildIdx int = leftChildIdx
		if leftChildIdx <= len(h.values)-1 && h.values[rightChildIdx] > h.values[leftChildIdx] {
			bigChildIdx = rightChildIdx
		}

		if h.values[bigChildIdx] <= x {
			break
		}
		h.values[idx] = h.values[bigChildIdx]
		idx = bigChildIdx
	}
	h.values[idx] = x
}

func main() {
	values := []int{5, 3, 2}
	heap := Heap{values}
	heap.push(10)
	fmt.Println(heap.toString())

	heap.pop()
	fmt.Println(heap.toString())
}
