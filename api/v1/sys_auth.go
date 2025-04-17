package v1

import (
	"fmt"
	"my-ops-admin/model/response"
	"my-ops-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	imgHeight, imgWidth, keyLong := 60, 240, 4
	driver := base64Captcha.NewDriverDigit(imgHeight, imgWidth, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	captchaKey, captchaBase64, _, err := cp.Generate()
	if err != nil {
		fmt.Println(err)
		utils.FailWithMessage("获取验证码失败", c)
		return
	}
	utils.OkWithData(response.Captcha{CaptchaKey: captchaKey, CaptchaBase64: captchaBase64}, c)
}
