package db

import "regexp"

func isUUID(input string) bool {
	// UUID pattern with hyphens
	uuidPattern := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

	return uuidPattern.MatchString(input)
}
