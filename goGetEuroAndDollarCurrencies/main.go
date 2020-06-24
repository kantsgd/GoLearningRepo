package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const URL = "https://www.tcmb.gov.tr/kurlar/today.xml"

type Currencies struct {
	XMLName    xml.Name   `xml:"Tarih_Date"`
	Currencies []Currency `xml:"Currency"`
}

type Currency struct {
	XMLName xml.Name `xml:"Currency"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"CurrencyName"`
	Buying  string   `xml:"ForexBuying"`
	Selling string   `xml:"ForexSelling"`
}

func getXmlFile(url string) (error, []byte) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	xmlfile, err := ioutil.ReadAll(resp.Body)

	return err, xmlfile
}

func fetchCurrencies() {
	err, xmlfile := getXmlFile(URL)
	if err != nil {
		panic(err)
	}

	var currencies Currencies
	xml.Unmarshal(xmlfile, &currencies)

	fmt.Println(currencies.Currencies[0].Name + " Buying:" + currencies.Currencies[0].Buying + " Selling: " + currencies.Currencies[0].Selling)
	fmt.Println(currencies.Currencies[3].Name + " Buying:" + currencies.Currencies[3].Buying + " Selling: " + currencies.Currencies[3].Selling)
}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			fetchCurrencies()
		}
	}
}
