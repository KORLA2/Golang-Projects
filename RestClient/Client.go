package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	URL = "https://youtube138.p.rapidapi.com"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HttpClient *http.Client
}

func NewClient(appKey string) *Client {
	return &Client{
		BaseURL: URL,
		apiKey:  appKey,
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type SuccessResponse struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v any) error {

	req = req.WithContext(ctx)
	fmt.Println("Sending Req", c)
	req.Header.Add("x-rapidapi-key", c.apiKey)
	req.Header.Add("x-rapidapi-host", c.BaseURL)
	req.Header.Set("content-type", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var errResponse ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errResponse); err != nil {
			return fmt.Errorf("error decoding error response: %v", err)
		}
	}

	fullResponse := SuccessResponse{
		Status: http.StatusOK,
		Data:   v,
	}

	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}
	return nil

}
