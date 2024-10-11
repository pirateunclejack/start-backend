package helper

import (
	"fmt"
	"log"
	"start-backend/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAllToken(user *model.User) string {
    JWT_SECRET := GetConfig().GetString("JWT_SECRET")
    if JWT_SECRET == "" {
        log.Printf("failed to get JWT_SECRET from config file")
        return ""
    }

    claims := model.UserClaims{
        Username: user.Username,
        Email: user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
		    NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(JWT_SECRET))
    if err != nil {
        log.Printf("failed to sign token: %v\n", err)
    }
    return tokenString
}

func ValidateToken(signed_token string) (claims *model.UserClaims, msg string) {
    JWT_SECRET := GetConfig().GetString("JWT_SECRET")
    if JWT_SECRET == "" {
        err := "failed to get JWT_SECRET from config file"
        return nil, err
    }

    token, err := jwt.ParseWithClaims(
        signed_token,
        &model.UserClaims{},
        func(token *jwt.Token) (interface{}, error){
            return []byte(JWT_SECRET), nil
        },
    )
    if err != nil {
        msg = err.Error()
        return
    }

    claims, ok := token.Claims.(*model.UserClaims)
    if !ok {
        msg = fmt.Sprintln("token is not valid")
        return
    }

    if claims.RegisteredClaims.ExpiresAt.Before(time.Now()) {
		msg = fmt.Sprintln("token is expired")
		return
	}

	return claims, msg
}
