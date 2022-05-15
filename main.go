package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	URL := "https://www.cannondale.com/ja-jp/bikes"
	//URL := "https://www.cannondale.com/ja-jp"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(URL),
		//chromedp.WaitVisible(`body > footer`),
		//chromedp.Click(`#MainNavigation > button`,chromedp.NodeVisible),
		//chromedp.Click(`#MainNavigation > div > div:nth-child(1) > ul.level-one.nav-with-subnavs > li:nth-child(1) > a`, chromedp.NodeVisible),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Text(`//*[@id="ProductGrid"]/div[2]/div/div[1]/div[1]/h3`, &res, chromedp.NodeVisible),
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
