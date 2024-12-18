package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\".*password.*\"`)

	var cnt int
	for _, l := range lines {
		cnt += len(re.FindAllString(strings.ToLower(l), -1))
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end\-of\-line\w+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +\w+`)

	for i, l := range lines {
		matches := re.FindStringSubmatch(l)
		if len(matches) > 0 {
			match := matches[0]
			username := strings.Fields(match)[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", username, l)
		}
	}
	return lines
}
