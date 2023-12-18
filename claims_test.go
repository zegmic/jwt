package jwt

import "testing"

func TestMergeNilClaims(t *testing.T) {
	c := PublicClaims{
		Issuer:   "wm",
		Subject:  "zegmic",
		Audience: "google",
		IssuedAt: 123,
		Name:     "Michal",
	}

	m := merge(c, nil)
	if len(m) != 5 {
		t.Error("only public claims are expected ")
	}
}

func TestMergeCustomClaims(t *testing.T) {
	c := PublicClaims{
		Issuer:   "wm",
		Subject:  "zegmic",
		Audience: "google",
		IssuedAt: 123,
		Name:     "Michal",
	}

	p := map[string]interface{}{
		"root": true,
	}

	m := merge(c, p)
	if len(m) != 6 {
		t.Error("public and custom claims are expected ")
	}
	if _, ok := m["root"]; !ok {
		t.Error("merged claims should contain custom claims")
	}
}

func TestMergeConflictingClaims(t *testing.T) {
	c := PublicClaims{
		Issuer:   "wm",
		Subject:  "zegmic",
		Audience: "google",
		IssuedAt: 123,
		Name:     "Michal",
	}

	p := map[string]interface{}{
		"aud": "Fake audience",
	}

	m := merge(c, p)
	if m["aud"] != "google" {
		t.Error("public claim expected to have a precedence over custom")
	}
}
