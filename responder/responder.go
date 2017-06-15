package responder

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	tickRate = time.Second
	liveness = time.Minute * 30
)

type Responder struct {
	Name     string
	lastTick uint64

	mu     sync.RWMutex
	once   sync.Once
	active bool
	failed bool
}

func (r *Responder) Update() {
	atomic.AddUint64(&r.lastTick, 1)
}

func (r *Responder) get() uint64 {
	return atomic.LoadUint64(&r.lastTick)
}

func (r *Responder) Active() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.active
}

func (r *Responder) Failed() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.failed
}

func (r *Responder) Ping() {
	r.once.Do(r.wait)
}

func (r *Responder) setactive() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.active = true
}

func (r *Responder) setfailed() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failed = true
}

func (r *Responder) reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.active = false
	r.failed = false
	r.once = sync.Once{}
}

func (r *Responder) resetFailed() {
	if !r.Failed() {
		return
	}
	r.reset()
}

// wait waits for responder activity
func (r *Responder) wait() {
	c := r.get()
	go func() {
		for i := 0; i < 300; i++ {
			if r.get() > c {
				r.setactive()
				time.Sleep(liveness)
				r.reset()
				return
			}
			time.Sleep(tickRate)
		}
		r.setfailed()
	}()
}
