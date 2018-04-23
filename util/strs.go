package util

func RemoveQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		return s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		return s[:len(s)-1]
	}
	return s
}
