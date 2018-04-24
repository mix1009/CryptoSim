// Copyright (c) 2018 Chun-Koo Park

package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var gDB *sql.DB = nil
var gDBUseCount int = 0

var dbName string
var dbUser string
var dbPassword string

func getDB() *sql.DB {

	if gDBUseCount > 1000 {
		gDBUseCount = 0
		gDB.Close()
		gDB = nil
	}

	if gDB != nil {
		return gDB
	}

	if len(dbName) == 0 {
		dbName = viper.GetString("database.name")
		dbUser = viper.GetString("database.user")
		dbPassword = viper.GetString("database.password")
	}

	var err error

	gDB, err = sql.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	gDBUseCount++

	return gDB
}

func dbInsert(pricedate string, no int, name string, symbol string, price float64, marketcap int, circulating_supply int, volume int) {
	db := getDB()

	db.Exec("INSERT INTO coinprice (pricedate,no,name,symbol,price,marketcap,circulating_supply,volume) VALUES (?,?,?,?,?,?,?,?)", pricedate, no, name, symbol, price, marketcap, circulating_supply, volume)
}

func dbInsertGlobalData(pricedate string, totalmarketcap int, bitcoinmarketcap int, volume int) {
	db := getDB()

	db.Exec("INSERT INTO globaldata (pricedate,totalmarketcap,bitcoinmarketcap,volume) VALUES (?,?,?,?)", pricedate, totalmarketcap, bitcoinmarketcap, volume)
}

func dbGetPriceForDateNo(pricedate string, no int) (string, float64) {
	db := getDB()

	rows, err := db.Query("SELECT symbol,price FROM coinprice WHERE pricedate=? AND no=?", pricedate, no)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var symbol string
		var price float64
		err := rows.Scan(&symbol, &price)
		if err != nil {
			panic(err.Error())
		}
		return symbol, price

	}
	return "", 0
}

func dbGetPriceForDateSymbol(pricedate string, symbol string) float64 {
	db := getDB()

	rows, err := db.Query("SELECT price FROM coinprice WHERE pricedate=? AND symbol=?", pricedate, symbol)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var price float64
		err := rows.Scan(&price)
		if err != nil {
			panic(err.Error())
		}
		return price

	}
	return 0
}
