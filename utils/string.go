package utils

//StringInSlice
func StringInSlice(content string, list []string) bool {
	for _, b := range list {
		if b == content {
			return true
		}
	}
	return false
}
