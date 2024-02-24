package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PrayerTime struct {
	Name string
	Time string
}

func main() {
	sourceUrl := "https://www.toptiershop.id/"

	resp, err := http.Get(sourceUrl)

	if err != nil {
		fmt.Println("error when http.Get", err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {

		fmt.Println("error when init new document", err)

	}

	// fmt.Println("result document", doc.Text())

	// rows := make([]PrayerTime, 0)

	doc.Find(".tab-content").Find(".row").Children().Each(func(i int, sel *goquery.Selection) {
		name := strings.ReplaceAll(strings.ReplaceAll(sel.Text(), " ", ""), "\n", "")
		image, _ := sel.Find("img").Attr("src")
		// rows = append(rows, *row)
		fmt.Println(name)
		fmt.Println(image)
	})

}
