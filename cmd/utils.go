package cmd

func checkNil(decoded interface{}, key string) string {
	val, ok := decoded.(map[string]interface{})[key]
	if ok && val != nil {
		return val.(string)
	}
	return ""
}
