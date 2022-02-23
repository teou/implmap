package implmap

import (
	"fmt"
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

	l := len(a)
	if l > 0 {
		fmt.Println(fmt.Sprintf("implmap append new type(%v) impl to name(%v) at index(%v), old array=%v", t, n, l, a))
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
func GetAll() map[string][]reflect.Type {
	l.RLock()
	defer l.RUnlock()
	mCopy := make(map[string][]reflect.Type)
	for key := range m {
		ret := []reflect.Type{}
		for _, t := range m[key] {
			if t == nil {
				continue
			}
			ret = append(ret, t)
			mCopy[key] = ret
		}
	}
	return mCopy
}
func isStructPtr(t reflect.Type) bool {
	return t != nil && t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
