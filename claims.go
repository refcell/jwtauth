package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// NewJWTClaims create a new JWT Claims object with the provided time as the
// issued at time. The jwt secret MUST be 32 bytes (256 bits) as defined by the
// go-ethereum [engine-api auth spec].
//
// [engine-api auth spec]: https://github.com/ethereum/execution-apis/blob/main/src/engine/authentication.md
func NewJWTClaims(jwtsecret [32]byte, curr time.Time) (string, error) {
	// Create a new JWT token with the provided time as the issued at time.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": &jwt.NumericDate{Time: curr},
	})
	// Create the signed token string
	s, err := token.SignedString(jwtsecret[:])
	if err != nil {
		return "", fmt.Errorf("failed to create JWT token: %w", err)
	}
	// Return the token
	return s, nil
}

// NewNoneJWTClaims create a new JWT Claims object with the provided time as the
// issued at time. The jwt secret MUST be 32 bytes (256 bits) as defined by the
// go-ethereum [engine-api auth spec].
//
// # This method uses the none signing method which is not recommended for
//
// [engine-api auth spec]: https://github.com/ethereum/execution-apis/blob/main/src/engine/authentication.md
func NewNoneJWTClaims(jwtsecret [32]byte, curr time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"iat": &jwt.NumericDate{Time: curr},
	})
	// Create the signed token string
	s, err := token.SignedString(jwtsecret[:])
	if err != nil {
		return "", fmt.Errorf("failed to create JWT token: %w", err)
	}
	// Return the token
	return s, nil
}
