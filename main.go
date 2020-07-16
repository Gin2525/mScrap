package main

import (
	"fmt"
	"mScraper/mailsender"
	"mScraper/scraper"
)

func main() {
	keyword := "十三機兵防衛圏"
	queries := map[string]string{
		"status_on_sale":                "1",
		"category_root":                 "5",
		"category_child":                "76",
		"category_grand_child%5B702%5D": "1"}

	urlStructure := scraper.SearchURL{Keyword: keyword, Queries: queries}
	items := scraper.FetchMercariItems(urlStructure)

	msg := urlStructure.RenderURL() + "\n\n"
	for _, item := range items {
		msg += item.String()
	}

	mail := new(mailsender.Mail)
	mail.From = "username@gmail.com"
	mail.Username = "username@gmail.com"
	mail.Password = "password"
	mail.To = "email address to send"
	mail.Sub = "title"
	mail.Msg = msg

	if err := mailsender.SendGmail(*mail); err != nil {
		fmt.Println(err)
		fmt.Println("failure...")
	} else {
		fmt.Println("success!")
	}
}
