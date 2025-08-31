package user

import (
	"os"
	"testing"
	"time"

	"github.com/7erryX/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/imroc/req/v3"
)

func TestGetCaptchaImg(t *testing.T) {
	c := req.C().ImpersonateChrome().SetTimeout(10 * time.Second)
	img, _, err := GetCapchaImage(c)
	if err != nil {
		t.Errorf("GetCapchaImage() failed: %v", err)
		return
	}

	t.Logf("GetCapchaImage() success: img size=%d", len(img))
	os.WriteFile("captcha.jpg", img, 0644)
}

func TestGetProfile(t *testing.T) {
	c := req.C().ImpersonateChrome().SetTimeout(10 * time.Second)
	c.SetCommonBearerAuthToken(config.TOKEN)
	_, err := GetProfile(c)
	if err != nil {
		t.Errorf("GetProfile() failed: %v", err)
		return
	}
}
