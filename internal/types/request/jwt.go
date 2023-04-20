package request

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go-template/global"
	"go-template/pkg/times"
	"time"
)

type BaseClaims struct {
	Uuid string
	ID   uint
	Name string
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(viper.GetString("jwt.signing-key")),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := times.ParseDuration(viper.GetString("jwt.expires-times"))
	ep, _ := times.ParseDuration(viper.GetString("jwt.buffer-times"))
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			// 签名生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			// 过期时间 7天  配置文件
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			// 签名的发行者
			Issuer: viper.GetString("jwt.issuer"),
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) RefreshToken(token string, claims CustomClaims) (string, error) {
	var err error
	dr, err := times.ParseDuration(viper.GetString("jwt.buffer-time"))
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
	newToken, err := j.CreateTokenByOldToken(token, claims)
	if err != nil {
		return "", err
	}
	_, err = j.ParseToken(newToken)
	if err != nil {
		return "", err
	}
	return newToken, err
}
