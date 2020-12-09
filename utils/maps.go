package utils

// Keys - Get all keys of a map
func Keys(m map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
