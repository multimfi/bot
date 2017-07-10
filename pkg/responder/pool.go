package responder

import (
	"errors"
	"sync"
	"time"
)

// for testing
var now = time.Now

// ErrNoResponse only returned on Pool.Get() failure.
var ErrNoResponse = errors.New("all responders failed")

// Pool is a pool of Responders mapped by name.
type Pool struct {
	poolMu sync.RWMutex
	pool   map[string]*Responder
}

// NewPool returns an empty Pool.
func NewPool() *Pool {
	return &Pool{
		pool: make(map[string]*Responder),
	}
}

// Add adds responder r to Pool if r is not already in pool.
func (p *Pool) Add(r *Responder) {
	p.poolMu.Lock()

	if _, exists := p.pool[r.Name]; exists {
		p.poolMu.Unlock()
		return
	}

	p.pool[r.Name] = r
	p.poolMu.Unlock()
}

// List returns a slice of current responders.
func (p *Pool) List() []*Responder {
	p.poolMu.RLock()

	ret := make([]*Responder, 0)

	for _, v := range p.pool {
		ret = append(ret, v)
	}

	p.poolMu.RUnlock()
	return ret
}

func get(w, l int) int {
	if w < 1 {
		return 0
	}

	if w <= l {
		return w - 1
	}

	w = w % l
	if w < 1 {
		w = l
	}
	return w - 1
}

func (p *Pool) get(name string, add bool) *Responder {
	p.poolMu.Lock()
	if r, exists := p.pool[name]; exists {
		p.poolMu.Unlock()
		return r
	}
	p.poolMu.Unlock()

	n := &Responder{Name: name}
	if add {
		p.Add(n)
	}

	return n
}

// Get returns ErrNoResponse when all responders are in failed state.
// All non existant elements of r are added to Pool.
func (p *Pool) Get(r []string) (*Responder, int, error) {
	_, w := now().ISOWeek()
	l := len(r)

	if l < 1 {
		return nil, 0, nil
	}

	for x := 0; x < l; x++ {
		i := get(w+x, l)
		u := p.get(r[i], true)

		if !u.Failed() {
			return u, i, nil
		}
	}

	return nil, 0, ErrNoResponse
}

// Update increments the tickcounter for name.
// Does not add name to pool.
func (p *Pool) Update(name string) {
	r := p.get(name, false)
	r.Update()
}

// ResetFailed resets failed state for name,
// returns true if name was in a failed state.
// Adds name to pool if it does not exist.
func (p *Pool) ResetFailed(name string) bool {
	r := p.get(name, true)
	return r.resetFailed()
}
