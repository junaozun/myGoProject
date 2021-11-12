package middleware

import (
	"net/http"
	"strings"
	"time"
	"weblogindemo/errno"
	"weblogindemo/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(userName string) (string, errno.RET) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			// 签发人
			Issuer: "webLoginDemo",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errno.ERROR
	}
	return token, errno.OK

}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errno.OK.RetCode()
	} else {
		return nil, errno.ERROR.RetCode()
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 规范
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusOK, errno.ERROR_TOKEN_Not_EXIST)
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, errno.ERROR_TOKEN_TYPE_ERR)
			c.Abort()
			return
		}
		key, err := CheckToken(checkToken[1])
		if err == errno.ERROR.RetCode() {
			c.JSON(http.StatusOK, errno.ERROR_TOKEN_WRONG)
			c.Abort()
			return
		}
		// 判断token是否过期了
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(http.StatusOK, errno.ERROR_TOKEN_RUNTIME)
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}

}
