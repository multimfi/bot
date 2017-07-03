package alert

import (
	"crypto/md5"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var bufPool = sync.Pool{}

func getBuf() []byte {
	b := bufPool.Get()
	if b == nil {
		return make([]byte, 0, 1024)
	}
	return b.([]byte)
}

func putBuf(b []byte) {
	b = b[:0]
	bufPool.Put(b)
}

func (a *Alert) names() []string {
	ret := make([]string, 0, len(a.Labels))

	for k := range a.Labels {
		ret = append(ret, k)
	}

	sort.Strings(ret)

	return ret
}

// Hash returns an md5 hash from alert labels.
func (a *Alert) Hash() [16]byte {
	b := getBuf()
	defer putBuf(b)

	for _, k := range a.names() {
		b = append(b, k...)
		b = append(b, a.Labels[k]...)
	}

	return md5.Sum(b)
}

func truncate(t time.Duration) time.Duration {
	return t - t%time.Second
}

func (a *Alert) Since() time.Duration {
	return truncate(time.Now().UTC().Sub(a.StartsAt))
}

func (a *Alert) Lasted() time.Duration {
	return truncate(a.EndsAt.Sub(a.StartsAt))
}

// SetCurrent sets the current expected responder for alert to i atomically.
func (a *Alert) SetCurrent(i int) {
	atomic.StoreInt32(&a.current, int32(i))
}

// Current returns the current expected responder for alert atomically.
func (a *Alert) Current() int {
	return int(atomic.LoadInt32(&a.current))
}

// AllFail returns true when SetAllFail has been called.
func (a *Alert) AllFail() bool {
	return atomic.LoadInt32(&a.allfail) == 1
}

// SetAllFails sets status to allfailed.
func (a *Alert) SetAllFail() {
	atomic.StoreInt32(&a.allfail, 1)
}
