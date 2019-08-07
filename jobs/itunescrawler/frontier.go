package itunescrawler

import (
	s "github.com/golang-collections/go-datastructures/set"
)

type Frontier struct {
	// Insert Queue
	I chan string
	// Select Queue
	S chan string
	// Set of strings processed till now
	set *s.Set
}

func NewFrontier(size int) *Frontier {
	frontier := &Frontier{
		I:   make(chan string),
		S:   make(chan string, size),
		set: s.New(),
	}

	go func(f *Frontier) {
		for i := range f.I {
			if !f.set.Exists(i) {
				f.set.Add(i)
				f.S <- i
			}
		}
	}(frontier)

	return frontier
}

func (f *Frontier) Ignore(s string) {
	f.set.Add(s)
}

func (f *Frontier) Clear() {
	f.set.Clear()
}
