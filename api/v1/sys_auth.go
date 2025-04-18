package v1

import (
	"my-ops-admin/global"
	"my-ops-admin/model/response"
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
	captchaKey, captchaBase64, _, err := cp.Generate()
	if err != nil {
		global.OPS_LOG.Error("获取验证码失败", zap.Error(err))
		utils.FailWithMessage("获取验证码失败", c)
		return
	}
	global.OPS_LOG.Debug("获取验证码成功")
	utils.OkWithData(response.Captcha{CaptchaKey: captchaKey, CaptchaBase64: captchaBase64}, c)
}
