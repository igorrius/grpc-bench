package tools

import "sync"

type Counter struct {
	mu    sync.Mutex
	count uint
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func (c *Counter) Clear() {
	c.mu.Lock()
	c.count = 0
	defer c.mu.Unlock()
}
