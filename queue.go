package main

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
