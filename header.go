package jwt

type header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

func generateHeader(alg SigningAlgorithm) header {
	return header{
		Algorithm: alg.Name(),
		Type:      "JWT",
	}
}
