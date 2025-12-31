package maps

type Map[K comparable, V any] interface {
	Put(K, V)
	Get(K) (V, bool)
	Del(K)
	Contains(K) bool
	Len() int
	Clear()
	Range(func(K, V))
}

func MergeMap(src, dst map[string]any) map[string]any {
	for key, srcValue := range src {
		if dstValue, ok := dst[key]; ok {
			srcValueMap, srcValueMapOk := srcValue.(map[string]any)
			dstValueMap, dstValueMapOk := dstValue.(map[string]any)
			if srcValueMapOk && dstValueMapOk {
				dst[key] = MergeMap(srcValueMap, dstValueMap)
			} else {
				dst[key] = srcValue
			}
		} else {
			dst[key] = srcValue
		}
	}
	return dst
}
