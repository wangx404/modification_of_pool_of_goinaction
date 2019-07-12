// Example provided with help from Fatih Arslan and Gabriel Aszalos.
// Package pool manages a user defined set of resources.
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages a set of resources that can be shared safely by
// multiple goroutines. The resource being managed must implement
// the io.Closer interface.
type Pool struct {
	m           sync.Mutex
	resources   chan io.Closer
	factory     func() (io.Closer, error)
	closed      bool
	maxSize     uint // maximum number of resources
  currentSize uint // current number of resouces
}

// ErrPoolClosed is returned when an Acquire returns on a
// closed pool.
var ErrPoolClosed = errors.New("Pool has been closed.")

// New creates a pool that manages resources. A pool requires a
// function that can allocate a new resource and the size of
// the pool.
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
		maxSize:   size,
	}, nil
}

// Acquire retrieves a resource	from the pool.
func (p *Pool) Acquire() (io.Closer, error) {
	// Check for a free resource.
	p.m.Lock()
	if p.currentSize < p.maxSize {
		log.Println("Acquire:", "New Resource")
		p.currentSize += 1
		p.m.Unlock()
		return p.factory()
	}
	p.m.Unlock()
  
	// No free resouce, wait for it
	r, _ := <-p.resources
	log.Println("Acquire:", "Shared Resource")
	// if you dont block codes under down, dead lock problem happens.
	//if !ok {
	//	return nil, ErrPoolClosed
	//}
	return r, nil
}

// Release places a new resource onto the pool.
func (p *Pool) Release(r io.Closer) {
	// Secure this operation with the Close operation.
	p.m.Lock()
	defer p.m.Unlock()

	// If the pool is closed, discard the resource.
	if p.closed {
		r.Close()
		return
	}

	select {
	// Attempt to place the new resource on the queue.
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	// If the queue is already at cap we close the resource.
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close will shutdown the pool and close all existing resources.
func (p *Pool) Close() {
	// Secure this operation with the Release operation.
	p.m.Lock()
	defer p.m.Unlock()

	// If the pool is already close, don't do anything.
	if p.closed {
		return
	}

	// Set the pool as closed.
	p.closed = true

	// Close the channel before we drain the channel of its
	// resources. If we don't do this, we will have a deadlock.
	close(p.resources)

	// Close the resources
	for r := range p.resources {
		r.Close()
	}
}
