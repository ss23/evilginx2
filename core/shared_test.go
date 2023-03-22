package core

import "testing"

func TestCheckSubFuzzy(t *testing.T) {
	if !checkSubFuzzy("sub.domain.com", "sub", "domain.com") {
		t.Error("Exact subdomains should pass")
	}

	if checkSubFuzzy("sub.otherdomain.com", "sub", "domain.com") {
		t.Error("Other domains should not pass")
	}

	if checkSubFuzzy("sub2.domain.com", "sub", "domain.com") {
		t.Error("Other subdomains should not pass")
	}

	if !checkSubFuzzy("sub-123.domain.com", "sub-*", "domain.com") {
		t.Error("Fuzzy match should pass")
	}

	if checkSubFuzzy("sub-123.otherdomain.com", "sub-*", "domain.com") {
		t.Error("Other domains should not pass")
	}

	if checkSubFuzzy("asub-3.domain.com", "sub-*", "domain.com") {
		t.Error("Other subdomains should not pass")
	}

	if !checkSubFuzzy("abc123.domain.com", "*", "domain.com") {
		t.Error("Full wildcard should work")
	}
}

func TestCheckActiveFuzzy(t *testing.T) {
	if !checkActiveFuzzy("test", []string{"test"}) {
		t.Error("Simple equal case should pass")
	}

	if checkActiveFuzzy("test2", []string{"test"}) {
		t.Error("Simple inequal case should not pass")
	}

	if !checkActiveFuzzy("sub.test.com", []string{"sub.test.com"}) {
		t.Error("Exact equal case should pass")
	}

	if checkActiveFuzzy("sub2.test.com", []string{"sub.test.com"}) {
		t.Error("Exact inequal case should not pass")
	}

	if !checkActiveFuzzy("sub-123.test.com", []string{"sub-*.test.com"}) {
		t.Error("Fuzzy equal case should pass")
	}

	if checkActiveFuzzy("blah.test.com", []string{"sub-*.test.com"}) {
		t.Error("Fuzzy inequal case should not pass")
	}
}

func TestCombineFuzzy(t *testing.T) {
	if combineFuzzy("orig-sub.domain.com", "orig-sub", "phish-sub", "gone.phishing.com") != "phish-sub.gone.phishing.com" {
		t.Error("Simple replacement should work")
	}

	if combineFuzzy("orig-test123.domain.com", "orig-*", "phish-*", "gone.phishing.com") != "phish-test123.gone.phishing.com" {
		t.Error("Fuzzy replacement should work")
	}
}
