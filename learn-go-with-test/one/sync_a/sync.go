package sync

import "sync"

type Counter struct{
	// これで書き込みをロックできる
	mu sync.Mutex
  value int
}

func (c *Counter) Inc() {
	// ロックかけられる
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}