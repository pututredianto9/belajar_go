package jwtAuth

import (
	"belajar/package/config"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt interface {
	GenerateToken(data *JwtData) (string, string, error)
	VerifyAccessToken(token string) (*JwtData, error)
	VerifyRefreshToken(token string) (string, error)
}

type Options struct {
	signingKey           string
	issuer               string
	accessTokenDuration  int
	refreshTokenDuration int
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.signingKey = cfg.JwtSigningKey
	opt.issuer = cfg.JwtIssuer
	opt.accessTokenDuration = cfg.JwtAccessTokenDuration
	opt.refreshTokenDuration = cfg.JwtRefreshTokenDuration

	return opt
}

// GenerateToken implements Jwt
func (o *Options) GenerateToken(data *JwtData) (string, string, error) {
	data.StandardClaims.ExpiresAt = time.Now().Local().Add(time.Second * time.Duration(o.accessTokenDuration)).Unix()
	acToken := jwt.NewWithClaims(jwt.SigningMethodHS512, data)
	accessToken, err := acToken.SignedString([]byte(o.signingKey))
	if err != nil {
		return "", "", err
	}

	data.StandardClaims.ExpiresAt = time.Now().Local().Add(time.Second * time.Duration(o.refreshTokenDuration)).Unix()
	rfToken := jwt.NewWithClaims(jwt.SigningMethodHS512, data.StandardClaims)
	refreshToken, err := rfToken.SignedString([]byte(o.signingKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// VerifyRefreshToken implements Jwt
func (o *Options) VerifyRefreshToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(o.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", err
	}

	return claims["jti"].(string), nil
}

// VerifyAccessToken implements Jwt
func (o *Options) VerifyAccessToken(token string) (*JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(o.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}

	jwtData := &JwtData{
		UserID:   claims["uid"].(string),
		UserName: claims["uname"].(string),
		Role:     claims["rol"].(string),
	}

	return jwtData, nil
}
