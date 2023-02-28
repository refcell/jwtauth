package main

import (
	"time"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

// TestEmptyJwtClaimsConfig tests constructing empty jwt claims config
func TestEmptyJwtClaimsConfig(t *testing.T) {
	claimsWithoutExpiry := jwt.MapClaims{}
	if err := claimsWithoutExpiry.Valid(); err != nil {
		t.Fatalf("claims without expiry should be valid")
	}
}

// TestJwtClaimsConfigNotIssuedNow checks we should not be able to use a token before issuance.
func TestJwtClaimsConfigNotIssuedNow(t *testing.T) {
	curr := time.Now()
	exp := curr.Add(24 * time.Hour)
	claimsWithoutExpiry := jwt.MapClaims{
		"iat": &jwt.NumericDate{Time: curr},
		"exp": &jwt.NumericDate{Time: exp},
	}
	err := claimsWithoutExpiry.Valid();
	if err.Error() != "Token used before issued" {
		t.Fatalf("claims without expiry should be valid")
	}
}
