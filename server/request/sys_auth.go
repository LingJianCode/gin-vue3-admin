package request

type Login struct {
	Username    string `form:"username"`    // 用户名
	Password    string `form:"password"`    // 密码
	CaptchaKey  string `form:"captchaKey"`  // 验证码id
	CaptchaCode string `form:"captchaCode"` // 验证码answer
}
