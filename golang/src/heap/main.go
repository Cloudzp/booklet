package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

func main() {
	myHeap := &MyHeap{}
	heap.Init(myHeap)

	for i := 0; i < 10; i++ {
		heap.Push(myHeap, randomEle())
	}

	//fmt.Println((*myHeap).element)
	for i := 0; i < 10; i++ {
		fmt.Println(heap.Pop(myHeap))
	}
}

func randomEle() *Test {
	return &Test{
		rand.Int(),
	}
}

type Test struct {
	P int
}

type MyHeap struct {
	element []*Test
}

/*
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

*/

/*
type frameQueue []*prioritizedFrame

func (fq frameQueue) Len() int {
	return len(fq)
}

func (fq frameQueue) Less(i, j int) bool {
	if fq[i].priority == fq[j].priority {
		return fq[i].insertId < fq[j].insertId
	}
	return fq[i].priority < fq[j].priority
}

func (fq frameQueue) Swap(i, j int) {
	fq[i], fq[j] = fq[j], fq[i]
}

func (fq *frameQueue) Push(x interface{}) {
	*fq = append(*fq, x.(*prioritizedFrame))
}

func (fq *frameQueue) Pop() interface{} {
	old := *fq
	n := len(old)
	*fq = old[0 : n-1]
	return old[n-1]
}


*/
func (h *MyHeap) Less(j, i int) bool {
	if h.element[i].P > h.element[j].P {
		return false
	}
	return true

}

func (h *MyHeap) Swap(i, j int) {
	//fmt.Println(i,j)

		h.element[i], h.element[j] = h.element[j], h.element[i]

}

func (h *MyHeap) Len() int {
	return len(h.element)
}

func (h *MyHeap) Push(x interface{}) {
	newData, ok := x.(*Test)
	if !ok {
		return
	}
	h.element = append(h.element, newData)
}

func (h *MyHeap) Pop() interface{} {
	n := len(h.element)
	fmt.Println("test", n)
	x := h.element[n-1]
	h.element = h.element[0:n-1]

	return x
}
