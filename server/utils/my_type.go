package utils

func GetType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case []int:
		return "[]int"
	case []string:
		return "[]string"
	case map[string]int:
		return "map[string]int"
	default:
		return "unknown"
	}
}
