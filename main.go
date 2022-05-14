package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

func main() {
	//URL := "https://www.cannondale.com/ja-jp/bikes/road"
	URL := "https://www.cannondale.com/ja-jp"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		chromedp.Click(`body > main > header:nth-child(1) > div.billboard-centered-lower > div > a`, chromedp.NodeVisible),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Text(`#BikeConfiguration > div.bike-configuration__inner > h1`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))
	log.Println(res)
}
