package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"utils"

	"github.com/gorilla/mux"
)

func FetchOpeningPrice() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		allParams := mux.Vars(r)

		param := allParams["symbol"]

		url := []string{"https://cloud.iexapis.com/beta/stock/", param, "/quote/ohlc?token=", os.Getenv("IEX_API_KEY")}

		openingPriceEndpoint := http.Client{

			Timeout: time.Second * 20,
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

		var f interface{}

		err = json.Unmarshal(body, &f)

		// this is how you would process the data on the
		// server- I"m going to send the whole JSON
		// back to the client app to process there

		// theJason := f.(map[string]interface{})

		// fmt.Println(theJason["previousClose"])

		// previousClose := theJason["previousClose"]

		// returnJason := make(map[string]interface{})

		// returnJason[param] = previousClose

		utils.ResponseJSON(w, f)
	}
}
