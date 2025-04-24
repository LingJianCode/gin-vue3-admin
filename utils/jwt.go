package utils

import (
	"errors"
	"math/rand"
	"my-ops-admin/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenValid            = errors.New("未知错误")
	ErrTokenExpired          = errors.New("token已过期")
	ErrTokenNotValidYet      = errors.New("token尚未激活")
	ErrTokenMalformed        = errors.New("这不是一个token")
	ErrTokenSignatureInvalid = errors.New("无效签名")
	ErrTokenInvalid          = errors.New("无法处理此token")
)

type MyCustomClaims struct {
	UserID   uint
	Username string
	Nickname string
	jwt.RegisteredClaims
}

// 签名密钥
const sign_key = "A3HbPp8vwz"

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func GenerateTokenUsingHs256(user *models.SysUser) (token string, claims MyCustomClaims, err error) {
	claims = MyCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MY_OPS_ADMIN",                                        // 签发者
			Subject:   user.Username,                                         // 签发对象
			Audience:  jwt.ClaimStrings{"webbrowser"},                        // 签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),         // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(-10 * time.Second)), // 最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                        // 签发时间
			ID:        randStr(10),                                           // wt ID, 类似于盐值
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(sign_key))
	return
}

func ParseTokenHs256(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sign_key), nil //返回签名密钥
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, ErrTokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, ErrTokenSignatureInvalid
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, ErrTokenNotValidYet
		default:
			return nil, ErrTokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, ErrTokenValid
}

// 从请求头中获取token
func ExtractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetUserID(c *gin.Context) (uint, error) {
	if claims, exists := c.Get("claims"); !exists {
		return 0, errors.New("从上下文获取用户信息失败")
	} else {
		waitUse := claims.(*MyCustomClaims)
		return waitUse.UserID, nil
	}
}
