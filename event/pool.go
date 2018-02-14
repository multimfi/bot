package event

import (
	"fmt"
	"io"
	"sync"
)

const Len = 1 << 8

type Value struct {
	Type uint32
	Gen  uint32
	Data []byte
}

type Pool struct {
	mu   sync.Mutex
	cond *sync.Cond
	v    []Value
	gen  uint32
}

func NewPool() *Pool {
	r := &Pool{
		v: make([]Value, Len),
	}
	r.cond = sync.NewCond(&r.mu)
	return r
}

func (e *Pool) Add(v []byte, t uint32) {
	e.mu.Lock()
	e.gen++

	e.v[e.gen%Len] = Value{
		Type: t,
		Gen:  e.gen,
		Data: v,
	}

	e.cond.Broadcast()
	e.mu.Unlock()
}

func (e *Pool) Get(gen uint32) Value {
	e.mu.Lock()
	ret := e.v[gen%Len]
	e.mu.Unlock()
	return ret
}

func (e *Pool) Gen() uint32 {
	e.mu.Lock()
	ret := e.gen
	e.mu.Unlock()
	return ret
}

func (e *Pool) Lock() uint32 {
	e.mu.Lock()
	return e.gen
}

func (e *Pool) Unlock() {
	e.mu.Unlock()
}

func (e *Pool) Next(cur uint32) uint32 {
	e.mu.Lock()
	ret := e.gen
	if cur != ret {
		e.mu.Unlock()
		return ret
	}
	for {
		e.cond.Wait()
		ret = e.gen
		if ret != cur {
			break
		}
	}
	e.mu.Unlock()
	return ret
}

func (e *Pool) String(w io.Writer) {
	e.mu.Lock()
	for k, v := range e.v {
		fmt.Fprintf(w, "%d: %q\n", k, v)
	}
	e.mu.Unlock()
}
