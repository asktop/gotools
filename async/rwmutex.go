package async

import "sync"

type RWMutex struct {
	sync.RWMutex
	safe bool
}

func New(unsafe ...bool) *RWMutex {
	mu := new(RWMutex)
	if len(unsafe) > 0 {
		mu.safe = !unsafe[0]
	} else {
		mu.safe = true
	}
	return mu
}

func (mu *RWMutex) IsSafe() bool {
	return mu.safe
}

func (mu *RWMutex) Lock() {
	if mu.safe {
		mu.RWMutex.Lock()
	}
}

func (mu *RWMutex) Unlock() {
	if mu.safe {
		mu.RWMutex.Unlock()
	}
}

func (mu *RWMutex) RLock() {
	if mu.safe {
		mu.RWMutex.RLock()
	}
}

func (mu *RWMutex) RUnlock() {
	if mu.safe {
		mu.RWMutex.RUnlock()
	}
}
