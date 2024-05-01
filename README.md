# go-anti-captcha
Golang AntiCaptcha API client

## Usage

```go
package main

import (
	"fmt"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha"
	"github.com/sitnikovik/go-anti-captcha/client/anticaptcha/response"
	"os"
	"time"
)

func main() {
	bb, _ := os.ReadFile("path_to_image")

	c := anticaptcha.NewClient("your_api_key")
	task, err := c.ImageToText(bb, 10*time.Second) // will wait solution for 10 sec or return error
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	// Solution for captcha image
	fmt.Println(task.Solution.Text)
}
```