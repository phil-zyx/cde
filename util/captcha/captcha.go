package captcha

import (
	"github.com/cde/util"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type Result struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 生成验证码
func GenerateCapcha(ctx *gin.Context) {
	parameters := base64Captcha.ConfigDigit{
		Height:     30,
		Width:      60,
		CaptchaLen: 4,
		MaxSkew:    0,
		DotCount:   2,
	}
	captchaId, captchaInstance := base64Captcha.GenerateCaptcha("", parameters)
	base64Blob := base64Captcha.CaptchaWriteToBase64Encoding(captchaInstance)
	captchaResult := &Result{
		Id:         captchaId,
		Base64Blob: base64Blob,
	}
	util.Success(ctx, "成功", captchaResult)
}

// 验证
func VerfiyCaptcha(idkey, verifyValue string) bool {
	return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}