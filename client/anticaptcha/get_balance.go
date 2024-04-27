package anticaptcha

import (
	"bytes"
	"encoding/json"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/errors"
	"net/http"

	http2 "github.com/sitnikovik/go-anti-captcha/internal/http"
)

// GetBalance returns current AntiCaptcha balance
func (c *client) GetBalance() (float64, error) {
	var err error

	body := struct {
		ClientKey string `json:"clientKey"`
	}{}
	body.ClientKey = c.token
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, baseUrl+EndpointGetBalance, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0.0, err
	}

	resp := struct {
		ErrorId int     `json:"errorId"`
		Balance float64 `json:"balance"`
	}{}
	err = http2.SendAndDecode(req, &resp)
	if err != nil {
		return 0.0, err
	}

	if resp.ErrorId != 0 {
		return 0.0, errors.ByErrorID(resp.ErrorId)
	}

	return resp.Balance, nil
}
