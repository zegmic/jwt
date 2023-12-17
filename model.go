package jwt

type header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type PublicClaims struct {
	Issuer   string
	Subject  string
	Audience string
	IssuedAt int
	Name     string
}
