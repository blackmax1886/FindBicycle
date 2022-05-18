package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
)

type product struct {
	URL, Name, Price string
}

func main() {
	URL := "https://www.cannondale.com/ja-jp/bikes"
	//URL := "https://www.cannondale.com/ja-jp"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(URL),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Text(`//*[@id="ProductGrid"]/div[2]/div/div[1]/div[1]/h3`, &res, chromedp.NodeVisible),
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	detailsSel := `//*[@class="product-details"]/`

	//var details []*cdp.Node
	//if err := chromedp.Run(ctx,chromedp.Nodes(detailsSel+``,&details)); err!= nil {
	//	log.Fatalf("could not get product infomation : %v", err)
	//}
	//result := make(map[string]product)
	//for i:=0; i < len(details); i++ {
	//
	//}

	var urls []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(detailsSel+`/parent::a`, &urls)); err != nil {
		log.Fatalf("could not get product infomation : %v", err)
	}
	for i := 0; i < len(urls); i++ {
		log.Printf("product link found : %s", urls[i].AttributeValue("href"))
	}
}
