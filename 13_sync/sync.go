package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// returns a new counter
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
