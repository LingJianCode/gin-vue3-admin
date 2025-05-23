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

var AuthApiApp = new(SysAuthApi)

type SysAuthApi struct{}

var store = base64Captcha.DefaultMemStore

func (a *SysAuthApi) Captcha(c *gin.Context) {
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
	utils.SuccessWithData(response.Captcha{CaptchaKey: captchaKey, CaptchaBase64: captchaBase64}, c)
}

func (a *SysAuthApi) Login(c *gin.Context) {
	var l request.Login
	err := c.ShouldBind(&l)
	if err != nil {
		global.OPS_LOGGER.Error("登录参数异常", zap.Error(err))
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if gin.Mode() == gin.DebugMode || l.CaptchaKey != "" && l.CaptchaCode != "" && store.Verify(l.CaptchaKey, l.CaptchaCode, true) {
		u := &models.SysUser{Username: l.Username, Password: l.Password}
		user, err := service.AuthServiceApp.Login(u)
		if err != nil {
			global.OPS_LOGGER.Error("用户名不存在或者密码错误!", zap.Error(err))
			utils.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		tokenNext(c, *user)
		return
	}
	utils.FailWithMessage("验证码错误", c)
}

func tokenNext(c *gin.Context, user models.SysUser) {
	token, claims, err := utils.GenerateTokenUsingHs256(&user)
	if err != nil {
		global.OPS_LOGGER.Error("获取token失败!", zap.Error(err))
		utils.FailWithMessage("获取token失败", c)
		return
	}
	utils.SuccessWithDetailed(response.LoginResponse{
		TokenType:   "Bearer",
		AccessToken: token,
		ExpiresIn:   claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}

func (a *SysAuthApi) Logout(c *gin.Context) {
	utils.Success(c)
}
