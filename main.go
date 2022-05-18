package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
)

type product struct {
	URL, Price string
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

	var names []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(detailsSel+`/h2/text()`, &names)); err != nil {
		log.Fatalf("could not get product names : %v", err)
	}
	log.Println(len(names))
	//log.Printf("NodeValue : %s", names[0].NodeValue)
	//log.Printf("NodeValue : %s", names[1].NodeValue)

	var subnames []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(detailsSel+`/h2/span/text()`, &subnames)); err != nil {
		log.Fatalf("could not get product subnames : %v", err)
	}
	log.Println(len(subnames))
	//log.Printf("NodeValue : %s", subnames[0].NodeValue)
	//log.Printf("NodeValue : %s", subnames[1].NodeValue)

	var prices []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(detailsSel+`/div/div/span/text()`, &prices)); err != nil {
		log.Fatalf("could not get product prices")
	}
	log.Println(len(prices))
	//log.Printf("NodeValue : %s", prices[0].NodeValue)
	//log.Printf("NodeValue : %s", prices[1].NodeValue)

	var urls []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(detailsSel+`/parent::a`, &urls)); err != nil {
		log.Fatalf("could not get product urls : %v", err)
	}
	log.Println(len(urls))

	products := make(map[string]product)
	for i := 0; i < len(names); i++ {
		name := names[i].NodeValue + subnames[i].NodeValue
		products[name] = product{
			URL:   urls[i].NodeValue,
			Price: prices[i].NodeValue,
		}
	}

	for k, v := range products {
		log.Printf("product %s (%s): '%s'", k, v.URL, v.Price)
	}
}
