package utils

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"strings"
)

type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

func CreateCode() (string, string) {
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 30,
		Width:  60,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	fmt.Println(idKeyC)
	return strings.Replace(base64stringC, "data:image/png;base64,", "", -1), idKeyC
}

func CheckCode(code string, idKey string) bool {
	return base64Captcha.VerifyCaptcha(idKey, code)
}
