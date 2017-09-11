package responder

import (
	"testing"
	"time"
)

type fakeTime struct {
	n time.Time
}

func (f *fakeTime) set(i int) {
	f.n = f.n.Add(time.Hour * 24 * 7 * time.Duration(i))
}

func (f *fakeTime) now() time.Time {
	return f.n
}

func TestPoolGet(t *testing.T) {
	t.Parallel()

	f := &fakeTime{}
	now = f.now

	p := NewPool()
	data := []string{"0", "1", "2", "3"}
	tc := []struct {
		n int
		w int
		e bool
	}{
		{0, 0, true},
		{3, 3, true},
		{2, 1, true},
		{1, 2, true},
		{3, 0, false},
		{2, 2, false},
		{1, 1, false},
	}
	for _, v := range tc {
		f.set(v.n)
		_, r, _ := p.Get(data)
		if v.e == (r != v.w) {
			t.Errorf("%d != %d (%v)", r, v.w, v.e)
		}
	}
}
