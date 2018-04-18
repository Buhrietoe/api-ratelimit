package limit

import (
	"container/list"
	"sync"
	"time"
)

// Rate defines simplified ratelimit
type Rate struct {
	Limit    int `json:"limit"`
	requests list.List
	mtx      sync.Mutex
}

// New creates a new rate
func New(limit int) *Rate {
	l := &Rate{
		Limit: limit,
	}

	l.requests.Init()
	return l
}

// Delay is a simple blocking rate limiter
func (r *Rate) Delay() {
	for {
		ok, delay := r.Check()
		if ok {
			break
		}
		time.Sleep(delay)
	}
}

// Check returns true if under limit
// returns false and delay if over limit
func (r *Rate) Check() (ok bool, delay time.Duration) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	now := time.Now()

	if l := r.requests.Len(); l < r.Limit {
		r.requests.PushBack(now)
		return true, 0
	}

	front := r.requests.Front()
	if diff := now.Sub(front.Value.(time.Time)); diff < time.Second {
		return false, time.Second - diff
	}

	front.Value = now
	r.requests.MoveToBack(front)
	return true, 0
}
