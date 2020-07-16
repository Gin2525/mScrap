package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	// "github.com/saintfish/chardet"
	// "golang.org/x/net/html/charset"
)

// MercariItem means item produced
type MercariItem struct {
	productName string
	price       int
}

// MercariURL means seach page in mercari
const MercariURL string = "https://www.mercari.com/jp/search/"

func appendKeyword(url, name string) string {
	return url + "?keyword=" + name + "&"
}

// 1:now on sale
func appendStatusOnSale(url string, status int) string {
	return url + "?status_on_sale=" + strconv.Itoa(status) + "&"
}

func fetchValueByName(name string) {

	url := appendStatusOnSale(appendKeyword(MercariURL, name), 1)
	res, _ := http.Get(url)
	// read
	buf, _ := ioutil.ReadAll(res.Body)
	bReader := bytes.NewReader(buf)
	doc, _ := goquery.NewDocumentFromReader(bReader)

	mercariItemsSelection := doc.Find(".items-box")
	mercariItems := make([]MercariItem, mercariItemsSelection.Length())

	mercariItemsSelection.Each(func(idx int, s *goquery.Selection) {
		pNames := mercariItemsSelection.Find("h3")
		pNames.Each(func(idx int, s *goquery.Selection) {
			mercariItems[idx].productName = s.Text()
		})
		mercariItemsSelection.Find(".items-box-price").Each(func(idx int, s *goquery.Selection) {
			mercariItems[idx].price = convNum(s.Text()[2:])
		})
	})

	for i, item := range mercariItems {
		fmt.Println(i, item)
	}
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

	fetchValueByName("十三機兵防衛圏")
}
