package core

import (
	"regexp"
	"strings"
)

func checkSubFuzzy(req, sub, domain string) bool {
	parts := strings.SplitN(req, ".", 2)
	if len(parts) == 2 && parts[1] != domain {
		return false
	}

	target_sub := parts[0]

	sub_regexpr_str := "^" + strings.ReplaceAll(sub, "*", "([^-._/]*)") + "$"

	sub_regexpr, err := regexp.Compile(sub_regexpr_str)
	if err != nil {
		return false
	}

	return sub_regexpr.MatchString(target_sub)
}

func checkActiveFuzzy(host string, activeHostnames []string) bool {
	for _, active := range activeHostnames {
		parts := strings.SplitN(active, ".", 2)
		if (len(parts) == 2 && checkSubFuzzy(host, parts[0], parts[1])) || host == active {
			return true
		}
	}

	return false
}

func combineFuzzy(original, origSubPattern, targetSubPattern, domain string) string {
	parts := strings.SplitN(original, ".", 2)
	origSub := parts[0]

	if !strings.ContainsRune(origSubPattern, '*') && (len(parts) == 2 && parts[1] != domain) {
		return combineHost(targetSubPattern, domain)
	}

	expr := regexp.MustCompile("^" + strings.ReplaceAll(origSubPattern, "*", "([^-._/]*)") + "$")
	matches := expr.FindStringSubmatch(origSub)
	if len(matches) <= 1 {
		return combineHost(targetSubPattern, domain)
	}

	for i := 1; i < len(matches); i++ {
		targetSubPattern = strings.Replace(targetSubPattern, "*", matches[i], 1)
	}

	return combineHost(targetSubPattern, domain)
}

func combineHost(sub string, domain string) string {
	if sub == "" {
		return domain
	}
	return sub + "." + domain
}

func stringExists(s string, sa []string) bool {
	for _, k := range sa {
		if s == k {
			return true
		}
	}
	return false
}

func intExists(i int, ia []int) bool {
	for _, k := range ia {
		if i == k {
			return true
		}
	}
	return false
}

func removeString(s string, sa []string) []string {
	for i, k := range sa {
		if s == k {
			return append(sa[:i], sa[i+1:]...)
		}
	}
	return sa
}

func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		ml := maxLen
		pre := s[:ml/2-1]
		suf := s[len(s)-(ml/2-2):]
		return pre + "..." + suf
	}
	return s
}
