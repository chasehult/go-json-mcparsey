package JSONHelper

import "strconv"

func IntRequired(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue int) (value int, isValid bool) {
	if maybeProductId, ok := origin[key]; ok {
		switch tempProductId := maybeProductId.(type) {
		case string:
			ProductId, err := strconv.Atoi(tempProductId)
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
			return ProductId, true
		case int:
			return tempProductId, true
		case float64:
			return int(tempProductId), true
		default:
			AppendWhenNotNil(requiredFields, key)
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}

	return defaultValue, false
}