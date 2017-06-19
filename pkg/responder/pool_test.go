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
	f := &fakeTime{}
	now = f.now

	p := NewPool()
	data := []string{"0", "1", "2", "3"}
	tc := []struct {
		n int
		w int
	}{
		{0, 0},
		{3, 3},
		{2, 1},
		{1, 2},
	}
	for _, v := range tc {
		f.set(v.n)
		_, r, _ := p.Get(data)
		if r != v.w {
			t.Errorf("%d != %d", r, v.w)
		}
	}
}
