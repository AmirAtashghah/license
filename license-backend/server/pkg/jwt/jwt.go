package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Config struct {
	JwtSecretKey string `koanf:"jwt_secret_key"`
}

type Service struct {
	cfg Config
}

func New(cfg Config) *Service {
	return &Service{cfg: cfg}
}

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (j *Service) GenerateJWTAccessToken(userId string) (string, error) {
	claims := &CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)), // todo change expire time
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenString.SignedString([]byte(j.cfg.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *Service) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.cfg.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func (j *Service) GetClaimsValue(claims jwt.MapClaims) map[string]string {
	m := make(map[string]string)
	userId := claims["userId"].(string)

	m["userId"] = userId

	return m
}
