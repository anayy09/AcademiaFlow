package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/anayy09/academiaflow-backend/configs"
)

type Claims struct {
    UserID   uint   `json:"user_id"`
    Email    string `json:"email"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func GenerateToken(userID uint, email, username string, config *configs.Config) (string, error) {
    claims := Claims{
        UserID:   userID,
        Email:    email,
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWT.ExpiresIn)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JWT.Secret))
}

func ValidateToken(tokenString string, config *configs.Config) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.JWT.Secret), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}