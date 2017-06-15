package responder

import (
	"errors"
	"sync"
	"time"
)

var ErrNoResponse = errors.New("all responders failed")

type Pool struct {
	poolMu sync.RWMutex
	pool   map[string]*Responder
}

func NewPool() *Pool {
	return &Pool{
		pool: make(map[string]*Responder),
	}
}

func (p *Pool) Add(r *Responder) {
	p.poolMu.Lock()

	if _, exists := p.pool[r.Name]; exists {
		p.poolMu.Unlock()
		return
	}

	p.pool[r.Name] = r
	p.poolMu.Unlock()
}

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
		return 1
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
func (p *Pool) Get(r []string) (*Responder, int, error) {
	_, w := time.Now().ISOWeek()
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

func (p *Pool) Update(name string) {
	r := p.get(name, false)
	r.Update()
}

func (p *Pool) ResetFailed(name string) bool {
	r := p.get(name, true)
	return r.resetFailed()
}
