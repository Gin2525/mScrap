package scraper

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	// MercariSerchBaseURL means seach page in mercari
	MercariSerchBaseURL string = "https://www.mercari.com/jp/search/"
)

// MercariItem means item produced
type MercariItem struct {
	productName string
	price       int
}

func (item MercariItem) String() string {
	price := strconv.Itoa(item.price)
	return "productname: " + item.productName + "\nprice: " + price + "\n\n"
}

// SearchURL have material to render url
type SearchURL struct {
	Keyword string
	Queries map[string]string
}

func buildURLStructure(keyword string, queries map[string]string) *SearchURL {
	return &SearchURL{keyword, queries}
}
// RenderURL render url mercari search page
func (urlStructure SearchURL) RenderURL() string {
	url := MercariSerchBaseURL + "?keyword=" + urlStructure.Keyword + "&"
	for k, v := range urlStructure.Queries {
		url += k + "=" + v + "&"
	}
	return url[:len(url)-1]
}

// FetchMercariItems can fetch mercari item by keyword and parameter queries
func FetchMercariItems(urlStructure SearchURL) []MercariItem {
	url := urlStructure.RenderURL()
	res, _ := http.Get(url)
	// read
	buf, _ := ioutil.ReadAll(res.Body)
	bReader := bytes.NewReader(buf)
	doc, _ := goquery.NewDocumentFromReader(bReader)
	mercariItemsSelection := doc.Find(".items-box")
	return exportItemsFromSelection(mercariItemsSelection)
}

func exportItemsFromSelection(mercariItemsSelection *goquery.Selection) []MercariItem {
	mercariItems := make([]MercariItem, mercariItemsSelection.Length())

	mercariItemsSelection.Each(func(idx int, s *goquery.Selection) {
		// get the product names from the selection
		pNames := mercariItemsSelection.Find(".items-box-name")
		pNames.Each(func(idx int, s *goquery.Selection) {
			mercariItems[idx].productName = s.Text()
		})
		// get the prices from the selection.
		prices := mercariItemsSelection.Find(".items-box-price")
		prices.Each(func(idx int, s *goquery.Selection) {
			mercariItems[idx].price = convNum(s.Text())
		})
	})

	return mercariItems
}

func convNum(s string) int {
	n := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if '0' <= s[i] && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else {
			continue
		}
	}
	return n
}
