package GoCommon_Captcha

import (
	"fmt"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 设置自带的store

type CaptchaStruct struct {
	store         base64Captcha.Store
	CaptchaId     string `json:"CaptchaId"`
	CaptchaB64    string `json:"CaptchaB64"`
	CaptchaAnswer string `json:"CaptchaAnswer"`
}

func NewCaptcha() CaptchaStruct {
	var store = base64Captcha.DefaultMemStore
	return CaptchaStruct{store: store}
}

// 生成验证码
func (captchaStruct *CaptchaStruct) CaptMake() (id, b64s, answer string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, captchaStruct.store)
	lid, lb64s, answer, lerr := captcha.Generate()
	return lid, lb64s, answer, lerr
}

// 验证captcha是否正确
func (captchaStruct *CaptchaStruct) CaptVerify(id string, capt string) bool {
	fmt.Println("id:" + id)
	fmt.Println("capt:" + capt)
	//if store.Verify(id, capt, true) {
	if captchaStruct.store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
	//return
}
