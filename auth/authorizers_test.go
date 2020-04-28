package auth

import (
	"testing"
)

func TestParseAndCheckToken(t *testing.T) {
	privateKey, publicKey := New()

	if token, err := NewToken(NewRoleClaims([]string{"admin"}), privateKey); err != nil {
		t.Fatal(err)
	} else {
		if claims, err := ParseAndCheckToken(token, publicKey); err != nil {
			t.Fatal(err)
		} else {
			t.Log(claims.String())
		}
	}
}
