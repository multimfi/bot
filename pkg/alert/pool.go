package alert

import (
	"sort"
	"sync"
)

// Pool stores alerts in a map by alert.Hash().
type Pool struct {
	alertsMu sync.RWMutex
	alerts   map[[16]byte]*Alert
}

// NewPool returns a new empty alert Pool.
func NewPool() *Pool {
	return &Pool{
		alerts: make(map[[16]byte]*Alert),
	}
}

// Reset clears the internal map of Pool.
func (p *Pool) Reset() bool {
	p.alertsMu.Lock()

	b := len(p.alerts)
	for k := range p.alerts {
		delete(p.alerts, k)
	}

	a := len(p.alerts)
	p.alertsMu.Unlock()

	return b != a
}

// List returns a slice sorted by StartsAt of current alerts.
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

// Add adds alert a to pool, if alert exists
// the current alert is returned.
// Returned bool indicates successful addition.
func (p *Pool) Add(a *Alert) (bool, *Alert) {
	h := a.Hash()

	p.alertsMu.Lock()
	if c, exists := p.alerts[h]; exists {
		p.alertsMu.Unlock()
		return !exists, c
	}

	p.alerts[h] = a
	p.alertsMu.Unlock()

	return true, a
}

// Remove deletes alert a from pool,
// returns true on success.
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
