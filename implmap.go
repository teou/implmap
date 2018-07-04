package implmap

import (
	"reflect"
	"sync"
)

var (
	m = make(map[string][]reflect.Type)
	l = &sync.RWMutex{}
)

//accept struct pointer only
func Add(n string, t reflect.Type) {
	if t == nil || n == "" || !isStructPtr(t) {
		return
	}
	l.Lock()
	defer l.Unlock()
	a, ok := m[n]
	if !ok || a == nil {
		a = []reflect.Type{}
	}
	a = append(a, t)
	m[n] = a
}

func Get(n string) []reflect.Type {
	if n == "" {
		return []reflect.Type{}
	}
	l.RLock()
	defer l.RUnlock()
	types := m[n]
	if types == nil {
		return []reflect.Type{}
	}
	ret := []reflect.Type{}
	for _, t := range types {
		if t == nil {
			continue
		}
		ret = append(ret, t)
	}
	return ret
}

func isStructPtr(t reflect.Type) bool {
	return t != nil && t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
