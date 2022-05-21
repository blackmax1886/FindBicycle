package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"strings"
	"time"
)

type product struct {
	Name, URL, Price string
}

func main() {
	URL := "https://www.cannondale.com/ja-jp/bikes"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	detailsSel := `//*[@class="product-details"]/h2`
	var nameNodes []*cdp.Node
	var priceNodes []*cdp.Node
	var urlNodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(URL),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Nodes(detailsSel, &nameNodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, nameNode := range nameNodes {
				err := dom.RequestChildNodes(nameNode.NodeID).WithDepth(-1).Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Sleep(time.Second),
		//chromedp.ActionFunc(func(ctx context.Context) error {
		//	printNodes(os.Stdout, nameNodes, "", "  ")
		//	return nil
		//}),

		chromedp.Nodes(detailsSel+`/following-sibling::div/div/span/text()`, &priceNodes),
		chromedp.Nodes(`//*[@class="content product product-card__link"]`, &urlNodes),
	},
	)
	if err != nil {
		log.Fatal(err)
	}

	var bikes []product
	for i, nameNode := range nameNodes {
		//printNodes(os.Stdout, nameNode.Children, "", "  ")

		bike := product{
			Name:  nameNode.Children[0].NodeValue,
			URL:   urlNodes[i].AttributeValue("href"),
			Price: priceNodes[i].NodeValue,
		}
		if nameNode.ChildNodeCount > 1 {
			bike.Name += nameNode.Children[1].Children[0].NodeValue
		}
		bikes = append(bikes, bike)
	}

	for _, bike := range bikes {
		log.Printf("\nModel Name: %s\nURL: %s\nPrice: %s", bike.Name, bike.URL, bike.Price)
	}
}

func printNodes(w io.Writer, nodes []*cdp.Node, padding, indent string) {
	// This will block until the chromedp listener closes the channel
	for _, node := range nodes {
		switch {
		case node.NodeName == "#text":
			fmt.Fprintf(w, "%s#text: %q\n", padding, node.NodeValue)
		default:
			fmt.Fprintf(w, "%s%s:\n", padding, strings.ToLower(node.NodeName))
			if n := len(node.Attributes); n > 0 {
				fmt.Fprintf(w, "%sattributes:\n", padding+indent)
				for i := 0; i < n; i += 2 {
					fmt.Fprintf(w, "%s%s: %q\n", padding+indent+indent, node.Attributes[i], node.Attributes[i+1])
				}
			}
		}
		if node.ChildNodeCount > 0 {
			fmt.Fprintf(w, "%schildren:\n", padding+indent)
			printNodes(w, node.Children, padding+indent+indent, indent)
		}
	}
}
