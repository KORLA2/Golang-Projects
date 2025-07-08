package main

import (
	"context"
	"fmt"
)

func main() {

	client := NewClient("28eb0e2f12mshc4cf2828411c787p1db064jsn5b2a21c91dbb")
	ctx := context.Background()
	res, err := client.GetThumbnail(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
