package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
	"my-ops-admin/service"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	imgHeight, imgWidth, keyLong := 60, 240, 4
	driver := base64Captcha.NewDriverDigit(imgHeight, imgWidth, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	captchaKey, captchaBase64, answer, err := cp.Generate()
	if err != nil {
		global.OPS_LOGGER.Error("获取验证码失败", zap.Error(err))
		utils.FailWithMessage("获取验证码失败", c)
		return
	}
	global.OPS_LOGGER.Debug("获取验证码成功", zap.String("captchaKey", captchaKey), zap.String("answer", answer))
	utils.OkWithData(response.Captcha{CaptchaKey: captchaKey, CaptchaBase64: captchaBase64}, c)
}

func Login(c *gin.Context) {
	var l request.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		global.OPS_LOGGER.Error("登录参数异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if l.CaptchaKey != "" && l.CaptchaCode != "" && store.Verify(l.CaptchaKey, l.CaptchaCode, true) {
		u := &models.SysUser{Username: l.Username, Password: l.Password}
		user, err := service.Login(u)
		if err != nil {
			global.OPS_LOGGER.Error("用户名不存在或者密码错误!", zap.Error(err))
			utils.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		TokenNext(c, *user)
		return
	}
	utils.FailWithMessage("验证码错误", c)
}

func TokenNext(c *gin.Context, user models.SysUser) {
	token, claims, err := utils.GenerateTokenUsingHs256(&user)
	if err != nil {
		global.OPS_LOGGER.Error("获取token失败!", zap.Error(err))
		utils.FailWithMessage("获取token失败", c)
		return
	}
	utils.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}

func Logout(c *gin.Context) {
	utils.Ok(c)
}
