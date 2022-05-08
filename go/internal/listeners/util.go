package listeners

func trimQuotes(value string) string {
	if value == "" {
		return value
	}
	if value[0] == '"' && value[len(value)-1] == '"' {
		return value[1 : len(value)-1]
	}
	return value
}
