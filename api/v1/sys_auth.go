package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/request"
	"my-ops-admin/response"
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
}
