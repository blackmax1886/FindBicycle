package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("www.cannondale.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".filter-and-sort", func(e *colly.HTMLElement) {
		fmt.Println("element found")
		fmt.Println(e.DOM.Html())
		//first := e.DOM.First()
		//card := e.DOM.Find("div > div > div:nth-child(1) > div.product-card__upper > div.card-inner")
		//fmt.Println(first,card)
		//link := e.ChildAttr("a.content.product.product-card__link","href")

		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//// Start scraping on https://hackerspaces.org
	c.Visit("https://www.cannondale.com/ja-jp/bikes/road")

}
