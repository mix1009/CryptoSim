// Copyright (c) 2018 Chun-Koo Park

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func trim(txt string) string {
	return strings.Trim(txt, " \t\n")
}

func trimSel(s *goquery.Selection) string {
	return trim(s.Text())
}

func toNum(s *goquery.Selection) int {
	str := s.Text()
	str = strings.Replace(str, "$", "", -1)
	str = strings.Replace(str, "*", "", -1)
	str = strings.Replace(str, ",", "", -1)

	val, err := strconv.Atoi(trim(str))

	if err != nil {
		return -1
	}

	return val
}

func toFloat(s *goquery.Selection) float64 {
	str := s.Text()
	str = strings.Replace(str, "$", "", -1)
	str = strings.Replace(str, "*", "", -1)
	str = strings.Replace(str, ",", "", -1)

	val, err := strconv.ParseFloat(trim(str), 0)

	if err != nil {
		return -1
	}

	return val
}

func parse_historical(filename string) {
	file, err := os.Open(filename)

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		//log.Fatal(err)
	}

	date := strings.Replace(filename, "data/", "", 1)

	idx := 0
	debugPrint := false

	totalMarketCap := 0
	bitcoinMarketCap := 0

	doc.Find("table[id=currencies-all] tbody tr").Each(func(i int, s *goquery.Selection) {
		idx += 1
		if debugPrint {
			fmt.Printf("---------------------\n")
		}

		num := 0
		name := ""
		symbol := ""
		marketcap := 0
		price := 0.0
		circulating_supply := 0
		volume := 0

		s.Find(".text-center").Each(func(j int, t *goquery.Selection) {
			num = toNum(t)
		})
		s.Find("a.currency-name-container").Each(func(j int, t *goquery.Selection) {
			name = trimSel(t)
		})
		s.Find(".col-symbol").Each(func(j int, t *goquery.Selection) {
			symbol = trimSel(t)
		})
		s.Find(".market-cap").Each(func(j int, t *goquery.Selection) {
			marketcap = toNum(t)
		})
		s.Find(".price").Each(func(j int, t *goquery.Selection) {
			price = toFloat(t)
		})
		s.Find(".circulating-supply").Each(func(j int, t *goquery.Selection) {
			circulating_supply = toNum(t)
		})
		s.Find(".volume").Each(func(j int, t *goquery.Selection) {
			volume = toNum(t)
		})

		if debugPrint {
			fmt.Printf("#=[%d]\n", num)
			fmt.Printf("name=[%s]\n", name)
			fmt.Printf("symbol=[%s]\n", symbol)
			fmt.Printf("price=[%g]\n", price)
			fmt.Printf("marketcap=[%d]\n", marketcap)
			fmt.Printf("circulating-supply=[%d]\n", circulating_supply)
			fmt.Printf("volume=[%d]\n", volume)
		}
		if price == 0.0 || symbol == "" {
			return
		}
		if symbol == "BTC" {
			bitcoinMarketCap = marketcap
		}

		dbInsert(date, num, name, symbol, price, marketcap, circulating_supply, volume)

	})
	doc.Find("span[id=total-marketcap]").Each(func(i int, s *goquery.Selection) {
		totalMarketCap = toNum(s)
	})

	//fmt.Printf("totalMarketCap = %d\n", totalMarketCap)
	//fmt.Printf("bitcoinMarketCap = %d\n", bitcoinMarketCap)

	if totalMarketCap > 0 && bitcoinMarketCap > 0 {
		totalVolume := 0
		dbInsertGlobalData(date, totalMarketCap, bitcoinMarketCap, totalVolume)
	}
}

func parse_all() {
	files, err := filepath.Glob("data/20??????")
	if err != nil {
		return
	}
	for _, file := range files {
		fmt.Printf("parsing %s\n", file)
		parse_historical(file)
	}
}

func insert_schema() {
	buf, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		return
	}

	str := string(buf)

	db := getDB()

	for _, s := range strings.Split(str, "\n\n") {
		s = strings.Trim(s, " \n")
		if len(s) == 0 {
			continue
		}

		_, err := db.Exec(s)

		if err != nil {
			fmt.Printf("error = %s\n", err)
		}
	}
}

func main() {
	loadConfig()
	insert_schema()
	//parse_historical("data/20180408")
	parse_all()
}
