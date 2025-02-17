package Optional

func MapStringInterface(origin map[string]interface{}, key string, defaultValue map[string]interface{}) map[string]interface{} {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case map[string]interface{}:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
