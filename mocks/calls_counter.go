package mocks

import "sync"

type CallsCounter struct {
	expectedCallsCount int
	mu                 sync.Mutex
	callsCount         int
}

func (c *CallsCounter) incCallsCount() {
	c.mu.Lock()
	c.callsCount += 1
	c.mu.Unlock()
}

func (c *CallsCounter) AssertCallsCount() bool {
	return c.callsCount == c.expectedCallsCount
}

func (c *CallsCounter) ExpectedCallsCount() int {
	return c.expectedCallsCount
}

func (c *CallsCounter) CallsCount() int {
	return c.callsCount
}
