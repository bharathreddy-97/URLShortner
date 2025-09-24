package persistence

import (
	"sync"
)

type SMap struct {
	m         sync.RWMutex
	data      map[string]string
	listOfIds map[string]struct{}
}

func GetSMap() *SMap {
	return &SMap{
		data:      make(map[string]string),
		listOfIds: map[string]struct{}{},
	}
}

func (s *SMap) Get(key string) string {
	s.m.Lock()
	defer s.m.Unlock()
	return s.data[key]
}

func (s *SMap) Set(key string, value string) {
	s.m.Lock()
	s.data[key] = value
	s.listOfIds[value] = struct{}{}
	s.m.Unlock()
}

func (s *SMap) CheckForId(id string) bool {
	s.m.Lock()
	defer s.m.Unlock()
	_, ok := s.listOfIds[id]
	return ok
}
