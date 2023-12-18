package jwt

import (
	"encoding/json"
	"fmt"
)

type SigningAlgorithm interface {
	Sign(payload string) []byte
	Name() string
}

func Generate(publicClaims PublicClaims, privateClaims map[string]interface{}, alg SigningAlgorithm) (string, error) {
	h := generateHeader(alg)
	headerJSON, err := json.Marshal(&h)
	if err != nil {
		return "", fmt.Errorf("failed to marshal header: %w", err)
	}
	encHeader := encodeBase64URL(headerJSON)

	claims := merge(publicClaims, privateClaims)
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("failed to marshal claims: %w", err)
	}
	encPayload := encodeBase64URL(payload)

	sign := alg.Sign(fmt.Sprintf("%s.%s", encHeader, encPayload))
	encSign := encodeBase64URL(sign)

	return fmt.Sprintf("%s.%s.%s", encHeader, encPayload, encSign), nil
}
