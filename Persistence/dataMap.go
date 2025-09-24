package persistence

import (
	"sync"
)

type SMap struct {
	m    sync.RWMutex
	data map[string]string
}

func GetSMap() *SMap {
	return &SMap{data: make(map[string]string)}
}

func (s *SMap) Get(key string) string {
	s.m.Lock()
	defer s.m.Unlock()
	return s.data[key]
}

func (s *SMap) Set(key string, value string) {
	s.m.Lock()
	s.data[key] = value
	s.m.Unlock()
}
