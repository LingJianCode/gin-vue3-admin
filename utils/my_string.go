package utils

func RemoveLastChar(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1]
}
