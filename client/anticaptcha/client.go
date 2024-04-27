package anticaptcha

import "github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response/task"

// baseUrl base AntiCaptcha url to
const baseUrl = "https://api.anti-captcha.com"

// Client implements client to work with AntiCaptcha API
type Client interface {
	// GetBalance returns current AntiCaptcha balance
	GetBalance() (float64, error)

	// ImageToText resolves image captcha and returns solution
	ImageToText(bb []byte, waitForSolution bool) (*task.Task, error)

	// GetTaskByID returns task result by id
	GetTaskByID(id int) (*task.Task, error)
}

// client is AntiCaptcha API client
type client struct {
	token string // AntiCaptcha auth key
}

// NewClient makes new AntiCaptcha API client
func NewClient(token string) Client {
	return &client{token: token}
}
