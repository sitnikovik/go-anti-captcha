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
)

func main() {
	bb, _ := os.ReadFile("path_to_image")

	c := anticaptcha.NewClient("your_api_key")
	task, err := c.ImageToText(bb, true)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	// Solution for captcha image
	fmt.Println(task.Solution.Text)
}
```
> You may make request to solve with no waiting for solution.
> Just pass `false` as second argument to `ImageToText` method, 
> and it will return `Task` object with `TaskID` field
> that you may use to get solution later by calling `GetTaskByid` method.