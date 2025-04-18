package securityUtils

import (
	"golang.org/x/exp/slices"
)

func IsAdmin(roles interface{}) bool {
	switch roles.(type) {
	case []int:
		r := roles.([]int)
		return slices.Contains(r, 1)
	case []string:
		r := roles.([]string)
		return slices.Contains(r, "admin")
	case int:
		return roles == 1
	default:
		return false
	}
}
