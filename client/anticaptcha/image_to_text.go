package anticaptcha

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response/task"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response/task/status"

	taskType "github.com/sitnikovik/go-anti-captcha/internal/anticaptcha/task_type"

	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"

	http2 "github.com/sitnikovik/go-anti-captcha/internal/http"
)

// imageToTextRequestInterval is an interval to do request to AntiCaptcha API
const imageToTextRequestInterval = 1 * time.Second

// ImageToText resolves image captcha and returns solution
func (c *client) ImageToText(bb []byte, timeoutInterval time.Duration) (*task.Task, error) {
	var err error

	body := struct {
		ClientKey string `json:"clientKey"`
		Task      struct {
			Type      string `json:"type"`
			Body      string `json:"body"`
			Phrase    bool   `json:"phrase"`
			Case      bool   `json:"case"`
			Numeric   bool   `json:"numeric"`
			Math      bool   `json:"math"`
			MinLength int    `json:"minLength"`
			MaxLength int    `json:"maxLength"`
		} `json:"task"`
	}{}
	body.ClientKey = c.token
	body.Task.Type = taskType.ImageToText
	body.Task.Body = base64.StdEncoding.EncodeToString(bb)

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, baseUrl+EndpointCreateTask, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var taskResult task.Task

	if err = http2.SendAndDecode(req, &taskResult); err != nil {
		return nil, err
	}

	if taskResult.ErrorID != 0 {
		return nil, errors.ByErrorID(taskResult.ErrorID)
	}

	if timeoutInterval <= 0 {
		return &taskResult, nil
	}

	ticker := time.NewTicker(imageToTextRequestInterval)
	timer := time.NewTimer(timeoutInterval)
	var tmpTask *task.Task
	for {
		select {
		case <-ticker.C:
			tmpTask, err = c.GetTaskByID(taskResult.TaskID)
			if err != nil {
				return nil, err
			}
			if tmpTask.Status == status.Ready {
				taskResult = *tmpTask
				ticker.Stop()

				return &taskResult, nil
			}
		case <-timer.C:
			return nil, errors.ErrResponseTimeout
		}
	}
}
