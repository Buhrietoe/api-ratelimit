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
func New(limit int) Rate {
	return Rate{
		Limit: limit,
	}
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
// Returns false and delay if over limit
// This implements a FIFO linked list of request times
// Oldest are at the front, new requests added to back
func (r *Rate) Check() (ok bool, delay time.Duration) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	now := time.Now()

	// add to request list if not yet at limit
	if l := r.requests.Len(); l < r.Limit {
		r.requests.PushBack(now)
		return true, 0
	}

	// check if oldest request was sooner than a second
	front := r.requests.Front()
	if diff := now.Sub(front.Value.(time.Time)); diff < time.Second {
		return false, time.Second - diff
	}

	// oldest request was longer ago than a second
	// update oldest and rotate to back of list
	front.Value = now
	r.requests.MoveToBack(front)
	return true, 0
}
