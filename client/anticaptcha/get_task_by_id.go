package anticaptcha

import (
	"bytes"
	"encoding/json"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response/task"
	"net/http"

	http2 "github.com/sitnikovik/go-anti-captcha/internal/http"
)

// GetTaskByID returns task result by id
func (c *client) GetTaskByID(id int) (*task.Task, error) {
	var err error

	body := struct {
		ClientKey string `json:"clientKey"`
		TaskID    int    `json:"taskId"`
	}{}
	body.ClientKey = c.token
	body.TaskID = id
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, baseUrl+EndpointGetTaskResult, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var t task.Task
	if err = http2.SendAndDecode(req, &t); err != nil {
		return nil, err
	}

	if t.ErrorID != 0 {
		return nil, errors.ByErrorID(t.ErrorID)
	}

	return &t, nil
}
