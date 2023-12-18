package jwt

type PublicClaims struct {
	Issuer   string
	Subject  string
	Audience string
	IssuedAt int
	Name     string
}

func merge(publicClaims PublicClaims, privateClaims map[string]interface{}) map[string]interface{} {
	m := privateClaims
	if m == nil {
		m = map[string]interface{}{}
	}
	m["iss"] = publicClaims.Issuer
	m["name"] = publicClaims.Name
	m["aud"] = publicClaims.Audience
	m["iat"] = publicClaims.IssuedAt
	m["sub"] = publicClaims.Subject

	return m
}
