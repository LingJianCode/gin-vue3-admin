package response

type Captcha struct {
	CaptchaKey    string `json:"captchaKey"`
	CaptchaBase64 string `json:"captchaBase64"`
}

type LoginResponse struct {
	TokenType    string `json:"tokenType"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
}
