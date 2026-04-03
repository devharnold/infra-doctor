package api

import (
	"net/http"
	"errors"
	"time"
)

func Check(url string) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return errors.New("non-200 response")
	}

	return nil
}