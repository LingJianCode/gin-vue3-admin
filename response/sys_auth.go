package response

import "my-ops-admin/models"

type Captcha struct {
	CaptchaKey    string `json:"captchaKey"`
	CaptchaBase64 string `json:"captchaBase64"`
}

type LoginResponse struct {
	User      models.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
