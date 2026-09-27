package table

import (
	"fmt"
	"regexp"
)

var uuidRegex = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)

func formatUuuid(s any) any {
	asString := fmt.Sprintf("%v", s)
	if !uuidRegex.MatchString(asString) {
		return s
	}
	return asString[:4] + "..." + asString[len(asString)-4:]
}
