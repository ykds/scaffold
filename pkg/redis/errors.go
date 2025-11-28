package redis

import "strings"

func IsNotFound(err error) bool {
	return strings.Contains(err.Error(), "redis: nil")
}
