package main

type Loc struct {
	col int
	row int
}

type Queue[T any] []T

func (q *Queue[T]) enqueue(t T) {
	*q = append(*q, t)
}

func (q *Queue[T]) dequeue() T {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}

func (q *Queue[T]) isEmpty() bool {
	return len(*q) == 0
}

type MaxHeap[T interface{ value() float64 }] []T

func (h MaxHeap[T]) Len() int { return len(h) }
func (h MaxHeap[T]) Less(i, j int) bool {
	return h[i].value() > h[j].value()
}
func (h MaxHeap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *MaxHeap[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type MinHeap[T interface{ value() float64 }] []T

func (h MinHeap[T]) Len() int { return len(h) }
func (h MinHeap[T]) Less(i, j int) bool {
	return h[i].value() > h[j].value()
}
func (h MinHeap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
