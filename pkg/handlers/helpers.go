package handlers

func removeMapValuesFromArrays(initialMap map[string][]string) map[string]string {
	flattenedMap := make(map[string]string)
	for key, value := range initialMap {
		flattenedMap[key] = value[0]
	}
	return flattenedMap
}
