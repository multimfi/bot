package alert

import (
	"sort"
	"sync"
)

type Pool struct {
	alertsMu sync.RWMutex
	alerts   map[uint32]*Alert
}

func NewPool() *Pool {
	return &Pool{
		alerts: make(map[uint32]*Alert),
	}
}

func (p *Pool) Reset() {
	p.alertsMu.Lock()
	for k := range p.alerts {
		delete(p.alerts, k)
	}
	p.alertsMu.Unlock()
}

func (p *Pool) List() []*Alert {
	p.alertsMu.RLock()

	r := make([]*Alert, 0)
	for _, v := range p.alerts {
		r = append(r, v)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].StartsAt.Before(r[j].StartsAt)
	})

	p.alertsMu.RUnlock()

	return r
}

func (p *Pool) Add(a *Alert) (bool, *Alert) {
	p.alertsMu.Lock()

	h := a.Hash()
	if c, exists := p.alerts[h]; exists {
		p.alertsMu.Unlock()
		return !exists, c
	}

	p.alerts[h] = a
	p.alertsMu.Unlock()

	return true, a
}

func (p *Pool) Remove(a *Alert) bool {
	p.alertsMu.Lock()

	h := a.Hash()
	if _, exists := p.alerts[h]; !exists {
		p.alertsMu.Unlock()
		return exists
	}

	delete(p.alerts, h)
	p.alertsMu.Unlock()

	return true
}
