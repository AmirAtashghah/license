package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"server/logger"
	"time"
)

const group = "jwt"

type Config struct { //
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
}

type Service struct {
	cfg Config
}

func New(cfg Config) *Service {
	return &Service{cfg: cfg}
}

type CustomClaims struct {
	UserID string `json:"userID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (j *Service) GenerateJWTAccessToken(userID, role string) (string, error) {
	claims := &CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)), // todo change expire time
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenString.SignedString([]byte(j.cfg.JwtSecretKey))
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		return "", err
	}

	return token, nil
}

func (j *Service) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.L().WithGroup(group).Error("error", "error", jwt.ErrSignatureInvalid.Error())

			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.cfg.JwtSecretKey), nil
	})
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.L().WithGroup(group).Error("error", "error", jwt.ErrTokenInvalidClaims.Error())

		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func (j *Service) GetClaimsValue(claims jwt.MapClaims) map[string]string {
	m := make(map[string]string)
	userID := claims["userID"]
	role := claims["role"]

	userIDStr, ok := userID.(string)
	if !ok {
		logger.L().WithGroup(group).Error("error", "error", "can not convert userId to string")

		return nil
	}
	roleStr, ok := role.(string)
	if !ok {
		logger.L().WithGroup(group).Error("error", "error", "can not convert userId to string")

		return nil
	}

	m["userID"] = userIDStr
	m["role"] = roleStr

	return m
}
