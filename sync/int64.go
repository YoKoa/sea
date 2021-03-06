package sync


import "sync/atomic"

type Int64 int64

func (a *Int64) Int64() int64 {
	return atomic.LoadInt64((*int64)(a))
}

func (a *Int64) AsInt() int {
	return int(a.Int64())
}

func (a *Int64) Set(v int64) {
	atomic.StoreInt64((*int64)(a), v)
}

func (a *Int64) CompareAndSwap(o, n int64) bool {
	return atomic.CompareAndSwapInt64((*int64)(a), o, n)
}

func (a *Int64) Swap(v int64) int64 {
	return atomic.SwapInt64((*int64)(a), v)
}

func (a *Int64) Add(v int64) int64 {
	return atomic.AddInt64((*int64)(a), v)
}

func (a *Int64) Sub(v int64) int64 {
	return a.Add(-v)
}

func (a *Int64) Incr() int64 {
	return a.Add(1)
}

func (a *Int64) Decr() int64 {
	return a.Add(-1)
}
