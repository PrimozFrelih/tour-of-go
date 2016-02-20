package mutex

import "sync"

type SafeStruct struct {
	m   map[string]int // NOTE: map is not thread-safe
	mux sync.RWMutex
}

func (s *SafeStruct) Increment(key string) {
	s.mux.Lock()
	defer s.mux.Unlock() // defer is always executed after when the method returns - like finally block

	s.m[key]++
}

func (s *SafeStruct) Value(key string) int {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.m[key]
}
