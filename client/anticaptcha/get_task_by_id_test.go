package anticaptcha

import (
	antiCaptchaErrors "github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"
	"github.com/sitnikovik/go-anti-captcha/internal/config"
	"testing"
)

// TestClient_GetTaskById tests GetTaskByID method with AntiCaptcha API
func TestClient_GetTaskById(t *testing.T) {
	t.Run("err no auth", func(t *testing.T) {
		c := NewClient("")
		task, err := c.GetTaskByID(0)
		if !antiCaptchaErrors.IsUnauthorized(err) {
			t.Errorf("GetTaskByID() error not about unauthorized (err: %v)", err)
			return
		}
		if task != nil {
			t.Errorf("GetTaskByID() task = %v, want nil", task)
		}
	})

	t.Run("err task not found", func(t *testing.T) {
		tkn := config.FromFile().AntiCaptcha.Token
		c := NewClient(tkn)
		task, err := c.GetTaskByID(0)
		if !antiCaptchaErrors.IsTaskNotFound(err) {
			t.Errorf("GetTaskByID() error not about task not found (err: %v)", err)
			return
		}
		if task != nil {
			t.Errorf("GetTaskByID() task = %v, want nil", task)
		}
	})
}
