package request

type Login struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	CaptchaKey  string `json:"captchaKey"`  // 验证码id
	CaptchaCode string `json:"captchaCode"` // 验证码answer
}
