package xcache

import "sync"

var syncCache sync.Map

func SyncLoad[T any](key string) (v T) {
	value, ok := syncCache.Load(key)
	if !ok {
		return
	}
	return value.(T)
}

func SyncStore(key string, value any) {
	syncCache.Store(key, value)
}

func SyncDelete(key string) {
	syncCache.Delete(key)
}
