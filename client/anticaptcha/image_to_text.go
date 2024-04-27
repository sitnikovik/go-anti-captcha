package anticaptcha

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response/task"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response/task/status"
	"net/http"
	"sync"

	taskType "github.com/sitnikovik/go-anti-captcha/internal/anticaptcha/task_type"

	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"

	http2 "github.com/sitnikovik/go-anti-captcha/internal/http"
)

// ImageToText resolves image captcha and returns solution
func (c *client) ImageToText(bb []byte, waitForSolution bool) (*task.Task, error) {
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

	if !waitForSolution {
		return &taskResult, nil
	}

	n := 5
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var tmpTask *task.Task
		defer wg.Done()
		for n > 0 {
			tmpTask, err = c.GetTaskByID(taskResult.TaskID)
			if err != nil {
				return
			}

			if tmpTask.Status == status.Ready {
				taskResult = *tmpTask
				return
			}
			n--
		}
	}()
	wg.Wait()

	if taskResult.ErrorID != 0 || taskResult.Solution.Text == "" {
		return nil, errors.ByErrorID(taskResult.ErrorID)
	}

	return &taskResult, err
}
