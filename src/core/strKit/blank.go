package strKit

// IsBlank
/*
("  \r\n ") => true
*/
func IsBlank(str string) bool {
	return IsEmpty(TrimSpace(str))
}

func BlankToDefault(str, def string) string {
	if IsBlank(str) {
		return def
	}
	return str
}
