package limit

import (
	"container/list"
	"sync"
)

type Rate struct {
	Limit    int `json:"limit"`
	mtx      sync.Mutex
	requests list.List
}
