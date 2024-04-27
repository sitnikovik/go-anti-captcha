package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SendAndDecode makes request and returns decoded response or error
func SendAndDecode(r *http.Request, v interface{}) error {
	c := http.Client{}
	var err error

	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, v); err != nil {
		return err
	}

	return nil
}
