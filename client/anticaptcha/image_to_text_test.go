package anticaptcha

import (
	"os"
	"testing"
	"time"

	antiCaptchaErrors "github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"

	"github.com/sitnikovik/go-anti-captcha/internal/config"
)

// TestClient_ImageToText tests ImageToText method with AntiCaptcha API
func TestClient_ImageToText(t *testing.T) {
	pathToCaptcha := "../../img/captcha.jpg"

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

	t.Run("ok with waiting for solution", func(t *testing.T) {
		bb, err := os.ReadFile(pathToCaptcha)
		if err != nil {
			t.Errorf("os.ReadFile() error = %v, want nil", err)
			return
		}

		tkn := config.FromFile().AntiCaptcha.Token
		c := NewClient(tkn)
		task, err := c.ImageToText(bb, 10*time.Second)
		if err != nil {
			t.Errorf("ImageToText() error = %v, want nil", err)
			return
		}

		ans := "76447"
		if task.Solution.Text != ans {
			t.Errorf("ImageToText() task.Solution.Text = %s, want %s", task.Solution.Text, ans)
		}
	})

	t.Run("ok without waiting for solution", func(t *testing.T) {
		bb, err := os.ReadFile(pathToCaptcha)
		if err != nil {
			t.Errorf("os.ReadFile() error = %v, want nil", err)
			return
		}

		tkn := config.FromFile().AntiCaptcha.Token
		c := NewClient(tkn)
		task, err := c.ImageToText(bb, 0)
		if err != nil {
			t.Errorf("ImageToText() error = %v, want nil", err)
			return
		}

		if task.TaskID == 0 {
			t.Errorf("ImageToText() task.TaskID = %d, want not 0", task.TaskID)
			return
		}
	})
}
