// Copyright (c) 2018 Chun-Koo Park

package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

func calculateGains(startDate string, endDate string, numAssets int) (float64, float64) {

	m := make(map[string]float64)

	totalIn := 0.0
	totalOut := 0.0
	debugPrint := false

	for i := 1; i <= numAssets; i++ {
		symbol, price := dbGetPriceForDateNo(startDate, i)
		fund := 1000.0
		if debugPrint {
			fmt.Printf("%d: %s %g coins=%g\n", i, symbol, price, fund/price)
		}
		if price == 0 {
			continue
		}

		m[symbol] = fund / price
		totalIn += fund
	}

	for k, coins := range m {
		price := dbGetPriceForDateSymbol(endDate, k)
		totalOut += coins * price
		if debugPrint {
			fmt.Printf("%s %g price*coin(%g)=%g\n", k, price, coins, coins*price)
		}
	}

	if debugPrint {
		fmt.Printf("# %s - %s : ", startDate, endDate)
		fmt.Printf("%.2f%% (%g -> %g)\n", (totalOut-totalIn)*100.0/totalIn, totalIn, totalOut)
	}

	return totalIn, totalOut
}

func simulate(weekInterval int, numAssets int, startDate string, endDate string) (float64, float64, float64, int) {
	files, err := filepath.Glob("data/20??????")
	if err != nil {
		return 0, 0, 0, 0
	}

	for i := 0; i < len(files); i++ {
		files[i] = strings.Replace(files[i], "data/", "", -1)
	}

	sort.Strings(files)

	diffWeeks := weekInterval

	delta := 0
	totalIn := 0.0
	totalOut := 0.0
	numSimulations := 0
	for true {
		startIdx := len(files) - 1 - diffWeeks - delta
		endIdx := startIdx + diffWeeks

		if startIdx < 0 {
			break
		}
		if len(startDate) > 0 && files[startIdx] < startDate {
			delta++
			continue
		}
		if len(endDate) > 0 && files[endIdx] > endDate {
			delta++
			continue
		}

		in, out := calculateGains(files[startIdx], files[endIdx], numAssets)
		totalIn += in
		totalOut += out
		numSimulations++

		delta++
	}

	//fmt.Printf("profit(%dw, #%d) = %.2f%%\n", weekInterval, numAssets, (totalOut-totalIn)*100.0/totalIn)
	//fmt.Printf("in/out = %g/%g\n", totalIn, totalOut)
	//fmt.Printf("%d\t %d\t %.2f%%\n", weekInterval, numAssets, (totalOut-totalIn)*100.0/totalIn)
	profitPercent := (totalOut - totalIn) * 100.0 / totalIn
	return profitPercent, totalIn, totalOut, numSimulations
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func xrange(start int, end int) []int {
	var arr []int

	for idx := start; idx <= end; idx++ {
		arr = append(arr, idx)
	}
	return arr
}

var wg sync.WaitGroup

func simulateGoRoutine(ret *float64, weekInterval int, numAssets int, startDate string, endDate string) {
	defer wg.Done()
	*ret, _, _, _ = simulate(weekInterval, numAssets, startDate, endDate)
}

func runSimulation(weekArray []int, numArray []int, startDate string, endDate string, outputFormat string) {
	useGoRoutine := false

	if outputFormat == "plotly" {
		for i := 0; i < len(numArray); i++ {
			fmt.Printf(",%d", numArray[i])
		}

		fmt.Printf("\n")
	}
	for _, weekInterval := range weekArray {
		result := make([]float64, len(numArray))

		if useGoRoutine {
			for idx, numAssets := range numArray {
				wg.Add(1)
				go simulateGoRoutine(&result[idx], weekInterval, numAssets, startDate, endDate)

				if idx%4 == 0 && idx > 0 {
					wg.Wait()
				}
			}
			wg.Wait()
		} else {
			for idx, numAssets := range numArray {
				result[idx], _, _, _ = simulate(weekInterval, numAssets, startDate, endDate)
			}
		}

		if outputFormat == "plotly" {
			fmt.Printf("%d", weekInterval)
			for _, v := range result {
				fmt.Printf(",%.2f", v)
			}
			fmt.Printf("\n")
		} else if outputFormat == "csv" {
			for i, v := range result {
				fmt.Printf("%d,%d,%.2f\n", weekInterval, numArray[i], v)
			}
		} else {
			for _, v := range result {
				fmt.Printf("%.2f\n", v)
			}
		}
	}

}

func main() {
	loadConfig()
	//simulate(4, 10, "20160101", "20161231")
	//simulate(4, 10, "20180101", "")

	weekArray := []int{1, 4, 13, 26, 52}
	//weekArray := xrange(1,26)
	//weekArray := xrange(1,52)

	numArray := []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 30, 40, 50, 75, 100}
	//numArray := xrange(1,200)

	startDate := "20170701"
	endDate := ""

	format := "plotly"
	//format := "csv"
	//format := ""

	runSimulation(weekArray, numArray, startDate, endDate, format)
}
