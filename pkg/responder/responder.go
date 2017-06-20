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

// Responder tracks responder state.
type Responder struct {
	Name     string
	lastTick uint64

	mu     sync.RWMutex
	once   sync.Once
	active bool
	failed bool
}

// Update increments an internal counter atomically.
func (r *Responder) Update() {
	atomic.AddUint64(&r.lastTick, 1)
}

func (r *Responder) get() uint64 {
	return atomic.LoadUint64(&r.lastTick)
}

// Active returns true if responder is considered active.
func (r *Responder) Active() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.active
}

// Failed returns true if responder is considered failed,
// state can only be reset with resetFailed().
func (r *Responder) Failed() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.failed
}

// Ping waits for a response from responder.
// Argument f is called on state change.
func (r *Responder) Ping(f func()) {
	r.once.Do(func() {
		r.wait(f)
	})
}

func (r *Responder) setactive() {
	r.mu.Lock()
	r.active = true
	r.mu.Unlock()
}

func (r *Responder) setfailed() {
	r.mu.Lock()
	r.failed = true
	r.mu.Unlock()
}

func (r *Responder) reset() {
	r.mu.Lock()
	r.active = false
	r.failed = false
	r.once = sync.Once{}
	r.mu.Unlock()
}

func (r *Responder) resetFailed() bool {
	if !r.Failed() {
		return false
	}
	r.reset()
	return true
}

// wait waits for responder activity
func (r *Responder) wait(f func()) {
	c := r.get()
	go func() {
		for i := 0; i < 300; i++ {
			time.Sleep(tickRate)
			if r.get() <= c {
				continue
			}

			r.setactive()
			f()

			time.Sleep(liveness)
			r.reset()

			f()
			return
		}

		r.setfailed()
		f()
	}()
}
