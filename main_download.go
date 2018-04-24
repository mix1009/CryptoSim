// Copyright (c) 2018 Chun-Koo Park

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func download_historical_page(date string) {
	os.MkdirAll("data", os.ModePerm)

	filename := fmt.Sprintf("data/%s", date)
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		// download
		fmt.Printf("%s: downloading to data directory\n", date)
		resp, err := http.Get("https://coinmarketcap.com/historical/" + date)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile(filename, body, 0644)
	} else {
		fmt.Printf("%s: already downloaded\n", date)
	}
}

func download_historical_index() {
	response, err := http.Get("https://coinmarketcap.com/historical/")
	if err != nil {
		return
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, "/historical/20") {

			date := strings.TrimRight(href, "/")
			date = strings.Replace(date, "/historical/", "", -1)
			//fmt.Printf("a = %s\n", href)
			if len(date) != 8 {
				return
			}
			//fmt.Printf("date = %s\n", date)
			download_historical_page(date)
		}
	})

}

func main() {
	download_historical_index()
}
