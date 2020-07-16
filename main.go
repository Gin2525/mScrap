package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

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

// SearchURL have material to render url
type SearchURL struct {
	keyword string
	queries map[string]string
}

func buildURLStructure(keyword string, queries map[string]string) *SearchURL {
	return &SearchURL{keyword, queries}
}

func (urlStructure SearchURL) renderURL() string {
	url := MercariSerchBaseURL + "?keyword=" + urlStructure.keyword + "&"
	for k, v := range urlStructure.queries {
		url += k + "=" + v + "&"
	}
	return url[:len(url)-1]
}

func fetchMercariItems(urlStructure SearchURL) []MercariItem {
	url := urlStructure.renderURL()
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

func main() {
	// url := "https://www.youtube.com/channel/UCkkxn2ldlFUMupTlXU8meAw/videos"

	// res, _ := http.Get(url)

	// // read
	// buf, _ := ioutil.ReadAll(res.Body)

	// // char code
	// det := chardet.NewTextDetector()
	// detRslt, _ := det.DetectBest(buf)
	// fmt.Println(detRslt.Charset)
	// // => EUC-JP

	// // convert char code
	// bReader := bytes.NewReader(buf)
	// reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	// // HTML parse
	// doc, _ := goquery.NewDocumentFromReader(reader)

	// rslt := doc.Find("title").Text()

	// fmt.Println(fetchMinimumValueByProductKey("203_4984995903644"))
	keyword := "十三機兵防衛圏"
	queries := map[string]string{
		"status_on_sale":                "1",
		"category_root":                 "5",
		"category_child":                "76",
		"category_grand_child%5B702%5D": "1"}
	urlStructure := SearchURL{keyword, queries}
	items := fetchMercariItems(urlStructure)
	for _,item := range items{
		fmt.Println(item)
	}
}
