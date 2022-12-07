package main

import "fmt"

type SlidingWindow[T any] struct {
	src  []T
	size int
	curr int
}

func NewSlidingWindow[T any](src []T, size int) (*SlidingWindow[T], error) {
	if size > len(src) {
		return nil, fmt.Errorf("size cannot be larger than source")
	}

	return &SlidingWindow[T]{src: src, size: size}, nil
}

func (w *SlidingWindow[T]) Next() (int, []T) {
	if w.curr+w.size > len(w.src) {
		return w.curr, nil
	}

	res := w.src[w.curr : w.curr+w.size]
	defer func() { w.curr++ }()
	return w.curr, res
}
