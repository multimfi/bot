package alert

import (
	"fmt"
	"hash/crc32"
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
	var ret []string

	for k := range a.Labels {
		ret = append(ret, k)
	}
	sort.Strings(ret)

	return ret
}

// Hash retuns a crc32 checksum from alert labels.
func (a *Alert) Hash() uint32 {
	b := getBuf()
	defer putBuf(b)

	for _, k := range a.names() {
		b = append(b, k...)
		b = append(b, a.Labels[k]...)
	}

	return crc32.ChecksumIEEE(b)
}

func truncate(t time.Duration) time.Duration {
	return t - t%time.Second
}

func (a *Alert) since() time.Duration {
	return time.Now().UTC().Sub(a.StartsAt)
}

func (a *Alert) lasted() time.Duration {
	return a.EndsAt.Sub(a.StartsAt)
}

// SetCurrent sets the current expected responder for alert to i atomically.
func (a *Alert) SetCurrent(i int32) {
	atomic.StoreInt32(&a.current, i)
}

// Current returns the current expected responder for alert atomically.
func (a *Alert) Current() int32 {
	return atomic.LoadInt32(&a.current)
}

func kdv(m map[string]string, k, v string) string {
	if r, ok := m[k]; ok {
		return r
	}
	return v
}

// String returns a consise summary for alert.
func (a *Alert) String() string {
	switch a.Status {
	case AlertFiring:
		return fmt.Sprintf(
			"%s: %s, %s: %s (since %s)",
			kdv(a.Labels, "instance", "none"),
			kdv(a.Labels, "job", "none"),
			kdv(a.Labels, "alertname", "unnamed"),
			kdv(a.Annotations, "summary", "none"),
			truncate(a.since()),
		)
	case AlertResolved:
		return fmt.Sprintf(
			"%s: %s, %s: %s (lasted %s)",
			kdv(a.Labels, "instance", "none"),
			kdv(a.Labels, "job", "none"),
			kdv(a.Labels, "alertname", "unnamed"),
			kdv(a.Annotations, "summary", "none"),
			truncate(a.lasted()),
		)
	}
	return "error: unknown status: " + a.Status
}
