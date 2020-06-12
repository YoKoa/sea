package concurrent

import "sync"

type LockFreeQueue struct {
	mu     sync.Mutex
	values []interface{}
}

func (q *LockFreeQueue) Add(v interface{}) (one bool) {
	q.mu.Lock()
	q.values = append(q.values, v)
	n := len(q.values)
	q.mu.Unlock()
	return n == 1
}
