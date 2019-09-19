package jwt

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

// AppClaims represent the claims parsed from JWT access token.
type AppClaims struct {
	ID     int
	Sub    string
	Scopes []string
}

// ParseClaims parses JWT claims into AppClaims.
func (c *AppClaims) ParseClaims(claims jwt.MapClaims) error {
	id, ok := claims["id"]
	if !ok {
		return errors.New("could not parse claim id")
	}
	c.ID = int(id.(float64))

	sub, ok := claims["sub"]
	if !ok {
		return errors.New("could not parse claim sub")
	}
	c.Sub = sub.(string)

	scp, ok := claims["scp"]
	if !ok {
		return errors.New("could not parse claims roles")
	}

	var scopes []string
	if scp != nil {
		for _, v := range scp.([]interface{}) {
			scopes = append(scopes, v.(string))
		}
	}
	c.Scopes = scopes

	return nil
}
