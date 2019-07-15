package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strconv"
	"encoding/json"
	"os"
	"strings"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	// consume list of all symbols

	jsonFile, err := os.Open("symbolHash.json")

	// if we os.Open returns an error then handle it
	
	if err != nil {
	    fmt.Println(err)
	}
	
	// fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	m := map[string]interface{}{}

	err = json.Unmarshal([]byte(byteValue), &m)
	
	// iterate through and make api for opening price of each, append to map

	message := make(chan float64)

	outputArray := make(map[string]float64)

	count := 0

	for k := range m {

		go APICall(k, message, outputArray)

		consoleOutput := <-message

		fmt.Println("output: ", k, consoleOutput)	

		count = count + 1

		if count % 50 == 0{

			fmt.Println(outputArray)
		}
	}

	// jsonErr := json.Unmarshal(body, &people1)

	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }

	// convert map to json and write file

}

func APICall(k string, message chan float64, outputArray map[string]float64 ) {
	
		url := []string{"https://cloud.iexapis.com/beta/stock/", k, "/quote/open?token=sk_5c1a1ec78f534b179da588b787245fe6"}

		openingPriceEndpoint := http.Client{

			Timeout: time.Second * 20, // Maximum of 2 secs
		}

		req, err := http.NewRequest(http.MethodGet, strings.Join(url, ""), nil)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "TTP-FP")

		res, getErr := openingPriceEndpoint.Do(req)

		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)

		if readErr != nil {
			log.Fatal(readErr)
		}

		formatted, _ := strconv.ParseFloat(string(body), 32) 

		outputArray[k] = formatted

		message <- formatted
}
