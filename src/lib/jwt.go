package lib

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var JWTSigninMethod = jwt.SigningMethodHS256
var JWTSignatureKey string = "ONLYGODKNOWS"

type JWT struct{}

type JWTClaims struct {
	jwt.StandardClaims
	Unique uuid.UUID `json:"unique"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
}

func (j JWT) GenerateToked(c JWTClaims) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: "survia-app",
		},
		Unique: c.Unique,
		Email:  c.Email,
		Role:   c.Role,
	}

	token := jwt.NewWithClaims(JWTSigninMethod, claims)

	signedToken, err := token.SignedString([]byte(JWTSignatureKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (JWT) Claim(auth string) (interface{}, error) {
	if auth == "" {
		return nil, ErrNoAuthorization
	}

	bearer := string(auth[0:7])
	token := string(auth[7:])

	if bearer != "Bearer " {
		return nil, ErrBearerType
	}

	v, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSignInMethod
		} else if method != JWTSigninMethod {
			return nil, ErrSignInMethod
		}
		return []byte(JWTSignatureKey), nil
	})

	if err != nil {
		return nil, ErrJWTParse
	}

	claim, ok := v.Claims.(jwt.MapClaims)
	if !ok || !v.Valid {
		return nil, ErrJWTClaims
	}
	return claim, nil
}
