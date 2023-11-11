package mqtt

import (
	"sync"
)

var (
	strMap = make(map[string]string)
	mu     sync.Mutex
)

func WriteToMap(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	strMap[key] = value
}

func ReadFromMap(key string) string {
	mu.Lock()
	defer mu.Unlock()
	return strMap[key]
}
