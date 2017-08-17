package utils

//CheckStringInSlice
func CheckStringInSlice(content string, list []string) bool {
	for _, b := range list {
		if b == content {
			return true
		}
	}
	return false
}
