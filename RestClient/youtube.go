package main

import (
	"context"

	"net/http"
)

type ChannelVideos struct {
	Contents []struct {
		Title      string      `json:"title,omitempty"`
		VideoID    string      `json:"videoId,omitempty"`
		Stats      Views       `json:"stats"`
		Thumbnails []Thumbnail `json:"thumbnails,omitempty"`
	} `json:"contents,omitempty"`
}

type Views struct {
	Views int `json:"views,omitempty"`
}

type Thumbnail struct {
	Height int    `json:"height,omitempty"`
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
}

func (c *Client) GetThumbnail(ctx context.Context) (*Thumbnail, error) {

	req, err := http.NewRequest("GET", c.BaseURL+"/channel/videos/?id=UCJ5v_MCY6GNUBTO8-D3XoAg&filter=videos_latest&hl=en&gl=US", nil)
	if err != nil {

		return nil, err
	}
	Thumb := Thumbnail{}
	if err := c.sendRequest(ctx, req, &Thumb); err != nil {
		return nil, err
	}
	return &Thumb, nil

}
