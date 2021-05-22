package middleware

import (
	"errors"
	"gin-vue-zhtai-server/model/response"
	"gin-vue-zhtai-server/utils"
	"gin-vue-zhtai-server/utils/message"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	Usename string `json:"usename"`
	jwt.StandardClaims
}

var (
	TokenExpired     error = errors.New("token已过期，请重新登录")
	TokenNotValidYet error = errors.New("token无效，请重新登录")
	TokenMalFormed   error = errors.New("token不正确, 请重新登录")
	TokenInvalid     error = errors.New("这不是一个token, 请重新登录")
)

//创建token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalFormed
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
		if claims, ok := token.Claims.(*MyClaims); ok {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这里jwt鉴权取头部信息 x-token 登录时返回token信息
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithMessage(message.TokenNotExit, c)
			c.Abort()
			return
		}
		checkToken := strings.Split(token, " ")
		if len(checkToken) == 0 {
			response.FailWithMessage(message.TokenTypeWrong, c)
			c.Abort()
			return
		}
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			response.FailWithMessage(message.TokenTypeWrong, c)
			c.Abort()
			return
		}
		j := NewJWT()
		// ParseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithMessage(message.TokenRunTime, c)
				c.Abort()
				return
			}
			// 其他错误
			response.FailWithMessage(message.ERROR, c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
