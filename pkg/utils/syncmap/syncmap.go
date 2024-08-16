package syncmap

import (
	"sync"
)

type SyncMap[key comparable, value any] struct {
	data sync.Map
}

func (a *SyncMap[T1, T2]) Load(key T1) (r T2, _ bool) {
	v, ok := a.data.Load(key)
	if !ok {
		return r, false
	}
	return v.(T2), true
}

func (a *SyncMap[T1, T2]) Store(key T1, value T2) {
	a.data.Store(key, value)
}

func (a *SyncMap[T1, T2]) LoadOrStore(key T1, value T2) (r T2, _ bool) {
	v, ok := a.data.LoadOrStore(key, value)
	return v.(T2), ok
}

func (a *SyncMap[T1, T2]) LoadAndDelete(key T1) (r T2, _ bool) {
	v, ok := a.data.LoadAndDelete(key)
	if !ok {
		return r, false
	}
	return v.(T2), true
}

func (a *SyncMap[T1, T2]) Delete(key T1) {
	a.data.Delete(key)
}

func (a *SyncMap[T1, T2]) Range(f func(key T1, value T2) bool) {
	a.data.Range(func(key, value any) bool {
		return f(key.(T1), value.(T2))
	})
}

func (a *SyncMap[key, T2]) ValueSlice() (r []T2) {
	for _, v := range a.data.Range {
		r = append(r, v.(T2))
	}
	return r
}

func (a *SyncMap[T1, T2]) Swap(x T1, b T2) (T2, bool) {
	v, ok := a.data.Swap(x, b)
	if ok {
		return v.(T2), true
	}

	return *new(T2), false
}

func (a *SyncMap[T1, T2]) CompareAndSwap(key T1, old T2, new T2) (swapped bool) {
	return a.data.CompareAndSwap(key, old, new)
}

func (a *SyncMap[T1, T2]) CompareAndDelete(key T1, old T2) (deleted bool) {
	return a.data.CompareAndDelete(key, old)
}

type Diff[K comparable, V any] struct {
	Rmoved bool
	Added  bool
	Modif  bool

	Key      K
	OldValue V
	NewValue V
}

func Differ[K comparable, V any](old, new *SyncMap[K, V], isSame func(v1, v2 V) bool) func(f func(Diff[K, V])) {
	return func(f func(Diff[K, V])) {
		for k, v1 := range old.Range {
			v2, ok := new.Load(k)
			if !ok {
				f(Diff[K, V]{Rmoved: true, Key: k, OldValue: v1})
				continue
			}

			if !isSame(v1, v2) {
				f(Diff[K, V]{Modif: true, Key: k, OldValue: v1, NewValue: v2})
			}
		}

		for k, v2 := range new.Range {
			_, ok := old.Load(k)
			if !ok {
				f(Diff[K, V]{Added: true, Key: k, NewValue: v2})
			}
		}
	}
}

func DiffMap[K comparable, V any](old, new map[K]V, isSame func(v1, v2 V) bool) func(f func(Diff[K, V])) {
	return func(f func(Diff[K, V])) {
		for k, v1 := range old {
			v2, ok := new[k]
			if !ok {
				f(Diff[K, V]{Rmoved: true, Key: k, OldValue: v1})
				continue
			}

			if !isSame(v1, v2) {
				f(Diff[K, V]{Modif: true, Key: k, OldValue: v1, NewValue: v2})
			}
		}

		for k, v2 := range new {
			_, ok := old[k]
			if !ok {
				f(Diff[K, V]{Added: true, Key: k, NewValue: v2})
			}
		}
	}
}
