package misc

import "os"

// Getenv works as Getenv but with default value d if not exit or empty.
func Getenv(key string, d ...string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	if len(d) != 0 {
		return d[0]
	}
	return ""
}
