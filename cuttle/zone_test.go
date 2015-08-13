package main

import (
	"testing"
)

func TestZone(t *testing.T) {
	var zone Zone
	var c1, c2 LimitController

	zone = *NewZone("*.github.com", true, "rps", 2)

	c1 = zone.GetController("www.github.com")
	c2 = zone.GetController("api.github.com")
	if c1 != c2 {
		t.Errorf("Shared zone should return shared controller.")
	}

	zone = *NewZone("*.github.com", false, "rps", 2)

	c1 = zone.GetController("www.github.com")
	c2 = zone.GetController("api.github.com")
	if c1 == c2 {
		t.Errorf("Non-shared zone should return individual controller.")
	}

	zone = *NewZone("*", false, "rps", 2)
	if !zone.MatchHost("github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*", "github.com", true)
	}

	zone = *NewZone("*.com", false, "rps", 2)
	if !zone.MatchHost("github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.com", "github.com", true)
	}
	if zone.MatchHost("github.org") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.com", "github.org", false)
	}

	zone = *NewZone("github.com", false, "rps", 2)
	if !zone.MatchHost("github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "github.com", "github.com", true)
	}
	if !zone.MatchHost("www.github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "github.com", "www.github.com", false)
	}

	zone = *NewZone("*.github.com", false, "rps", 2)
	if !zone.MatchHost("www.github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.github.com", "www.github.com", true)
	}
	if zone.MatchHost("github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.github.com", "github.com", false)
	}
	if zone.MatchHost("hubgit.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.github.com", "hubgit.com", false)
	}

	zone = *NewZone("*.*.github.com", false, "rps", 2)
	if !zone.MatchHost("x.www.github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.*.github.com", "x.www.github.com", true)
	}
	if !zone.MatchHost("www.github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.*.github.com", "www.github.com", false)
	}
	if zone.MatchHost("github.com") {
		t.Errorf("zone(%s).MatchHost(%s) should be %s", "*.*.github.com", "github.com", false)
	}
}
