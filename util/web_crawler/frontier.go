package web_crawler

import (
	"github.com/golang-collections/go-datastructures/set"
)

type Frontier struct {
	// Input channel
	I chan string
	// Output channel
	O chan string
	// Set of strings processed till now
	set *set.Set
}

func NewFrontier(size int) *Frontier {
	frontier := &Frontier{
		I:   make(chan string, 1000),
		O:   make(chan string, size),
		set: set.New(),
	}

	go func() {
		for i := range frontier.I {
			if !frontier.set.Exists(i) {
				frontier.set.Add(i)
				frontier.O <- i
			}
		}
	}()

	return frontier
}

// Ignore input
func (f *Frontier) Ignore(s string) {
	f.set.Add(s)
}

// Clear all values received till now
func (f *Frontier) Clear() {
	f.set.Clear()
}
