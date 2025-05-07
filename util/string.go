package util

import (
	"regexp"
	"strings"
)

// converts struct field name ex: PublishedYear to published_year
func CamelToSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")

	return strings.ToLower(snake)
}
