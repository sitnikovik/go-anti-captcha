package anticaptcha

import (
	antiCaptchaErrors "github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"
	"testing"

	"github.com/sitnikovik/go-anti-captcha/internal/config"
)

// TestClient_GetBalance tests GetBalance method with AntiCaptcha API
func TestClient_GetBalance(t *testing.T) {
	t.Run("err no auth", func(t *testing.T) {
		c := NewClient("")
		_, err := c.GetBalance()
		if !antiCaptchaErrors.IsUnauthorized(err) {
			t.Errorf("GetTaskByID() error not about unauthorized (err: %v)", err)
			return
		}
	})

	t.Run("ok", func(t *testing.T) {
		tkn := config.FromFile().AntiCaptcha.Token
		c := NewClient(tkn)
		_, err := c.GetBalance()
		if err != nil {
			t.Errorf("GetBalance() error = %v, want nil", err)
			return
		}
	})
}
