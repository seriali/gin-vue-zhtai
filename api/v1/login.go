package v1

import (
	"gin-vue-zhtai-server/middleware"
	"gin-vue-zhtai-server/model"
	"gin-vue-zhtai-server/model/request"
	"gin-vue-zhtai-server/model/response"
	"gin-vue-zhtai-server/utils/message"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func LoginController(c *gin.Context) {
	var L request.LoginParams
	_ = c.ShouldBindJSON(&L)
	code := model.CheckLogin(L.Username, L.Password)
	if code == message.SUCCESS {
		data := &model.User{Username: L.Username, Password: L.Password}
		if err, u := model.Login(data); err != nil {
			response.FailWithMessage(message.LoginFail, c)
		} else {
			setToken(c, *u)
		}

	} else {
		response.FailWithMessage(code, c)
	}
}

type DataResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}

// 设置token
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Usename: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "GinBlog",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(message.ERROR, c)
		/*c.JSON(http.StatusOK, gin.H{
			"status": message.ERROR,
			"message": message.GetMsg(message.ERROR),
			"token": token,
		})*/
	}
	res := DataResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}
	response.OkWithDetailed(message.SUCCESS, res, c)
	/*c.JSON(http.StatusOK, gin.H{
		"status": message.SUCCESS,
		"data": res,
		"message": message.GetMsg(message.SUCCESS),
	})*/
	return
}
