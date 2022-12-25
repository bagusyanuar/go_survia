package lib

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	ErrBearerType      = errors.New("invalid bearer type")
	ErrSignInMethod    = errors.New("invalid signin method")
	ErrJWTClaims       = errors.New("invalid jwt claim")
	ErrJWTParse        = errors.New("invalid parse jwt")
	ErrNoAuthorization = errors.New("unauthorized")
	ErrInvalidPassword = errors.New("password did not match")
	ErrInvalidRole     = errors.New("role did not match")
	ErrBadRequest      = errors.New("bad request")
)

func ErrorSignIn(err error) Response {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: "user not found!",
		}
	} else if errors.Is(err, ErrInvalidPassword) {
		return Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: "password did not match",
		}
	} else if errors.Is(err, ErrInvalidRole) {
		return Response{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: "role did not match",
		}
	} else {
		return Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "error while sign in " + err.Error(),
		}
	}
}

func ErrorMessageValidation(err error) []map[string]interface{} {
	results := []map[string]interface{}{}
	for _, err := range err.(validator.ValidationErrors) {
		tmp := map[string]interface{}{
			"key":     err.Field(),
			"message": err.Tag(),
		}
		results = append(results, tmp)
	}
	return results
}

func ValidateRequest(r interface{}) (m []map[string]interface{}, err error) {
	v := validator.New()
	if e := v.Struct(r); e != nil {
		return ErrorMessageValidation(e), e
	}
	return []map[string]interface{}{}, nil
}
