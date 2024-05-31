package toggltrack

import (
	"regexp"
)

type issue struct {
	Number      string
	Description string
	IsValid     bool
}

func indexExists(s []string, i int) bool {
	return i >= 0 && i < len(s)
}

func newIssue(apiDesc string) issue {
	i := issue{}

	var re = regexp.MustCompile(`^#?([0-9]+)(?: - )?(.*)$`)
	m := re.FindStringSubmatch(apiDesc)
	i.IsValid = len(m) > 0

	if !i.IsValid {
		return i
	}

	if indexExists(m, 1) && m[1] != "" {
		i.Number = m[1]
	}

	if indexExists(m, 2) && m[2] != "" {
		i.Description = m[2]
	}

	return i
}
