package jwt_test

import (
	"encoding/base64"
	jwt "github.com/jamesbartlettwm/wm-test-michal-zeglarski"
	"strings"
	"testing"
)

func TestGeneratePublicClaims(t *testing.T) {
	alg := jwt.NewHMACSHA256Signer("s3cr37")
	claims := jwt.PublicClaims{
		Issuer:   "wm",
		Subject:  "zegmic",
		Audience: "external",
		IssuedAt: 1516239022,
		Name:     "Michal",
	}
	token, err := jwt.Generate(claims, nil, alg)
	if err != nil {
		t.Errorf("failed generating a token: %v", err)
	}
	exp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleHRlcm5hbCIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoid20iLCJuYW1lIjoiTWljaGFsIiwic3ViIjoiemVnbWljIn0.wIFjCNn860PmOVY2DxRjmBD--CLgFcfJ6lJ_aF3fR6o"
	if token != exp {
		t.Errorf("invalid token. Expected %s got %s", exp, token)
	}
}

func TestGenerateCustomClaims(t *testing.T) {
	alg := jwt.NewHMACSHA256Signer("s3cr37")
	claims := jwt.PublicClaims{
		Issuer:   "wm",
		Subject:  "zegmic",
		Audience: "external",
		IssuedAt: 1516239022,
		Name:     "Michal",
	}
	priv := map[string]interface{}{
		"root": true,
	}
	token, err := jwt.Generate(claims, priv, alg)
	if err != nil {
		t.Errorf("failed generating a token: %v", err)
	}
	exp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleHRlcm5hbCIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoid20iLCJuYW1lIjoiTWljaGFsIiwicm9vdCI6dHJ1ZSwic3ViIjoiemVnbWljIn0.zQWbyyECGqrjwytQRY1yhm4saVCqewfrxaoyqS4aSMU"
	if token != exp {
		t.Errorf("invalid token. Expected %s got %s", exp, token)
	}
}

func TestGenerateConflictingClaims(t *testing.T) {
	alg := jwt.NewHMACSHA256Signer("s3cr37")
	claims := jwt.PublicClaims{
		Issuer:   "wm",
		Subject:  "zegmic",
		Audience: "external",
		IssuedAt: 1516239022,
		Name:     "Michal",
	}
	priv := map[string]interface{}{
		"iss": "Fake Issuer",
	}
	token, err := jwt.Generate(claims, priv, alg)
	if err != nil {
		t.Errorf("failed generating a token: %v", err)
	}

	parts := strings.Split(token, ".")
	payload, err := decode(parts[1])
	if err != nil {
		t.Errorf("failed decoding a token: %v", err)
	}

	if strings.Contains(string(payload), "Fake Issuer") {
		t.Errorf("public claim should not be overwritten by private claim")
	}
}

func decode(data string) ([]byte, error) {
	data = strings.Replace(data, "-", "+", -1)
	data = strings.Replace(data, "_", "/", -1)

	switch len(data) % 4 {
	case 0:
	case 2:
		data += "=="
	case 3:
		data += "="
	}

	return base64.StdEncoding.DecodeString(data)
}
