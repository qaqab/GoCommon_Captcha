package GoCommon_Captcha

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_Captcha"
)

func TestGenerateCaptcha(t *testing.T) {
	var CaptchaStruct = GoCommon_Captcha.NewCaptcha()
	id, b64s, answer, err := CaptchaStruct.CaptMake()
	fmt.Println(id)
	fmt.Println(b64s)
	fmt.Println(answer)
	fmt.Println(err)
	fmt.Println(CaptchaStruct.CaptVerify(id, answer))

}

func TestCheckCaptcha(t *testing.T) {
	var CaptchaStruct = GoCommon_Captcha.NewCaptcha()

	id := "mqu10BbqAFxESavQZO9E"
	answer := "v9wl"
	if CaptchaStruct.CaptVerify(id, answer) {
		fmt.Println("OK")
	} else {
		fmt.Println("NO")
	}
}
