package persistence

import (
	"regexp"
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

func ReturnOriginalURL(dataMap *SMap, id string) string {
	dataMap.m.Lock()
	defer dataMap.m.Unlock()
	for key, value := range dataMap.data {
		if value == id {
			return key
		}
	}
	return ""
}

// func (s *SMap) GetTopMetrics() map[string]int64 {
// 	s.m.Lock()
// 	defer s.m.Unlock()

// 	keys := make([]string, 0, len(s.data))
// 	for k := range s.data {
// 		keys = append(keys, k)
// 	}

// 	return metricsMap
// }

func returnHost(url string) string {
	r := regexp.MustCompile(`/`)

	return r.Split(url, -1)[2]
}
